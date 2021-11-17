package hubspot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	oauthTokenPath = "oauth/v1/token"

	GrantTypeRefreshToken = "refresh_token"
)

type OAuthTokenRetriever interface {
	RetrieveToken() (*OAuthToken, error)
}

type OAuthTokenManager struct {
	oauthPath string

	HTTPClient *http.Client
	Config     *OAuthConfig
	Token      *OAuthToken
}

var _ OAuthTokenRetriever = (*OAuthTokenManager)(nil)

func (otm *OAuthTokenManager) RetrieveToken() (*OAuthToken, error) {
	if otm.Token.valid() {
		return otm.Token, nil
	}

	tokenByte, err := otm.fetchTokenFromHubSpot()
	if err != nil {
		return nil, err
	}

	if len(tokenByte) == 0 {
		return nil, errors.New("missing authorization token")
	}

	return otm.refreshToken(tokenByte)
}

func (otm *OAuthTokenManager) fetchTokenFromHubSpot() (tokenByte []byte, err error) {
	if err := otm.Config.valid(); err != nil {
		return nil, err
	}

	res, err := otm.HTTPClient.PostForm(otm.oauthPath, otm.Config.convertToFormData())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if isErrorStatusCode(res.StatusCode) {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to authorize: %s", string(b))
	}

	return ioutil.ReadAll(res.Body)
}

func (otm *OAuthTokenManager) refreshToken(tokenByte []byte) (newToken *OAuthToken, err error) {
	newToken = new(OAuthToken)
	if err := json.Unmarshal(tokenByte, newToken); err != nil {
		return nil, err
	}
	newToken.setExpiry()

	if !newToken.valid() {
		return nil, errors.New("invalid authorization token")
	}
	otm.Token = newToken

	return otm.Token, nil
}

type OAuthConfig struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

func (oc *OAuthConfig) convertToFormData() url.Values {
	form := url.Values{}
	form.Set("grant_type", oc.GrantType)
	form.Set("client_id", oc.ClientID)
	form.Set("client_secret", oc.ClientSecret)
	form.Set("refresh_token", oc.RefreshToken)
	return form
}

func (oc *OAuthConfig) valid() error {
	var missingList []string
	if oc.GrantType == "" {
		missingList = append(missingList, "GrantType")
	}

	if oc.ClientID == "" {
		missingList = append(missingList, "ClientID")
	}

	if oc.ClientSecret == "" {
		missingList = append(missingList, "ClientSecret")
	}

	if oc.RefreshToken == "" {
		missingList = append(missingList, "RefreshToken")
	}

	if len(missingList) > 0 {
		return fmt.Errorf("missing required options: %v", strings.Join(missingList, ","))
	}

	return nil
}

type OAuthToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int       `json:"expires_in"`
	Expiry       time.Time `json:"-"`
}

// timeNow is time.Now, but it has been redefined as a test variable.
var timeNow = time.Now

func (ot *OAuthToken) setExpiry() {
	// To prevent last minute expirations, the expiration date will be accelerated by 10 minutes.
	ot.Expiry = timeNow().Add(time.Duration(ot.ExpiresIn) * time.Second).Add(-10 * time.Minute)
}

func (ot *OAuthToken) valid() bool {
	return ot != nil && ot.AccessToken != "" && ot.RefreshToken != "" && ot.expired()
}

func (ot *OAuthToken) expired() bool {
	if ot.Expiry.IsZero() {
		return false
	}
	return ot.Expiry.After(timeNow())
}

package hubspot

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// For test only

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

type MockConfig struct {
	Status int
	Header http.Header
	Body   []byte
}

// Client

func NewMockClient(conf *MockConfig) *Client {
	cli := &Client{
		HTTPClient: NewMockHTTPClient(conf),
		baseURL:    defaultBaseURL,
		apiVersion: defaultAPIVersion,
	}
	cli.CRM = newCRM(cli)
	SetPrivateAppToken("token")(cli)

	return cli
}

func NewMockHTTPClient(conf *MockConfig) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: conf.Status,
				Body:       io.NopCloser(bytes.NewBuffer(conf.Body)),
				Header:     conf.Header,
			}
		}),
	}
}

// OAuth

func NewMockOAuthTokenRetriever() OAuthTokenRetriever {
	return &MockOAuthTokenManager{}
}

type MockOAuthTokenManager struct{}

func (otm *MockOAuthTokenManager) RetrieveToken() (*OAuthToken, error) {
	return &OAuthToken{
		AccessToken:  "test_access_token",
		RefreshToken: "test_access_token",
		ExpiresIn:    21600,
		Expiry:       time.Date(9999, 12, 31, 0, 0, 0, 0, time.Local),
	}, nil
}

func MockTimeNow() func() {
	mockTime := time.Date(2020, 12, 31, 12, 0, 0, 0, time.UTC)
	timeNow = func() time.Time { return mockTime }
	return func() { timeNow = time.Now }
}

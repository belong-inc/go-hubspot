package hubspot_test

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/courtyard-nft/go-hubspot/"
	"github.com/google/go-cmp/cmp"
)

func TestOAuthTokenManager_RetrieveToken(t *testing.T) {
	f := hubspot.MockTimeNow()
	defer f()

	type fields struct {
		HTTPClient *http.Client
		Config     *hubspot.OAuthConfig
		Token      *hubspot.OAuthToken
	}
	tests := []struct {
		name    string
		fields  fields
		want    *hubspot.OAuthToken
		wantErr error
	}{
		{
			name: "Success",
			fields: fields{
				HTTPClient: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Body:   []byte(`{"access_token": "test_access_token","refresh_token": "test_refresh_token","expires_in": 21600}`),
				}),
				Config: &hubspot.OAuthConfig{
					GrantType:    "grant_type",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
			want: &hubspot.OAuthToken{
				AccessToken:  "test_access_token",
				RefreshToken: "test_refresh_token",
				ExpiresIn:    21600,
				Expiry:       time.Date(2020, 12, 31, 17, 50, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "Invalid got token",
			fields: fields{
				HTTPClient: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Body:   []byte(`{"access_token": "","refresh_token": "","expires_in": 0}`),
				}),
				Config: &hubspot.OAuthConfig{
					GrantType:    "grant_type",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
			want:    nil,
			wantErr: errors.New("invalid authorization token"),
		},
		{
			name: "Missing oauth config",
			fields: fields{
				HTTPClient: nil,
				Config:     &hubspot.OAuthConfig{},
			},
			want:    nil,
			wantErr: errors.New("missing required options: GrantType,ClientID,ClientSecret,RefreshToken"),
		},
		{
			name: "Error response 4XX",
			fields: fields{
				HTTPClient: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Body:   []byte(`{"message": "bad request"}`),
				}),
				Config: &hubspot.OAuthConfig{
					GrantType:    "grant_type",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
			want:    nil,
			wantErr: errors.New(`failed to authorize: {"message": "bad request"}`),
		},
		{
			name: "Error response 5XX",
			fields: fields{
				HTTPClient: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusInternalServerError,
					Body:   []byte(`{"message": "internal server error"}`),
				}),
				Config: &hubspot.OAuthConfig{
					GrantType:    "grant_type",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
			want:    nil,
			wantErr: errors.New(`failed to authorize: {"message": "internal server error"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			otm := &hubspot.OAuthTokenManager{
				HTTPClient: tt.fields.HTTPClient,
				Config:     tt.fields.Config,
				Token:      tt.fields.Token,
			}
			got, err := otm.RetrieveToken()
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("RetrieveToken() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("RetrieveToken() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestOAuthTokenManager_fetchTokenFromHubSpot(t *testing.T) {
	type fields struct {
		HTTPClient *http.Client
		Config     *hubspot.OAuthConfig
		Token      *hubspot.OAuthToken
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr error
	}{
		{
			name: "Success",
			fields: fields{
				HTTPClient: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Body:   []byte(`{"access_token": "test_access_token","refresh_token": "test_refresh_token","expires_in": 21600}`),
				}),
				Config: &hubspot.OAuthConfig{
					GrantType:    "grant_type",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
			want:    []byte(`{"access_token": "test_access_token","refresh_token": "test_refresh_token","expires_in": 21600}`),
			wantErr: nil,
		},
		{
			name: "Missing oauth config",
			fields: fields{
				HTTPClient: nil,
				Config:     &hubspot.OAuthConfig{},
			},
			want:    nil,
			wantErr: errors.New("missing required options: GrantType,ClientID,ClientSecret,RefreshToken"),
		},
		{
			name: "Error response 4XX",
			fields: fields{
				HTTPClient: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Body:   []byte(`{"message": "bad request"}`),
				}),
				Config: &hubspot.OAuthConfig{
					GrantType:    "grant_type",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
			want:    nil,
			wantErr: errors.New(`failed to authorize: {"message": "bad request"}`),
		},
		{
			name: "Error response 5XX",
			fields: fields{
				HTTPClient: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusInternalServerError,
					Body:   []byte(`{"message": "internal server error"}`),
				}),
				Config: &hubspot.OAuthConfig{
					GrantType:    "grant_type",
					ClientID:     "client_id",
					ClientSecret: "client_secret",
					RefreshToken: "refresh_token",
				},
			},
			want:    nil,
			wantErr: errors.New(`failed to authorize: {"message": "internal server error"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			otm := &hubspot.OAuthTokenManager{
				HTTPClient: tt.fields.HTTPClient,
				Config:     tt.fields.Config,
				Token:      tt.fields.Token,
			}
			got, err := hubspot.ExportFetchTokenFromHubSpot(otm)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("fetchTokenFromHubSpot() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("fetchTokenFromHubSpot() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestOAuthTokenManager_refreshToken(t *testing.T) {
	f := hubspot.MockTimeNow()
	defer f()

	type fields struct {
		HTTPClient *http.Client
		Config     *hubspot.OAuthConfig
		Token      *hubspot.OAuthToken
	}
	type args struct {
		tokenByte []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.OAuthToken
		wantErr error
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				tokenByte: []byte(`{"access_token": "test_access_token","refresh_token": "test_refresh_token","expires_in": 21600}`),
			},
			want: &hubspot.OAuthToken{
				AccessToken:  "test_access_token",
				RefreshToken: "test_refresh_token",
				ExpiresIn:    21600,
				Expiry:       time.Date(2020, 12, 31, 17, 50, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name:   "Missing new access token",
			fields: fields{},
			args: args{
				tokenByte: []byte(`{"access_token": "","refresh_token": "test_refresh_token","expires_in": 21600}`),
			},
			want:    nil,
			wantErr: errors.New("invalid authorization token"),
		},
		{
			name:   "Invalid new token json",
			fields: fields{},
			args: args{
				tokenByte: []byte(`{"invalid_field": "test"}`),
			},
			want:    nil,
			wantErr: errors.New("invalid authorization token"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			otm := &hubspot.OAuthTokenManager{
				HTTPClient: tt.fields.HTTPClient,
				Config:     tt.fields.Config,
				Token:      tt.fields.Token,
			}
			got, err := hubspot.ExportRefreshToken(otm, tt.args.tokenByte)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("refreshToken() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("refreshToken() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestOAuthConfig_convertToFormData(t *testing.T) {
	type fields struct {
		GrantType    string
		ClientID     string
		ClientSecret string
		RefreshToken string
	}
	tests := []struct {
		name   string
		fields fields
		want   url.Values
	}{
		{
			name: "Success all field exists",
			fields: fields{
				GrantType:    "refresh_token",
				ClientID:     "w38dgh4e93rg9wjs",
				ClientSecret: "0jhfds9309rqlq08",
				RefreshToken: "sdfs3r-saefkj3r8-shgslay",
			},
			want: url.Values{
				"grant_type":    []string{"refresh_token"},
				"client_id":     []string{"w38dgh4e93rg9wjs"},
				"client_secret": []string{"0jhfds9309rqlq08"},
				"refresh_token": []string{"sdfs3r-saefkj3r8-shgslay"},
			},
		},
		{
			name: "Success missing grant type fields",
			fields: fields{
				ClientID:     "w38dgh4e93rg9wjs",
				ClientSecret: "0jhfds9309rqlq08",
				RefreshToken: "sdfs3r-saefkj3r8-shgslay",
			},
			want: url.Values{
				"grant_type":    []string{""},
				"client_id":     []string{"w38dgh4e93rg9wjs"},
				"client_secret": []string{"0jhfds9309rqlq08"},
				"refresh_token": []string{"sdfs3r-saefkj3r8-shgslay"},
			},
		},
		{
			name: "Success missing client id fields",
			fields: fields{
				GrantType:    "refresh_token",
				ClientSecret: "0jhfds9309rqlq08",
				RefreshToken: "sdfs3r-saefkj3r8-shgslay",
			},
			want: url.Values{
				"grant_type":    []string{"refresh_token"},
				"client_id":     []string{""},
				"client_secret": []string{"0jhfds9309rqlq08"},
				"refresh_token": []string{"sdfs3r-saefkj3r8-shgslay"},
			},
		},
		{
			name: "Success missing client secret fields",
			fields: fields{
				GrantType:    "refresh_token",
				ClientID:     "w38dgh4e93rg9wjs",
				RefreshToken: "sdfs3r-saefkj3r8-shgslay",
			},
			want: url.Values{
				"grant_type":    []string{"refresh_token"},
				"client_id":     []string{"w38dgh4e93rg9wjs"},
				"client_secret": []string{""},
				"refresh_token": []string{"sdfs3r-saefkj3r8-shgslay"},
			},
		},
		{
			name: "Success missing refresh token fields",
			fields: fields{
				GrantType:    "refresh_token",
				ClientID:     "w38dgh4e93rg9wjs",
				ClientSecret: "0jhfds9309rqlq08",
			},
			want: url.Values{
				"grant_type":    []string{"refresh_token"},
				"client_id":     []string{"w38dgh4e93rg9wjs"},
				"client_secret": []string{"0jhfds9309rqlq08"},
				"refresh_token": []string{""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oc := &hubspot.OAuthConfig{
				GrantType:    tt.fields.GrantType,
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
				RefreshToken: tt.fields.RefreshToken,
			}
			got := hubspot.ExportConvertToFormData(oc)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("convertToFormData() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestOAuthConfig_valid(t *testing.T) {
	type fields struct {
		GrantType    string
		ClientID     string
		ClientSecret string
		RefreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "Valid",
			fields: fields{
				GrantType:    "refresh_token",
				ClientID:     "w38dgh4e93rg9wjs",
				ClientSecret: "0jhfds9309rqlq08",
				RefreshToken: "sdfs3r-saefkj3r8-shgslay",
			},
			wantErr: nil,
		},
		{
			name:    "Invalid missing all fields",
			fields:  fields{},
			wantErr: errors.New("missing required options: GrantType,ClientID,ClientSecret,RefreshToken"),
		},
		{
			name: "Invalid missing grant type",
			fields: fields{
				ClientID:     "w38dgh4e93rg9wjs",
				ClientSecret: "0jhfds9309rqlq08",
				RefreshToken: "sdfs3r-saefkj3r8-shgslay",
			},
			wantErr: errors.New("missing required options: GrantType"),
		},
		{
			name: "Invalid missing client id",
			fields: fields{
				GrantType:    "refresh_token",
				ClientSecret: "0jhfds9309rqlq08",
				RefreshToken: "sdfs3r-saefkj3r8-shgslay",
			},
			wantErr: errors.New("missing required options: ClientID"),
		},
		{
			name: "Invalid missing client secret",
			fields: fields{
				GrantType:    "refresh_token",
				ClientID:     "w38dgh4e93rg9wjs",
				RefreshToken: "sdfs3r-saefkj3r8-shgslay",
			},
			wantErr: errors.New("missing required options: ClientSecret"),
		},
		{
			name: "Invalid missing refresh token",
			fields: fields{
				GrantType:    "refresh_token",
				ClientID:     "w38dgh4e93rg9wjs",
				ClientSecret: "0jhfds9309rqlq08",
			},
			wantErr: errors.New("missing required options: RefreshToken"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oc := &hubspot.OAuthConfig{
				GrantType:    tt.fields.GrantType,
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
				RefreshToken: tt.fields.RefreshToken,
			}
			err := hubspot.ExportConfigValid(oc)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("valid() error mismatch: want %s got %s", tt.wantErr, err)
			}
		})
	}
}

func TestOAuthToken_setExpiry(t *testing.T) {
	f := hubspot.MockTimeNow()
	defer f()

	type fields struct {
		AccessToken  string
		RefreshToken string
		ExpiresIn    int
		Expiry       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "Success",
			fields: fields{
				ExpiresIn: 21600,
			},
			want: time.Date(2020, 12, 31, 17, 50, 0, 0, time.UTC),
		},
		{
			name: "Success ExpiresIn 0 sec",
			fields: fields{
				ExpiresIn: 0,
			},
			want: time.Date(2020, 12, 31, 11, 50, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ot := &hubspot.OAuthToken{
				AccessToken:  tt.fields.AccessToken,
				RefreshToken: tt.fields.RefreshToken,
				ExpiresIn:    tt.fields.ExpiresIn,
				Expiry:       tt.fields.Expiry,
			}
			hubspot.ExportSetExpiry(ot)
			if diff := cmp.Diff(tt.want, ot.Expiry); diff != "" {
				t.Errorf("setExpiry() mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestOAuthToken_valid(t *testing.T) {
	f := hubspot.MockTimeNow()
	defer f()

	tests := []struct {
		name   string
		fields *hubspot.OAuthToken
		want   bool
	}{
		{
			name: "Valid",
			fields: &hubspot.OAuthToken{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				ExpiresIn:    21600,
				Expiry:       time.Date(2020, 12, 31, 12, 10, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name:   "Invalid token undefined",
			fields: nil,
			want:   false,
		},
		{
			name: "Invalid no access token",
			fields: &hubspot.OAuthToken{
				AccessToken:  "",
				RefreshToken: "refresh_token",
				ExpiresIn:    21600,
				Expiry:       time.Date(2020, 12, 31, 12, 10, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "Invalid no refresh token",
			fields: &hubspot.OAuthToken{
				AccessToken:  "access_token",
				RefreshToken: "",
				ExpiresIn:    21600,
				Expiry:       time.Date(2020, 12, 31, 11, 50, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "Invalid token expired",
			fields: &hubspot.OAuthToken{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				ExpiresIn:    21600,
				Expiry:       time.Date(2020, 12, 31, 11, 50, 0, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hubspot.ExportTokenValid(tt.fields)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("valid() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestOAuthToken_expired(t *testing.T) {
	f := hubspot.MockTimeNow()
	defer f()

	type fields struct {
		AccessToken  string
		RefreshToken string
		ExpiresIn    int
		Expiry       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Token enabled",
			fields: fields{
				Expiry: time.Date(2020, 12, 31, 12, 10, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "Token expired",
			fields: fields{
				Expiry: time.Date(2020, 12, 31, 11, 50, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "Token expiry is zero value",
			fields: fields{
				Expiry: time.Time{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ot := &hubspot.OAuthToken{
				AccessToken:  tt.fields.AccessToken,
				RefreshToken: tt.fields.RefreshToken,
				ExpiresIn:    tt.fields.ExpiresIn,
				Expiry:       tt.fields.Expiry,
			}
			got := hubspot.ExportExpired(ot)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("expired() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

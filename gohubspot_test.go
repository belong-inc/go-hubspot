package hubspot_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	hubspot "github.com/belong-inc/go-hubspot"
	"github.com/google/go-cmp/cmp"
)

var (
	closeDate  = hubspot.HsTime(time.Date(2019, 12, 7, 16, 50, 6, 678000000, time.UTC))
	modifyDate = hubspot.HsTime(time.Date(2019, 12, 7, 16, 50, 6, 678000000, time.UTC))
	createdAt  = hubspot.HsTime(time.Date(2019, 10, 30, 3, 30, 17, 883000000, time.UTC))
	updatedAt  = hubspot.HsTime(time.Date(2019, 12, 7, 16, 50, 6, 678000000, time.UTC))
	archivedAt = hubspot.HsTime(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
)

var cmpTimeOption = cmp.AllowUnexported(hubspot.HsTime{})

func TestNewClient(t *testing.T) {
	type args struct {
		setAuthMethod hubspot.AuthMethod
		opts          []hubspot.Option
	}
	type settings struct {
		client     *http.Client
		baseURL    *url.URL
		apiVersion string
		authMethod hubspot.AuthMethod
	}
	tests := []struct {
		name     string
		args     args
		settings settings
		wantErr  error
	}{
		{
			name: "Success new client and set api key",
			args: args{
				setAuthMethod: hubspot.SetAPIKey("key"),
			},
			settings: settings{
				client:     http.DefaultClient,
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetAPIKey("key"),
			},
			wantErr: nil,
		},
		{
			name: "Success new client and set oauth",
			args: args{
				setAuthMethod: hubspot.SetOAuth(&hubspot.OAuthConfig{
					GrantType:    "grant-type",
					ClientID:     "client-id",
					ClientSecret: "client-secret",
					RefreshToken: "refresh-token",
				}),
			},
			settings: settings{
				client:     http.DefaultClient,
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetOAuth(&hubspot.OAuthConfig{
					GrantType:    "grant-type",
					ClientID:     "client-id",
					ClientSecret: "client-secret",
					RefreshToken: "refresh-token",
				}),
			},
			wantErr: nil,
		},
		{
			name: "Success new client and set private app token",
			args: args{
				setAuthMethod: hubspot.SetPrivateAppToken("token"),
			},
			settings: settings{
				client:     http.DefaultClient,
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetPrivateAppToken("token"),
			},
			wantErr: nil,
		},
		{
			name: "Success new client with custom http client",
			args: args{
				setAuthMethod: hubspot.SetPrivateAppToken("token"),
				opts:          []hubspot.Option{hubspot.WithHTTPClient(&http.Client{Timeout: 100 * time.Second})},
			},
			settings: settings{
				client:     &http.Client{Timeout: 100 * time.Second},
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetPrivateAppToken("token"),
			},
			wantErr: nil,
		},
		{
			name: "Success new client with custom base url",
			args: args{
				setAuthMethod: hubspot.SetPrivateAppToken("token"),
				opts:          []hubspot.Option{hubspot.WithBaseURL(&url.URL{Scheme: "http", Host: "example.com"})},
			},
			settings: settings{
				client:     http.DefaultClient,
				baseURL:    &url.URL{Scheme: "http", Host: "example.com"},
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetPrivateAppToken("token"),
			},
			wantErr: nil,
		},
		{
			name: "Success new client with custom api version",
			args: args{
				setAuthMethod: hubspot.SetPrivateAppToken("token"),
				opts:          []hubspot.Option{hubspot.WithAPIVersion("v0")},
			},
			settings: settings{
				client:     http.DefaultClient,
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: "v0",
				authMethod: hubspot.SetPrivateAppToken("token"),
			},
			wantErr: nil,
		},
		{
			name: "Failed new client because not set auth method",
			args: args{
				setAuthMethod: nil,
			},
			settings: settings{},
			wantErr:  errors.New("the authentication method is not set"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want *hubspot.Client = nil
			if tt.wantErr == nil {
				want = &hubspot.Client{
					HTTPClient: tt.settings.client,
				}
				want.ExportSetAPIVersion(tt.settings.apiVersion)
				want.ExportSetBaseURL(tt.settings.baseURL)
				want.CRM = hubspot.ExportNewCRM(want)
				want.Marketing = hubspot.ExportNewMarketing(want)
				want.Conversation = hubspot.ExportNewConversation(want)
				tt.settings.authMethod(want)
			}

			got, err := hubspot.NewClient(tt.args.setAuthMethod, tt.args.opts...)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("NewClient() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if !reflect.DeepEqual(want, got) {
				t.Errorf("NewClient() response mismatch: want %v got %v", want, got)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	type body struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	type settings struct {
		client     *http.Client
		baseURL    *url.URL
		apiVersion string
		authMethod hubspot.AuthMethod
		crm        *hubspot.CRM
		marketing  *hubspot.Marketing
	}
	type args struct {
		method string
		path   string
		body   *body
		option *hubspot.RequestQueryOption
	}
	type wantReq struct {
		method string
		url    string
		header http.Header
		body   []byte
	}
	tests := []struct {
		name     string
		settings settings
		args     args
		want     wantReq
		wantErr  error
	}{
		{
			name: "Created http.GET request with APIKey",
			settings: settings{
				client: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Body: []byte(`{"id":"001","name":"example"}`),
				}),
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetAPIKey("key"),
				crm:        nil,
			},
			args: args{
				method: http.MethodGet,
				path:   "objects/test",
				body:   nil,
				option: nil,
			},
			want: wantReq{
				method: http.MethodGet,
				url:    "https://api.hubapi.com/objects/test?hapikey=key",
				body:   nil,
				header: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
			wantErr: nil,
		},
		{
			name: "Created http.POST request with APIKey",
			settings: settings{
				client:     hubspot.NewMockHTTPClient(&hubspot.MockConfig{}),
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetAPIKey("key"),
				crm:        nil,
			},
			args: args{
				method: http.MethodPost,
				path:   "objects/test",
				body: &body{
					ID:   "001",
					Name: "example",
				},
				option: nil,
			},
			want: wantReq{
				method: http.MethodPost,
				url:    "https://api.hubapi.com/objects/test?hapikey=key",
				body:   []byte(`{"id":"001","name":"example"}`),
				header: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
			wantErr: nil,
		},
		{
			name: "Created http.POST request with OAuth",
			settings: settings{
				client: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Body:   []byte(`{"access_token": "test_access_token","refresh_token": "test_refresh_token","expires_in": 21600}`),
				}),
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetOAuth(&hubspot.OAuthConfig{
					GrantType:    "grant-type",
					ClientID:     "client-id",
					ClientSecret: "client-secret",
					RefreshToken: "refresh-token",
				}),
				crm: nil,
			},
			args: args{
				method: http.MethodPost,
				path:   "objects/test",
				body: &body{
					ID:   "001",
					Name: "example",
				},
				option: nil,
			},
			want: wantReq{
				method: http.MethodPost,
				url:    "https://api.hubapi.com/objects/test",
				body:   []byte(`{"id":"001","name":"example"}`),
				header: http.Header{
					"Content-Type":  []string{"application/json"},
					"Authorization": []string{"Bearer test_access_token"},
				},
			},
			wantErr: nil,
		},
		{
			name: "Created http.POST request with PrivateApp token",
			settings: settings{
				client:     hubspot.NewMockHTTPClient(&hubspot.MockConfig{}),
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetPrivateAppToken("token"),
				crm:        nil,
			},
			args: args{
				method: http.MethodPost,
				path:   "objects/test",
				body: &body{
					ID:   "001",
					Name: "example",
				},
				option: nil,
			},
			want: wantReq{
				method: http.MethodPost,
				url:    "https://api.hubapi.com/objects/test",
				body:   []byte(`{"id":"001","name":"example"}`),
				header: http.Header{
					"Content-Type":  []string{"application/json"},
					"Authorization": []string{"Bearer token"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &hubspot.Client{
				HTTPClient: tt.settings.client,
			}
			c.ExportSetAPIVersion(tt.settings.apiVersion)
			c.ExportSetBaseURL(tt.settings.baseURL)
			tt.settings.authMethod(c)

			var (
				got *http.Request
				err error
			)
			// Because we are using interface, we need to explicitly specify nil, which has no type information.
			if tt.args.body == nil {
				got, err = c.NewRequest(tt.args.method, tt.args.path, nil, tt.args.option, "application/json")
			} else {
				got, err = c.NewRequest(tt.args.method, tt.args.path, tt.args.body, tt.args.option, "application/json")
			}
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("NewRequest() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if tt.want.method != got.Method {
				t.Errorf("NewRequest() method mismatch: want %s got %s", tt.want.method, got.Method)
				return
			}
			if tt.want.url != got.URL.String() {
				t.Errorf("NewRequest() url mismatch: want %s got %s", tt.want.url, got.URL.String())
				return
			}
			if !reflect.DeepEqual(tt.want.header, got.Header) {
				t.Errorf("NewRequest() header mismatch: want %s got %s", tt.want.header, got.Header)
				return
			}
			b, _ := ioutil.ReadAll(got.Body)
			if string(tt.want.body) != string(b) {
				t.Errorf("NewRequest() body mismatch: want %s got %s", string(tt.want.body), string(b))
			}
		})
	}
}

func TestClient_CreateAndDo(t *testing.T) {
	type settings struct {
		client     *http.Client
		baseURL    *url.URL
		apiVersion string
		authMethod hubspot.AuthMethod
		crm        *hubspot.CRM
	}
	type args struct {
		method   string
		relPath  string
		data     interface{}
		option   *hubspot.RequestQueryOption
		resource interface{}
	}
	tests := []struct {
		name     string
		settings settings
		args     args
		want     interface{}
		wantErr  error
	}{
		{
			name: "Success GET request",
			settings: settings{
				client: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"createdAt":"2019-10-30T03:30:17.883Z","archived":false,"id":"512","properties":{"amount":"1500.00","closedate":"2019-12-07T16:50:06.678Z","createdate":"2019-10-30T03:30:17.883Z","dealname":"Custom data integrations","dealstage":"presentation scheduled","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","hubspot_owner_id":"910901","pipeline":"default"},"updatedAt":"2019-12-07T16:50:06.678Z"}`),
				}),
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetPrivateAppToken("test"),
				crm:        nil,
			},
			args: args{
				method:   http.MethodGet,
				relPath:  "crm/v3/objects/deals",
				data:     nil,
				option:   nil,
				resource: &hubspot.ResponseResource{Properties: &hubspot.Deal{}},
			},
			want: &hubspot.ResponseResource{
				ID:       "512",
				Archived: false,
				Properties: &hubspot.Deal{
					Amount:           hubspot.NewString("1500.00"),
					DealName:         hubspot.NewString("Custom data integrations"),
					DealStage:        hubspot.NewString("presentation scheduled"),
					DealOwnerID:      hubspot.NewString("910901"),
					PipeLine:         hubspot.NewString("default"),
					CreateDate:       &createdAt,
					CloseDate:        &closeDate,
					LastModifiedDate: &modifyDate,
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Success POST request",
			settings: settings{
				client: hubspot.NewMockHTTPClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"createdAt":"2019-10-30T03:30:17.883Z","archived":false,"id":"512","properties":{"amount":"1500.00","closedate":"2019-12-07T16:50:06.678Z","createdate":"2019-10-30T03:30:17.883Z","dealname":"Custom data integrations","dealstage":"presentation scheduled","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","hubspot_owner_id":"910901","pipeline":"default"},"updatedAt":"2019-12-07T16:50:06.678Z"}`),
				}),
				baseURL:    hubspot.ExportBaseURL,
				apiVersion: hubspot.ExportAPIVersion,
				authMethod: hubspot.SetPrivateAppToken("test"),
				crm:        nil,
			},
			args: args{
				method:  http.MethodGet,
				relPath: "crm/v3/objects/deals",
				data: &hubspot.Deal{
					Amount:      hubspot.NewString("1500.00"),
					DealName:    hubspot.NewString("Custom data integrations"),
					DealStage:   hubspot.NewString("presentation scheduled"),
					DealOwnerID: hubspot.NewString("910901"),
					PipeLine:    hubspot.NewString("default"),
				},
				option:   nil,
				resource: &hubspot.ResponseResource{Properties: &hubspot.Deal{}},
			},
			want: &hubspot.ResponseResource{
				ID:       "512",
				Archived: false,
				Properties: &hubspot.Deal{
					Amount:           hubspot.NewString("1500.00"),
					DealName:         hubspot.NewString("Custom data integrations"),
					DealStage:        hubspot.NewString("presentation scheduled"),
					DealOwnerID:      hubspot.NewString("910901"),
					PipeLine:         hubspot.NewString("default"),
					CreateDate:       &createdAt,
					CloseDate:        &closeDate,
					LastModifiedDate: &modifyDate,
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &hubspot.Client{
				HTTPClient: tt.settings.client,
			}
			c.ExportSetAPIVersion(tt.settings.apiVersion)
			c.ExportSetBaseURL(tt.settings.baseURL)
			tt.settings.authMethod(c)

			err := c.CreateAndDo(tt.args.method, tt.args.relPath, "application/json", tt.args.data, tt.args.option, tt.args.resource)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("CreateAndDo() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, tt.args.resource, cmpTimeOption); diff != "" {
				t.Errorf("CreateAndDo() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestCheckResponseError(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Response StatusOK",
			args: args{
				r: &http.Response{
					StatusCode: http.StatusOK,
				},
			},
			wantErr: nil,
		},
		{
			name: "Response BadRequest",
			args: args{
				r: &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`))),
				},
			},
			wantErr: &hubspot.APIError{
				HTTPStatusCode: http.StatusBadRequest,
				Message:        "Invalid input (details will vary based on the error)",
				CorrelationID:  "aeb5f871-7f07-4993-9211-075dc63e7cbf",
				Category:       hubspot.ValidationError,
				Links: hubspot.ErrLinks{
					KnowledgeBase: "https://www.hubspot.com/products/service/knowledge-base",
				},
			},
		},
		{
			name: "Response BadRequest with error details",
			args: args{
				r: &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(`{"status": "error","message": "Property values were not valid: [{\"isValid\":false,\"message\":\"Email address bcooper@example.con is invalid\",\"error\":\"INVALID_EMAIL\",\"name\":\"email\"}]","correlationId":"aeb5f871-7f07-4993-9211-075dc63e7cbf","category":"VALIDATION_ERROR"}`))),
				},
			},
			wantErr: &hubspot.APIError{
				HTTPStatusCode: http.StatusBadRequest,
				Message:        `Property values were not valid: [{"isValid":false,"message":"Email address bcooper@example.con is invalid","error":"INVALID_EMAIL","name":"email"}]`,
				CorrelationID:  "aeb5f871-7f07-4993-9211-075dc63e7cbf",
				Category:       hubspot.ValidationError,
				Status:         "error",
				Details: []hubspot.ErrDetail{
					{
						IsValid: false,
						Message: "Email address bcooper@example.con is invalid",
						Error:   hubspot.InvalidEmailError,
						Name:    "email",
					},
				},
			},
		},
		{
			name: "Response BadRequest with unexpected error details",
			args: args{
				r: &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(`{"status": "error","message": "Property values were not valid: [{\"isValid\":false,\"message\":\"Email address bcooper@example.con is invalid\",\"error\":\"INVALID_EMAIL\",\"name\":\"email\"},{'json':unexpected}]","correlationId":"aeb5f871-7f07-4993-9211-075dc63e7cbf","category":"VALIDATION_ERROR"}`))),
				},
			},
			wantErr: &hubspot.APIError{
				HTTPStatusCode: http.StatusBadRequest,
				Message:        `Property values were not valid: [{"isValid":false,"message":"Email address bcooper@example.con is invalid","error":"INVALID_EMAIL","name":"email"},{'json':unexpected}]`,
				CorrelationID:  "aeb5f871-7f07-4993-9211-075dc63e7cbf",
				Category:       hubspot.ValidationError,
				Status:         "error",
				Details: []hubspot.ErrDetail{
					{
						IsValid: false,
						Message: "Email address bcooper@example.con is invalid",
						Error:   hubspot.InvalidEmailError,
						Name:    "email",
					},
					{
						IsValid: false,
						Message: `unable to read error detail {'json':unexpected}: invalid character '\'' looking for beginning of object key string`,
						Error:   hubspot.UnknownDetailError,
						Name:    "unknown",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := hubspot.CheckResponseError(tt.args.r)
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("CheckResponseError() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

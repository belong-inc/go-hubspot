package hubspot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultAPIVersion = "v3"
)

var (
	defaultBaseURL = &url.URL{Scheme: "https", Host: "api.hubapi.com"}

	// jsonPattern matches the JSON pattern.
	// Used to extracts the error details contained in the error message.
	jsonPattern = regexp.MustCompile(`{[\s\S]*?}`)
)

// Client manages communication with the HubSpot API.
type Client struct {
	HTTPClient *http.Client

	baseURL    *url.URL
	apiVersion string

	authenticator Authenticator

	CRM       *CRM
	Marketing *Marketing
}

// RequestPayload is common request structure for HubSpot APIs.
type RequestPayload struct {
	Properties interface{} `json:"properties,omitempty"`
}

// ResponseResource is common response structure for HubSpot APIs.
type ResponseResource struct {
	ID           string        `json:"id,omitempty"`
	Archived     bool          `json:"archived,omitempty"`
	Associations *Associations `json:"associations,omitempty"`
	Properties   interface{}   `json:"properties,omitempty"`
	CreatedAt    *HsTime       `json:"createdAt,omitempty"`
	UpdatedAt    *HsTime       `json:"updatedAt,omitempty"`
	ArchivedAt   *HsTime       `json:"archivedAt,omitempty"`
}

// NewClient returns a new HubSpot API client with APIKey or OAuthConfig.
// HubSpot officially recommends authentication with OAuth.
// e.g. hubspot.NewClient(hubspot.SetPrivateAppToken("key"))
func NewClient(setAuthMethod AuthMethod, opts ...Option) (*Client, error) {
	if setAuthMethod == nil {
		return nil, errors.New("the authentication method is not set")
	}

	c := &Client{
		HTTPClient: http.DefaultClient,
		baseURL:    defaultBaseURL,
		apiVersion: defaultAPIVersion,
	}

	// Set the authentication method specified by the argument.
	// Authentication method is either APIKey or OAuth.
	setAuthMethod(c)

	for _, o := range opts {
		o(c)
	}

	// Since the baseURL and apiVersion may change, initialize the service after applying the options.
	c.CRM = newCRM(c)
	c.Marketing = newMarketing(c)

	return c, nil
}

// NewRequest creates an API request.
// After creating a request, add the authentication information according to the method specified in NewClient().
func (c *Client) NewRequest(method, path string, body, option interface{}, contentType string) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	var requestBody []byte = nil

	if body != nil {
		if !strings.HasPrefix(contentType, "multipart") {
			requestBody, err = json.Marshal(body)
			if err != nil {
				return nil, err
			}
		} else {
			var ok bool
			requestBody, ok = body.([]byte)
			if !ok {
				return nil, fmt.Errorf("error in typecasting mime body to []byte")
			}
		}
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	// Parse query options
	if option != nil {
		q, err := query.Values(option)
		if err != nil {
			return nil, err
		}
		for k, values := range u.Query() {
			for _, v := range values {
				q.Add(k, v)
			}
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	// Configure authentication settings using the method specified during NewClient().
	if err := c.authenticator.SetAuthentication(req); err != nil {
		return nil, err
	}

	return req, nil
}

// CreateAndDo performs a web request to HubSpot.
// The `data`, `options` and `resource` arguments are optional and only relevant in certain situations.
// If the data argument is non-nil, it will be used as the body of the request for POST and PUT requests.
// The options argument is used for specifying request options such as search parameters.
// The resource argument is marshalled data returned from HubSpot.
// If the resource contains a pointer to data, the data will be overwritten with the content of the response.
func (c *Client) CreateAndDo(method, relPath, contentType string, data, option, resource interface{}) error {
	if strings.HasPrefix(relPath, "/") {
		relPath = strings.TrimLeft(relPath, "/")
	}

	req, err := c.NewRequest(method, relPath, data, option, contentType)
	if err != nil {
		return err
	}

	_, err = c.doGetHeaders(req, resource)
	if err != nil {
		return err
	}

	return nil
}

// doGetHeaders executes a request, decoding the response into `v` and also returns any response headers.
// FIXME: Add optional retry process
func (c *Client) doGetHeaders(req *http.Request, v interface{}) (http.Header, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resErr := CheckResponseError(resp); resErr != nil {
		return nil, resErr
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, err
		}
	}
	return resp.Header, nil
}

// CheckResponseError checks the response, and in case of error, maps it to the error structure.
func CheckResponseError(r *http.Response) error {
	if !isErrorStatusCode(r.StatusCode) {
		return nil
	}

	hubspotErr := &APIError{
		HTTPStatusCode: r.StatusCode,
	}

	if r.Body != nil {
		if err := json.NewDecoder(r.Body).Decode(hubspotErr); err != nil {
			return &APIError{
				HTTPStatusCode: r.StatusCode,
				Message:        fmt.Sprintf("unable to read response from hubspot: %s", err),
			}
		}
		// HubSpot contain error details in the error message, so we need to extract them with a regexp.
		if details := jsonPattern.FindAllString(hubspotErr.Message, -1); len(details) != 0 {
			for _, detail := range details {
				var errDetail ErrDetail
				if err := json.Unmarshal([]byte(detail), &errDetail); err != nil {
					errDetail = ErrDetail{
						IsValid: false,
						Message: fmt.Sprintf("unable to read error detail %s: %s", detail, err),
						Error:   UnknownDetailError,
						Name:    "unknown",
					}
				}
				hubspotErr.Details = append(hubspotErr.Details, errDetail)
			}
		}
	}

	return hubspotErr
}

func isErrorStatusCode(code int) bool {
	// If status code is more than 400, return true
	return http.StatusBadRequest <= code
}

// Get performs a GET request for the given path and saves the result in the given resource.
func (c *Client) Get(path string, resource interface{}, option interface{}) error {
	return c.CreateAndDo(http.MethodGet, path, "application/json", nil, option, resource)
}

// Post performs a POST request for the given path and saves the result in the given resource.
func (c *Client) Post(path string, data, resource interface{}) error {
	return c.CreateAndDo(http.MethodPost, path, "application/json", data, nil, resource)
}

// Put performs a PUT request for the given path and saves the result in the given resource.
func (c *Client) Put(path string, data, resource interface{}) error {
	return c.CreateAndDo(http.MethodPut, path, "application/json", data, nil, resource)
}

// Patch performs a PATCH request for the given path and saves the result in the given resource.
func (c *Client) Patch(path string, data, resource interface{}) error {
	return c.CreateAndDo(http.MethodPatch, path, "application/json", data, nil, resource)
}

// Delete performs a DELETE request for the given path.
func (c *Client) Delete(path string) error {
	return c.CreateAndDo(http.MethodDelete, path, "application/json", nil, nil, nil)
}

func (c *Client) PostMultipart(path, boundary string, data, resource interface{}) error {
	mimeType := fmt.Sprintf("multipart/form-data; boundary=%s", boundary)
	return c.CreateAndDo(http.MethodPost, path, mimeType, data, nil, resource)
}

package hubspot

import "net/url"

// For test only

var (
	ExportAPIVersion = defaultAPIVersion

	ExportBaseURL         = defaultBaseURL
	ExportContactBasePath = contactBasePath
	ExportDealBasePath    = dealBasePath
)

var (
	ExportNewCRM       = newCRM
	ExportNewMarketing = newMarketing

	ExportSetupProperties = (*RequestQueryOption).setupProperties

	ExportFetchTokenFromHubSpot = (*OAuthTokenManager).fetchTokenFromHubSpot
	ExportRefreshToken          = (*OAuthTokenManager).refreshToken

	ExportConvertToFormData = (*OAuthConfig).convertToFormData
	ExportConfigValid       = (*OAuthConfig).valid

	ExportSetExpiry  = (*OAuthToken).setExpiry
	ExportTokenValid = (*OAuthToken).valid
	ExportExpired    = (*OAuthToken).expired
)

func (c *Client) ExportGetAPIVersion() string {
	return c.apiVersion
}

func (c *Client) ExportGetBaseURL() *url.URL {
	return c.baseURL
}

func (c *Client) ExportSetAPIVersion(version string) {
	c.apiVersion = version
}

func (c *Client) ExportSetBaseURL(url *url.URL) {
	c.baseURL = url
}

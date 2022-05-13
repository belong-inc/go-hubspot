package hubspot

import (
	"fmt"

	"github.com/belong-inc/go-hubspot/legacy"
)

const (
	// NOTE: THe API of marketing email is not migrated to v3 yet.
	marketingEmailBasePath = ""
)

type (
	// Statistics is response from marketing email statistics API v1 as of now.
	Statistics = legacy.StatisticsResponse
	// BulkStatisticsResponse is response from marketing email statistics API v1 as of now.
	// This contains list of Statistics.
	BulkStatisticsResponse = legacy.BulkResponseResource
)

// MarketingEmailService is an interface of marketing email endpoints of the HubSpot API.
// As of May 2022, HubSpot provides only API v1 therefore the implementation is based on document in
// https://legacydocs.hubspot.com/docs/methods/cms_email/get-the-statistics-for-a-marketing-email.
type MarketingEmailService interface {
	GetStatistics(emailID int, statistics interface{}, option *RequestQueryOption) (*ResponseResource, error)
	ListStatistics(statistics interface{}, option *RequestQueryOption) (*ResponseResource, error)
}

type MarketingEmailOp struct {
	marketingEmailBasePath string
	client                 *Client
	legacyAPIHelper        legacy.MarketingEmailHelper
}

var _ MarketingEmailService = (*MarketingEmailOp)(nil)

// NewMarketingEmail creates a new MarketingEmailService.
func NewMarketingEmail(client *Client) MarketingEmailService {
	return &MarketingEmailOp{
		client:          client,
		legacyAPIHelper: legacy.NewMarketingEmailHelper(),
	}
}

// GetStatistics get a Statistics for given emailID.
func (m *MarketingEmailOp) GetStatistics(emailID int, resource interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	if err := m.client.Get(m.legacyAPIHelper.GetStatisticsPath()+fmt.Sprintf("/%d", emailID), resource, option.setupProperties(defaultContactFields)); err != nil {
		return nil, err
	}
	return &ResponseResource{Properties: resource}, nil
}

// ListStatistics get a list of Statistics.
func (m MarketingEmailOp) ListStatistics(resource interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	if err := m.client.Get(m.legacyAPIHelper.GetStatisticsPath(), resource, option); err != nil {
		return nil, err
	}
	return &ResponseResource{Properties: resource}, nil
}

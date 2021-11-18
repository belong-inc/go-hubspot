package hubspot

const (
	dealBasePath = "deals"
)

// DealService is an interface of deal endpoints of the HubSpot API.
// HubSpot deal can be used to manage transactions.
// It can also be associated with other CRM objects such as contact and company.
// Reference: https://developers.hubspot.com/docs/api/crm/deals
type DealService interface {
	Get(dealID string, deal interface{}, option *RequestQueryOption) (*ResponseResource, error)
	Create(deal interface{}) (*ResponseResource, error)
	Update(dealID string, deal interface{}) (*ResponseResource, error)
	AssociateAnotherObj(dealID string, conf *AssociationConfig) (*ResponseResource, error)
}

// DealServiceOp handles communication with the product related methods of the HubSpot API.
type DealServiceOp struct {
	dealPath string
	client   *Client
}

var _ DealService = (*DealServiceOp)(nil)

// Deal represents a HubSpot deal.
type Deal struct {
	Amount                  *HsStr `json:"amount,omitempty"`
	AmountInCompanyCurrency *HsStr `json:"amount_in_home_currency,omitempty"`
	AnnualContractValue     *HsStr `json:"hs_acv,omitempty"`
	AnnualRecurringRevenue  *HsStr `json:"hs_arr,omitempty"`
	ClosedLostReason        *HsStr `json:"closed_lost_reason,omitempty"`
	ClosedWonReason         *HsStr `json:"closed_won_reason,omitempty"`
	DealDescription         *HsStr `json:"description,omitempty"`
	DealName                *HsStr `json:"dealname,omitempty"`
	DealOwnerID             *HsStr `json:"hubspot_owner_id,omitempty"`
	DealStage               *HsStr `json:"dealstage,omitempty"`
	DealType                *HsStr `json:"dealtype,omitempty"`
	ForecastAmount          *HsStr `json:"hs_forecast_amount,omitempty"`
	ForecastCategory        *HsStr `json:"hs_forecast_category,omitempty"`
	ForecastProbability     *HsStr `json:"hs_forecast_probability,omitempty"`
	MonthlyRecurringRevenue *HsStr `json:"hs_mrr,omitempty"`
	NextStep                *HsStr `json:"hs_next_step,omitempty"`
	NumberOfContacts        *HsStr `json:"num_associated_contacts,omitempty"`
	NumberOfSalesActivities *HsStr `json:"num_notes,omitempty"`
	NumberOfTimesContacted  *HsStr `json:"num_contacted_notes,omitempty"`
	ObjectID                *HsStr `json:"hs_object_id,omitempty"`
	PipeLine                *HsStr `json:"pipeline,omitempty"`
	TeamID                  *HsStr `json:"hubspot_team_id,omitempty"`
	TotalContractValue      *HsStr `json:"hs_tcv,omitempty"`

	CreateDate        *HsTime `json:"createdate,omitempty"`
	CloseDate         *HsTime `json:"closedate,omitempty"`
	LastActivityDate  *HsTime `json:"notes_last_updated,omitempty"`
	LastContacted     *HsTime `json:"notes_last_contacted,omitempty"`
	LastModifiedDate  *HsTime `json:"hs_lastmodifieddate,omitempty"`
	NextActivityDate  *HsTime `json:"notes_next_activity_date,omitempty"`
	OwnerAssignedDate *HsTime `json:"hubspot_owner_assigneddate,omitempty"`
}

var defaultDealFields = []string{
	"amount",
	"amount_in_home_currency",
	"hs_acv",
	"hs_arr",
	"closed_lost_reason",
	"closed_won_reason",
	"description",
	"dealname",
	"hubspot_owner_id",
	"dealstage",
	"dealtype",
	"hs_forecast_amount",
	"hs_forecast_category",
	"hs_forecast_probability",
	"hs_mrr",
	"hs_next_step",
	"num_associated_contacts",
	"num_notes",
	"num_contacted_notes",
	"hs_object_id",
	"hubspot_owner_assigneddate",
	"pipeline",
	"hubspot_team_id",
	"hs_tcv",
	"createdate",
	"closedate",
	"notes_last_updated",
	"notes_last_contacted",
	"hs_lastmodifieddate",
	"notes_next_activity_date",
}

// Get gets a deal.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *DealServiceOp) Get(dealID string, deal interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: deal}
	if err := s.client.Get(s.dealPath+"/"+dealID, resource, option.setupProperties(defaultDealFields)); err != nil {
		return nil, err
	}
	return resource, nil
}

// Create creates a new deal.
// In order to bind the created content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Deal in your own structure.
func (s *DealServiceOp) Create(deal interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: deal}
	resource := &ResponseResource{Properties: deal}
	if err := s.client.Post(s.dealPath, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Update updates a deal.
// In order to bind the updated content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Deal in your own structure.
func (s *DealServiceOp) Update(dealID string, deal interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: deal}
	resource := &ResponseResource{Properties: deal}
	if err := s.client.Patch(s.dealPath+"/"+dealID, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// AssociateAnotherObj associates Deal with another HubSpot objects.
// If you want to associate a custom object, please use a defined value in HubSpot.
func (s *DealServiceOp) AssociateAnotherObj(dealID string, conf *AssociationConfig) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: &Deal{}}
	if err := s.client.Put(s.dealPath+"/"+dealID+"/"+conf.makeAssociationPath(), nil, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

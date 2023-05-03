package hubspot

import "fmt"

const (
	crmTicketsBasePath = "tickets"
)

// CrmTicketsServivce is an interface of CRM tickets endpoints of the HubSpot API.
// Reference: https://developers.hubspot.com/docs/api/crm/tickets
type CrmTicketsServivce interface {
	List(option *RequestQueryOption) (*CrmTicketsList, error)
	Get(ticketId string, option *RequestQueryOption) (*CrmTicket, error)
	Create(reqData *CrmTicketCreateRequest) (*CrmTicket, error)
	Archive(ticketId string) error
	Update(ticketId string, reqData *CrmTicketUpdateRequest) (*CrmTicket, error)
	Search(reqData *CrmTicketSearchRequest) (*CrmTicketsList, error)
}

// CrmTicketsServivceOp handles communication with the CRM tickets endpoints of the HubSpot API.
type CrmTicketsServivceOp struct {
	client         *Client
	crmTicketsPath string
}

var _ CrmTicketsServivce = (*CrmTicketsServivceOp)(nil)

type CrmTicket struct {
	Id                    *HsStr                 `json:"id,omitempty"`
	Properties            map[string]interface{} `json:"properties,omitempty"`
	PropertiesWithHistory map[string]interface{} `json:"propertiesWithHistory,omitempty"`
	CreatedAt             *HsTime                `json:"createdAt,omitempty"`
	UpdatedAt             *HsTime                `json:"updatedAt,omitempty"`
	Archived              *HsBool                `json:"archived,omitempty"`
	ArchivedAt            *HsTime                `json:"archivedAt,omitempty"`
}

type CrmTicketsPagingData struct {
	After *HsStr `json:"after,omitempty"`
	Link  *HsStr `json:"link,omitempty"`
}

type CrmTicketsPaging struct {
	Next *CrmTicketsPagingData `json:"next,omitempty"`
}

type CrmTicketsList struct {
	Total   *HsInt            `json:"total,omitempty"`
	Results []*CrmTicket      `json:"results"`
	Paging  *CrmTicketsPaging `json:"paging,omitempty"`
}

func (s *CrmTicketsServivceOp) List(option *RequestQueryOption) (*CrmTicketsList, error) {
	var resource CrmTicketsList
	if err := s.client.Get(s.crmTicketsPath, &resource, option); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmTicketsServivceOp) Get(ticketId string, option *RequestQueryOption) (*CrmTicket, error) {
	var resource CrmTicket
	path := fmt.Sprintf("%s/%s", s.crmTicketsPath, ticketId)
	if err := s.client.Get(path, &resource, option); err != nil {
		return nil, err
	}
	return &resource, nil
}

type CrmTicketAssociationTarget struct {
	Id *HsStr `json:"id,omitempty"`
}

type CrmTicketAssociationType struct {
	AssociationCategory *HsStr `json:"associationCategory,omitempty"`
	AssociationTypeId   *HsInt `json:"associationTypeId,omitempty"`
}

type CrmTicketAssociation struct {
	To    CrmTicketAssociationTarget `json:"to,omitempty"`
	Types []CrmTicketAssociationType `json:"type,omitempty"`
}

type CrmTicketCreateRequest struct {
	Properties   map[string]interface{}  `json:"properties,omitempty"`
	Associations []*CrmTicketAssociation `json:"associations,omitempty"`
}

type CrmTicketUpdateRequest = CrmTicketCreateRequest

func (s *CrmTicketsServivceOp) Create(reqData *CrmTicketCreateRequest) (*CrmTicket, error) {
	var resource CrmTicket
	if err := s.client.Post(s.crmTicketsPath, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmTicketsServivceOp) Archive(ticketId string) error {
	path := fmt.Sprintf("%s/%s", s.crmTicketsPath, ticketId)
	return s.client.Delete(path, nil)
}

func (s *CrmTicketsServivceOp) Update(ticketId string, reqData *CrmTicketUpdateRequest) (*CrmTicket, error) {
	var resource CrmTicket
	path := fmt.Sprintf("%s/%s", s.crmTicketsPath, ticketId)
	if err := s.client.Patch(path, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

type CrmTicketSearchFilter struct {
	Value        *HsStr   `json:"value,omitempty"`
	HighValue    *HsStr   `json:"highValue,omitempty"`
	Values       []*HsStr `json:"values,omitempty"`
	PropertyName *HsStr   `json:"propertyName,omitempty"`
	Operator     *HsStr   `json:"operator,omitempty"`
}

type CrmTicketSearchFilterGroup struct {
	Filters    []*CrmTicketSearchFilter `json:"filters,omitempty"`
	Sorts      []*HsStr                 `json:"sorts,omitempty"`
	Query      *HsStr                   `json:"query"`
	Properties []*HsStr                 `json:"properties,omitempty"`
	Limit      *HsInt                   `json:"limit,omitempty"`
	After      *HsInt                   `json:"after,omitempty"`
}

type CrmTicketSearchRequest struct {
	FilterGroups []*CrmTicketSearchFilterGroup `json:"filterGroups,omitempty"`
}

func (s *CrmTicketsServivceOp) Search(reqData *CrmTicketSearchRequest) (*CrmTicketsList, error) {
	var resource CrmTicketsList
	path := fmt.Sprintf("%s/search", s.crmTicketsPath)
	if err := s.client.Post(path, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

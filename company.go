package hubspot

const (
	companyBasePath = "companies"
)

// CompanyService is an interface of company endpoints of the HubSpot API.
// HubSpot companies store information about organizations.
// It can also be associated with other CRM objects such as deal and contact.
// Reference: https://developers.hubspot.com/docs/api/crm/companies
type CompanyService interface {
	Get(companyID string, company interface{}, option *RequestQueryOption) (*ResponseResource, error)
	Create(company interface{}) (*ResponseResource, error)
	Update(companyID string, company interface{}) (*ResponseResource, error)
	Delete(companyID string) error
	AssociateAnotherObj(companyID string, conf *AssociationConfig) (*ResponseResource, error)
	SearchByDomain(domain string) (*CompanySearchResponse, error)
	SearchByName(name string) (*CompanySearchResponse, error)
}

// CompanyServiceOp handles communication with the product related methods of the HubSpot API.
type CompanyServiceOp struct {
	companyPath string
	client      *Client
}

var _ CompanyService = (*CompanyServiceOp)(nil)

// Get gets a Company.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *CompanyServiceOp) Get(companyID string, company interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: company}
	if err := s.client.Get(s.companyPath+"/"+companyID, resource, option.setupProperties(defaultCompanyFields)); err != nil {
		return nil, err
	}
	return resource, nil
}

// Create creates a new company.
// In order to bind the created content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Company in your own structure.
func (s *CompanyServiceOp) Create(company interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: company}
	resource := &ResponseResource{Properties: company}
	if err := s.client.Post(s.companyPath, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Update updates a company.
// In order to bind the updated content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Company in your own structure.
func (s *CompanyServiceOp) Update(companyID string, company interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: company}
	resource := &ResponseResource{Properties: company}
	if err := s.client.Patch(s.companyPath+"/"+companyID, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Delete deletes a company.
func (s *CompanyServiceOp) Delete(companyID string) error {
	return s.client.Delete(s.companyPath+"/"+companyID, nil)
}

// AssociateAnotherObj associates Company with another HubSpot objects.
// If you want to associate a custom object, please use a defined value in HubSpot.
func (s *CompanyServiceOp) AssociateAnotherObj(companyID string, conf *AssociationConfig) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: &Company{}}
	if err := s.client.Put(s.companyPath+"/"+companyID+"/"+conf.makeAssociationPath(), nil, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// CompanySearchResponse represents the response from searching companies.
type CompanySearchResponse struct {
	Total   int64           `json:"total"`
	Results []CompanyResult `json:"results"`
}

type CompanyResult struct {
	ID         string  `json:"id"`
	Properties Company `json:"properties"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
	Archived   bool    `json:"archived"`
}

// SearchByDomain searches for a company by domain.
func (s *CompanyServiceOp) SearchByDomain(domain string) (*CompanySearchResponse, error) {
	req := &ContactSearchRequest{
		FilterGroups: []FilterGroup{
			{
				Filters: []Filter{
					{
						PropertyName: "domain",
						Operator:     "EQ",
						Value:        domain,
					},
				},
			},
		},
	}

	resource := &CompanySearchResponse{}
	if err := s.client.Post(s.companyPath+"/search", req, resource); err != nil {
		return nil, err
	}

	return resource, nil
}

// SearchByName searches for a company by name.
func (s *CompanyServiceOp) SearchByName(name string) (*CompanySearchResponse, error) {
	req := &ContactSearchRequest{
		FilterGroups: []FilterGroup{
			{
				Filters: []Filter{
					{
						PropertyName: "name",
						Operator:     "EQ",
						Value:        name,
					},
				},
			},
		},
	}

	resource := &CompanySearchResponse{}
	if err := s.client.Post(s.companyPath+"/search", req, resource); err != nil {
		return nil, err
	}

	return resource, nil
}

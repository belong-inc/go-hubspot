package hubspot

import (
	"fmt"
)

const (
	crmSchemasPath = "schemas"
)

type CrmSchemaAssociation struct {
	FromObjectTypeId   *HsStr  `json:"fromObjectTypeId"`
	ToObjectTypeId     *HsStr  `json:"toObjectTypeId"`
	Name               *HsStr  `json:"name"`
	ID                 *HsStr  `json:"id"`
	CreatedAt          *HsTime `json:"createdAt"`
	UpdatedAt          *HsTime `json:"updatedAt"`
	Cardinality        *HsStr  `json:"cardinality"`
	InverseCardinality *HsStr  `json:"inverseCardinality"`
}

type CrmSchemaLabels struct {
	Singular *HsStr `json:"singular"`
	Plural   *HsStr `json:"plural"`
}

type CrmSchemasList struct {
	Results []*CrmSchema
}

type CrmSchema struct {
	Labels                     *CrmSchemaLabels        `json:"labels,omitempty"`
	PrimaryDisplayProperty     *HsStr                  `json:"primaryDisplayProperty,omitempty"`
	Archived                   *HsBool                 `json:"archived,omitempty"`
	ID                         *HsStr                  `json:"id,omitempty"`
	FullyQualifiedName         *HsStr                  `json:"fullyQualifiedName,omitempty"`
	CreatedAt                  *HsTime                 `json:"createdAt,omitempty"`
	UpdatedAt                  *HsTime                 `json:"updatedAt,omitempty"`
	ObjectTypeId               *HsStr                  `json:"objectTypeId,omitempty"`
	Properties                 []*CrmProperty          `json:"properties,omitempty"`
	Associations               []*CrmSchemaAssociation `json:"associations,omitempty"`
	Name                       *HsStr                  `json:"name,omitempty"`
	MetaType                   *HsStr                  `json:"metaType,omitempty"`
	RequiredProperties         []*HsStr                `json:"requiredProperties,omitempty"`
	Restorable                 *HsBool                 `json:"restorable,omitempty"`
	SearchableProperties       []*HsStr                `json:"searchableProperties,omitempty"`
	SecondaryDisplayProperties []*HsStr                `json:"secondaryDisplayProperties,omitempty"`
	PortalId                   *HsInt                  `json:"portalId"`
}

// CrmSchemasService is an interface of CRM schemas endpoints of the HubSpot API.
// Reference: https://developers.hubspot.com/docs/api/crm/crm-custom-objects
type CrmSchemasService interface {
	List() (*CrmSchemasList, error)
	Create(reqData interface{}) (*CrmSchema, error)
	Get(objectType string) (*CrmSchema, error)
	Delete(objectType string, option *RequestQueryOption) error
	Update(objectType string, reqData interface{}) (*CrmSchema, error)
}

// CrmSchemasServiceOp handles communication with the CRM schemas endpoint.
type CrmSchemasServiceOp struct {
	client         *Client
	crmSchemasPath string
}

var _ CrmSchemasService = (*CrmSchemasServiceOp)(nil)

func (s *CrmSchemasServiceOp) List() (*CrmSchemasList, error) {
	var resource CrmSchemasList
	if err := s.client.Get(s.crmSchemasPath, &resource, nil); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmSchemasServiceOp) Create(reqData interface{}) (*CrmSchema, error) {
	var resource CrmSchema
	if err := s.client.Post(s.crmSchemasPath, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmSchemasServiceOp) Get(objectType string) (*CrmSchema, error) {
	var resource CrmSchema
	path := fmt.Sprintf("%s/%s", s.crmSchemasPath, objectType)
	if err := s.client.Get(path, &resource, nil); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmSchemasServiceOp) Delete(objectType string, option *RequestQueryOption) error {
	path := fmt.Sprintf("%s/%s", s.crmSchemasPath, objectType)
	return s.client.Delete(path, option)
}

func (s *CrmSchemasServiceOp) Update(objectType string, reqData interface{}) (*CrmSchema, error) {
	var resource CrmSchema
	path := fmt.Sprintf("%s/%s", s.crmSchemasPath, objectType)
	if err := s.client.Patch(path, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

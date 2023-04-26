package hubspot

import (
	"fmt"
)

const (
	crmPropertiesPath = "properties"
)

type CrmPropertiesList struct {
	Results []*CrmProperty `json:"results,omitempty"`
}

type CrmProperty struct {
	UpdatedAt            *HsTime                      `json:"updatedAt,omitempty"`
	CreatedAt            *HsTime                      `json:"createdAt,omitempty"`
	ArchivedAt           *HsTime                      `json:"archivedAt,omitempty"`
	Name                 *HsStr                       `json:"name,omitempty"`
	Label                *HsStr                       `json:"label,omitempty"`
	Type                 *HsStr                       `json:"type,omitempty"`
	FieldType            *HsStr                       `json:"fieldType,omitempty"`
	Description          *HsStr                       `json:"description,omitempty"`
	GroupName            *HsStr                       `json:"groupName,omitempty"`
	Options              []*CrmPropertyOptions        `json:"options,omitempty"`
	CreatedUserId        *HsStr                       `json:"createdUserId,omitempty"`
	UpdatedUserId        *HsStr                       `json:"updatedUserId,omitempty"`
	ReferencedObjectType *HsStr                       `json:"referencedObjectType,omitempty"`
	DisplayOrder         *HsInt                       `json:"displayOrder,omitempty"`
	Calculated           *HsBool                      `json:"calculated,omitempty"`
	ExternalOptions      *HsBool                      `json:"externalOptions,omitempty"`
	Archived             *HsBool                      `json:"archived,omitempty"`
	HasUniqueValue       *HsBool                      `json:"hasUniqueValue,omitempty"`
	Hidden               *HsBool                      `json:"hidden,omitempty"`
	HubspotDefined       *HsBool                      `json:"hubspotDefined,omitempty"`
	ShowCurrencySymbol   *HsBool                      `json:"showCurrencySymbol,omitempty"`
	ModificationMetaData *CrmPropertyModificationMeta `json:"modificationMetadata,omitempty"`
	FormField            *HsBool                      `json:"formField,omitempty"`
	CalculationFormula   *HsStr                       `json:"calculationFormula,omitempty"`
}

type CrmPropertyModificationMeta struct {
	Archivable       *HsBool `json:"archivable,omitempty"`
	ReadOnlyDefition *HsBool `json:"readOnlyDefinition,omitempty"`
	ReadOnlyValue    *HsBool `json:"readOnlyValue,omitempty"`
	ReadOnlyOptions  *HsBool `json:"readOnlyOptions,omitempty"`
}

type CrmPropertyOptions struct {
	Label        *HsStr  `json:"label,omitempty"`
	Value        *HsStr  `json:"value,omitempty"`
	Description  *HsStr  `json:"description,omitempty"`
	DisplayOrder *HsInt  `json:"displayOrder,omitempty"`
	Hidden       *HsBool `json:"hidden,omitempty"`
}

// CrmPropertiesService is an interface of CRM properties endpoints of the HubSpot API.
// Reference: https://developers.hubspot.com/docs/api/crm/properties
type CrmPropertiesService interface {
	List(objectType string) (*CrmPropertiesList, error)
	Create(objectType string, reqData interface{}) (*CrmProperty, error)
	Get(objectType string, propertyName string) (*CrmProperty, error)
	Delete(objectType string, propertyName string) error
	Update(objectType string, propertyName string, reqData interface{}) (*CrmProperty, error)
}

// CrmPropertiesServiceOp handles communication with the CRM properties endpoint.
type CrmPropertiesServiceOp struct {
	client            *Client
	crmPropertiesPath string
}

var _ CrmPropertiesService = (*CrmPropertiesServiceOp)(nil)

func (s *CrmPropertiesServiceOp) List(objectType string) (*CrmPropertiesList, error) {
	var resource CrmPropertiesList
	path := fmt.Sprintf("%s/%s", s.crmPropertiesPath, objectType)
	if err := s.client.Get(path, &resource, nil); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmPropertiesServiceOp) Get(objectType, propertyName string) (*CrmProperty, error) {
	var resource CrmProperty
	path := fmt.Sprintf("%s/%s/%s", s.crmPropertiesPath, objectType, propertyName)
	if err := s.client.Get(path, &resource, nil); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmPropertiesServiceOp) Create(objectType string, reqData interface{}) (*CrmProperty, error) {
	var resource CrmProperty
	path := fmt.Sprintf("%s/%s", s.crmPropertiesPath, objectType)
	if err := s.client.Post(path, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmPropertiesServiceOp) Delete(objectType string, propertyName string) error {
	path := fmt.Sprintf("%s/%s/%s", s.crmPropertiesPath, objectType, propertyName)
	return s.client.Delete(path, nil)
}

func (s *CrmPropertiesServiceOp) Update(objectType string, propertyName string, reqData interface{}) (*CrmProperty, error) {
	var resource CrmProperty
	path := fmt.Sprintf("%s/%s/%s", s.crmPropertiesPath, objectType, propertyName)
	if err := s.client.Patch(path, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

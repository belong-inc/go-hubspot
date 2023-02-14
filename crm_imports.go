package hubspot

import (
	"fmt"
)

const (
	crmImportsBasePath = "imports"
)

// CrmImportsService is an interface of CRM bulk import endpoints of the HubSpot API.
// Reference: https://developers.hubspot.com/docs/api/crm/imports
type CrmImportsService interface {
	Active(option *CrmActiveImportOptions) (interface{}, error)
	Get(int64) (interface{}, error)
	Cancel(int64) (interface{}, error)
	Errors(int64, *CrmImportErrorsOptions) (interface{}, error)
	Start(*CrmImportConfig) (interface{}, error)
}

// CrmImportsServiceOp handles communication with the bulk CRM import endpoints of the HubSpot API.
type CrmImportsServiceOp struct {
	client         *Client
	crmImportsPath string
}

var _ CrmImportsService = (*CrmImportsServiceOp)(nil)

type CrmImportErrorsOptions struct {
	After string `url:"after,omitempty"`
	Limit int    `url:"limit,omitempty"`
}

func (s *CrmImportsServiceOp) Errors(importId int64, option *CrmImportErrorsOptions) (interface{}, error) {
	resource := make(map[string]interface{})
	path := fmt.Sprintf("%s/%d/errors", s.crmImportsPath, importId)
	if err := s.client.Get(path, &resource, option); err != nil {
		return nil, err
	}
	return resource, nil
}

type CrmActiveImportOptions struct {
	Before string `url:"before,omitempty"`
	After  string `url:"after,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

func (s *CrmImportsServiceOp) Active(option *CrmActiveImportOptions) (interface{}, error) {
	resource := make(map[string]interface{})
	if err := s.client.Get(s.crmImportsPath, &resource, option); err != nil {
		return nil, err
	}
	return resource, nil
}

func (s *CrmImportsServiceOp) Get(importId int64) (interface{}, error) {
	resource := make(map[string]interface{})
	path := fmt.Sprintf("%s/%d", s.crmImportsPath, importId)
	if err := s.client.Get(path, &resource, nil); err != nil {
		return nil, err
	}
	return resource, nil
}

func (s *CrmImportsServiceOp) Cancel(importId int64) (interface{}, error) {
	resource := make(map[string]interface{})
	path := fmt.Sprintf("%s/%d/cancel", s.crmImportsPath, importId)
	if err := s.client.Post(path, &resource, nil); err != nil {
		return nil, err
	}
	return resource, nil
}

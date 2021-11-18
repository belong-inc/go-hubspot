package hubspot

import (
	"fmt"
)

// Associate associates HubSpot objects like Deal and Contact.
// Object type and IDs are required for making the association.
// Using a custom object is also possible.
// Reference: https://developers.hubspot.com/docs/api/crm/associations

const (
	associationBasePath = "associations"
)

// ObjectType is the name used in object association.
type ObjectType string

// Default Object types
const (
	ObjectTypeContact ObjectType = "contacts"
	ObjectTypeDeal    ObjectType = "deals"
)

// AssociationType is the name of the key used to associate the objects together.
type AssociationType string

// Default association types
// Reference: https://legacydocs.hubspot.com/docs/methods/crm-associations/crm-associations-overview
const (
	AssociationTypeContactToCompany    AssociationType = "contact_to_company"
	AssociationTypeContactToDeal       AssociationType = "contact_to_deal"
	AssociationTypeContactToEngagement AssociationType = "contact_to_engagement"
	AssociationTypeContactToTicket     AssociationType = "contact_to_ticket"

	AssociationTypeDealToContact    AssociationType = "deal_to_contact"
	AssociationTypeDealToCompany    AssociationType = "deal_to_company"
	AssociationTypeDealToEngagement AssociationType = "deal_to_engagement"
	AssociationTypeDealToLineItem   AssociationType = "deal_to_line_item"
	AssociationTypeDealToTicket     AssociationType = "deal_to_ticket"
)

type AssociationConfig struct {
	ToObject   ObjectType
	ToObjectID string
	Type       AssociationType
}

func (c *AssociationConfig) makeAssociationPath() string {
	return fmt.Sprintf("%s/%s/%s/%s", associationBasePath, c.ToObject, c.ToObjectID, c.Type)
}

type Associations struct {
	Contacts struct {
		Results []AssociationResult `json:"results"`
	} `json:"contacts"`
	Deals struct {
		Results []AssociationResult `json:"results"`
	} `json:"deals"`
}

type AssociationResult struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

package hubspot

const (
	contactBasePath = "contacts"
)

// ContactService is an interface of contact endpoints of the HubSpot API.
// HubSpot contacts store information about individuals.
// It can also be associated with other CRM objects such as deal and company.
// Reference: https://developers.hubspot.com/docs/api/crm/contacts
type ContactService interface {
	Get(contactID string, contact interface{}, option *RequestQueryOption) (*ResponseResource, error)
	Create(contact interface{}) (*ResponseResource, error)
	Update(contactID string, contact interface{}) (*ResponseResource, error)
	Delete(contactID string) error
	AssociateAnotherObj(contactID string, conf *AssociationConfig) (*ResponseResource, error)
	SearchByEmail(email string) (*ContactSearchResponse, error)
}

// ContactServiceOp handles communication with the product related methods of the HubSpot API.
type ContactServiceOp struct {
	contactPath string
	client      *Client
}

var _ ContactService = (*ContactServiceOp)(nil)

// ContactSearchRequest represents the request body for searching contacts.
type ContactSearchRequest struct {
	FilterGroups []FilterGroup `json:"filterGroups"`
}

// FilterGroup represents a group of filters.
type FilterGroup struct {
	Filters []Filter `json:"filters"`
}

// Filter represents a single filter.
type Filter struct {
	PropertyName string   `json:"propertyName"`
	Operator     string   `json:"operator"`
	Values       []string `json:"values,omitempty"`
	Value        string   `json:"value,omitempty"`
}

// ContactSearchResponse represents the response from searching contacts.
type ContactSearchResponse struct {
	Results []Contact `json:"results"`
}

type Contact struct {
	Address                                     *HsStr  `json:"address,omitempty"`
	AnnualRevenue                               *HsStr  `json:"annualrevenue,omitempty"`
	City                                        *HsStr  `json:"city,omitempty"`
	CloseDate                                   *HsTime `json:"closedate,omitempty"`
	Company                                     *HsStr  `json:"company,omitempty"`
	CompanySize                                 *HsStr  `json:"company_size,omitempty"`
	Country                                     *HsStr  `json:"country,omitempty"`
	CreateDate                                  *HsTime `json:"createdate,omitempty"`
	CurrentlyInWorkflow                         *HsStr  `json:"currentlyinworkflow,omitempty"`
	DateOfBirth                                 *HsStr  `json:"date_of_birth,omitempty"`
	DaysToClose                                 *HsStr  `json:"days_to_close,omitempty"`
	Degree                                      *HsStr  `json:"degree,omitempty"`
	Email                                       *HsStr  `json:"email,omitempty"`
	EngagementsLastMeetingBooked                *HsTime `json:"engagements_last_meeting_booked,omitempty"`
	EngagementsLastMeetingBookedCampaign        *HsStr  `json:"engagements_last_meeting_booked_campaign,omitempty"`
	EngagementsLastMeetingBookedMedium          *HsStr  `json:"engagements_last_meeting_booked_medium,omitempty"`
	EngagementsLastMeetingBookedSource          *HsStr  `json:"engagements_last_meeting_booked_source,omitempty"`
	Fax                                         *HsStr  `json:"fax,omitempty"`
	FieldOfStudy                                *HsStr  `json:"field_of_study,omitempty"`
	FirstConversionDate                         *HsTime `json:"first_conversion_date,omitempty"`
	FirstConversionEventName                    *HsStr  `json:"first_conversion_event_name,omitempty"`
	FirstDealCreatedDate                        *HsTime `json:"first_deal_created_date,omitempty"`
	FirstName                                   *HsStr  `json:"firstname,omitempty"`
	Gender                                      *HsStr  `json:"gender,omitempty"`
	GraduationDate                              *HsStr  `json:"graduation_date,omitempty"`
	HsAnalyticsAveragePageViews                 *HsStr  `json:"hs_analytics_average_page_views,omitempty"`
	HsAnalyticsFirstReferrer                    *HsStr  `json:"hs_analytics_first_referrer,omitempty"`
	HsAnalyticsFirstTimestamp                   *HsTime `json:"hs_analytics_first_timestamp,omitempty"`
	HsAnalyticsFirstTouchConvertingCampaign     *HsStr  `json:"hs_analytics_first_touch_converting_campaign,omitempty"`
	HsAnalyticsFirstURL                         *HsStr  `json:"hs_analytics_first_url,omitempty"`
	HsAnalyticsFirstVisitTimestamp              *HsTime `json:"hs_analytics_first_visit_timestamp,omitempty"`
	HsAnalyticsLastReferrer                     *HsStr  `json:"hs_analytics_last_referrer,omitempty"`
	HsAnalyticsLastTimestamp                    *HsTime `json:"hs_analytics_last_timestamp,omitempty"`
	HsAnalyticsLastTouchConvertingCampaign      *HsStr  `json:"hs_analytics_last_touch_converting_campaign,omitempty"`
	HsAnalyticsLastURL                          *HsStr  `json:"hs_analytics_last_url,omitempty"`
	HsAnalyticsLastVisitTimestamp               *HsTime `json:"hs_analytics_last_visit_timestamp,omitempty"`
	HsAnalyticsNumEventCompletions              *HsStr  `json:"hs_analytics_num_event_completions,omitempty"`
	HsAnalyticsNumPageViews                     *HsStr  `json:"hs_analytics_num_page_views,omitempty"`
	HsAnalyticsNumVisits                        *HsStr  `json:"hs_analytics_num_visits,omitempty"`
	HsAnalyticsRevenue                          *HsStr  `json:"hs_analytics_revenue,omitempty"`
	HsAnalyticsSource                           *HsStr  `json:"hs_analytics_source,omitempty"`
	HsAnalyticsSourceData1                      *HsStr  `json:"hs_analytics_source_data_1,omitempty"`
	HsAnalyticsSourceData2                      *HsStr  `json:"hs_analytics_source_data_2,omitempty"`
	HsBuyingRole                                *HsStr  `json:"hs_buying_role,omitempty"`
	HsContentMembershipEmailConfirmed           HsBool  `json:"hs_content_membership_email_confirmed,omitempty"`
	HsContentMembershipNotes                    *HsStr  `json:"hs_content_membership_notes,omitempty"`
	HsContentMembershipRegisteredAt             *HsTime `json:"hs_content_membership_registered_at,omitempty"`
	HsContentMembershipRegistrationDomainSentTo *HsStr  `json:"hs_content_membership_registration_domain_sent_to,omitempty"`
	HsContentMembershipRegistrationEmailSentAt  *HsTime `json:"hs_content_membership_registration_email_sent_at,omitempty"`
	HsContentMembershipStatus                   *HsStr  `json:"hs_content_membership_status,omitempty"`
	HsCreateDate                                *HsTime `json:"hs_createdate,omitempty"`
	HsEmailBadAddress                           HsBool  `json:"hs_email_bad_address,omitempty"`
	HsEmailBounce                               *HsStr  `json:"hs_email_bounce,omitempty"`
	HsEmailClick                                *HsStr  `json:"hs_email_click,omitempty"`
	HsEmailClickDate                            *HsTime `json:"hs_email_first_click_date,omitempty"`
	HsEmailDelivered                            *HsStr  `json:"hs_email_delivered,omitempty"`
	HsEmailDomain                               *HsStr  `json:"hs_email_domain,omitempty"`
	HsEmailFirstOpenDate                        *HsTime `json:"hs_email_first_open_date,omitempty"`
	HsEmailFirstSendDate                        *HsTime `json:"hs_email_first_send_date,omitempty"`
	HsEmailHardBounceReasonEnum                 *HsStr  `json:"hs_email_hard_bounce_reason_enum,omitempty"`
	HsEmailLastClickDate                        *HsTime `json:"hs_email_last_click_date,omitempty"`
	HsEmailLastEmailName                        *HsStr  `json:"hs_email_last_email_name,omitempty"`
	HsEmailLastOpenDate                         *HsTime `json:"hs_email_last_open_date,omitempty"`
	HsEmailLastSendDate                         *HsTime `json:"hs_email_last_send_date,omitempty"`
	HsEmailOpen                                 *HsStr  `json:"hs_email_open,omitempty"`
	HsEmailOpenDate                             *HsTime `json:"hs_email_open_date,omitempty"`
	HsEmailOptOut                               HsBool  `json:"hs_email_optout,omitempty"`
	HsEmailOptOut6766004                        *HsStr  `json:"hs_email_optout_6766004,omitempty"`
	HsEmailOptOut6766098                        *HsStr  `json:"hs_email_optout_6766098,omitempty"`
	HsEmailOptOut6766099                        *HsStr  `json:"hs_email_optout_6766099,omitempty"`
	HsEmailOptOut6766130                        *HsStr  `json:"hs_email_optout_6766130,omitempty"`
	HsEmailQuarantined                          HsBool  `json:"hs_email_quarantined,omitempty"`
	HsEmailSendsSinceLastEngagement             *HsStr  `json:"hs_email_sends_since_last_engagement,omitempty"`
	HsEmailConfirmationStatus                   *HsStr  `json:"hs_emailconfirmationstatus,omitempty"`
	HsFeedbackLastNpsFollowUp                   *HsStr  `json:"hs_feedback_last_nps_follow_up,omitempty"`
	HsFeedbackLastNpsRating                     *HsStr  `json:"hs_feedback_last_nps_rating,omitempty"`
	HsFeedbackLastSurveyDate                    *HsTime `json:"hs_feedback_last_survey_date,omitempty"`
	HsIPTimezone                                *HsStr  `json:"hs_ip_timezone,omitempty"`
	HsIsUnworked                                *HsStr  `json:"hs_is_unworked,omitempty"`
	HsLanguage                                  *HsStr  `json:"hs_language,omitempty"`
	HsLastSalesActivityTimestamp                *HsTime `json:"hs_last_sales_activity_timestamp,omitempty"`
	HsLeadStatus                                *HsStr  `json:"hs_lead_status,omitempty"`
	HsLifeCycleStageCustomerDate                *HsTime `json:"hs_lifecyclestage_customer_date,omitempty"`
	HsLifeCycleStageEvangelistDate              *HsTime `json:"hs_lifecyclestage_evangelist_date,omitempty"`
	HsLifeCycleStageLeadDate                    *HsTime `json:"hs_lifecyclestage_lead_date,omitempty"`
	HsLifeCycleStageMarketingQualifiedLeadDate  *HsTime `json:"hs_lifecyclestage_marketingqualifiedlead_date,omitempty"`
	HsLifeCycleStageOpportunityDate             *HsTime `json:"hs_lifecyclestage_opportunity_date,omitempty"`
	HsLifeCycleStageOtherDate                   *HsTime `json:"hs_lifecyclestage_other_date,omitempty"`
	HsLifeCycleStageSalesQualifiedLeadDate      *HsTime `json:"hs_lifecyclestage_salesqualifiedlead_date,omitempty"`
	HsLifeCycleStageSubscriberDate              *HsTime `json:"hs_lifecyclestage_subscriber_date,omitempty"`
	HsMarketableReasonID                        *HsStr  `json:"hs_marketable_reason_id,omitempty"`
	HsMarketableReasonType                      *HsStr  `json:"hs_marketable_reason_type,omitempty"`
	HsMarketableStatus                          *HsStr  `json:"hs_marketable_status,omitempty"`
	HsMarketableUntilRenewal                    *HsStr  `json:"hs_marketable_until_renewal,omitempty"`
	HsObjectID                                  *HsStr  `json:"hs_object_id,omitempty"`
	HsPersona                                   *HsStr  `json:"hs_persona,omitempty"`
	HsPredictiveContactScoreV2                  *HsStr  `json:"hs_predictivecontactscore_v2,omitempty"`
	HsPredictiveScoringTier                     *HsStr  `json:"hs_predictivescoringtier,omitempty"`
	HsSalesEmailLastClicked                     *HsTime `json:"hs_sales_email_last_clicked,omitempty"`
	HsSalesEmailLastOpened                      *HsTime `json:"hs_sales_email_last_opened,omitempty"`
	HsSalesEmailLastReplied                     *HsTime `json:"hs_sales_email_last_replied,omitempty"`
	HsSequencesIsEnrolled                       HsBool  `json:"hs_sequences_is_enrolled,omitempty"`
	HubspotOwnerAssignedDate                    *HsTime `json:"hubspot_owner_assigneddate,omitempty"`
	HubspotOwnerID                              *HsStr  `json:"hubspot_owner_id,omitempty"`
	HubspotTeamID                               *HsStr  `json:"hubspot_team_id,omitempty"`
	HubspotScore                                *HsStr  `json:"hubspotscore,omitempty"`
	Industry                                    *HsStr  `json:"industry,omitempty"`
	IPCity                                      *HsStr  `json:"ip_city,omitempty"`
	IPCountry                                   *HsStr  `json:"ip_country,omitempty"`
	IPCountryCode                               *HsStr  `json:"ip_country_code,omitempty"`
	IPState                                     *HsStr  `json:"ip_state,omitempty"`
	IPStateCode                                 *HsStr  `json:"ip_state_code,omitempty"`
	JobFunction                                 *HsStr  `json:"job_function,omitempty"`
	JobTitle                                    *HsStr  `json:"jobtitle,omitempty"`
	LastModifiedDate                            *HsTime `json:"lastmodifieddate,omitempty"`
	LastName                                    *HsStr  `json:"lastname,omitempty"`
	LifeCycleStage                              *HsStr  `json:"lifecyclestage,omitempty"`
	MaritalStatus                               *HsStr  `json:"marital_status,omitempty"`
	Message                                     *HsStr  `json:"message,omitempty"`
	MilitaryStatus                              *HsStr  `json:"military_status,omitempty"`
	MobilePhone                                 *HsStr  `json:"mobilephone,omitempty"`
	NotesLastContacted                          *HsTime `json:"notes_last_contacted,omitempty"`
	NotesLastUpdated                            *HsTime `json:"notes_last_updated,omitempty"`
	NotesNextActivityDate                       *HsTime `json:"notes_next_activity_date,omitempty"`
	NumAssociatedDeals                          *HsStr  `json:"num_associated_deals,omitempty"`
	NumContactedNotes                           *HsStr  `json:"num_contacted_notes,omitempty"`
	NumNotes                                    *HsStr  `json:"num_notes,omitempty"`
	NumUniqueConversionEvents                   *HsStr  `json:"num_unique_conversion_events,omitempty"`
	NumEmployees                                *HsStr  `json:"numemployees,omitempty"`
	RecentConversionDate                        *HsTime `json:"recent_conversion_date,omitempty"`
	RecentConversionEventName                   *HsStr  `json:"recent_conversion_event_name,omitempty"`
	RecentDealAmount                            *HsStr  `json:"recent_deal_amount,omitempty"`
	RecentDealCloseDate                         *HsTime `json:"recent_deal_close_date,omitempty"`
	RelationshipStatus                          *HsStr  `json:"relationship_status,omitempty"`
	Salutation                                  *HsStr  `json:"salutation,omitempty"`
	School                                      *HsStr  `json:"school,omitempty"`
	Seniority                                   *HsStr  `json:"seniority,omitempty"`
	StartDate                                   *HsStr  `json:"start_date,omitempty"`
	State                                       *HsStr  `json:"state,omitempty"`
	TotalRevenue                                *HsStr  `json:"total_revenue,omitempty"`
	Website                                     *HsStr  `json:"website,omitempty"`
	WorkEmail                                   *HsStr  `json:"work_email,omitempty"`
	Zip                                         *HsStr  `json:"zip,omitempty"`
}

var defaultContactFields = []string{
	"address",
	"annualrevenue",
	"city",
	"closedate",
	"company",
	"company_size",
	"country",
	"createdate",
	"currentlyinworkflow",
	"date_of_birth",
	"days_to_close",
	"degree",
	"email",
	"engagements_last_meeting_booked",
	"engagements_last_meeting_booked_campaign",
	"engagements_last_meeting_booked_medium",
	"engagements_last_meeting_booked_source",
	"fax",
	"field_of_study",
	"first_conversion_date",
	"first_conversion_event_name",
	"first_deal_created_date",
	"firstname",
	"gender",
	"graduation_date",
	"hs_analytics_average_page_views",
	"hs_analytics_first_referrer",
	"hs_analytics_first_timestamp",
	"hs_analytics_first_touch_converting_campaign",
	"hs_analytics_first_url",
	"hs_analytics_first_visit_timestamp",
	"hs_analytics_last_referrer",
	"hs_analytics_last_timestamp",
	"hs_analytics_last_touch_converting_campaign",
	"hs_analytics_last_url",
	"hs_analytics_last_visit_timestamp",
	"hs_analytics_num_event_completions",
	"hs_analytics_num_page_views",
	"hs_analytics_num_visits",
	"hs_analytics_revenue",
	"hs_analytics_source",
	"hs_analytics_source_data_1",
	"hs_analytics_source_data_2",
	"hs_buying_role",
	"hs_content_membership_email_confirmed",
	"hs_content_membership_notes",
	"hs_content_membership_registered_at",
	"hs_content_membership_registration_domain_sent_to",
	"hs_content_membership_registration_email_sent_at",
	"hs_content_membership_status",
	"hs_createdate",
	"hs_email_bad_address",
	"hs_email_bounce",
	"hs_email_click",
	"hs_email_first_click_date",
	"hs_email_delivered",
	"hs_email_domain",
	"hs_email_first_open_date",
	"hs_email_first_send_date",
	"hs_email_hard_bounce_reason_enum",
	"hs_email_last_click_date",
	"hs_email_last_email_name",
	"hs_email_last_open_date",
	"hs_email_last_send_date",
	"hs_email_open",
	"hs_email_open_date",
	"hs_email_optout",
	"hs_email_optout_6766004",
	"hs_email_optout_6766098",
	"hs_email_optout_6766099",
	"hs_email_optout_6766130",
	"hs_email_quarantined",
	"hs_email_sends_since_last_engagement",
	"hs_emailconfirmationstatus",
	"hs_feedback_last_nps_follow_up",
	"hs_feedback_last_nps_rating",
	"hs_feedback_last_survey_date",
	"hs_ip_timezone",
	"hs_is_unworked",
	"hs_language",
	"hs_last_sales_activity_timestamp",
	"hs_lead_status",
	"hs_lifecyclestage_customer_date",
	"hs_lifecyclestage_evangelist_date",
	"hs_lifecyclestage_lead_date",
	"hs_lifecyclestage_marketingqualifiedlead_date",
	"hs_lifecyclestage_opportunity_date",
	"hs_lifecyclestage_other_date",
	"hs_lifecyclestage_salesqualifiedlead_date",
	"hs_lifecyclestage_subscriber_date",
	"hs_marketable_reason_id",
	"hs_marketable_reason_type",
	"hs_marketable_status",
	"hs_marketable_until_renewal",
	"hs_object_id",
	"hs_persona",
	"hs_predictivecontactscore_v2",
	"hs_predictivescoringtier",
	"hs_sales_email_last_clicked",
	"hs_sales_email_last_opened",
	"hs_sales_email_last_replied",
	"hs_sequences_is_enrolled",
	"hubspot_owner_assigneddate",
	"hubspot_owner_id",
	"hubspot_team_id",
	"hubspotscore",
	"industry",
	"ip_city",
	"ip_country",
	"ip_country_code",
	"ip_state",
	"ip_state_code",
	"job_function",
	"jobtitle",
	"lastmodifieddate",
	"lastname",
	"lifecyclestage",
	"marital_status",
	"message",
	"military_status",
	"mobilephone",
	"notes_last_contacted",
	"notes_last_updated",
	"notes_next_activity_date",
	"num_associated_deals",
	"num_contacted_notes",
	"num_notes",
	"num_unique_conversion_events",
	"numemployees",
	"recent_conversion_date",
	"recent_conversion_event_name",
	"recent_deal_amount",
	"recent_deal_close_date",
	"relationship_status",
	"salutation",
	"school",
	"seniority",
	"start_date",
	"state",
	"total_revenue",
	"website",
	"work_email",
	"zip",
}

// Get gets a contact.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *ContactServiceOp) Get(contactID string, contact interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: contact}
	if err := s.client.Get(s.contactPath+"/"+contactID, resource, option.setupProperties(defaultContactFields)); err != nil {
		return nil, err
	}
	return resource, nil
}

// Create creates a new contact.
// In order to bind the created content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Contact in your own structure.
func (s *ContactServiceOp) Create(contact interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: contact}
	resource := &ResponseResource{Properties: contact}
	if err := s.client.Post(s.contactPath, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Update updates a contact.
// In order to bind the updated content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Contact in your own structure.
func (s *ContactServiceOp) Update(contactID string, contact interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: contact}
	resource := &ResponseResource{Properties: contact}
	if err := s.client.Patch(s.contactPath+"/"+contactID, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Delete deletes a contact.
func (s *ContactServiceOp) Delete(contactID string) error {
	return s.client.Delete(s.contactPath+"/"+contactID, nil)
}

// AssociateAnotherObj associates Contact with another HubSpot objects.
// If you want to associate a custom object, please use a defined value in HubSpot.
func (s *ContactServiceOp) AssociateAnotherObj(contactID string, conf *AssociationConfig) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: &Contact{}}
	if err := s.client.Put(s.contactPath+"/"+contactID+"/"+conf.makeAssociationPath(), nil, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// SearchByEmail searches for a contact by email.
func (s *ContactServiceOp) SearchByEmail(email string) (*ContactSearchResponse, error) {
	req := &ContactSearchRequest{
		FilterGroups: []FilterGroup{
			{
				Filters: []Filter{
					{
						PropertyName: "email",
						Operator:     "EQ",
						Value:        email,
					},
				},
			},
		},
	}

	resource := &ContactSearchResponse{}
	if err := s.client.Post(s.contactPath+"/search", req, resource); err != nil {
		return nil, err
	}

	return resource, nil
}

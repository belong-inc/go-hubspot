package hubspot_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/belong-inc/go-hubspot"
	"github.com/google/go-cmp/cmp"
)

func TestContactServiceOp_Create(t *testing.T) {
	contact := &hubspot.Contact{
		Email:       hubspot.NewString("hubspot@example.com"),
		FirstName:   hubspot.NewString("Bryan"),
		LastName:    hubspot.NewString("Cooper"),
		MobilePhone: hubspot.NewString("(877) 929-0687"),
		Website:     hubspot.NewString("example.com"),
	}

	type fields struct {
		contactPath string
		client      *hubspot.Client
	}
	type args struct {
		contact interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully create a contact",
			fields: fields{
				contactPath: hubspot.ExportContactBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusCreated,
					Header: http.Header{},
					Body:   []byte(`{"id":"contact001","properties":{"createdate":"2019-10-30T03:30:17.883Z","email":"hubspot@example.com","firstname":"Bryan","hs_is_unworked":"true","lastmodifieddate":"2019-12-07T16:50:06.678Z","lastname":"Cooper","mobilephone":"(877) 929-0687","website":"https://example.com"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","archived":false}`),
				}),
			},
			args: args{
				contact: contact,
			},
			want: &hubspot.ResponseResource{
				ID:       "contact001",
				Archived: false,
				Properties: &hubspot.Contact{
					Email:            hubspot.NewString("hubspot@example.com"),
					FirstName:        hubspot.NewString("Bryan"),
					HsIsUnworked:     hubspot.NewString("true"),
					LastName:         hubspot.NewString("Cooper"),
					MobilePhone:      hubspot.NewString("(877) 929-0687"),
					Website:          hubspot.NewString("https://example.com"),
					CreateDate:       &createdAt,
					LastModifiedDate: &modifyDate,
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Received invalid request",
			fields: fields{
				contactPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				contact: contact,
			},
			want: nil,
			wantErr: &hubspot.APIError{
				HTTPStatusCode: http.StatusBadRequest,
				Message:        "Invalid input (details will vary based on the error)",
				CorrelationID:  "aeb5f871-7f07-4993-9211-075dc63e7cbf",
				Category:       "VALIDATION_ERROR",
				Links: hubspot.ErrLinks{
					KnowledgeBase: "https://www.hubspot.com/products/service/knowledge-base",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.CRM.Contact.Create(tt.args.contact)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("Create() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
				t.Errorf("Create() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestContactServiceOp_Update(t *testing.T) {
	contact := &hubspot.Contact{
		Email:       hubspot.NewString("hubspot@example.com"),
		FirstName:   hubspot.NewString("Bryan"),
		LastName:    hubspot.NewString("Cooper"),
		MobilePhone: hubspot.NewString("(877) 929-0687"),
		Website:     hubspot.NewString("example.com"),
	}

	type fields struct {
		contactPath string
		client      *hubspot.Client
	}
	type args struct {
		contactID string
		contact   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully update a contact",
			fields: fields{
				contactPath: hubspot.ExportContactBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusCreated,
					Header: http.Header{},
					Body:   []byte(`{"id":"contact001","properties":{"createdate":"2019-10-30T03:30:17.883Z","email":"hubspot@example.com","firstname":"Bryan","hs_is_unworked":"true","lastmodifieddate":"2019-12-07T16:50:06.678Z","lastname":"Cooper","mobilephone":"(877) 929-0687","website":"https://example.com"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","archived":false}`),
				}),
			},
			args: args{
				contact: contact,
			},
			want: &hubspot.ResponseResource{
				ID:       "contact001",
				Archived: false,
				Properties: &hubspot.Contact{
					Email:            hubspot.NewString("hubspot@example.com"),
					FirstName:        hubspot.NewString("Bryan"),
					HsIsUnworked:     hubspot.NewString("true"),
					LastName:         hubspot.NewString("Cooper"),
					MobilePhone:      hubspot.NewString("(877) 929-0687"),
					Website:          hubspot.NewString("https://example.com"),
					CreateDate:       &createdAt,
					LastModifiedDate: &modifyDate,
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Received invalid request",
			fields: fields{
				contactPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				contactID: "contact001",
				contact:   contact,
			},
			want: nil,
			wantErr: &hubspot.APIError{
				HTTPStatusCode: http.StatusBadRequest,
				Message:        "Invalid input (details will vary based on the error)",
				CorrelationID:  "aeb5f871-7f07-4993-9211-075dc63e7cbf",
				Category:       "VALIDATION_ERROR",
				Links: hubspot.ErrLinks{
					KnowledgeBase: "https://www.hubspot.com/products/service/knowledge-base",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.CRM.Contact.Update(tt.args.contactID, tt.args.contact)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("Update() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
				t.Errorf("Update() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestContactServiceOp_Get(t *testing.T) {
	type CustomFields struct {
		hubspot.Contact
		CustomName string          `json:"custom_name,omitempty"`
		CustomDate *hubspot.HsTime `json:"custom_date,omitempty"`
	}

	type fields struct {
		contactPath string
		client      *hubspot.Client
	}
	type args struct {
		contactID string
		contact   interface{}
		option    *hubspot.RequestQueryOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully get a contact",
			fields: fields{
				contactPath: hubspot.ExportContactBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"id":"contact001","properties":{"address":"address","annualrevenue":"0.0","city":"chiyodaku","closedate":"2019-12-07T16:50:06.678Z","company":"company","company_size":"","country":"Japan","createdate":"2019-10-30T03:30:17.883Z","currentlyinworkflow":"workflow","date_of_birth":"1990-01-01","days_to_close":"10","degree":"degree","email":"hubspot@example.com","engagements_last_meeting_booked":"2019-10-30T03:30:17.883Z","engagements_last_meeting_booked_campaign":"campaign","engagements_last_meeting_booked_medium":"medium","engagements_last_meeting_booked_source":"source","fax":"0123456789","field_of_study":"field_of_study","first_conversion_date":"","first_conversion_event_name":"event_name","first_deal_created_date":"","firstname":"Bryan","gender":"male","graduation_date":"2019-12-31","hs_analytics_average_page_views":"0","hs_analytics_first_referrer":"first_referrer","hs_analytics_first_timestamp":"2019-10-30T03:30:17.883Z","hs_analytics_first_touch_converting_campaign":"campaign","hs_analytics_first_url":"https://example.com","hs_analytics_first_visit_timestamp":"2019-10-30T03:30:17.883Z","hs_analytics_last_referrer":null,"hs_analytics_last_timestamp":null,"hs_analytics_last_touch_converting_campaign":null,"hs_analytics_last_url":"https://example.com","hs_analytics_last_visit_timestamp":null,"hs_analytics_num_event_completions":"0","hs_analytics_num_page_views":"0","hs_analytics_num_visits":"0","hs_analytics_revenue":"0.0","hs_analytics_source":"OFFLINE","hs_analytics_source_data_1":"API","hs_analytics_source_data_2":"PublicObjectResource","hs_buying_role":"role","hs_content_membership_email_confirmed":null,"hs_content_membership_notes":null,"hs_content_membership_registered_at":null,"hs_content_membership_registration_domain_sent_to":null,"hs_content_membership_registration_email_sent_at":null,"hs_content_membership_status":null,"hs_createdate":null,"hs_email_bad_address":"true","hs_email_bounce":null,"hs_email_click":null,"hs_email_delivered":null,"hs_email_domain":"example.com","hs_email_first_click_date":null,"hs_email_first_open_date":null,"hs_email_first_send_date":null,"hs_email_hard_bounce_reason_enum":null,"hs_email_last_click_date":null,"hs_email_last_email_name":null,"hs_email_last_open_date":null,"hs_email_last_send_date":null,"hs_email_open":null,"hs_email_optout":null,"hs_email_optout_6766004":null,"hs_email_optout_6766098":null,"hs_email_optout_6766099":null,"hs_email_optout_6766130":null,"hs_email_quarantined":null,"hs_email_sends_since_last_engagement":null,"hs_emailconfirmationstatus":"status","hs_feedback_last_nps_follow_up":null,"hs_feedback_last_nps_rating":null,"hs_feedback_last_survey_date":null,"hs_ip_timezone":null,"hs_is_unworked":"true","hs_language":null,"hs_last_sales_activity_timestamp":null,"hs_lead_status":null,"hs_lifecyclestage_customer_date":null,"hs_lifecyclestage_evangelist_date":null,"hs_lifecyclestage_lead_date":null,"hs_lifecyclestage_marketingqualifiedlead_date":null,"hs_lifecyclestage_opportunity_date":null,"hs_lifecyclestage_other_date":null,"hs_lifecyclestage_salesqualifiedlead_date":null,"hs_lifecyclestage_subscriber_date":null,"hs_marketable_reason_id":null,"hs_marketable_reason_type":null,"hs_marketable_status":null,"hs_marketable_until_renewal":null,"hs_object_id":"contact001","hs_persona":null,"hs_predictivecontactscore_v2":"1.02","hs_predictivescoringtier":"tier_1","hs_sales_email_last_clicked":null,"hs_sales_email_last_opened":null,"hs_sales_email_last_replied":null,"hs_sequences_is_enrolled":null,"hubspot_owner_assigneddate":null,"hubspot_owner_id":null,"hubspot_team_id":null,"hubspotscore":null,"industry":null,"ip_city":null,"ip_country":null,"ip_country_code":null,"ip_state":null,"ip_state_code":null,"job_function":null,"jobtitle":null,"lastmodifieddate":"2019-12-07T16:50:06.678Z","lastname":"Cooper","lifecyclestage":null,"marital_status":null,"message":"message","military_status":null,"mobilephone":"(877) 929-0687","notes_last_contacted":null,"notes_last_updated":null,"notes_next_activity_date":"","num_associated_deals":null,"num_contacted_notes":null,"num_notes":null,"num_unique_conversion_events":"0","numemployees":null,"recent_conversion_date":"2019-10-30T03:30:17.883Z","recent_conversion_event_name":"event","recent_deal_amount":null,"recent_deal_close_date":null,"relationship_status":null,"salutation":null,"school":"school","seniority":null,"start_date":null,"state":"state","total_revenue":"0.0","website":"https://example.com","work_email":null,"zip":"1000001"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","archived":false}`),
				}),
			},
			args: args{
				contactID: "contact001",
				contact:   &hubspot.Contact{},
			},
			want: &hubspot.ResponseResource{
				ID:       "contact001",
				Archived: false,
				Properties: &hubspot.Contact{
					Address:                                 hubspot.NewString("address"),
					AnnualRevenue:                           hubspot.NewString("0.0"),
					City:                                    hubspot.NewString("chiyodaku"),
					CloseDate:                               &closeDate,
					Company:                                 hubspot.NewString("company"),
					CompanySize:                             hubspot.BlankStr,
					Country:                                 hubspot.NewString("Japan"),
					CreateDate:                              &createdAt,
					CurrentlyInWorkflow:                     hubspot.NewString("workflow"),
					DateOfBirth:                             hubspot.NewString("1990-01-01"),
					DaysToClose:                             hubspot.NewString("10"),
					Degree:                                  hubspot.NewString("degree"),
					Email:                                   hubspot.NewString("hubspot@example.com"),
					EngagementsLastMeetingBooked:            &createdAt,
					EngagementsLastMeetingBookedCampaign:    hubspot.NewString("campaign"),
					EngagementsLastMeetingBookedMedium:      hubspot.NewString("medium"),
					EngagementsLastMeetingBookedSource:      hubspot.NewString("source"),
					Fax:                                     hubspot.NewString("0123456789"),
					FieldOfStudy:                            hubspot.NewString("field_of_study"),
					FirstConversionDate:                     &hubspot.HsTime{},
					FirstConversionEventName:                hubspot.NewString("event_name"),
					FirstDealCreatedDate:                    &hubspot.HsTime{},
					FirstName:                               hubspot.NewString("Bryan"),
					Gender:                                  hubspot.NewString("male"),
					GraduationDate:                          hubspot.NewString("2019-12-31"),
					HsAnalyticsAveragePageViews:             hubspot.NewString("0"),
					HsAnalyticsFirstReferrer:                hubspot.NewString("first_referrer"),
					HsAnalyticsFirstTimestamp:               &createdAt,
					HsAnalyticsFirstTouchConvertingCampaign: hubspot.NewString("campaign"),
					HsAnalyticsFirstURL:                     hubspot.NewString("https://example.com"),
					HsAnalyticsFirstVisitTimestamp:          &createdAt,
					HsAnalyticsLastReferrer:                 nil,
					HsAnalyticsLastTimestamp:                nil,
					HsAnalyticsLastTouchConvertingCampaign:  nil,
					HsAnalyticsLastURL:                      hubspot.NewString("https://example.com"),
					HsAnalyticsLastVisitTimestamp:           nil,
					HsAnalyticsNumEventCompletions:          hubspot.NewString("0"),
					HsAnalyticsNumPageViews:                 hubspot.NewString("0"),
					HsAnalyticsNumVisits:                    hubspot.NewString("0"),
					HsAnalyticsRevenue:                      hubspot.NewString("0.0"),
					HsAnalyticsSource:                       hubspot.NewString("OFFLINE"),
					HsAnalyticsSourceData1:                  hubspot.NewString("API"),
					HsAnalyticsSourceData2:                  hubspot.NewString("PublicObjectResource"),
					HsBuyingRole:                            hubspot.NewString("role"),
					HsContentMembershipEmailConfirmed:       false,
					HsContentMembershipNotes:                nil,
					HsContentMembershipRegisteredAt:         nil,
					HsContentMembershipRegistrationDomainSentTo: nil,
					HsContentMembershipRegistrationEmailSentAt:  nil,
					HsContentMembershipStatus:                   nil,
					HsCreateDate:                                nil,
					HsEmailBadAddress:                           true,
					HsEmailBounce:                               nil,
					HsEmailClick:                                nil,
					HsEmailClickDate:                            nil,
					HsEmailDelivered:                            nil,
					HsEmailDomain:                               hubspot.NewString("example.com"),
					HsEmailFirstOpenDate:                        nil,
					HsEmailFirstSendDate:                        nil,
					HsEmailHardBounceReasonEnum:                 nil,
					HsEmailLastClickDate:                        nil,
					HsEmailLastEmailName:                        nil,
					HsEmailLastOpenDate:                         nil,
					HsEmailLastSendDate:                         nil,
					HsEmailOpen:                                 nil,
					HsEmailOpenDate:                             nil,
					HsEmailOptOut:                               false,
					HsEmailOptOut6766004:                        nil,
					HsEmailOptOut6766098:                        nil,
					HsEmailOptOut6766099:                        nil,
					HsEmailOptOut6766130:                        nil,
					HsEmailQuarantined:                          false,
					HsEmailSendsSinceLastEngagement:             nil,
					HsEmailConfirmationStatus:                   hubspot.NewString("status"),
					HsFeedbackLastNpsFollowUp:                   nil,
					HsFeedbackLastNpsRating:                     nil,
					HsFeedbackLastSurveyDate:                    nil,
					HsIPTimezone:                                nil,
					HsIsUnworked:                                hubspot.NewString("true"),
					HsLanguage:                                  nil,
					HsLastSalesActivityTimestamp:                nil,
					HsLeadStatus:                                nil,
					HsLifeCycleStageCustomerDate:                nil,
					HsLifeCycleStageEvangelistDate:              nil,
					HsLifeCycleStageLeadDate:                    nil,
					HsLifeCycleStageMarketingQualifiedLeadDate:  nil,
					HsLifeCycleStageOpportunityDate:             nil,
					HsLifeCycleStageOtherDate:                   nil,
					HsLifeCycleStageSalesQualifiedLeadDate:      nil,
					HsLifeCycleStageSubscriberDate:              nil,
					HsMarketableReasonID:                        nil,
					HsMarketableReasonType:                      nil,
					HsMarketableStatus:                          nil,
					HsMarketableUntilRenewal:                    nil,
					HsObjectID:                                  hubspot.NewString("contact001"),
					HsPersona:                                   nil,
					HsPredictiveContactScoreV2:                  hubspot.NewString("1.02"),
					HsPredictiveScoringTier:                     hubspot.NewString("tier_1"),
					HsSalesEmailLastClicked:                     nil,
					HsSalesEmailLastOpened:                      nil,
					HsSalesEmailLastReplied:                     nil,
					HsSequencesIsEnrolled:                       false,
					HubspotOwnerAssignedDate:                    nil,
					HubspotOwnerID:                              nil,
					HubspotTeamID:                               nil,
					HubspotScore:                                nil,
					Industry:                                    nil,
					IPCity:                                      nil,
					IPCountry:                                   nil,
					IPCountryCode:                               nil,
					IPState:                                     nil,
					IPStateCode:                                 nil,
					JobFunction:                                 nil,
					JobTitle:                                    nil,
					LastModifiedDate:                            &modifyDate,
					LastName:                                    hubspot.NewString("Cooper"),
					LifeCycleStage:                              nil,
					MaritalStatus:                               nil,
					Message:                                     hubspot.NewString("message"),
					MilitaryStatus:                              nil,
					MobilePhone:                                 hubspot.NewString("(877) 929-0687"),
					NotesLastContacted:                          nil,
					NotesLastUpdated:                            nil,
					NotesNextActivityDate:                       &hubspot.HsTime{},
					NumAssociatedDeals:                          nil,
					NumContactedNotes:                           nil,
					NumNotes:                                    nil,
					NumUniqueConversionEvents:                   hubspot.NewString("0"),
					NumEmployees:                                nil,
					RecentConversionDate:                        &createdAt,
					RecentConversionEventName:                   hubspot.NewString("event"),
					RecentDealAmount:                            nil,
					RecentDealCloseDate:                         nil,
					RelationshipStatus:                          nil,
					Salutation:                                  nil,
					School:                                      hubspot.NewString("school"),
					Seniority:                                   nil,
					StartDate:                                   nil,
					State:                                       hubspot.NewString("state"),
					TotalRevenue:                                hubspot.NewString("0.0"),
					Website:                                     hubspot.NewString("https://example.com"),
					WorkEmail:                                   nil,
					Zip:                                         hubspot.NewString("1000001"),
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Successfully get a deal with custom fields",
			fields: fields{
				contactPath: hubspot.ExportContactBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"id":"contact001","properties":{"address":"address","annualrevenue":"0.0","city":"chiyodaku","closedate":"2019-12-07T16:50:06.678Z","company":"company","company_size":"","country":"Japan","createdate":"2019-10-30T03:30:17.883Z","currentlyinworkflow":"workflow","date_of_birth":"1990-01-01","days_to_close":"10","degree":"degree","email":"hubspot@example.com","engagements_last_meeting_booked":"2019-10-30T03:30:17.883Z","engagements_last_meeting_booked_campaign":"campaign","engagements_last_meeting_booked_medium":"medium","engagements_last_meeting_booked_source":"source","fax":"0123456789","field_of_study":"field_of_study","first_conversion_date":"","first_conversion_event_name":"event_name","first_deal_created_date":"","firstname":"Bryan","gender":"male","graduation_date":"2019-12-31","hs_analytics_average_page_views":"0","hs_analytics_first_referrer":"first_referrer","hs_analytics_first_timestamp":"2019-10-30T03:30:17.883Z","hs_analytics_first_touch_converting_campaign":"campaign","hs_analytics_first_url":"https://example.com","hs_analytics_first_visit_timestamp":"2019-10-30T03:30:17.883Z","hs_analytics_last_referrer":null,"hs_analytics_last_timestamp":null,"hs_analytics_last_touch_converting_campaign":null,"hs_analytics_last_url":"https://example.com","hs_analytics_last_visit_timestamp":null,"hs_analytics_num_event_completions":"0","hs_analytics_num_page_views":"0","hs_analytics_num_visits":"0","hs_analytics_revenue":"0.0","hs_analytics_source":"OFFLINE","hs_analytics_source_data_1":"API","hs_analytics_source_data_2":"PublicObjectResource","hs_buying_role":"role","hs_content_membership_email_confirmed":null,"hs_content_membership_notes":null,"hs_content_membership_registered_at":null,"hs_content_membership_registration_domain_sent_to":null,"hs_content_membership_registration_email_sent_at":null,"hs_content_membership_status":null,"hs_createdate":null,"hs_email_bad_address":"false","hs_email_bounce":null,"hs_email_click":null,"hs_email_delivered":null,"hs_email_domain":"example.com","hs_email_first_click_date":null,"hs_email_first_open_date":null,"hs_email_first_send_date":null,"hs_email_hard_bounce_reason_enum":null,"hs_email_last_click_date":null,"hs_email_last_email_name":null,"hs_email_last_open_date":null,"hs_email_last_send_date":null,"hs_email_open":null,"hs_email_optout":null,"hs_email_optout_6766004":null,"hs_email_optout_6766098":null,"hs_email_optout_6766099":null,"hs_email_optout_6766130":null,"hs_email_quarantined":null,"hs_email_sends_since_last_engagement":null,"hs_emailconfirmationstatus":"status","hs_feedback_last_nps_follow_up":null,"hs_feedback_last_nps_rating":null,"hs_feedback_last_survey_date":null,"hs_ip_timezone":null,"hs_is_unworked":"true","hs_language":null,"hs_last_sales_activity_timestamp":null,"hs_lead_status":null,"hs_lifecyclestage_customer_date":null,"hs_lifecyclestage_evangelist_date":null,"hs_lifecyclestage_lead_date":null,"hs_lifecyclestage_marketingqualifiedlead_date":null,"hs_lifecyclestage_opportunity_date":null,"hs_lifecyclestage_other_date":null,"hs_lifecyclestage_salesqualifiedlead_date":null,"hs_lifecyclestage_subscriber_date":null,"hs_marketable_reason_id":null,"hs_marketable_reason_type":null,"hs_marketable_status":null,"hs_marketable_until_renewal":null,"hs_object_id":"contact001","hs_persona":null,"hs_predictivecontactscore_v2":"1.02","hs_predictivescoringtier":"tier_1","hs_sales_email_last_clicked":null,"hs_sales_email_last_opened":null,"hs_sales_email_last_replied":null,"hs_sequences_is_enrolled":null,"hubspot_owner_assigneddate":null,"hubspot_owner_id":null,"hubspot_team_id":null,"hubspotscore":null,"industry":null,"ip_city":null,"ip_country":null,"ip_country_code":null,"ip_state":null,"ip_state_code":null,"job_function":null,"jobtitle":null,"lastmodifieddate":"2019-12-07T16:50:06.678Z","lastname":"Cooper","lifecyclestage":null,"marital_status":null,"message":"message","military_status":null,"mobilephone":"(877) 929-0687","notes_last_contacted":null,"notes_last_updated":null,"notes_next_activity_date":"","num_associated_deals":null,"num_contacted_notes":null,"num_notes":null,"num_unique_conversion_events":"0","numemployees":null,"recent_conversion_date":"2019-10-30T03:30:17.883Z","recent_conversion_event_name":"event","recent_deal_amount":null,"recent_deal_close_date":null,"relationship_status":null,"salutation":null,"school":"school","seniority":null,"start_date":null,"state":"state","total_revenue":"0.0","website":"https://example.com","work_email":null,"zip":"1000001","custom_name":"hubspot custom","custom_date":"2019-12-07T16:50:06.678Z"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","archived":false}`),
				}),
			},
			args: args{
				contactID: "contact001",
				contact:   &CustomFields{},
				option: &hubspot.RequestQueryOption{
					CustomProperties: []string{
						"custom_name",
						"custom_date",
					},
				},
			},
			want: &hubspot.ResponseResource{
				ID:       "contact001",
				Archived: false,
				Properties: &CustomFields{
					Contact: hubspot.Contact{
						Address:                                 hubspot.NewString("address"),
						AnnualRevenue:                           hubspot.NewString("0.0"),
						City:                                    hubspot.NewString("chiyodaku"),
						CloseDate:                               &closeDate,
						Company:                                 hubspot.NewString("company"),
						CompanySize:                             hubspot.BlankStr,
						Country:                                 hubspot.NewString("Japan"),
						CreateDate:                              &createdAt,
						CurrentlyInWorkflow:                     hubspot.NewString("workflow"),
						DateOfBirth:                             hubspot.NewString("1990-01-01"),
						DaysToClose:                             hubspot.NewString("10"),
						Degree:                                  hubspot.NewString("degree"),
						Email:                                   hubspot.NewString("hubspot@example.com"),
						EngagementsLastMeetingBooked:            &createdAt,
						EngagementsLastMeetingBookedCampaign:    hubspot.NewString("campaign"),
						EngagementsLastMeetingBookedMedium:      hubspot.NewString("medium"),
						EngagementsLastMeetingBookedSource:      hubspot.NewString("source"),
						Fax:                                     hubspot.NewString("0123456789"),
						FieldOfStudy:                            hubspot.NewString("field_of_study"),
						FirstConversionDate:                     &hubspot.HsTime{},
						FirstConversionEventName:                hubspot.NewString("event_name"),
						FirstDealCreatedDate:                    &hubspot.HsTime{},
						FirstName:                               hubspot.NewString("Bryan"),
						Gender:                                  hubspot.NewString("male"),
						GraduationDate:                          hubspot.NewString("2019-12-31"),
						HsAnalyticsAveragePageViews:             hubspot.NewString("0"),
						HsAnalyticsFirstReferrer:                hubspot.NewString("first_referrer"),
						HsAnalyticsFirstTimestamp:               &createdAt,
						HsAnalyticsFirstTouchConvertingCampaign: hubspot.NewString("campaign"),
						HsAnalyticsFirstURL:                     hubspot.NewString("https://example.com"),
						HsAnalyticsFirstVisitTimestamp:          &createdAt,
						HsAnalyticsLastReferrer:                 nil,
						HsAnalyticsLastTimestamp:                nil,
						HsAnalyticsLastTouchConvertingCampaign:  nil,
						HsAnalyticsLastURL:                      hubspot.NewString("https://example.com"),
						HsAnalyticsLastVisitTimestamp:           nil,
						HsAnalyticsNumEventCompletions:          hubspot.NewString("0"),
						HsAnalyticsNumPageViews:                 hubspot.NewString("0"),
						HsAnalyticsNumVisits:                    hubspot.NewString("0"),
						HsAnalyticsRevenue:                      hubspot.NewString("0.0"),
						HsAnalyticsSource:                       hubspot.NewString("OFFLINE"),
						HsAnalyticsSourceData1:                  hubspot.NewString("API"),
						HsAnalyticsSourceData2:                  hubspot.NewString("PublicObjectResource"),
						HsBuyingRole:                            hubspot.NewString("role"),
						HsContentMembershipEmailConfirmed:       false,
						HsContentMembershipNotes:                nil,
						HsContentMembershipRegisteredAt:         nil,
						HsContentMembershipRegistrationDomainSentTo: nil,
						HsContentMembershipRegistrationEmailSentAt:  nil,
						HsContentMembershipStatus:                   nil,
						HsCreateDate:                                nil,
						HsEmailBadAddress:                           false,
						HsEmailBounce:                               nil,
						HsEmailClick:                                nil,
						HsEmailClickDate:                            nil,
						HsEmailDelivered:                            nil,
						HsEmailDomain:                               hubspot.NewString("example.com"),
						HsEmailFirstOpenDate:                        nil,
						HsEmailFirstSendDate:                        nil,
						HsEmailHardBounceReasonEnum:                 nil,
						HsEmailLastClickDate:                        nil,
						HsEmailLastEmailName:                        nil,
						HsEmailLastOpenDate:                         nil,
						HsEmailLastSendDate:                         nil,
						HsEmailOpen:                                 nil,
						HsEmailOpenDate:                             nil,
						HsEmailOptOut:                               false,
						HsEmailOptOut6766004:                        nil,
						HsEmailOptOut6766098:                        nil,
						HsEmailOptOut6766099:                        nil,
						HsEmailOptOut6766130:                        nil,
						HsEmailQuarantined:                          false,
						HsEmailSendsSinceLastEngagement:             nil,
						HsEmailConfirmationStatus:                   hubspot.NewString("status"),
						HsFeedbackLastNpsFollowUp:                   nil,
						HsFeedbackLastNpsRating:                     nil,
						HsFeedbackLastSurveyDate:                    nil,
						HsIPTimezone:                                nil,
						HsIsUnworked:                                hubspot.NewString("true"),
						HsLanguage:                                  nil,
						HsLastSalesActivityTimestamp:                nil,
						HsLeadStatus:                                nil,
						HsLifeCycleStageCustomerDate:                nil,
						HsLifeCycleStageEvangelistDate:              nil,
						HsLifeCycleStageLeadDate:                    nil,
						HsLifeCycleStageMarketingQualifiedLeadDate:  nil,
						HsLifeCycleStageOpportunityDate:             nil,
						HsLifeCycleStageOtherDate:                   nil,
						HsLifeCycleStageSalesQualifiedLeadDate:      nil,
						HsLifeCycleStageSubscriberDate:              nil,
						HsMarketableReasonID:                        nil,
						HsMarketableReasonType:                      nil,
						HsMarketableStatus:                          nil,
						HsMarketableUntilRenewal:                    nil,
						HsObjectID:                                  hubspot.NewString("contact001"),
						HsPersona:                                   nil,
						HsPredictiveContactScoreV2:                  hubspot.NewString("1.02"),
						HsPredictiveScoringTier:                     hubspot.NewString("tier_1"),
						HsSalesEmailLastClicked:                     nil,
						HsSalesEmailLastOpened:                      nil,
						HsSalesEmailLastReplied:                     nil,
						HsSequencesIsEnrolled:                       false,
						HubspotOwnerAssignedDate:                    nil,
						HubspotOwnerID:                              nil,
						HubspotTeamID:                               nil,
						HubspotScore:                                nil,
						Industry:                                    nil,
						IPCity:                                      nil,
						IPCountry:                                   nil,
						IPCountryCode:                               nil,
						IPState:                                     nil,
						IPStateCode:                                 nil,
						JobFunction:                                 nil,
						JobTitle:                                    nil,
						LastModifiedDate:                            &modifyDate,
						LastName:                                    hubspot.NewString("Cooper"),
						LifeCycleStage:                              nil,
						MaritalStatus:                               nil,
						Message:                                     hubspot.NewString("message"),
						MilitaryStatus:                              nil,
						MobilePhone:                                 hubspot.NewString("(877) 929-0687"),
						NotesLastContacted:                          nil,
						NotesLastUpdated:                            nil,
						NotesNextActivityDate:                       &hubspot.HsTime{},
						NumAssociatedDeals:                          nil,
						NumContactedNotes:                           nil,
						NumNotes:                                    nil,
						NumUniqueConversionEvents:                   hubspot.NewString("0"),
						NumEmployees:                                nil,
						RecentConversionDate:                        &createdAt,
						RecentConversionEventName:                   hubspot.NewString("event"),
						RecentDealAmount:                            nil,
						RecentDealCloseDate:                         nil,
						RelationshipStatus:                          nil,
						Salutation:                                  nil,
						School:                                      hubspot.NewString("school"),
						Seniority:                                   nil,
						StartDate:                                   nil,
						State:                                       hubspot.NewString("state"),
						TotalRevenue:                                hubspot.NewString("0.0"),
						Website:                                     hubspot.NewString("https://example.com"),
						WorkEmail:                                   nil,
						Zip:                                         hubspot.NewString("1000001"),
					},
					CustomName: "hubspot custom",
					CustomDate: &updatedAt,
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Received invalid request",
			fields: fields{
				contactPath: hubspot.ExportContactBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				contactID: "contact001",
				contact:   &hubspot.Contact{},
			},
			want: nil,
			wantErr: &hubspot.APIError{
				HTTPStatusCode: http.StatusBadRequest,
				Message:        "Invalid input (details will vary based on the error)",
				CorrelationID:  "aeb5f871-7f07-4993-9211-075dc63e7cbf",
				Category:       "VALIDATION_ERROR",
				Links: hubspot.ErrLinks{
					KnowledgeBase: "https://www.hubspot.com/products/service/knowledge-base",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.CRM.Contact.Get(tt.args.contactID, tt.args.contact, tt.args.option)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("Get() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
				t.Errorf("Get() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

func TestContactServiceOp_AssociateAnotherObj(t *testing.T) {
	type fields struct {
		contactPath string
		client      *hubspot.Client
	}
	type args struct {
		contactID string
		conf      *hubspot.AssociationConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully associated object",
			fields: fields{
				contactPath: hubspot.ExportContactBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"id":"contact001","properties":{"createdate":"2019-10-30T03:30:17.883Z","hs_object_id":"contact001","lastmodifieddate":"2019-12-07T16:50:06.678Z"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","associations":{"deals":{"results":[{"id":"deal001","type":"contact_to_deal"}]}},"archived":false,"archivedAt":"2020-01-01T00:00:00.000Z"}`),
				}),
			},
			args: args{
				contactID: "contact001",
				conf: &hubspot.AssociationConfig{
					ToObject:   hubspot.ObjectTypeDeal,
					ToObjectID: "deal001",
					Type:       hubspot.AssociationTypeContactToDeal,
				},
			},
			want: &hubspot.ResponseResource{
				ID:       "contact001",
				Archived: false,
				Associations: &hubspot.Associations{
					Deals: struct {
						Results []hubspot.AssociationResult `json:"results"`
					}{
						Results: []hubspot.AssociationResult{
							{ID: "deal001", Type: string(hubspot.AssociationTypeContactToDeal)},
						},
					},
				},
				Properties: &hubspot.Contact{
					HsObjectID:       hubspot.NewString("contact001"),
					CreateDate:       &createdAt,
					LastModifiedDate: &updatedAt,
				},
				CreatedAt:  &createdAt,
				UpdatedAt:  &updatedAt,
				ArchivedAt: &archivedAt,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.CRM.Contact.AssociateAnotherObj(tt.args.contactID, tt.args.conf)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("AssociateAnotherObj() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
				t.Errorf("AssociateAnotherObj() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

package hubspot_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/courtyard-nft/go-hubspot/"
	"github.com/google/go-cmp/cmp"
)

func TestDealServiceOp_Create(t *testing.T) {
	deal := &hubspot.Deal{
		Amount:      hubspot.NewString("1500.00"),
		DealName:    hubspot.NewString("Custom data integrations"),
		DealStage:   hubspot.NewString("presentation scheduled"),
		DealOwnerID: hubspot.NewString("910901"),
		PipeLine:    hubspot.NewString("default"),
		CloseDate:   &closeDate,
	}

	type fields struct {
		dealPath string
		client   *hubspot.Client
	}
	type args struct {
		deal interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully create a deal",
			fields: fields{
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusCreated,
					Header: http.Header{},
					Body:   []byte(`{"createdAt":"2019-10-30T03:30:17.883Z","archived":false,"id":"512","properties":{"amount":"1500.00","closedate":"2019-12-07T16:50:06.678Z","createdate":"2019-10-30T03:30:17.883Z","dealname":"Custom data integrations","dealstage":"presentation scheduled","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","hubspot_owner_id":"910901","pipeline":"default"},"updatedAt":"2019-12-07T16:50:06.678Z"}`),
				}),
			},
			args: args{
				deal: deal,
			},
			want: &hubspot.ResponseResource{
				ID:       "512",
				Archived: false,
				Properties: &hubspot.Deal{
					Amount:           hubspot.NewString("1500.00"),
					DealName:         hubspot.NewString("Custom data integrations"),
					DealStage:        hubspot.NewString("presentation scheduled"),
					DealOwnerID:      hubspot.NewString("910901"),
					PipeLine:         hubspot.NewString("default"),
					CreateDate:       &createdAt,
					CloseDate:        &closeDate,
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
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				deal: deal,
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
			got, err := tt.fields.client.CRM.Deal.Create(tt.args.deal)
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

func TestDealServiceOp_Update(t *testing.T) {
	deal := &hubspot.Deal{
		Amount:      hubspot.NewString("1500.00"),
		DealName:    hubspot.NewString("Custom data integrations"),
		DealStage:   hubspot.NewString("presentation scheduled"),
		DealOwnerID: hubspot.NewString("910901"),
		PipeLine:    hubspot.NewString("default"),
		CloseDate:   &closeDate,
	}

	type fields struct {
		dealPath string
		client   *hubspot.Client
	}
	type args struct {
		dealID string
		deal   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully update a deal",
			fields: fields{
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"createdAt": "2019-10-30T03:30:17.883Z","archived": false,"id": "512","properties": {"amount": "1500.00","closedate": "2019-12-07T16:50:06.678Z","createdate": "2019-10-30T03:30:17.883Z","dealname": "Custom data integrations","dealstage": "presentation scheduled","hs_lastmodifieddate": "2019-12-07T16:50:06.678Z","hubspot_owner_id": "910901","pipeline": "default"},"updatedAt": "2019-12-07T16:50:06.678Z"}`),
				}),
			},
			args: args{
				dealID: "512",
				deal:   deal,
			},
			want: &hubspot.ResponseResource{
				ID:       "512",
				Archived: false,
				Properties: &hubspot.Deal{
					Amount:           hubspot.NewString("1500.00"),
					DealName:         hubspot.NewString("Custom data integrations"),
					DealStage:        hubspot.NewString("presentation scheduled"),
					DealOwnerID:      hubspot.NewString("910901"),
					PipeLine:         hubspot.NewString("default"),
					CreateDate:       &createdAt,
					CloseDate:        &closeDate,
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
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				dealID: "512",
				deal:   deal,
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
			got, err := tt.fields.client.CRM.Deal.Update(tt.args.dealID, tt.args.deal)
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

func TestDealServiceOp_Get(t *testing.T) {
	type CustomFields struct {
		hubspot.Deal
		CustomName string          `json:"custom_name,omitempty"`
		CustomDate *hubspot.HsTime `json:"custom_date,omitempty"`
	}

	type fields struct {
		dealPath string
		client   *hubspot.Client
	}
	type args struct {
		dealID string
		deal   interface{}
		option *hubspot.RequestQueryOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully get a deal",
			fields: fields{
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"id":"512","properties":{"amount":"1500.00","amount_in_home_currency":"1500.00","closed_lost_reason":"lost reason","closed_won_reason":"won reason","closedate":"2019-12-07T16:50:06.678Z","createdate":"2019-10-30T03:30:17.883Z","dealname":"Custom data integrations","dealstage":"presentation scheduled","dealtype":null,"description":"test hubspot deal record","hs_acv":"1000.00","hs_arr":null,"hs_forecast_amount":"1500.00","hs_forecast_probability":"900101","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","hs_mrr":"100.00","hs_next_step":"tel","hs_object_id":"512","hs_tcv":"0","hubspot_owner_assigneddate":null,"hubspot_owner_id":"910901","hubspot_team_id":null,"notes_last_contacted":null,"notes_last_updated":"2019-12-07T16:50:06.678Z","notes_next_activity_date":"","num_associated_contacts":"0","num_contacted_notes":"0","num_notes":"1","pipeline":"default"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","archived":false}`),
				}),
			},
			args: args{
				dealID: "512",
				deal:   &hubspot.Deal{},
			},
			want: &hubspot.ResponseResource{
				ID:       "512",
				Archived: false,
				Properties: &hubspot.Deal{
					Amount:                  hubspot.NewString("1500.00"),
					AmountInCompanyCurrency: hubspot.NewString("1500.00"),
					AnnualContractValue:     hubspot.NewString("1000.00"),
					AnnualRecurringRevenue:  nil,
					ClosedLostReason:        hubspot.NewString("lost reason"),
					ClosedWonReason:         hubspot.NewString("won reason"),
					DealDescription:         hubspot.NewString("test hubspot deal record"),
					DealName:                hubspot.NewString("Custom data integrations"),
					DealOwnerID:             hubspot.NewString("910901"),
					DealStage:               hubspot.NewString("presentation scheduled"),
					DealType:                nil,
					ForecastAmount:          hubspot.NewString("1500.00"),
					ForecastCategory:        nil,
					ForecastProbability:     hubspot.NewString("900101"),
					MonthlyRecurringRevenue: hubspot.NewString("100.00"),
					NextStep:                hubspot.NewString("tel"),
					NumberOfContacts:        hubspot.NewString("0"),
					NumberOfSalesActivities: hubspot.NewString("1"),
					NumberOfTimesContacted:  hubspot.NewString("0"),
					ObjectID:                hubspot.NewString("512"),
					PipeLine:                hubspot.NewString("default"),
					TeamID:                  nil,
					TotalContractValue:      hubspot.NewString("0"),
					CreateDate:              &createdAt,
					CloseDate:               &closeDate,
					LastActivityDate:        &updatedAt,
					LastContacted:           nil,
					LastModifiedDate:        &modifyDate,
					NextActivityDate:        &hubspot.HsTime{}, // FIXME: when empty string do not initialize.
					OwnerAssignedDate:       nil,
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Successfully get a deal with custom fields",
			fields: fields{
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"id":"512","properties":{"amount":"1500.00","amount_in_home_currency":"1500.00","closed_lost_reason":"lost reason","closed_won_reason":"won reason","closedate":"2019-12-07T16:50:06.678Z","createdate":"2019-10-30T03:30:17.883Z","dealname":"Custom data integrations","dealstage":"presentation scheduled","dealtype":null,"description":"test hubspot deal record","hs_acv":"1000.00","hs_arr":null,"hs_forecast_amount":"1500.00","hs_forecast_probability":"900101","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","hs_mrr":"100.00","hs_next_step":"tel","hs_object_id":"512","hs_tcv":"0","hubspot_owner_assigneddate":null,"hubspot_owner_id":"910901","hubspot_team_id":null,"notes_last_contacted":null,"notes_last_updated":"2019-12-07T16:50:06.678Z","notes_next_activity_date":"","num_associated_contacts":"0","num_contacted_notes":"0","num_notes":"1","pipeline":"default","custom_name":"hubspot custom","custom_date":"2019-12-07T16:50:06.678Z"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","archived":false}`),
				}),
			},
			args: args{
				dealID: "512",
				deal:   &CustomFields{},
				option: &hubspot.RequestQueryOption{
					CustomProperties: []string{
						"custom_name",
						"custom_date",
					},
				},
			},
			want: &hubspot.ResponseResource{
				ID:       "512",
				Archived: false,
				Properties: &CustomFields{
					Deal: hubspot.Deal{
						Amount:                  hubspot.NewString("1500.00"),
						AmountInCompanyCurrency: hubspot.NewString("1500.00"),
						AnnualContractValue:     hubspot.NewString("1000.00"),
						AnnualRecurringRevenue:  nil,
						ClosedLostReason:        hubspot.NewString("lost reason"),
						ClosedWonReason:         hubspot.NewString("won reason"),
						DealDescription:         hubspot.NewString("test hubspot deal record"),
						DealName:                hubspot.NewString("Custom data integrations"),
						DealOwnerID:             hubspot.NewString("910901"),
						DealStage:               hubspot.NewString("presentation scheduled"),
						DealType:                nil,
						ForecastAmount:          hubspot.NewString("1500.00"),
						ForecastCategory:        nil,
						ForecastProbability:     hubspot.NewString("900101"),
						MonthlyRecurringRevenue: hubspot.NewString("100.00"),
						NextStep:                hubspot.NewString("tel"),
						NumberOfContacts:        hubspot.NewString("0"),
						NumberOfSalesActivities: hubspot.NewString("1"),
						NumberOfTimesContacted:  hubspot.NewString("0"),
						ObjectID:                hubspot.NewString("512"),
						PipeLine:                hubspot.NewString("default"),
						TeamID:                  nil,
						TotalContractValue:      hubspot.NewString("0"),
						CreateDate:              &createdAt,
						CloseDate:               &closeDate,
						LastActivityDate:        &updatedAt,
						LastContacted:           nil,
						LastModifiedDate:        &modifyDate,
						NextActivityDate:        &hubspot.HsTime{}, // FIXME: when empty string do not initialize.
						OwnerAssignedDate:       nil,
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
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				dealID: "deal001",
				deal:   &hubspot.Deal{},
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
			got, err := tt.fields.client.CRM.Deal.Get(tt.args.dealID, tt.args.deal, tt.args.option)
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

func TestDealServiceOp_AssociateAnotherObj(t *testing.T) {
	type fields struct {
		dealPath string
		client   *hubspot.Client
	}
	type args struct {
		dealID string
		conf   *hubspot.AssociationConfig
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
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"id":"512","properties":{"createdate":"2019-10-30T03:30:17.883Z","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","hs_object_id":"512"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","associations":{"contacts":{"results":[{"id":"20074","type":"deal_to_contact"}]}},"archived":false,"archivedAt":"2020-01-01T00:00:00.000Z"}`),
				}),
			},
			args: args{
				dealID: "512",
				conf: &hubspot.AssociationConfig{
					ToObject:   hubspot.ObjectTypeContact,
					ToObjectID: "20074",
					Type:       hubspot.AssociationTypeDealToContact,
				},
			},
			want: &hubspot.ResponseResource{
				ID:       "512",
				Archived: false,
				Associations: &hubspot.Associations{
					Contacts: struct {
						Results []hubspot.AssociationResult `json:"results"`
					}{
						Results: []hubspot.AssociationResult{
							{ID: "20074", Type: "deal_to_contact"},
						},
					},
				},
				Properties: &hubspot.Deal{
					ObjectID:         hubspot.NewString("512"),
					CreateDate:       &createdAt,
					LastModifiedDate: &updatedAt,
				},
				CreatedAt:  &createdAt,
				UpdatedAt:  &updatedAt,
				ArchivedAt: &archivedAt,
			},
			wantErr: nil,
		},
		{
			name: "Received unable association type error",
			fields: fields{
				dealPath: hubspot.ExportDealBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"status":"error","message":"test is not a valid association type between deals and contacts","correlationId":"correlation_id","context":{"type":["test"],"fromObjectType":["deals"],"toObjectType":["contacts"]},"category":"VALIDATION_ERROR","subCategory":"crm.associations.INVALID_ASSOCIATION_TYPE"}`),
				}),
			},
			args: args{
				dealID: "512",
				conf: &hubspot.AssociationConfig{
					ToObject:   hubspot.ObjectTypeContact,
					ToObjectID: "20074",
					Type:       "test",
				},
			},
			want: nil,
			wantErr: &hubspot.APIError{
				HTTPStatusCode: http.StatusBadRequest,
				Status:         "error",
				Message:        "test is not a valid association type between deals and contacts",
				CorrelationID:  "correlation_id",
				Context: hubspot.ErrContext{
					Type:           []string{"test"},
					FromObjectType: []string{string(hubspot.ObjectTypeDeal)},
					ToObjectType:   []string{string(hubspot.ObjectTypeContact)},
				},
				Category:    "VALIDATION_ERROR",
				SubCategory: "crm.associations.INVALID_ASSOCIATION_TYPE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.CRM.Deal.AssociateAnotherObj(tt.args.dealID, tt.args.conf)
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

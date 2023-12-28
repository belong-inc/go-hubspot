package hubspot_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/belong-inc/go-hubspot"
	"github.com/google/go-cmp/cmp"
)

func TestCompanyServiceOp_Create(t *testing.T) {
	company := &hubspot.Company{
		City:     hubspot.NewString("Cambridge"),
		Domain:   hubspot.NewString("biglytics.net"),
		Industry: hubspot.NewString("Technology"),
		Name:     hubspot.NewString("Biglytics"),
		Phone:    hubspot.NewString("(877)929-0687"),
		State:    hubspot.NewString("Massachusetts"),
	}

	type fields struct {
		companyPath string
		client      *hubspot.Client
	}
	type args struct {
		company interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully create a company",
			fields: fields{
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusCreated,
					Header: http.Header{},
					Body:   []byte(`{"id":"company001","properties":{"city":"Cambridge","createdate":"2019-10-30T03:30:17.883Z","domain":"biglytics.net","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","industry":"Technology","name":"Biglytics","phone":"(877)929-0687","state":"Massachusetts"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","archived":false}`),
				}),
			},
			args: args{
				company: company,
			},
			want: &hubspot.ResponseResource{
				ID:       "company001",
				Archived: false,
				Properties: &hubspot.Company{
					City:               hubspot.NewString("Cambridge"),
					Createdate:         &createdAt,
					Domain:             hubspot.NewString("biglytics.net"),
					HsLastmodifieddate: &updatedAt,
					Industry:           hubspot.NewString("Technology"),
					Name:               hubspot.NewString("Biglytics"),
					Phone:              hubspot.NewString("(877)929-0687"),
					State:              hubspot.NewString("Massachusetts"),
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Received invalid request",
			fields: fields{
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				company: company,
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
			got, err := tt.fields.client.CRM.Company.Create(tt.args.company)
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

func TestCompanyServiceOp_Update(t *testing.T) {
	company := &hubspot.Company{
		City:     hubspot.NewString("Cambridge"),
		Domain:   hubspot.NewString("biglytics.net"),
		Industry: hubspot.NewString("Technology"),
		Name:     hubspot.NewString("Biglytics"),
		Phone:    hubspot.NewString("(877)929-0687"),
		State:    hubspot.NewString("Massachusetts"),
	}

	type fields struct {
		companyPath string
		client      *hubspot.Client
	}
	type args struct {
		companyID string
		company   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantErr error
	}{
		{
			name: "Successfully update a company",
			fields: fields{
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusCreated,
					Header: http.Header{},
					Body:   []byte(`{"id":"company001","properties":{"city":"Cambridge","createdate":"2019-10-30T03:30:17.883Z","domain":"biglytics.net","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","industry":"Technology","name":"Biglytics","phone":"(877)929-0687","state":"Massachusetts"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","archived":false}`),
				}),
			},
			args: args{
				company: company,
			},
			want: &hubspot.ResponseResource{
				ID:       "company001",
				Archived: false,
				Properties: &hubspot.Company{
					City:               hubspot.NewString("Cambridge"),
					Createdate:         &createdAt,
					Domain:             hubspot.NewString("biglytics.net"),
					HsLastmodifieddate: &updatedAt,
					Industry:           hubspot.NewString("Technology"),
					Name:               hubspot.NewString("Biglytics"),
					Phone:              hubspot.NewString("(877)929-0687"),
					State:              hubspot.NewString("Massachusetts"),
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Received invalid request",
			fields: fields{
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				companyID: "company001",
				company:   company,
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
			got, err := tt.fields.client.CRM.Company.Update(tt.args.companyID, tt.args.company)
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

func TestCompanyServiceOp_Get(t *testing.T) {
	type CustomFields struct {
		hubspot.Company
		CustomName *hubspot.HsStr  `json:"custom_name,omitempty"`
		CustomDate *hubspot.HsTime `json:"custom_date,omitempty"`
	}

	type fields struct {
		companyPath string
		client      *hubspot.Client
	}
	type args struct {
		companyID string
		company   interface{}
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
			name: "Successfully get a company",
			fields: fields{
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"id":"company001","properties":{"annualrevenue":"1000000","city":"Cambridge","createdate":"2019-10-30T03:30:17.883Z","domain":"biglytics.net","hs_lastmodifieddate":"2019-12-07T16:50:06.678Z","industry":"Technology","name":"Biglytics","phone":"(877)929-0687","state":"Massachusetts","custom_name":"biglytics","custom_date":"2019-10-30T03:30:17.883Z","hs_created_by_user_id":"","twitterfollowers":1000},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z"}`),
				}),
			},
			args: args{
				companyID: "company001",
				company:   &CustomFields{},
				option: &hubspot.RequestQueryOption{
					Properties: []string{
						"custom_name",
						"custom_date",
					},
				},
			},
			want: &hubspot.ResponseResource{
				ID:       "company001",
				Archived: false,
				Properties: &CustomFields{
					Company: hubspot.Company{
						Annualrevenue:      hubspot.NewInt(1000000),
						City:               hubspot.NewString("Cambridge"),
						Createdate:         &createdAt,
						Domain:             hubspot.NewString("biglytics.net"),
						HsLastmodifieddate: &updatedAt,
						Industry:           hubspot.NewString("Technology"),
						Name:               hubspot.NewString("Biglytics"),
						Phone:              hubspot.NewString("(877)929-0687"),
						State:              hubspot.NewString("Massachusetts"),
						HsCreatedByUserId:  hubspot.NewInt(0),
						Twitterfollowers:   hubspot.NewInt(1000),
					},
					CustomName: hubspot.NewString("biglytics"),
					CustomDate: &createdAt,
				},
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			wantErr: nil,
		},
		{
			name: "Received invalid request",
			fields: fields{
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				companyID: "company001",
				company:   &hubspot.Company{},
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
			got, err := tt.fields.client.CRM.Company.Get(tt.args.companyID, tt.args.company, tt.args.option)
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

func TestCompanyServiceOp_Delete(t *testing.T) {
	type fields struct {
		companyPath string
		client      *hubspot.Client
	}
	type args struct {
		companyID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Successfully delete a company",
			fields: fields{
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusNoContent,
					Header: http.Header{},
				}),
			},
			args: args{
				companyID: "company001",
			},
			wantErr: nil,
		},
		{
			name: "Received invalid request",
			fields: fields{
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				companyID: "company001",
			},
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
			err := tt.fields.client.CRM.Company.Delete(tt.args.companyID)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("Delete() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
		})
	}
}

func TestCompanyServiceOp_AssociateAnotherObj(t *testing.T) {
	type fields struct {
		companyPath string
		client      *hubspot.Client
	}
	type args struct {
		companyID string
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
				companyPath: hubspot.ExportCompanyBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"id":"company001","properties":{"createdate":"2019-10-30T03:30:17.883Z","hs_object_id":"company001","lastmodifieddate":"2019-12-07T16:50:06.678Z"},"createdAt":"2019-10-30T03:30:17.883Z","updatedAt":"2019-12-07T16:50:06.678Z","associations":{"deals":{"results":[{"id":"deal001","type":"company_to_deal"}]}},"archived":false,"archivedAt":"2020-01-01T00:00:00.000Z"}`),
				}),
			},
			args: args{
				companyID: "company001",
				conf: &hubspot.AssociationConfig{
					ToObject:   hubspot.ObjectTypeDeal,
					ToObjectID: "deal001",
					Type:       hubspot.AssociationTypeCompanyToDeal,
				},
			},
			want: &hubspot.ResponseResource{
				ID:       "company001",
				Archived: false,
				Associations: &hubspot.Associations{
					Deals: struct {
						Results []hubspot.AssociationResult `json:"results"`
					}{
						Results: []hubspot.AssociationResult{
							{ID: "deal001", Type: string(hubspot.AssociationTypeCompanyToDeal)},
						},
					},
				},
				Properties: &hubspot.Contact{
					HsObjectID:       hubspot.NewString("company001"),
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
			got, err := tt.fields.client.CRM.Contact.AssociateAnotherObj(tt.args.companyID, tt.args.conf)
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

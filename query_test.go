package hubspot_test

import (
	"testing"

	"github.com/belong-inc/go-hubspot"
	"github.com/google/go-cmp/cmp"
)

func TestRequestQueryOption_setupProperties(t *testing.T) {
	type fields struct {
		Properties           []string
		CustomProperties     []string
		Associations         []string
		PaginateAssociations bool
		Archived             bool
		IDProperty           string
	}
	type args struct {
		defaultFields []string
	}
	tests := []struct {
		name   string
		fields *fields
		args   args
		want   *hubspot.RequestQueryOption
	}{
		{
			name: "Success with all fields",
			fields: &fields{
				CustomProperties:     []string{"tel", "occupation"},
				Associations:         []string{"deals"},
				PaginateAssociations: true,
				Archived:             true,
				IDProperty:           "id",
			},
			args: args{
				defaultFields: []string{"id", "name", "age"},
			},
			want: &hubspot.RequestQueryOption{
				Properties: []string{
					"id",
					"name",
					"age",
					"tel",
					"occupation",
				},
				CustomProperties:     []string{"tel", "occupation"},
				Associations:         []string{"deals"},
				PaginateAssociations: true,
				Archived:             true,
				IDProperty:           "id",
			},
		},
		{
			name: "Success with custom properties field only",
			fields: &fields{
				CustomProperties: []string{"tel", "occupation"},
			},
			args: args{
				defaultFields: []string{"id", "name", "age"},
			},
			want: &hubspot.RequestQueryOption{
				Properties: []string{
					"id",
					"name",
					"age",
					"tel",
					"occupation",
				},
				CustomProperties:     []string{"tel", "occupation"},
				Associations:         nil,
				PaginateAssociations: false,
				Archived:             false,
				IDProperty:           "",
			},
		},
		{
			name:   "Success option is nil",
			fields: nil,
			args: args{
				defaultFields: []string{"id", "name", "age"},
			},
			want: &hubspot.RequestQueryOption{
				Properties: []string{
					"id",
					"name",
					"age",
				},
				CustomProperties:     nil,
				Associations:         nil,
				PaginateAssociations: false,
				Archived:             false,
				IDProperty:           "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var o *hubspot.RequestQueryOption = nil
			if tt.fields != nil {
				o = &hubspot.RequestQueryOption{
					Properties:           tt.fields.Properties,
					CustomProperties:     tt.fields.CustomProperties,
					Associations:         tt.fields.Associations,
					PaginateAssociations: tt.fields.PaginateAssociations,
					Archived:             tt.fields.Archived,
					IDProperty:           tt.fields.IDProperty,
				}
			}
			got := hubspot.ExportSetupProperties(o, tt.args.defaultFields)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("setupProperties() response mismatch (-want +got):%s", diff)
			}
		})
	}
}

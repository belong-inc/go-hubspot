package hubspot

// RequestQueryOption is a set of options to be specified in the query when making a Get request.
// RequestQueryOption.Properties will be overwritten internally, so do not specify it.
// If you want to get the custom fields as well, specify the field names in RequestQueryOption.CustomProperties.
// Items with no value set will be ignored.
type RequestQueryOption struct {
	Properties           []string `url:"properties,comma,omitempty"`
	CustomProperties     []string `url:"-"`
	Associations         []string `url:"associations,comma,omitempty"`
	PaginateAssociations bool     `url:"paginateAssociations,omitempty"` // HubSpot defaults false
	Archived             bool     `url:"archived,omitempty"`             // HubSpot defaults false
	IDProperty           string   `url:"idProperty,omitempty"`
}

// setupProperties sets the property to get.
// RequestQueryOption.Properties will be overwritten.
// If RequestQueryOption is nil, only the default properties will be set.
func (o *RequestQueryOption) setupProperties(defaultFields []string) *RequestQueryOption {
	opts := RequestQueryOption{}
	if o != nil {
		opts = *o
	}
	opts.Properties = append(defaultFields, opts.CustomProperties...)
	return &opts
}

type BulkRequestQueryOption struct {
	// Properties sets a comma separated list of the properties to be returned in the response.
	Properties []string `url:"properties,comma,omitempty"`
	// Limit is the maximum number of results to display per page.
	Limit int `url:"limit,comma,omitempty"`
	// After is the paging cursor token of the last successfully read resource will be returned as the paging.next.after.
	After string `url:"after,omitempty"`

	// Offset is used to get the next page of results.
	// Available only in API v1.
	Offset string `url:"offset,omitempty"`
	// orderBy is used to order by a particular field value.
	// Use a negative value to sort in descending order.
	// Available only in API v1.
	OrderBy string `url:"orderBy,omitempty"`
}

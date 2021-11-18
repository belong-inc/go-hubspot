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

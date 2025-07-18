package hubspot

// Operator is search for HubSpot API filters.
// It defines the operators that can be used in search filters for various HubSpot objects such as contacts, companies, deals, and tickets.
// https://developers.hubspot.com/docs/guides/api/crm/search#filter-search-results
type Operator string

const (
	LT               Operator = "LT"                 // Less than the specified value.
	LTE              Operator = "LTE"                // Less than or equal to the specified value.
	GT               Operator = "GT"                 // Greater than the specified value.
	GTE              Operator = "GTE"                // Greater than or equal to the specified value.
	EQ               Operator = "EQ"                 // Equal to the specified value.
	NEQ              Operator = "NEQ"                // Not equal to the specified value.
	Between          Operator = "BETWEEN"            // Within the specified range. In your request, use key-value pairs to set highValue and value. Refer to the example below the table.
	IN               Operator = "IN"                 // Included within the specified list. Searches by exact match. In your request, include the list values in a values array. When searching a string property with this operator, values must be lowercase. Refer to the example below the table.
	NotIN            Operator = "NOT_IN"             // Not included within the specified list. In your request, include the list values in a values array. When searching a string property with this operator, values must be lowercase.
	HasProperty      Operator = "HAS_PROPERTY"       // Has a value for the specified property.
	NotHasProperty   Operator = "NOT_HAS_PROPERTY"   // Doesn't have a value for the specified property.
	ContainsToken    Operator = "CONTAINS_TOKEN"     // Contains a token. In your request, you can use wildcards (*) to complete a partial search. For example, use the value *@hubspot.com to retrieve contacts with a HubSpot email address.
	NotContainsToken Operator = "NOT_CONTAINS_TOKEN" // Doesn't contain a token.
)

// SearchOptions represents the options for searching HubSpot objects.
type SearchOptions struct {
	FilterGroups []FilterGroup `json:"filterGroups,omitempty"` // A list of filter groups, each containing one or more filters.
	Sorts        []Sort        `json:"sorts,omitempty"`
	Query        string        `json:"query,omitempty"`
	Properties   []string      `json:"properties,omitempty"`
	Limit        int           `json:"limit,omitempty"` // The maximum number of entries per page is 200, the default is 10.
	After        int           `json:"after,omitempty"` // The offset for pagination, used to retrieve the next page of results.
}

// FilterGroup represents a group of filters for HubSpot search requests.
type FilterGroup struct {
	Filters []Filter `json:"filters,omitempty"`
}

// Filter represents a single filter for HubSpot search requests.
type Filter struct {
	PropertyName string   `json:"propertyName"`
	Operator     Operator `json:"operator"`
	Values       []HsStr  `json:"values,omitempty"`
	Value        *HsStr   `json:"value,omitempty"`
	HighValue    *HsStr   `json:"highValue,omitempty"`
}

type SortDirection string

const (
	Asc  SortDirection = "ASCENDING"
	Desc SortDirection = "DESCENDING"
)

type Sort struct {
	PropertyName string        `json:"propertyName"`
	Direction    SortDirection `json:"direction"`
}

package legacy

type BulkResponseResource struct {
	Limit      int         `json:"limit,omitempty"`
	Objects    interface{} `json:"objects,omitempty"`
	Offset     int         `json:"offset,omitempty"`
	Total      int         `json:"total,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
}

package hubspot

import "fmt"

const (
	// ValidationError is the APIError.Category.
	// This is returned by HubSpot when the HTTP Status is 400.
	// In this case, the verification details error will be included in Details
	ValidationError = "VALIDATION_ERROR"

	// InvalidEmailError is the value of ErrDetail.Error when an error occurs in the Email validation.
	InvalidEmailError = "INVALID_EMAIL"
	// UnknownDetailError is the value set by go-hubspot when extraction the error details failed.
	UnknownDetailError = "UNKNOWN_DETAIL"
)

type APIError struct {
	HTTPStatusCode int         `json:"-"`
	Status         string      `json:"status,omitempty"`
	Message        string      `json:"message,omitempty"`
	CorrelationID  string      `json:"correlationId,omitempty"`
	Context        ErrContext  `json:"context,omitempty"`
	Category       string      `json:"category,omitempty"`
	SubCategory    string      `json:"subCategory,omitempty"`
	Links          ErrLinks    `json:"links,omitempty"`
	Details        []ErrDetail `json:"details,omitempty"`
}

type ErrDetail struct {
	IsValid bool   `json:"isValid,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Name    string `json:"name,omitempty"`
}

type ErrContext struct {
	ID             []string `json:"id,omitempty"`
	Type           []string `json:"type,omitempty"`
	ObjectType     []string `json:"objectType,omitempty"`
	FromObjectType []string `json:"fromObjectType,omitempty"`
	ToObjectType   []string `json:"toObjectType,omitempty"`
}

type ErrLinks struct {
	APIKey        string `json:"api key,omitempty"`
	KnowledgeBase string `json:"knowledge-base,omitempty"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%d: %s", e.HTTPStatusCode, e.Message)
}

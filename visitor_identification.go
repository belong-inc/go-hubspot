package hubspot

const (
	visitorIdentificationBasePath = "/conversations/v3/visitor-identification"
)

// VisitorIdentificationServiceOp handles communication with the Visitor Identification related methods of the HubSpot API.
type VisitorIdentificationServiceOp struct {
	client *Client
}

type IdentificationTokenResponse struct {
	Token string `json:"token"`
}

type IdentificationTokenRequest struct {
	FirstName string
	LastName  string
	Email     string
}

// VisitorIdentificationService is an interface of visitor identification endpoints of the HubSpot API.
type VisitorIdentificationService interface {
	GenerateIdentificationToken(request IdentificationTokenRequest) (*IdentificationTokenResponse, error)
}

// Create creates a Identification Token.
// In order to bind the created content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.request in your own structure.
func (s *VisitorIdentificationServiceOp) GenerateIdentificationToken(request IdentificationTokenRequest) (*IdentificationTokenResponse, error) {
	response := &IdentificationTokenResponse{}
	path := visitorIdentificationBasePath + "/tokens/create"
	if err := s.client.Post(path, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

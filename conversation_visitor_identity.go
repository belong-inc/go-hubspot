package hubspot

const (
	visitorIdentificationBasePath = "/conversations/v3/visitor-identification"
)

type IdentificationTokenResponse struct {
	Token string `json:"token"`
}

type IdentificationTokenRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type VisitorIdentificationService interface {
	GenerateIdentificationToken(option IdentificationTokenRequest) (*IdentificationTokenResponse, error)
}

type VisitorIdentificationServiceOp struct {
	client *Client
}

var _ VisitorIdentificationService = (*VisitorIdentificationServiceOp)(nil)

func (s *VisitorIdentificationServiceOp) GenerateIdentificationToken(option IdentificationTokenRequest) (*IdentificationTokenResponse, error) {
	response := &IdentificationTokenResponse{}
	path := visitorIdentificationBasePath + "/tokens/create"
	if err := s.client.Post(path, option, response); err != nil {
		return nil, err
	}
	return response, nil
}

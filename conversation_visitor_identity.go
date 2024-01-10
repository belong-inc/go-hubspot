package hubspot

type IdentificationTokenResponse struct {
	Token string `json:"token"`
}

type IdentificationTokenRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type VisitorIdentificationService interface {
	GenerateIdentificationToken(option IdentificationTokenRequest) (*IdentificationTokenResponse, error)
}

type VisitorIdentificationServiceOp struct {
	client   *Client
	basePath string
}

var _ VisitorIdentificationService = (*VisitorIdentificationServiceOp)(nil)

func (s *VisitorIdentificationServiceOp) GenerateIdentificationToken(option IdentificationTokenRequest) (*IdentificationTokenResponse, error) {
	response := &IdentificationTokenResponse{}
	path := s.basePath + "/tokens/create"
	if err := s.client.Post(path, option, response); err != nil {
		return nil, err
	}
	return response, nil
}

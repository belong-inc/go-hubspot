package hubspot

type IDentificationTokenResponse struct {
	Token string `json:"token"`
}

type IDentificationTokenRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type VisitorIDentificationService interface {
	GenerateIDentificationToken(option IDentificationTokenRequest) (*IDentificationTokenResponse, error)
}

type VisitorIDentificationServiceOp struct {
	client   *Client
	basePath string
}

var _ VisitorIDentificationService = (*VisitorIDentificationServiceOp)(nil)

func (s *VisitorIDentificationServiceOp) GenerateIDentificationToken(option IDentificationTokenRequest) (*IDentificationTokenResponse, error) {
	response := &IDentificationTokenResponse{}
	path := s.basePath + "/tokens/create"
	if err := s.client.Post(path, option, response); err != nil {
		return nil, err
	}
	return response, nil
}

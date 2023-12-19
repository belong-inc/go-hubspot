package hubspot

import (
	"os"
	"testing"
)

type IdentificationTokenRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func TestGenerateIdentificationToken(t *testing.T) {
	//t.SkipNow() // Remove this line when you're ready to run the test

	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	request := &IdentificationTokenRequest{
		Email:     "test@example.com",
		FirstName: "Test",
	}

	response, err := cli.VisitorIdentification.GenerateIdentificationToken(request)
	if err != nil {
		t.Error(err)
	}

	if response.Token == "" {
		t.Errorf("expected response.Token to be not empty")
	}
}

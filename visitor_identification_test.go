package hubspot

import (
	"os"
	"testing"
)

func TestGenerateIdentificationToken(t *testing.T) {
	t.SkipNow()

	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	request := IdentificationTokenRequest{
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@example.com",
	}

	response, err := cli.VisitorIdentification.GenerateIdentificationToken(request)
	if err != nil {
		t.Error(err)
	}

	if response.Token == "" {
		t.Errorf("expected response.Token to be not empty")
	}
}

package hubspot

const (
	visitorIdentificationBasePath = "/conversations/v3/visitor-identification"
)

type Conversation struct {
	VisitorIdentification VisitorIdentificationService
}

func newConversation(c *Client) *Conversation {
	return &Conversation{
		VisitorIdentification: &VisitorIdentificationServiceOp{
			client:   c,
			basePath: visitorIdentificationBasePath,
		},
	}
}

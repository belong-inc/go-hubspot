package hubspot

const (
	visitorIDentificationBasePath = "/conversations/v3/visitor-identification"
)

type Conversation struct {
	VisitorIDentification VisitorIDentificationService
}

func newConversation(c *Client) *Conversation {
	return &Conversation{
		VisitorIDentification: &VisitorIDentificationServiceOp{
			client:   c,
			basePath: visitorIDentificationBasePath,
		},
	}
}

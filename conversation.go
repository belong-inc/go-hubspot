package hubspot

type Conversation struct {
	VisitorIdentification      VisitorIdentificationService
	IdentificationTokenRequest IdentificationTokenRequest
}

func newConversation(c *Client) *Conversation {
	return &Conversation{
		VisitorIdentification: &VisitorIdentificationServiceOp{
			client: c,
		},
	}
}

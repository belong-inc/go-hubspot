package hubspot

const (
	marketingBasePath = "marketing"
)

type Marketing struct {
	Email MarketingEmailService
}

func newMarketing(c *Client) *Marketing {
	return &Marketing{Email: NewMarketingEmail(c)}
}

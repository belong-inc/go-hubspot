package hubspot

import "fmt"

const (
	marketingBasePath = "marketing"
)

type Marketing struct {
	Email         MarketingEmailService
	Transactional TransactionalServiceOp
}

// "/marketing/v3/transactional/single-email/send"
func newMarketing(c *Client) *Marketing {
	return &Marketing{
		Email: NewMarketingEmail(c),
		Transactional: TransactionalServiceOp{
			client:            c,
			transactionalPath: fmt.Sprintf("%s/%s/%s", marketingBasePath, c.apiVersion, transactionalBasePath),
		},
	}
}

package hubspot

import "fmt"

const transactionalBasePath = "transactional"

// DealService is an interface of deal endpoints of the HubSpot API.
// HubSpot deal can be used to manage transactions.
// It can also be associated with other CRM objects such as contact and company.
// Reference: https://developers.hubspot.com/docs/api/crm/deals
type TransactionalService interface {
	SendSingleEmail(props *SendSingleEmailProperties) (*SendSingleEmailResponse, error)
}

// DealServiceOp handles communication with the product related methods of the HubSpot API.
type TransactionalServiceOp struct {
	transactionalPath string
	client            *Client
}

var _ TransactionalService = (*TransactionalServiceOp)(nil)

type SendSingleEmailMessage struct {
	To      string   `json:"to"`
	From    string   `json:"from,omitempty"`
	SendId  string   `json:"sendId,omitempty"`
	ReplyTo []string `json:"replyTo,omitempty"`
	Cc      string   `json:"cc,omitempty"`
	Bcc     string   `json:"bcc,omitempty"`
}

type SendSingleEmailProperties struct {
	EmailId           string                  `json:"emailId"`
	Message           *SendSingleEmailMessage `json:"message"`
	ContactProperties *Contact                `json:"contactProperties,omitempty"`
	CustomProperties  interface{}             `json:"customProperties,omitempty"`
}

type SendSingleEmailResponse struct {
	RequestedAt string `json:"requestedAt"`
	StatusId    string `json:"statusId"`
	Status      string `json:"status"`
}

func (s *TransactionalServiceOp) SendSingleEmail(props *SendSingleEmailProperties) (*SendSingleEmailResponse, error) {
	resource := &SendSingleEmailResponse{}
	if err := s.client.Post(fmt.Sprintf("%s/single-email/send", s.transactionalPath), props, resource); err != nil {
		return nil, err
	}
	return resource, nil

}

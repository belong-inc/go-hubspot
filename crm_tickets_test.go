package hubspot

import (
	"fmt"
	"os"
	"testing"
)

func TestListTickets(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	opt := &RequestQueryOption{}
	opt.Properties = []string{"Content"}
	res, err := cli.CRM.Tickets.List(opt)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", res.Results[0])
}

func TestGetCrmTicket(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	opt := &RequestQueryOption{}
	opt.Properties = []string{"Content", "associated_contact_lifecycle_stage", "hubspot_owner_id"}
	res, err := cli.CRM.Tickets.Get("1594949554", opt)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestCreateCrmTicket(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	props := make(map[string]interface{})
	props["hs_pipeline"] = "30440034"
	props["hs_pipeline_stage"] = "69304142"
	props["hubspot_owner_id"] = "301296186"
	props["hs_ticket_priority"] = "LOW"
	props["content"] = "this would be some content"
	props["subject"] = "testing, please ignore"
	req := &CrmTicketCreateRequest{
		Properties: props,
	}
	res, err := cli.CRM.Tickets.Create(req)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestDeleteCrmTicket(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	err := cli.CRM.Tickets.Archive("1594967688")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateCrmTicket(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	props := make(map[string]interface{})
	props["hs_ticket_priority"] = "HIGH"
	req := &CrmTicketCreateRequest{
		Properties: props,
	}
	res, err := cli.CRM.Tickets.Update("1594957134", req)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestSearchCrmTicket(t *testing.T) {
	t.SkipNow()

	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))

	req := &CrmTicketSearchRequest{
		FilterGroups: []*CrmTicketSearchFilterGroup{
			{
				Filters: []*CrmTicketSearchFilter{
					{
						Value:        NewString("LOW"),
						PropertyName: NewString("hs_ticket_priority"),
						Operator:     NewString("EQ"),
					},
				},
			},
		},
	}

	res, err := cli.CRM.Tickets.Search(req)
	if err != nil {
		t.Error(err)
	}

	for _, ticket := range res.Results {
		fmt.Printf("%+v\n", ticket)
	}
}

package hubspot

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestListCrmProperties(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	// Use crm_schemas:TestCreate() to generate this...
	res, err := cli.CRM.Properties.List("cars")
	if err != nil {
		t.Error(err)
	}

	if len(res.Results) < 1 {
		t.Error("expected len(res.Results) to be > 1")
	}
}

func TestGetCrmProperty(t *testing.T) {
	t.SkipNow()

	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	// Use crm_schemas:TestCreate() to generate this...
	res, err := cli.CRM.Properties.Get("cars", "model")
	if err != nil {
		t.Error(err)
	}
	if *res.Name != "model" {
		t.Errorf("expected res.Name to be model, got %s", res.Name)
	}
}

func TestCreateProperty(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	newProp := &CrmProperty{
		Name:      NewString("mileage"),
		Label:     NewString("Mileage Label"),
		Type:      NewString("number"),
		FieldType: NewString("number"),
		GroupName: NewString("cars_information"),
	}

	_, err := cli.CRM.Properties.Create("cars", newProp)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateProperty(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))

	updateProp := make(map[string]interface{})
	updateProp["label"] = fmt.Sprintf("Updated Label %s", time.Now().String())

	res, err := cli.CRM.Properties.Update("cars", "mileage", &updateProp)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Label != updateProp["label"] {
		t.Errorf("expected res.Label to be %s, got %s", updateProp["label"], res.Label)
	}
}

func TestDeleteProperty(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	err := cli.CRM.Properties.Delete("cars", "mileage")
	if err != nil {
		t.Error(err)
	}
}

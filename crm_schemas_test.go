package hubspot

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetSchemas(t *testing.T) {
	t.SkipNow()

	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	res, err := cli.CRM.Schemas.List()
	if err != nil {
		t.Error(err)
	}
	if len(res.Results) < 1 {
		t.Errorf("expected results to have some results")
	}
}

func TestCreateSchema(t *testing.T) {
	t.SkipNow()
	// this example is from Hubspot..
	exampleJSON := `
	{
	  "name": "cars",
	  "labels": {
		"singular": "Car",
		"plural": "Cars"
	  },
	  "primaryDisplayProperty": "model",
	  "secondaryDisplayProperties": [
		 "make"
	],
	  "searchableProperties": [
		 "year",
		 "make",
		 "vin",
		 "model"
	],
	  "requiredProperties": [
		 "year",
		 "make",
		 "vin",
		 "model"
	  ],
	  "properties": [
		{
		  "name": "condition",
		  "label": "Condition",
		  "type": "enumeration",
		  "fieldType": "select",
		  "options": [
			{
			  "label": "New",
			  "value": "new"
			},
			{
			  "label": "Used",
			  "value": "used"
			}
		  ]
		},
		{
		  "name": "date_received",
		  "label": "Date received",
		  "type": "date",
		  "fieldType": "date"
		},
		{
		  "name": "year",
		  "label": "Year",
		  "type": "number",
		  "fieldType": "number"
		},
		{
		  "name": "make",
		  "label": "Make",
		  "type": "string",
		  "fieldType": "text"
		},
		{
		  "name": "model",
		  "label": "Model",
		  "type": "string",
		  "fieldType": "text"
		},
		{
		  "name": "vin",
		  "label": "VIN",
		  "type": "string",
		  "hasUniqueValue": true,
		  "fieldType": "text"
		},
		{
		  "name": "price",
		  "label": "Price",
		  "type": "number",
		  "fieldType": "number"
		},
		{
		  "name": "notes",
		  "label": "Notes",
		  "type": "string",
		  "fieldType": "text"
		}
	  ],
	  "associatedObjects": [
		"CONTACT"
	  ]
	}
	`

	req := make(map[string]interface{})
	err := json.Unmarshal([]byte(exampleJSON), &req)
	if err != nil {
		t.Error(err)
	}

	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	res, err := cli.CRM.Schemas.Create(req)
	if err != nil {
		t.Error(err)
	}

	if *res.Name != "cars" {
		t.Errorf("expected post schema result to have an id field")
	}
}

func TestGetSchema(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	res, err := cli.CRM.Schemas.Get("cars")
	if err != nil {
		t.Error(err)
	}
	if *res.Name != "cars" {
		t.Errorf("expected post schema result to have an id field")
	}
}

func TestDeleteSchema(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))

	res, err := cli.CRM.Schemas.Get("cars")
	if err != nil {
		t.Error(err)
	}

	err = cli.CRM.Schemas.Delete(string(*res.FullyQualifiedName), &RequestQueryOption{Archived: false})
	if err != nil {
		t.Error(err)
	}

	err = cli.CRM.Schemas.Delete(string(*res.FullyQualifiedName), &RequestQueryOption{Archived: true})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateSchema(t *testing.T) {
	t.SkipNow()
	// Note: You need to wait some time after calling create before calling Update as it'll return an error message.
	req := make(map[string]interface{})
	req["primaryDisplayProperty"] = "year"

	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))

	res, err := cli.CRM.Schemas.Get("cars")
	if err != nil {
		t.Error(err)
		return
	}

	if *res.PrimaryDisplayProperty != "model" {
		t.Error("expected primaryDisplayProperty to be model before update")
		return
	}

	res, err = cli.CRM.Schemas.Update(string(*res.FullyQualifiedName), req)
	if err != nil {
		t.Error(err)
		return
	}

	if *res.PrimaryDisplayProperty != "year" {
		t.Errorf("expected primaryDisplayProperty to be year after update, got %s", res.PrimaryDisplayProperty)
	}
}

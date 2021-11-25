# go-hubspot
[![godoc](https://godoc.org/github.com/belong-inc/go-hubspot?status.svg)](https://pkg.go.dev/github.com/belong-inc/go-hubspot)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

HubSpot Go Library that works with [HubSpot API v3](https://developers.hubspot.com/docs/api/overview).  
HubSpot officially supports client library of Node.js, PHP, Ruby, and Python but not Go.

Note: go-hubspot currently doesn't cover all the APIs but mainly implemented CRM APIs. Implemented APIs are used in
production.

# Install

```shell
$ go get github.com/belong-inc/go-hubspot
```

# Usage

## Authentication

### API key

You should take api key in advance. Follow steps
in [here](https://knowledge.hubspot.com/integrations/how-do-i-get-my-hubspot-api-key).

```go
// Initialize hubspot client with apikey
client, _ := hubspot.NewClient(hubspot.SetAPIKey("YOUR_API_KEY"))
```

### OAuth

You should take refresh token in advance. Follow steps
in [here](https://developers.hubspot.com/docs/api/working-with-oauth).

```go
// Initialize hubspot client with OAuth refresh token.
client, _ := hubspot.NewClient(hubspot.SetOAuth(&hubspot.OAuthConfig{
    GrantType:    hubspot.GrantTypeRefreshToken,
    ClientID:     "YOUR_CLIENT_ID",
    ClientSecret: "YOUR_CLIENT_SECRET",
    RefreshToken: "YOUR_REFRESH_TOKEN",
}))
```

## API call

### Get contact

```go
// Initialize hubspot client with auth method.
client, _ := hubspot.NewClient(hubspot.SetAPIKey("YOUR_API_KEY"))

// Get a Contact object whose id is `contactID`.
// Contact instance needs to be provided as to bind response value.
res, _ := client.CRM.Contact.Get("contactID", &hubspot.Contact{}, nil)

// Type assertion to convert `interface` to `hubspot.Contact`.
contact, ok := res.Properties.(*hubspot.Contact)
if !ok {
    return errors.New("unable to type assertion")
}

// Use contact fields.
fmt.Println(contact.FirstName, contact.LastName)
```

---

### Create contact

```go
// Initialize hubspot client with auth method.
client, _ := hubspot.NewClient(hubspot.SetAPIKey("YOUR_API_KEY"))

// Create request payload.
req := &hubspot.Contact{
    Email:       hubspot.NewString("hubspot@example.com"),
    FirstName:   hubspot.NewString("Bryan"),
    LastName:    hubspot.NewString("Cooper"),
    MobilePhone: hubspot.NewString("(877) 929-0687"),
    Website:     hubspot.NewString("example.com"),
    Zip:         nil,
}

// Call create contact api.
res, _ := client.CRM.Contact.Create(req)

// Type assertion to convert `interface` to `hubspot.Contact`.
contact, ok := res.Properties.(*hubspot.Contact)
if !ok {
    return errors.New("unable to assert type")
}

// Use contact fields.
fmt.Println(contact.FirstName, contact.LastName)
```

---

### Associate objects

```go
// Initialize hubspot client with auth method.
client, _ := hubspot.NewClient(hubspot.SetAPIKey("YOUR_API_KEY"))

// Call associate api.
client.CRM.Contact.AssociateAnotherObj("contact001", &hubspot.AssociationConfig{
    ToObject:   hubspot.ObjectTypeDeal,
    ToObjectID: "dealID",
    Type:       hubspot.AssociationTypeContactToDeal,
})
```

## API call using custom fields

Custom fields are added out of existing object such as Deal or Contact.  
Therefor a new struct needs to be created which contain default fields and additional custom field, and set to Properties field of a request.

### Get deal with custom fields.

```go
type CustomDeal struct {
	hubspot.Deal // embed default fields.
	CustomA string `json:"custom_a,omitempty"`
	CustomB string `json:"custom_b,omitempty"`
}

// Initialize hubspot client with auth method.
client, _ := hubspot.NewClient(hubspot.SetAPIKey("YOUR_API_KEY"))

// Get a Deal object whose id is `dealID`.
// CustomDeal instance needs to be provided as to bind response value contained custom fields.
res, _ := client.CRM.Deal.Get("dealID", &CustomDeal{}, &hubspot.RequestQueryOption{
    CustomProperties: []string{
        "custom_a",
        "custom_b",
    },
})

// Type assertion to convert `interface` to `CustomDeal`.
customDeal, ok := res.Properties.(*CustomDeal)
if !ok {
    return errors.New("unable to type assertion")
}

// Use custom deal fields.
fmt.Println(customDeal.CustomA, customDeal.CustomB)
```

---

### Create deal with custom properties.

```go
type CustomDeal struct {
	hubspot.Deal // embed default fields.
	CustomA string `json:"custom_a,omitempty"`
	CustomB string `json:"custom_b,omitempty"`
}

// Initialize hubspot client with auth method.
client, _ := hubspot.NewClient(hubspot.SetAPIKey("YOUR_API_KEY"))

req := &CustomDeal{
    Deal: hubspot.Deal{
        Amount:      hubspot.NewString("1000"),
        DealName:    hubspot.NewString("deal001"),
        DealStage:   hubspot.NewString("stageA"),
        DealOwnerID: hubspot.NewString("admin"),
        PipeLine:    hubspot.NewString("default"),
    },
    CustomA: "custom field A",
    CustomB: "custom field B",
}

// Call create deal api with custom struct.
res, _ := client.CRM.Deal.Create(req)

// Type assertion to convert `interface` to `CustomDeal`.
customDeal, ok := res.Properties.(*CustomDeal)
if !ok {
    return errors.New("unable to type assertion")
}

// Use custom deal fields.
fmt.Println(customDeal.CustomA, customDeal.CustomB)
```

# API availability

|Category     | API     | Availability |
|-------------|---------|--------------|
|CRM          | Deal    |  Available |
|CRM          | Contact |  Available |
|CMS          | All     |  Not Implemented |
|Conversations| All     |  Not Implemented |
|Events       | All     |  Not Implemented |
|Marketing    | All     |  Not Implemented |
|Files        | All     |  Not Implemented |
|Settings     | All     |  Not Implemented |
|Webhooks     | All     |  Not Implemented |

# Authentication availability

|Type         | Availability |
|-------------|--------------|
|API key      | Available |
|OAuth        | Available |
|Private apps | Not Implemented |

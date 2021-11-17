go-hubspot
---
[![godoc](https://godoc.org/github.com/belong-inc/go-hubspot?status.svg)](https://pkg.go.dev/github.com/belong-inc/go-hubspot)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Hubspot Go Library that works with [Hubspot API v3](https://developers.hubspot.com/docs/api/overview). HubSpot
officially supports client library of Node.js, PHP, Ruby, and Python but not Go.

Note: go-hubspot currently doesn't cover all the APIs but mainly implemented CRM APIs. Implemented APIs are used in
production.

# Install

`$ go get github.com/belong-inc/go-hubspot`

# Usage

## Authentication

### API key

### OAuth

You should take refresh token in advance. Follow steps
in [here](https://developers.hubspot.com/docs/api/working-with-oauth).

## API call

TODO: put example

### Using custom fields

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


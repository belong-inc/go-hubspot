package hubspot_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/belong-inc/go-hubspot"
)

func TestWithAPIVersion(t *testing.T) {
	want := "v0"
	c, _ := hubspot.NewClient(hubspot.SetPrivateAppToken("token", "secret"), hubspot.WithAPIVersion(want))
	if want != c.ExportGetAPIVersion() {
		t.Errorf("WithAPIVersion() result mismatch: want %s got %s", want, c.ExportGetAPIVersion())
	}
}

func TestWithHTTPClient(t *testing.T) {
	want := &http.Client{Timeout: 10 * time.Second}
	c, _ := hubspot.NewClient(hubspot.SetPrivateAppToken("token", "secret"), hubspot.WithHTTPClient(want))
	if diff := cmp.Diff(want, c.HTTPClient); diff != "" {
		t.Errorf("WithHTTPClient() result mismatch: (-want +got):%s", diff)
	}
}

func TestWithBaseURL(t *testing.T) {
	want := &url.URL{Scheme: "http", Host: "example.com"}
	c, _ := hubspot.NewClient(hubspot.SetPrivateAppToken("token", "secret"), hubspot.WithBaseURL(want))
	if diff := cmp.Diff(want, c.ExportGetBaseURL()); diff != "" {
		t.Errorf("WithBaseURL() result mismatch: (-want +got):%s", diff)
	}
}

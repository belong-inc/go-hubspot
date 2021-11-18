package hubspot

import (
	"encoding/json"
	"time"
)

type HsStr string

// BlankStr should be used to include empty string in HubSpot fields.
// This is because fields set to `nil` will be ignored by omitempty.
var BlankStr = NewString("")

// NewString returns pointer HsStr(string).
// Use it to convert values that are not held in variables.
// Make sure to use BlankStr for empty string.
func NewString(s string) *HsStr {
	h := HsStr(s)
	return &h
}

// String implemented Stringer.
func (hs HsStr) String() string {
	return string(hs)
}

type HsBool bool

// UnmarshalJSON implemented json.Unmarshaler.
// This is because there are cases where the Time value returned by HubSpot is null or "true" / "false".
func (hb *HsBool) UnmarshalJSON(b []byte) error {
	s := string(b)

	*hb = false
	if s == "true" || s == `"true"` {
		*hb = true
	}
	return nil
}

type HsTime time.Time

// UnmarshalJSON implemented json.Unmarshaler.
// This is because there are cases where the Time value returned by HubSpot is null or empty string.
// The time.Time does not support Parse with empty string.
func (ht *HsTime) UnmarshalJSON(b []byte) error {
	// FIXME: Initialization is performed on empty string.
	if len(b) == 0 || string(b) == `""` {
		return nil
	}
	v := &time.Time{}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	*ht = HsTime(*v)
	return nil
}

// String implemented Stringer.
// If the value is zero, it will be displayed as `<nil>`.
func (ht *HsTime) String() string {
	v := time.Time(*ht)
	if v.IsZero() {
		return "<nil>"
	}
	return v.String()
}

// ToTime convert HsTime to time.Time.
// If the value is zero, it will be return nil.
func (ht *HsTime) ToTime() *time.Time {
	v := time.Time(*ht)
	if v.IsZero() {
		return nil
	}
	return &v
}

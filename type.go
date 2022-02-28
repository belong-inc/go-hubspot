package hubspot

import (
	"encoding/json"
	"time"
)

// HsStr is defined to identify HubSpot's empty string from null.
// If you want to set a HubSpot's value, use NewString(), if null, use `nil` in the request field.
type HsStr string

// BlankStr should be used to include empty string in HubSpot fields.
// This is because fields set to `nil` will be ignored by omitempty.
var BlankStr = NewString("")

// NewString returns pointer HsStr(string).
// Make sure to use BlankStr for empty string.
func NewString(s string) *HsStr {
	v := HsStr(s)
	return &v
}

// String implemented Stringer.
func (hs *HsStr) String() string {
	if hs == nil {
		return ""
	}
	return string(*hs)
}

// HsBool is defined to marshal the HubSpot boolean fields of `true`, `"true"`, and so on, into a bool type.
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

// HsTime is defined to identify HubSpot time fields with null and empty string.
// If you want to set a HubSpot's value, use NewTime(), if null, use `nil` in the request field.
type HsTime time.Time

// NewTime returns pointer HsTime(time.Time).
func NewTime(t time.Time) *HsTime {
	v := HsTime(t)
	return &v
}

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
func (ht *HsTime) String() string {
	if ht == nil {
		return "nil"
	}
	v := time.Time(*ht)
	if v.IsZero() {
		return ""
	}
	return v.String()
}

// ToTime convert HsTime to time.Time.
// If the value is zero, it will be return nil.
func (ht *HsTime) ToTime() *time.Time {
	if ht == nil {
		return nil
	}
	v := time.Time(*ht)
	if v.IsZero() {
		return nil
	}
	return &v
}

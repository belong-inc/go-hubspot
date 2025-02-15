package hubspot_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/belong-inc/go-hubspot"
)

func TestHsStr_String(t *testing.T) {
	tests := []struct {
		name string
		hs   *hubspot.HsStr
		want string
	}{
		{
			name: "Success",
			hs:   hubspot.NewString("text"),
			want: "text",
		},
		{
			name: "Success case of nil receiver",
			hs:   nil,
			want: "",
		},
		{
			name: "Success case of empty string",
			hs:   hubspot.NewString(""),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hs.String(); got != tt.want {
				t.Errorf("HsStr.String() mismatch: want %v, got = %v", tt.want, got)
			}
		})
	}
}

func TestHsBool_Boolean(t *testing.T) {
	tests := []struct {
		input    bool
		expected bool
	}{
		{true, true},
		{false, false},
	}

	for _, test := range tests {
		result := hubspot.NewBoolean(test.input)
		if *result != hubspot.HsBool(test.expected) {
			t.Errorf("NewBoolean(%v) = %v; want %v", test.input, *result, test.expected)
		}
	}
}

var testDate = time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC)

func TestHsTime_String(t *testing.T) {
	tests := []struct {
		name string
		ht   *hubspot.HsTime
		want string
	}{
		{
			name: "Success",
			ht:   hubspot.NewTime(testDate),
			want: "2022-02-28 00:00:00 +0000 UTC",
		},
		{
			name: "Success case of nil receiver",
			ht:   nil,
			want: "nil",
		},
		{
			name: "Success case of zero value",
			ht:   &hubspot.HsTime{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ht.String(); got != tt.want {
				t.Errorf("HsTime.String() mismatch: want %v, got = %v", tt.want, got)
			}
		})
	}
}

func TestHsTime_ToTime(t *testing.T) {
	tests := []struct {
		name string
		ht   *hubspot.HsTime
		want *time.Time
	}{
		{
			name: "Success",
			ht:   hubspot.NewTime(testDate),
			want: &testDate,
		},
		{
			name: "Success case of nil receiver",
			ht:   nil,
			want: nil,
		},
		{
			name: "Success case of zero value",
			ht:   &hubspot.HsTime{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ht.ToTime()
			if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
				t.Errorf("ToTime() mismatch (-want +got):%s", diff)
			}
		})
	}
}

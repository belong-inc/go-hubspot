package hubspot_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/belong-inc/go-hubspot"
	"github.com/google/go-cmp/cmp"
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

func TestHsTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *hubspot.HsTime
		wantErr bool
	}{
		{
			name:  "Unix timestamp in milliseconds",
			input: `1645920000000`, // 2022-02-27 00:00:00 UTC
			want:  hubspot.NewTime(time.Date(2022, time.February, 27, 0, 0, 0, 0, time.UTC)),
		},
		{
			name:  "Unix timestamp with fractional milliseconds",
			input: `1645920123456`, // 2022-02-27 00:02:03.456 UTC
			want:  hubspot.NewTime(time.Date(2022, time.February, 27, 0, 2, 3, 456000000, time.UTC)),
		},
		{
			name:  "ISO 8601 date string",
			input: `"2022-02-28T00:00:00Z"`,
			want:  hubspot.NewTime(time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC)),
		},
		{
			name:  "ISO 8601 date string with timezone",
			input: `"2022-02-28T15:30:45+09:00"`,
			want:  hubspot.NewTime(time.Date(2022, time.February, 28, 6, 30, 45, 0, time.UTC)),
		},
		{
			name:  "Date string without timezone (RFC3339)",
			input: `"2022-02-28T00:00:00.000Z"`,
			want:  hubspot.NewTime(time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC)),
		},
		{
			name:  "Empty string",
			input: `""`,
			want:  &hubspot.HsTime{},
		},
		{
			name:  "Null value",
			input: `null`,
			want:  &hubspot.HsTime{},
		},
		{
			name:  "Zero unix timestamp",
			input: `0`,
			want:  hubspot.NewTime(time.Unix(0, 0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ht hubspot.HsTime
			if err := json.Unmarshal([]byte(tt.input), &ht); err != nil {
				t.Errorf("UnmarshalJSON() unexpected error: %v", err)
				return
			}
			want := time.Time(*tt.want)
			got := time.Time(ht)
			if !want.Equal(got) {
				t.Errorf("UnmarshalJSON() mismatch: want %v, got %v", want, got)
			}
		})
	}
}

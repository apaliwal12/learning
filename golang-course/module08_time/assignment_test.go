package main

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		input   string
		want    time.Time
		wantErr bool
	}{
		{"2024-10-31", time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC), false},
		{"2000-01-01", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), false},
		{"invalid", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseDate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("ParseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatUSDate(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC), "10/31/2024"},
		{time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), "01/01/2000"},
		{time.Date(1999, 12, 5, 0, 0, 0, 0, time.UTC), "12/05/1999"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatUSDate(tt.input); got != tt.want {
				t.Errorf("FormatUSDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		input time.Time
		want  bool
	}{
		{time.Date(2024, 2, 2, 0, 0, 0, 0, time.UTC), false},  // Friday
		{time.Date(2024, 2, 3, 0, 0, 0, 0, time.UTC), true},   // Saturday
		{time.Date(2024, 2, 4, 0, 0, 0, 0, time.UTC), true},   // Sunday
		{time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC), false},  // Monday
	}

	for _, tt := range tests {
		t.Run(tt.input.Weekday().String(), func(t *testing.T) {
			if got := IsWeekend(tt.input); got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name      string
		birthdate time.Time
		today     time.Time
		want      int
	}{
		{
			name:      "Birthday already passed",
			birthdate: time.Date(2000, 5, 10, 0, 0, 0, 0, time.UTC),
			today:     time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC),
			want:      24,
		},
		{
			name:      "Birthday is today",
			birthdate: time.Date(2000, 10, 31, 0, 0, 0, 0, time.UTC),
			today:     time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC),
			want:      24,
		},
		{
			name:      "Birthday has not passed yet",
			birthdate: time.Date(2000, 12, 10, 0, 0, 0, 0, time.UTC),
			today:     time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC),
			want:      23,
		},
		{
			name:      "Leap year baby - Feb 28",
			birthdate: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			today:     time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC),
			want:      22,
		},
		{
			name:      "Leap year baby - Mar 1",
			birthdate: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			today:     time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
			want:      23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateAge(tt.birthdate, tt.today); got != tt.want {
				t.Errorf("CalculateAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

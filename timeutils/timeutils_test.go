package timeutils

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name     string
		t        time.Time
		layout   string
		expected string
	}{
		{"default layout", now, "", ""}, // We can't test exact value due to time
		{"custom layout", time.Date(2025, 1, 15, 14, 30, 0, 0, time.UTC), "2006-01-02", "2025-01-15"},
		{"datetime layout", time.Date(2025, 1, 15, 14, 30, 0, 0, time.UTC), DefaultDateTimeFormat, "2025-01-15 14:30:00"},
		{"date layout", time.Date(2025, 1, 15, 14, 30, 0, 0, time.UTC), DefaultDateFormat, "2025-01-15"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FormatTime(test.t, test.layout)
			if test.layout != "" && result != test.expected {
				t.Errorf("FormatTime(%v, %q) = %q; want %q", test.t, test.layout, result, test.expected)
			}
		})
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		name      string
		timeStr   string
		layout    string
		shouldErr bool
	}{
		{"with layout", "2025-01-15", "2006-01-02", false},
		{"default layout date", "2025-01-15 14:30:00", "", false},
		{"default layout date only", "2025-01-15", "", false},
		{"invalid format", "invalid", "2006-01-02", true},
		{"RFC3339", "2025-01-15T14:30:00Z", "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ParseTime(test.timeStr, test.layout)
			if test.shouldErr && err == nil {
				t.Errorf("ParseTime(%q, %q) expected error but got none", test.timeStr, test.layout)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("ParseTime(%q, %q) unexpected error: %v", test.timeStr, test.layout, err)
			}
			if !test.shouldErr && result.IsZero() && test.layout != "" {
				t.Errorf("ParseTime(%q, %q) returned zero time", test.timeStr, test.layout)
			}
		})
	}
}

func TestAddDays(t *testing.T) {
	base := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		t        time.Time
		days     int
		expected time.Time
	}{
		{"add 1 day", base, 1, time.Date(2025, 1, 16, 12, 0, 0, 0, time.UTC)},
		{"add 7 days", base, 7, time.Date(2025, 1, 22, 12, 0, 0, 0, time.UTC)},
		{"subtract 1 day", base, -1, time.Date(2025, 1, 14, 12, 0, 0, 0, time.UTC)},
		{"add 0 days", base, 0, base},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := AddDays(test.t, test.days)
			if !result.Equal(test.expected) {
				t.Errorf("AddDays(%v, %d) = %v; want %v", test.t, test.days, result, test.expected)
			}
		})
	}
}

func TestAddMonths(t *testing.T) {
	base := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		t        time.Time
		months   int
		expected time.Time
	}{
		{"add 1 month", base, 1, time.Date(2025, 2, 15, 12, 0, 0, 0, time.UTC)},
		{"add 12 months", base, 12, time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)},
		{"subtract 1 month", base, -1, time.Date(2024, 12, 15, 12, 0, 0, 0, time.UTC)},
		{"add 0 months", base, 0, base},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := AddMonths(test.t, test.months)
			if !result.Equal(test.expected) {
				t.Errorf("AddMonths(%v, %d) = %v; want %v", test.t, test.months, result, test.expected)
			}
		})
	}
}

func TestAddYears(t *testing.T) {
	base := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		t        time.Time
		years    int
		expected time.Time
	}{
		{"add 1 year", base, 1, time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)},
		{"add 5 years", base, 5, time.Date(2030, 1, 15, 12, 0, 0, 0, time.UTC)},
		{"subtract 1 year", base, -1, time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)},
		{"add 0 years", base, 0, base},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := AddYears(test.t, test.years)
			if !result.Equal(test.expected) {
				t.Errorf("AddYears(%v, %d) = %v; want %v", test.t, test.years, result, test.expected)
			}
		})
	}
}

func TestDaysBetween(t *testing.T) {
	t1 := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 1, 20, 12, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{"5 days apart", t1, t2, 5},
		{"reverse order", t2, t1, -5},
		{"same day", t1, t1, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DaysBetween(test.t1, test.t2)
			if result != test.expected {
				t.Errorf("DaysBetween(%v, %v) = %d; want %d", test.t1, test.t2, result, test.expected)
			}
		})
	}
}

func TestHoursBetween(t *testing.T) {
	t1 := time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 1, 15, 15, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{"5 hours apart", t1, t2, 5},
		{"reverse order", t2, t1, -5},
		{"same time", t1, t1, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HoursBetween(test.t1, test.t2)
			if result != test.expected {
				t.Errorf("HoursBetween(%v, %v) = %d; want %d", test.t1, test.t2, result, test.expected)
			}
		})
	}
}

func TestMinutesBetween(t *testing.T) {
	t1 := time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 1, 15, 10, 30, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{"30 minutes apart", t1, t2, 30},
		{"reverse order", t2, t1, -30},
		{"same time", t1, t1, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MinutesBetween(test.t1, test.t2)
			if result != test.expected {
				t.Errorf("MinutesBetween(%v, %v) = %d; want %d", test.t1, test.t2, result, test.expected)
			}
		})
	}
}

func TestTimeAgo(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name string
		t    time.Time
	}{
		{"recent time", now.Add(-30 * time.Second)},
		{"minutes ago", now.Add(-30 * time.Minute)},
		{"hours ago", now.Add(-5 * time.Hour)},
		{"days ago", now.Add(-7 * 24 * time.Hour)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TimeAgo(test.t)
			if result == "" {
				t.Errorf("TimeAgo(%v) returned empty string", test.t)
			}
		})
	}
}

func TestTimeAgoEn(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name string
		t    time.Time
	}{
		{"recent time", now.Add(-30 * time.Second)},
		{"minutes ago", now.Add(-30 * time.Minute)},
		{"hours ago", now.Add(-5 * time.Hour)},
		{"days ago", now.Add(-7 * 24 * time.Hour)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TimeAgoEn(test.t)
			if result == "" {
				t.Errorf("TimeAgoEn(%v) returned empty string", test.t)
			}
		})
	}
}

func TestToday(t *testing.T) {
	result := Today()
	now := time.Now()
	expected := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	
	if !result.Equal(expected) {
		t.Errorf("Today() = %v; want %v", result, expected)
	}
}

func TestYesterday(t *testing.T) {
	yesterday := Yesterday()
	today := Today()
	expected := AddDays(today, -1)
	
	if !yesterday.Equal(expected) {
		t.Errorf("Yesterday() = %v; want %v", yesterday, expected)
	}
}

func TestTomorrow(t *testing.T) {
	tomorrow := Tomorrow()
	today := Today()
	expected := AddDays(today, 1)
	
	if !tomorrow.Equal(expected) {
		t.Errorf("Tomorrow() = %v; want %v", tomorrow, expected)
	}
}

func TestThisWeek(t *testing.T) {
	result := ThisWeek()
	if result.Weekday() != time.Monday || result.Hour() != 0 || result.Minute() != 0 {
		t.Errorf("ThisWeek() should return Monday 00:00:00, got %v", result)
	}
}

func TestThisMonth(t *testing.T) {
	result := ThisMonth()
	now := time.Now()
	expected := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	
	if !result.Equal(expected) {
		t.Errorf("ThisMonth() = %v; want %v", result, expected)
	}
}

func TestThisYear(t *testing.T) {
	result := ThisYear()
	now := time.Now()
	expected := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	
	if !result.Equal(expected) {
		t.Errorf("ThisYear() = %v; want %v", result, expected)
	}
}

func TestStartOfDay(t *testing.T) {
	t1 := time.Date(2025, 1, 15, 14, 30, 45, 123456789, time.UTC)
	result := StartOfDay(t1)
	expected := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	
	if !result.Equal(expected) {
		t.Errorf("StartOfDay(%v) = %v; want %v", t1, result, expected)
	}
}

func TestEndOfDay(t *testing.T) {
	t1 := time.Date(2025, 1, 15, 14, 30, 45, 123456789, time.UTC)
	result := EndOfDay(t1)
	expected := time.Date(2025, 1, 15, 23, 59, 59, 999999999, time.UTC)
	
	if !result.Equal(expected) {
		t.Errorf("EndOfDay(%v) = %v; want %v", t1, result, expected)
	}
}

func TestStartOfWeek(t *testing.T) {
	// Test with a known day
	t1 := time.Date(2025, 1, 15, 14, 30, 0, 0, time.UTC) // Wednesday
	result := StartOfWeek(t1)
	
	if result.Weekday() != time.Monday || result.Hour() != 0 || result.Minute() != 0 {
		t.Errorf("StartOfWeek(%v) should return Monday 00:00:00, got %v", t1, result)
	}
}

func TestEndOfWeek(t *testing.T) {
	t1 := time.Date(2025, 1, 15, 14, 30, 0, 0, time.UTC)
	result := EndOfWeek(t1)
	start := StartOfWeek(t1)
	expected := EndOfDay(AddDays(start, 6))
	
	if !result.Equal(expected) {
		t.Errorf("EndOfWeek(%v) = %v; want %v", t1, result, expected)
	}
}

func TestStartOfMonth(t *testing.T) {
	t1 := time.Date(2025, 1, 15, 14, 30, 0, 0, time.UTC)
	result := StartOfMonth(t1)
	expected := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	
	if !result.Equal(expected) {
		t.Errorf("StartOfMonth(%v) = %v; want %v", t1, result, expected)
	}
}

func TestEndOfMonth(t *testing.T) {
	t1 := time.Date(2025, 1, 15, 14, 30, 0, 0, time.UTC)
	result := EndOfMonth(t1)
	expected := time.Date(2025, 1, 31, 23, 59, 59, 999999999, time.UTC)
	
	if !result.Equal(expected) {
		t.Errorf("EndOfMonth(%v) = %v; want %v", t1, result, expected)
	}
}

func TestStartOfYear(t *testing.T) {
	t1 := time.Date(2025, 6, 15, 14, 30, 0, 0, time.UTC)
	result := StartOfYear(t1)
	expected := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	
	if !result.Equal(expected) {
		t.Errorf("StartOfYear(%v) = %v; want %v", t1, result, expected)
	}
}

func TestEndOfYear(t *testing.T) {
	t1 := time.Date(2025, 6, 15, 14, 30, 0, 0, time.UTC)
	result := EndOfYear(t1)
	expected := time.Date(2025, 12, 31, 23, 59, 59, 999999999, time.UTC)
	
	if !result.Equal(expected) {
		t.Errorf("EndOfYear(%v) = %v; want %v", t1, result, expected)
	}
}

func TestToTimezone(t *testing.T) {
	utcTime := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name      string
		t         time.Time
		timezone  string
		shouldErr bool
	}{
		{"valid timezone", utcTime, "Asia/Shanghai", false},
		{"UTC timezone", utcTime, "UTC", false},
		{"invalid timezone", utcTime, "Invalid/Zone", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ToTimezone(test.t, test.timezone)
			if test.shouldErr && err == nil {
				t.Errorf("ToTimezone(%v, %q) expected error but got none", test.t, test.timezone)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("ToTimezone(%v, %q) unexpected error: %v", test.t, test.timezone, err)
			}
			if !test.shouldErr && result.IsZero() {
				t.Errorf("ToTimezone(%v, %q) returned zero time", test.t, test.timezone)
			}
		})
	}
}

func TestToUTC(t *testing.T) {
	shanghaiLoc, _ := time.LoadLocation("Asia/Shanghai")
	localTime := time.Date(2025, 1, 15, 20, 0, 0, 0, shanghaiLoc)
	result := ToUTC(localTime)
	
	if result.Location() != time.UTC {
		t.Errorf("ToUTC(%v) location = %v; want UTC", localTime, result.Location())
	}
}

func TestIsToday(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name     string
		t        time.Time
		expected bool
	}{
		{"current time", now, true},
		{"start of today", Today(), true},
		{"yesterday", Yesterday(), false},
		{"tomorrow", Tomorrow(), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsToday(test.t)
			if result != test.expected {
				t.Errorf("IsToday(%v) = %v; want %v", test.t, result, test.expected)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	weekday := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC) // Wednesday
	saturday := time.Date(2025, 1, 18, 12, 0, 0, 0, time.UTC)
	sunday := time.Date(2025, 1, 19, 12, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		t        time.Time
		expected bool
	}{
		{"weekday", weekday, false},
		{"saturday", saturday, true},
		{"sunday", sunday, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsWeekend(test.t)
			if result != test.expected {
				t.Errorf("IsWeekend(%v) = %v; want %v", test.t, result, test.expected)
			}
		})
	}
}

func TestIsWeekday(t *testing.T) {
	weekday := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC) // Wednesday
	saturday := time.Date(2025, 1, 18, 12, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		t        time.Time
		expected bool
	}{
		{"weekday", weekday, true},
		{"saturday", saturday, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsWeekday(test.t)
			if result != test.expected {
				t.Errorf("IsWeekday(%v) = %v; want %v", test.t, result, test.expected)
			}
		})
	}
}

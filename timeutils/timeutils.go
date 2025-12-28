// Package timeutils 提供时间/日期相关的工具函数
// Package timeutils provides time/date utility functions
package timeutils

import (
	"fmt"
	"time"
)

// 常用的时间格式常量
// Common time format constants
const (
	// DefaultDateTimeFormat 默认日期时间格式 (2006-01-02 15:04:05)
	DefaultDateTimeFormat = "2006-01-02 15:04:05"
	// DefaultDateFormat 默认日期格式 (2006-01-02)
	DefaultDateFormat = "2006-01-02"
	// DefaultTimeFormat 默认时间格式 (15:04:05)
	DefaultTimeFormat = "15:04:05"
	// RFC3339Format RFC3339 格式 (2006-01-02T15:04:05Z07:00)
	RFC3339Format = time.RFC3339
	// ISO8601Format ISO8601 格式 (2006-01-02T15:04:05.000Z)
	ISO8601Format = "2006-01-02T15:04:05.000Z"
)

// FormatTime 格式化时间为指定格式的字符串
//
// 参数 / Parameters:
//   - t: 要格式化的时间 / time to format
//   - layout: 时间格式布局，如果为空则使用默认格式 (2006-01-02 15:04:05) / time layout, use default if empty
//
// 返回值 / Returns:
//   - string: 格式化后的时间字符串 / formatted time string
//
// 示例 / Example:
//   FormatTime(time.Now(), "2006-01-02") // "2025-01-15"
//
// FormatTime formats a time to a string with specified layout
func FormatTime(t time.Time, layout string) string {
	if layout == "" {
		layout = DefaultDateTimeFormat
	}
	return t.Format(layout)
}

// ParseTime 解析时间字符串为 time.Time
//
// 参数 / Parameters:
//   - timeStr: 时间字符串 / time string
//   - layout: 时间格式布局，如果为空则尝试常用格式 / time layout, try common formats if empty
//
// 返回值 / Returns:
//   - time.Time: 解析后的时间 / parsed time
//   - error: 如果解析失败则返回错误 / error if parsing fails
//
// 示例 / Example:
//   t, _ := ParseTime("2025-01-15", "2006-01-02")
//
// ParseTime parses a time string to time.Time
func ParseTime(timeStr, layout string) (time.Time, error) {
	if layout == "" {
		// 尝试常用格式
		layouts := []string{
			DefaultDateTimeFormat,
			DefaultDateFormat,
			RFC3339Format,
			ISO8601Format,
			time.RFC3339Nano,
			time.RFC822,
			time.RFC822Z,
			time.RFC1123,
			time.RFC1123Z,
		}
		for _, l := range layouts {
			if t, err := time.Parse(l, timeStr); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("无法解析时间字符串: %s", timeStr)
	}
	return time.Parse(layout, timeStr)
}

// AddDays 在指定时间上增加指定天数
//
// 参数 / Parameters:
//   - t: 基准时间 / base time
//   - days: 要增加的天数（可以是负数） / days to add (can be negative)
//
// 返回值 / Returns:
//   - time.Time: 增加天数后的时间 / time after adding days
//
// 示例 / Example:
//   tomorrow := AddDays(time.Now(), 1)
//
// AddDays adds specified days to a time
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths 在指定时间上增加指定月数
//
// 参数 / Parameters:
//   - t: 基准时间 / base time
//   - months: 要增加的月数（可以是负数） / months to add (can be negative)
//
// 返回值 / Returns:
//   - time.Time: 增加月数后的时间 / time after adding months
//
// 示例 / Example:
//   nextMonth := AddMonths(time.Now(), 1)
//
// AddMonths adds specified months to a time
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears 在指定时间上增加指定年数
//
// 参数 / Parameters:
//   - t: 基准时间 / base time
//   - years: 要增加的年数（可以是负数） / years to add (can be negative)
//
// 返回值 / Returns:
//   - time.Time: 增加年数后的时间 / time after adding years
//
// 示例 / Example:
//   nextYear := AddYears(time.Now(), 1)
//
// AddYears adds specified years to a time
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// DaysBetween 计算两个时间之间的天数差
//
// 参数 / Parameters:
//   - t1: 第一个时间 / first time
//   - t2: 第二个时间 / second time
//
// 返回值 / Returns:
//   - int: 天数差（t2 - t1） / days difference (t2 - t1)
//
// 示例 / Example:
//   days := DaysBetween(time.Now(), futureTime)
//
// DaysBetween calculates the number of days between two times
func DaysBetween(t1, t2 time.Time) int {
	duration := t2.Sub(t1)
	return int(duration.Hours() / 24)
}

// HoursBetween 计算两个时间之间的小时数差
//
// 参数 / Parameters:
//   - t1: 第一个时间 / first time
//   - t2: 第二个时间 / second time
//
// 返回值 / Returns:
//   - int: 小时数差（t2 - t1） / hours difference (t2 - t1)
//
// 示例 / Example:
//   hours := HoursBetween(time.Now(), futureTime)
//
// HoursBetween calculates the number of hours between two times
func HoursBetween(t1, t2 time.Time) int {
	duration := t2.Sub(t1)
	return int(duration.Hours())
}

// MinutesBetween 计算两个时间之间的分钟数差
//
// 参数 / Parameters:
//   - t1: 第一个时间 / first time
//   - t2: 第二个时间 / second time
//
// 返回值 / Returns:
//   - int: 分钟数差（t2 - t1） / minutes difference (t2 - t1)
//
// 示例 / Example:
//   minutes := MinutesBetween(time.Now(), futureTime)
//
// MinutesBetween calculates the number of minutes between two times
func MinutesBetween(t1, t2 time.Time) int {
	duration := t2.Sub(t1)
	return int(duration.Minutes())
}

// TimeAgo 返回相对时间描述（如"2小时前"）
//
// 参数 / Parameters:
//   - t: 要计算的时间 / time to calculate
//
// 返回值 / Returns:
//   - string: 相对时间描述 / relative time description
//
// 示例 / Example:
//   TimeAgo(time.Now().Add(-2 * time.Hour)) // "2小时前"
//
// TimeAgo returns a relative time description (e.g., "2 hours ago")
func TimeAgo(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	if duration < time.Minute {
		return "刚刚"
	} else if duration < time.Hour {
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d分钟前", minutes)
	} else if duration < 24*time.Hour {
		hours := int(duration.Hours())
		return fmt.Sprintf("%d小时前", hours)
	} else if duration < 30*24*time.Hour {
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d天前", days)
	} else if duration < 365*24*time.Hour {
		months := int(duration.Hours() / (30 * 24))
		return fmt.Sprintf("%d个月前", months)
	} else {
		years := int(duration.Hours() / (365 * 24))
		return fmt.Sprintf("%d年前", years)
	}
}

// TimeAgoEn 返回英文相对时间描述（如"2 hours ago"）
//
// 参数 / Parameters:
//   - t: 要计算的时间 / time to calculate
//
// 返回值 / Returns:
//   - string: 英文相对时间描述 / English relative time description
//
// 示例 / Example:
//   TimeAgoEn(time.Now().Add(-2 * time.Hour)) // "2 hours ago"
//
// TimeAgoEn returns an English relative time description (e.g., "2 hours ago")
func TimeAgoEn(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	if duration < time.Minute {
		return "just now"
	} else if duration < time.Hour {
		minutes := int(duration.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	} else if duration < 24*time.Hour {
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else if duration < 30*24*time.Hour {
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	} else if duration < 365*24*time.Hour {
		months := int(duration.Hours() / (30 * 24))
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	} else {
		years := int(duration.Hours() / (365 * 24))
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	}
}

// Today 返回今天的开始时间（00:00:00）
//
// 返回值 / Returns:
//   - time.Time: 今天的开始时间 / start of today
//
// 示例 / Example:
//   startOfToday := Today()
//
// Today returns the start of today (00:00:00)
func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// Yesterday 返回昨天的开始时间（00:00:00）
//
// 返回值 / Returns:
//   - time.Time: 昨天的开始时间 / start of yesterday
//
// 示例 / Example:
//   startOfYesterday := Yesterday()
//
// Yesterday returns the start of yesterday (00:00:00)
func Yesterday() time.Time {
	return AddDays(Today(), -1)
}

// Tomorrow 返回明天的开始时间（00:00:00）
//
// 返回值 / Returns:
//   - time.Time: 明天的开始时间 / start of tomorrow
//
// 示例 / Example:
//   startOfTomorrow := Tomorrow()
//
// Tomorrow returns the start of tomorrow (00:00:00)
func Tomorrow() time.Time {
	return AddDays(Today(), 1)
}

// ThisWeek 返回本周的开始时间（周一 00:00:00）
//
// 返回值 / Returns:
//   - time.Time: 本周的开始时间 / start of this week
//
// 示例 / Example:
//   startOfWeek := ThisWeek()
//
// ThisWeek returns the start of this week (Monday 00:00:00)
func ThisWeek() time.Time {
	now := time.Now()
	weekday := int(now.Weekday())
	// Go中 Sunday = 0, Monday = 1, ..., Saturday = 6
	// 转换为 Monday = 0, ..., Sunday = 6
	if weekday == 0 {
		weekday = 7
	}
	daysFromMonday := weekday - 1
	return time.Date(now.Year(), now.Month(), now.Day()-daysFromMonday, 0, 0, 0, 0, now.Location())
}

// ThisMonth 返回本月的开始时间（1号 00:00:00）
//
// 返回值 / Returns:
//   - time.Time: 本月的开始时间 / start of this month
//
// 示例 / Example:
//   startOfMonth := ThisMonth()
//
// ThisMonth returns the start of this month (1st 00:00:00)
func ThisMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// ThisYear 返回本年的开始时间（1月1日 00:00:00）
//
// 返回值 / Returns:
//   - time.Time: 本年的开始时间 / start of this year
//
// 示例 / Example:
//   startOfYear := ThisYear()
//
// ThisYear returns the start of this year (Jan 1st 00:00:00)
func ThisYear() time.Time {
	now := time.Now()
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
}

// StartOfDay 返回指定日期当天的开始时间（00:00:00）
//
// 参数 / Parameters:
//   - t: 指定时间 / specified time
//
// 返回值 / Returns:
//   - time.Time: 当天的开始时间 / start of the day
//
// 示例 / Example:
//   start := StartOfDay(time.Now())
//
// StartOfDay returns the start of the day (00:00:00) for the specified time
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay 返回指定日期当天的结束时间（23:59:59.999999999）
//
// 参数 / Parameters:
//   - t: 指定时间 / specified time
//
// 返回值 / Returns:
//   - time.Time: 当天的结束时间 / end of the day
//
// 示例 / Example:
//   end := EndOfDay(time.Now())
//
// EndOfDay returns the end of the day (23:59:59.999999999) for the specified time
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek 返回指定日期所在周的开始时间（周一 00:00:00）
//
// 参数 / Parameters:
//   - t: 指定时间 / specified time
//
// 返回值 / Returns:
//   - time.Time: 所在周的开始时间 / start of the week
//
// 示例 / Example:
//   start := StartOfWeek(time.Now())
//
// StartOfWeek returns the start of the week (Monday 00:00:00) for the specified time
func StartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	daysFromMonday := weekday - 1
	return time.Date(t.Year(), t.Month(), t.Day()-daysFromMonday, 0, 0, 0, 0, t.Location())
}

// EndOfWeek 返回指定日期所在周的结束时间（周日 23:59:59.999999999）
//
// 参数 / Parameters:
//   - t: 指定时间 / specified time
//
// 返回值 / Returns:
//   - time.Time: 所在周的结束时间 / end of the week
//
// 示例 / Example:
//   end := EndOfWeek(time.Now())
//
// EndOfWeek returns the end of the week (Sunday 23:59:59.999999999) for the specified time
func EndOfWeek(t time.Time) time.Time {
	start := StartOfWeek(t)
	return EndOfDay(AddDays(start, 6))
}

// StartOfMonth 返回指定日期所在月的开始时间（1号 00:00:00）
//
// 参数 / Parameters:
//   - t: 指定时间 / specified time
//
// 返回值 / Returns:
//   - time.Time: 所在月的开始时间 / start of the month
//
// 示例 / Example:
//   start := StartOfMonth(time.Now())
//
// StartOfMonth returns the start of the month (1st 00:00:00) for the specified time
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 返回指定日期所在月的结束时间（最后一天 23:59:59.999999999）
//
// 参数 / Parameters:
//   - t: 指定时间 / specified time
//
// 返回值 / Returns:
//   - time.Time: 所在月的结束时间 / end of the month
//
// 示例 / Example:
//   end := EndOfMonth(time.Now())
//
// EndOfMonth returns the end of the month (last day 23:59:59.999999999) for the specified time
func EndOfMonth(t time.Time) time.Time {
	// 获取下个月的第一天，然后减去1纳秒
	nextMonth := StartOfMonth(t).AddDate(0, 1, 0)
	return nextMonth.Add(-time.Nanosecond)
}

// StartOfYear 返回指定日期所在年的开始时间（1月1日 00:00:00）
//
// 参数 / Parameters:
//   - t: 指定时间 / specified time
//
// 返回值 / Returns:
//   - time.Time: 所在年的开始时间 / start of the year
//
// 示例 / Example:
//   start := StartOfYear(time.Now())
//
// StartOfYear returns the start of the year (Jan 1st 00:00:00) for the specified time
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 返回指定日期所在年的结束时间（12月31日 23:59:59.999999999）
//
// 参数 / Parameters:
//   - t: 指定时间 / specified time
//
// 返回值 / Returns:
//   - time.Time: 所在年的结束时间 / end of the year
//
// 示例 / Example:
//   end := EndOfYear(time.Now())
//
// EndOfYear returns the end of the year (Dec 31st 23:59:59.999999999) for the specified time
func EndOfYear(t time.Time) time.Time {
	nextYear := StartOfYear(t).AddDate(1, 0, 0)
	return nextYear.Add(-time.Nanosecond)
}

// ToTimezone 将时间转换为指定时区
//
// 参数 / Parameters:
//   - t: 要转换的时间 / time to convert
//   - timezone: 时区名称（如"Asia/Shanghai", "UTC", "America/New_York"） / timezone name
//
// 返回值 / Returns:
//   - time.Time: 转换后的时间 / converted time
//   - error: 如果时区无效则返回错误 / error if timezone is invalid
//
// 示例 / Example:
//   t, _ := ToTimezone(time.Now(), "Asia/Shanghai")
//
// ToTimezone converts a time to the specified timezone
func ToTimezone(t time.Time, timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("无效的时区: %s", timezone)
	}
	return t.In(loc), nil
}

// ToUTC 将时间转换为UTC时区
//
// 参数 / Parameters:
//   - t: 要转换的时间 / time to convert
//
// 返回值 / Returns:
//   - time.Time: UTC时间 / UTC time
//
// 示例 / Example:
//   utcTime := ToUTC(time.Now())
//
// ToUTC converts a time to UTC timezone
func ToUTC(t time.Time) time.Time {
	return t.UTC()
}

// IsToday 判断时间是否为今天
//
// 参数 / Parameters:
//   - t: 要判断的时间 / time to check
//
// 返回值 / Returns:
//   - bool: 是否为今天 / whether it is today
//
// 示例 / Example:
//   isToday := IsToday(someTime)
//
// IsToday checks if the time is today
func IsToday(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// IsWeekend 判断时间是否为周末
//
// 参数 / Parameters:
//   - t: 要判断的时间 / time to check
//
// 返回值 / Returns:
//   - bool: 是否为周末（周六或周日） / whether it is weekend (Saturday or Sunday)
//
// 示例 / Example:
//   isWeekend := IsWeekend(time.Now())
//
// IsWeekend checks if the time is weekend
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday 判断时间是否为工作日
//
// 参数 / Parameters:
//   - t: 要判断的时间 / time to check
//
// 返回值 / Returns:
//   - bool: 是否为工作日（周一到周五） / whether it is weekday (Monday to Friday)
//
// 示例 / Example:
//   isWeekday := IsWeekday(time.Now())
//
// IsWeekday checks if the time is weekday
func IsWeekday(t time.Time) bool {
	return !IsWeekend(t)
}

package main

import (
	"fmt"
	"time"

	"github.com/Rodert/go-commons/timeutils"
)

func main() {
	fmt.Println("=== Go Commons Time Utils Demo ===")
	fmt.Println()

	now := time.Now()
	fmt.Println("当前时间 / Current Time:", timeutils.FormatTime(now, timeutils.DefaultDateTimeFormat))
	fmt.Println()

	fmt.Println("时间格式化 / Time Formatting:")
	fmt.Println("默认格式:", timeutils.FormatTime(now, ""))
	fmt.Println("日期格式:", timeutils.FormatTime(now, timeutils.DefaultDateFormat))
	fmt.Println("时间格式:", timeutils.FormatTime(now, timeutils.DefaultTimeFormat))
	fmt.Println("RFC3339:", timeutils.FormatTime(now, timeutils.RFC3339Format))
	fmt.Println()

	fmt.Println("时间计算 / Time Calculations:")
	fmt.Println("明天:", timeutils.FormatTime(timeutils.Tomorrow(), timeutils.DefaultDateFormat))
	fmt.Println("昨天:", timeutils.FormatTime(timeutils.Yesterday(), timeutils.DefaultDateFormat))
	fmt.Println("加7天:", timeutils.FormatTime(timeutils.AddDays(now, 7), timeutils.DefaultDateTimeFormat))
	fmt.Println("加1个月:", timeutils.FormatTime(timeutils.AddMonths(now, 1), timeutils.DefaultDateTimeFormat))
	fmt.Println("加1年:", timeutils.FormatTime(timeutils.AddYears(now, 1), timeutils.DefaultDateTimeFormat))
	fmt.Println()

	fmt.Println("时间差计算 / Time Difference:")
	yesterday := timeutils.Yesterday()
	fmt.Println("昨天到现在的天数差:", timeutils.DaysBetween(yesterday, now))
	fmt.Println("昨天到现在的小时数差:", timeutils.HoursBetween(yesterday, now))
	fmt.Println("昨天到现在的分钟数差:", timeutils.MinutesBetween(yesterday, now))
	fmt.Println()

	fmt.Println("相对时间 / Relative Time:")
	fmt.Println("2小时前:", timeutils.TimeAgo(now.Add(-2*time.Hour)))
	fmt.Println("2天前:", timeutils.TimeAgo(now.Add(-2*24*time.Hour)))
	fmt.Println("2个月前:", timeutils.TimeAgo(now.Add(-2*30*24*time.Hour)))
	fmt.Println("2小时前 (英文):", timeutils.TimeAgoEn(now.Add(-2*time.Hour)))
	fmt.Println()

	fmt.Println("时间范围 / Time Ranges:")
	fmt.Println("今天开始:", timeutils.FormatTime(timeutils.Today(), timeutils.DefaultDateTimeFormat))
	fmt.Println("本周开始:", timeutils.FormatTime(timeutils.ThisWeek(), timeutils.DefaultDateTimeFormat))
	fmt.Println("本月开始:", timeutils.FormatTime(timeutils.ThisMonth(), timeutils.DefaultDateTimeFormat))
	fmt.Println("本年开始:", timeutils.FormatTime(timeutils.ThisYear(), timeutils.DefaultDateTimeFormat))
	fmt.Println()

	fmt.Println("时间边界 / Time Boundaries:")
	testTime := time.Date(2025, 6, 15, 14, 30, 45, 0, time.UTC)
	fmt.Println("测试时间:", timeutils.FormatTime(testTime, timeutils.DefaultDateTimeFormat))
	fmt.Println("当天开始:", timeutils.FormatTime(timeutils.StartOfDay(testTime), timeutils.DefaultDateTimeFormat))
	fmt.Println("当天结束:", timeutils.FormatTime(timeutils.EndOfDay(testTime), timeutils.DefaultDateTimeFormat))
	fmt.Println("本周开始:", timeutils.FormatTime(timeutils.StartOfWeek(testTime), timeutils.DefaultDateTimeFormat))
	fmt.Println("本周结束:", timeutils.FormatTime(timeutils.EndOfWeek(testTime), timeutils.DefaultDateTimeFormat))
	fmt.Println("本月开始:", timeutils.FormatTime(timeutils.StartOfMonth(testTime), timeutils.DefaultDateTimeFormat))
	fmt.Println("本月结束:", timeutils.FormatTime(timeutils.EndOfMonth(testTime), timeutils.DefaultDateTimeFormat))
	fmt.Println("本年开始:", timeutils.FormatTime(timeutils.StartOfYear(testTime), timeutils.DefaultDateTimeFormat))
	fmt.Println("本年结束:", timeutils.FormatTime(timeutils.EndOfYear(testTime), timeutils.DefaultDateTimeFormat))
	fmt.Println()

	fmt.Println("时区转换 / Timezone Conversion:")
	utcTime := timeutils.ToUTC(now)
	fmt.Println("UTC时间:", timeutils.FormatTime(utcTime, timeutils.DefaultDateTimeFormat))
	shanghaiTime, err := timeutils.ToTimezone(now, "Asia/Shanghai")
	if err == nil {
		fmt.Println("上海时间:", timeutils.FormatTime(shanghaiTime, timeutils.DefaultDateTimeFormat))
	}
	fmt.Println()

	fmt.Println("时间判断 / Time Checks:")
	fmt.Println("是否今天:", timeutils.IsToday(now))
	fmt.Println("是否今天 (昨天):", timeutils.IsToday(yesterday))
	fmt.Println("是否周末:", timeutils.IsWeekend(now))
	fmt.Println("是否工作日:", timeutils.IsWeekday(now))
}

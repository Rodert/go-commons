package main

import (
	"fmt"
	"github.com/Rodert/go-commons/stringutils"
)

func main() {
	fmt.Println("=== Go Commons String Utils Demo ===")

	fmt.Println("基本字符串检查 / Basic String Checks:")
	fmt.Println("IsEmpty:", stringutils.IsEmpty(""), stringutils.IsEmpty("hello"))
	fmt.Println("IsNotEmpty:", stringutils.IsNotEmpty(""), stringutils.IsNotEmpty("hello"))
	fmt.Println("IsBlank:", stringutils.IsBlank("  \t\n"), stringutils.IsBlank("hello"))
	fmt.Println("IsNotBlank:", stringutils.IsNotBlank("  \t\n"), stringutils.IsNotBlank("hello"))
	fmt.Println()

	fmt.Println("字符串修改 / String Modification:")
	fmt.Println("Trim:", stringutils.Trim("  hello  "))
	fmt.Println("TrimToEmpty:", stringutils.TrimToEmpty("  world  "))
	fmt.Println("Truncate:", stringutils.Truncate("abcdefghijk", 5))
	fmt.Println("TruncateWithSuffix:", stringutils.TruncateWithSuffix("abcdef", 4, ".."))
	fmt.Println("Capitalize:", stringutils.Capitalize("hello"))
	fmt.Println("Uncapitalize:", stringutils.Uncapitalize("Hello"))
	fmt.Println()

	fmt.Println("字符串填充 / String Padding:")
	fmt.Println("PadLeft:", stringutils.PadLeft("42", 5, '0'))
	fmt.Println("PadRight:", stringutils.PadRight("42", 5, '0'))
	fmt.Println("PadCenter:", stringutils.PadCenter("42", 5, '0'))
	fmt.Println()

	fmt.Println("字符串搜索 / String Search:")
	fmt.Println("ContainsAny:", stringutils.ContainsAny("gopher", "go", "java"))
	fmt.Println("ContainsAll:", stringutils.ContainsAll("gopher", "go", "ph"))
	fmt.Println("CountMatches:", stringutils.CountMatches("ababababa", "ab"))
	fmt.Println()

	fmt.Println("字符串转换 / String Conversion:")
	fmt.Println("Reverse:", stringutils.Reverse("hello"))
	fmt.Println("SwapCase:", stringutils.SwapCase("Hello World"))
	fmt.Println("Repeat:", stringutils.Repeat("ab", 3))
}
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/stringutils"
)

func main() {
	fmt.Println("IsBlank:", stringutils.IsBlank("  \t\n"))
	fmt.Println("Trim:", stringutils.Trim("  hello  "))
	fmt.Println("TruncateWithSuffix:", stringutils.TruncateWithSuffix("abcdef", 4, ".."))
	fmt.Println("PadLeft:", stringutils.PadLeft("42", 5, '0'))
	fmt.Println("ContainsAny:", stringutils.ContainsAny("gopher", "go", "java"))
} 
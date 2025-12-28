package convertutils

import (
	"reflect"
	"testing"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  int
		shouldErr bool
	}{
		{"valid number", "123", 123, false},
		{"negative number", "-123", -123, false},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := StringToInt(test.input)
			if test.shouldErr && err == nil {
				t.Errorf("StringToInt(%q) expected error but got none", test.input)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("StringToInt(%q) unexpected error: %v", test.input, err)
			}
			if !test.shouldErr && result != test.expected {
				t.Errorf("StringToInt(%q) = %d; want %d", test.input, result, test.expected)
			}
		})
	}
}

func TestStringToFloat64(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  float64
		shouldErr bool
	}{
		{"valid float", "123.45", 123.45, false},
		{"negative float", "-123.45", -123.45, false},
		{"invalid string", "abc", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := StringToFloat64(test.input)
			if test.shouldErr && err == nil {
				t.Errorf("StringToFloat64(%q) expected error but got none", test.input)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("StringToFloat64(%q) unexpected error: %v", test.input, err)
			}
			if !test.shouldErr && result != test.expected {
				t.Errorf("StringToFloat64(%q) = %f; want %f", test.input, result, test.expected)
			}
		})
	}
}

func TestStringToBool(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  bool
		shouldErr bool
	}{
		{"true", "true", true, false},
		{"false", "false", false, false},
		{"1", "1", true, false},
		{"0", "0", false, false},
		{"invalid", "abc", false, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := StringToBool(test.input)
			if test.shouldErr && err == nil {
				t.Errorf("StringToBool(%q) expected error but got none", test.input)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("StringToBool(%q) unexpected error: %v", test.input, err)
			}
			if !test.shouldErr && result != test.expected {
				t.Errorf("StringToBool(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"positive", 123, "123"},
		{"negative", -123, "-123"},
		{"zero", 0, "0"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntToString(test.input)
			if result != test.expected {
				t.Errorf("IntToString(%d) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

func TestFloat64ToString(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		prec     int
		expected string
	}{
		{"two decimals", 123.456, 2, "123.46"},
		{"zero decimals", 123.456, 0, "123"},
		{"five decimals", 123.456789, 5, "123.45679"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Float64ToString(test.input, test.prec)
			if result != test.expected {
				t.Errorf("Float64ToString(%f, %d) = %q; want %q", test.input, test.prec, result, test.expected)
			}
		})
	}
}

func TestBoolToString(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected string
	}{
		{"true", true, "true"},
		{"false", false, "false"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BoolToString(test.input)
			if result != test.expected {
				t.Errorf("BoolToString(%v) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		expected  int
		shouldErr bool
	}{
		{"int", 123, 123, false},
		{"int64", int64(123), 123, false},
		{"float64", 123.45, 123, false},
		{"string", "123", 123, false},
		{"invalid", "abc", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ToInt(test.input)
			if test.shouldErr && err == nil {
				t.Errorf("ToInt(%v) expected error but got none", test.input)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("ToInt(%v) unexpected error: %v", test.input, err)
			}
			if !test.shouldErr && result != test.expected {
				t.Errorf("ToInt(%v) = %d; want %d", test.input, result, test.expected)
			}
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"int", 123, "123"},
		{"int64", int64(123), "123"},
		{"float64", 123.45, "123.45"},
		{"bool", true, "true"},
		{"string", "hello", "hello"},
		{"bytes", []byte("hello"), "hello"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToString(test.input)
			if result != test.expected {
				t.Errorf("ToString(%v) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

func TestDeepCopy(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	src := &User{Name: "John", Age: 30}
	var dst User

	err := DeepCopy(src, &dst)
	if err != nil {
		t.Errorf("DeepCopy(%v, &dst) unexpected error: %v", src, err)
	}

	if dst.Name != src.Name {
		t.Errorf("DeepCopy name = %q; want %q", dst.Name, src.Name)
	}
	if dst.Age != src.Age {
		t.Errorf("DeepCopy age = %d; want %d", dst.Age, src.Age)
	}

	// 修改源对象，目标对象不应受影响
	src.Name = "Jane"
	if dst.Name != "John" {
		t.Errorf("DeepCopy modified destination when source changed")
	}

	// 测试切片
	srcSlice := []int{1, 2, 3}
	var dstSlice []int
	err = DeepCopy(&srcSlice, &dstSlice)
	if err != nil {
		t.Errorf("DeepCopy slice unexpected error: %v", err)
	}
	if !reflect.DeepEqual(srcSlice, dstSlice) {
		t.Errorf("DeepCopy slice = %v; want %v", dstSlice, srcSlice)
	}
}

package configutils

import (
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()
	if config == nil {
		t.Errorf("NewConfig() = nil, want non-nil")
	}
	if config.data == nil {
		t.Errorf("NewConfig().data = nil, want non-nil")
	}
}

func TestSetAndGet(t *testing.T) {
	config := NewConfig()

	// 测试简单键
	// Test simple key
	config.Set("key", "value")
	value, exists := config.Get("key")
	if !exists {
		t.Errorf("Get() = false, want true")
	}
	if value != "value" {
		t.Errorf("Get() = %v, want 'value'", value)
	}

	// 测试嵌套键
	// Test nested key
	config.Set("database.host", "localhost")
	value, exists = config.Get("database.host")
	if !exists {
		t.Errorf("Get() = false, want true")
	}
	if value != "localhost" {
		t.Errorf("Get() = %v, want 'localhost'", value)
	}

	// 测试不存在的键
	// Test non-existent key
	_, exists = config.Get("nonexistent")
	if exists {
		t.Errorf("Get() = true, want false")
	}
}

func TestGetString(t *testing.T) {
	config := NewConfig()

	// 测试存在的键
	// Test existing key
	config.Set("name", "test")
	result := config.GetString("name", "default")
	if result != "test" {
		t.Errorf("GetString() = %v, want 'test'", result)
	}

	// 测试不存在的键
	// Test non-existent key
	result = config.GetString("nonexistent", "default")
	if result != "default" {
		t.Errorf("GetString() = %v, want 'default'", result)
	}

	// 测试类型转换
	// Test type conversion
	config.Set("number", 123)
	result = config.GetString("number", "default")
	if result != "123" {
		t.Errorf("GetString() = %v, want '123'", result)
	}
}

func TestGetInt(t *testing.T) {
	config := NewConfig()

	// 测试整数
	// Test integer
	config.Set("port", int64(3306))
	result := config.GetInt("port", 8080)
	if result != 3306 {
		t.Errorf("GetInt() = %v, want 3306", result)
	}

	// 测试浮点数转换
	// Test float conversion
	config.Set("port2", 3306.0)
	result = config.GetInt("port2", 8080)
	if result != 3306 {
		t.Errorf("GetInt() = %v, want 3306", result)
	}

	// 测试字符串转换
	// Test string conversion
	config.Set("port3", "3306")
	result = config.GetInt("port3", 8080)
	if result != 3306 {
		t.Errorf("GetInt() = %v, want 3306", result)
	}

	// 测试默认值
	// Test default value
	result = config.GetInt("nonexistent", 8080)
	if result != 8080 {
		t.Errorf("GetInt() = %v, want 8080", result)
	}
}

func TestGetFloat(t *testing.T) {
	config := NewConfig()

	// 测试浮点数
	// Test float
	config.Set("ratio", 0.5)
	result := config.GetFloat("ratio", 0.0)
	if result != 0.5 {
		t.Errorf("GetFloat() = %v, want 0.5", result)
	}

	// 测试整数转换
	// Test integer conversion
	config.Set("ratio2", 1)
	result = config.GetFloat("ratio2", 0.0)
	if result != 1.0 {
		t.Errorf("GetFloat() = %v, want 1.0", result)
	}

	// 测试默认值
	// Test default value
	result = config.GetFloat("nonexistent", 0.0)
	if result != 0.0 {
		t.Errorf("GetFloat() = %v, want 0.0", result)
	}
}

func TestGetBool(t *testing.T) {
	config := NewConfig()

	// 测试布尔值
	// Test boolean
	config.Set("debug", true)
	result := config.GetBool("debug", false)
	if result != true {
		t.Errorf("GetBool() = %v, want true", result)
	}

	// 测试字符串转换
	// Test string conversion
	config.Set("enabled", "true")
	result = config.GetBool("enabled", false)
	if result != true {
		t.Errorf("GetBool() = %v, want true", result)
	}

	// 测试默认值
	// Test default value
	result = config.GetBool("nonexistent", false)
	if result != false {
		t.Errorf("GetBool() = %v, want false", result)
	}
}

func TestGetStringSlice(t *testing.T) {
	config := NewConfig()

	// 测试切片
	// Test slice
	config.Set("hosts", []interface{}{"host1", "host2", "host3"})
	result := config.GetStringSlice("hosts", []string{})
	if len(result) != 3 {
		t.Errorf("GetStringSlice() length = %v, want 3", len(result))
	}

	// 测试逗号分隔的字符串
	// Test comma-separated string
	config.Set("hosts2", "host1,host2,host3")
	result = config.GetStringSlice("hosts2", []string{})
	if len(result) != 3 {
		t.Errorf("GetStringSlice() length = %v, want 3", len(result))
	}

	// 测试默认值
	// Test default value
	result = config.GetStringSlice("nonexistent", []string{"default"})
	if len(result) != 1 || result[0] != "default" {
		t.Errorf("GetStringSlice() = %v, want ['default']", result)
	}
}

func TestHas(t *testing.T) {
	config := NewConfig()

	config.Set("key", "value")
	if !config.Has("key") {
		t.Errorf("Has() = false, want true")
	}

	if config.Has("nonexistent") {
		t.Errorf("Has() = true, want false")
	}
}

func TestLoadFromJSONString(t *testing.T) {
	config := NewConfig()
	jsonStr := `{"name":"test","age":30,"database":{"host":"localhost","port":3306}}`

	err := config.LoadFromJSONString(jsonStr)
	if err != nil {
		t.Errorf("LoadFromJSONString() error = %v, want nil", err)
	}

	if config.GetString("name", "") != "test" {
		t.Errorf("GetString('name') = %v, want 'test'", config.GetString("name", ""))
	}

	if config.GetInt("age", 0) != 30 {
		t.Errorf("GetInt('age') = %v, want 30", config.GetInt("age", 0))
	}

	if config.GetString("database.host", "") != "localhost" {
		t.Errorf("GetString('database.host') = %v, want 'localhost'", config.GetString("database.host", ""))
	}
}

func TestLoadFromEnv(t *testing.T) {
	// 设置测试环境变量
	// Set test environment variables
	os.Setenv("TEST_CONFIG_NAME", "test")
	os.Setenv("TEST_CONFIG_PORT", "3306")
	os.Setenv("TEST_CONFIG_DEBUG", "true")
	defer func() {
		os.Unsetenv("TEST_CONFIG_NAME")
		os.Unsetenv("TEST_CONFIG_PORT")
		os.Unsetenv("TEST_CONFIG_DEBUG")
	}()

	config := NewConfig()
	config.LoadFromEnv("TEST_CONFIG_")

	if config.GetString("name", "") != "test" {
		t.Errorf("GetString('name') = %v, want 'test'", config.GetString("name", ""))
	}

	if config.GetInt("port", 0) != 3306 {
		t.Errorf("GetInt('port') = %v, want 3306", config.GetInt("port", 0))
	}

	if !config.GetBool("debug", false) {
		t.Errorf("GetBool('debug') = %v, want true", config.GetBool("debug", false))
	}
}

func TestMerge(t *testing.T) {
	config1 := NewConfig()
	config1.Set("key1", "value1")
	config1.Set("nested.key", "value1")

	config2 := NewConfig()
	config2.Set("key2", "value2")
	config2.Set("nested.key", "value2")

	config1.Merge(config2)

	if config1.GetString("key1", "") != "value1" {
		t.Errorf("Merge() key1 = %v, want 'value1'", config1.GetString("key1", ""))
	}

	if config1.GetString("key2", "") != "value2" {
		t.Errorf("Merge() key2 = %v, want 'value2'", config1.GetString("key2", ""))
	}

	// 合并时，config2的值会覆盖config1的值
	// When merging, config2's value will overwrite config1's value
	if config1.GetString("nested.key", "") != "value2" {
		t.Errorf("Merge() nested.key = %v, want 'value2'", config1.GetString("nested.key", ""))
	}
}

func TestSetDefaults(t *testing.T) {
	config := NewConfig()

	// 设置一个已存在的值
	// Set an existing value
	config.Set("key1", "existing")

	// 设置默认值
	// Set defaults
	defaults := map[string]interface{}{
		"key1": "default1",
		"key2": "default2",
	}
	config.SetDefaults(defaults)

	// 已存在的值不应被覆盖
	// Existing value should not be overwritten
	if config.GetString("key1", "") != "existing" {
		t.Errorf("SetDefaults() key1 = %v, want 'existing'", config.GetString("key1", ""))
	}

	// 不存在的值应被设置
	// Non-existent value should be set
	if config.GetString("key2", "") != "default2" {
		t.Errorf("SetDefaults() key2 = %v, want 'default2'", config.GetString("key2", ""))
	}
}

func TestValidate(t *testing.T) {
	config := NewConfig()
	config.Set("port", int64(3306))

	// 测试验证通过
	// Test validation pass
	err := config.Validate("port", func(v interface{}) bool {
		port, ok := v.(int64)
		return ok && port > 0 && port < 65536
	})
	if err != nil {
		t.Errorf("Validate() error = %v, want nil", err)
	}

	// 测试验证失败
	// Test validation fail
	err = config.Validate("port", func(v interface{}) bool {
		return false
	})
	if err == nil {
		t.Errorf("Validate() error = nil, want non-nil")
	}

	// 测试不存在的键
	// Test non-existent key
	err = config.Validate("nonexistent", func(v interface{}) bool {
		return true
	})
	if err == nil {
		t.Errorf("Validate() error = nil, want non-nil")
	}
}

func TestUnmarshal(t *testing.T) {
	config := NewConfig()
	jsonStr := `{"name":"test","age":30}`
	config.LoadFromJSONString(jsonStr)

	type TestConfig struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var cfg TestConfig
	err := config.Unmarshal(&cfg)
	if err != nil {
		t.Errorf("Unmarshal() error = %v, want nil", err)
	}

	if cfg.Name != "test" {
		t.Errorf("Unmarshal() Name = %v, want 'test'", cfg.Name)
	}

	if cfg.Age != 30 {
		t.Errorf("Unmarshal() Age = %v, want 30", cfg.Age)
	}
}

func TestKeys(t *testing.T) {
	config := NewConfig()
	config.Set("key1", "value1")
	config.Set("key2", "value2")
	config.Set("nested.key", "value")

	keys := config.Keys()
	if len(keys) != 3 {
		t.Errorf("Keys() length = %v, want 3", len(keys))
	}

	// 检查是否包含所有键
	// Check if all keys are included
	expectedKeys := map[string]bool{
		"key1":      true,
		"key2":      true,
		"nested.key": true,
	}

	for _, key := range keys {
		if !expectedKeys[key] {
			t.Errorf("Keys() contains unexpected key: %v", key)
		}
	}
}

func TestClear(t *testing.T) {
	config := NewConfig()
	config.Set("key", "value")

	config.Clear()

	if config.Has("key") {
		t.Errorf("Clear() key still exists")
	}

	if len(config.Keys()) != 0 {
		t.Errorf("Clear() keys length = %v, want 0", len(config.Keys()))
	}
}

func TestAll(t *testing.T) {
	config := NewConfig()
	config.Set("key1", "value1")
	config.Set("key2", "value2")

	all := config.All()

	if len(all) != 2 {
		t.Errorf("All() length = %v, want 2", len(all))
	}

	// 修改返回的map不应影响原始配置
	// Modifying returned map should not affect original config
	all["key3"] = "value3"

	if config.Has("key3") {
		t.Errorf("All() modification affected original config")
	}
}

func TestLoadConfigFromJSON(t *testing.T) {
	// 创建临时JSON文件
	// Create temporary JSON file
	tmpFile := "/tmp/test_config.json"
	jsonContent := `{"name":"test","port":3306}`
	err := os.WriteFile(tmpFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile)

	config, err := LoadConfigFromJSON(tmpFile)
	if err != nil {
		t.Errorf("LoadConfigFromJSON() error = %v, want nil", err)
	}

	if config.GetString("name", "") != "test" {
		t.Errorf("LoadConfigFromJSON() name = %v, want 'test'", config.GetString("name", ""))
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	// 设置测试环境变量
	// Set test environment variables
	os.Setenv("APP_NAME", "testapp")
	os.Setenv("APP_PORT", "8080")
	defer func() {
		os.Unsetenv("APP_NAME")
		os.Unsetenv("APP_PORT")
	}()

	config := LoadConfigFromEnv("APP_")

	if config.GetString("name", "") != "testapp" {
		t.Errorf("LoadConfigFromEnv() name = %v, want 'testapp'", config.GetString("name", ""))
	}

	if config.GetInt("port", 0) != 8080 {
		t.Errorf("LoadConfigFromEnv() port = %v, want 8080", config.GetInt("port", 0))
	}
}

func TestNestedConfig(t *testing.T) {
	config := NewConfig()

	// 设置嵌套配置
	// Set nested configuration
	config.Set("app.name", "myapp")
	config.Set("app.version", "1.0.0")
	config.Set("database.host", "localhost")
	config.Set("database.port", 3306)

	// 测试获取嵌套值
	// Test getting nested values
	if config.GetString("app.name", "") != "myapp" {
		t.Errorf("GetString('app.name') = %v, want 'myapp'", config.GetString("app.name", ""))
	}

	if config.GetString("database.host", "") != "localhost" {
		t.Errorf("GetString('database.host') = %v, want 'localhost'", config.GetString("database.host", ""))
	}

	if config.GetInt("database.port", 0) != 3306 {
		t.Errorf("GetInt('database.port') = %v, want 3306", config.GetInt("database.port", 0))
	}
}

func TestParseValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{"boolean true", "true", true},
		{"boolean false", "false", false},
		{"boolean TRUE", "TRUE", true},
		{"boolean FALSE", "FALSE", false},
		{"integer", "123", int64(123)},
		{"float", "123.45", 123.45},
		{"string", "hello", "hello"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseValue(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseValue(%q) = %v (%T), want %v (%T)", tt.input, result, result, tt.expected, tt.expected)
			}
		})
	}
}


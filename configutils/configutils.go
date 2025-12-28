// Package configutils 提供配置管理相关的工具函数
// Package configutils provides configuration management utility functions
package configutils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config 配置接口，用于存储和管理配置
// Config interface for storing and managing configuration
type Config struct {
	data map[string]interface{}
}

// NewConfig 创建新的配置对象
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - *Config: 新的配置对象 / new config object
//
// 示例 / Example:
//   config := NewConfig()
//
// NewConfig creates a new config object
func NewConfig() *Config {
	return &Config{
		data: make(map[string]interface{}),
	}
}

// LoadFromJSON 从JSON文件加载配置
//
// 参数 / Parameters:
//   - filepath: JSON文件路径 / JSON file path
//
// 返回值 / Returns:
//   - error: 如果加载失败则返回错误 / error if loading fails
//
// 示例 / Example:
//   err := config.LoadFromJSON("config.json")
//
// LoadFromJSON loads configuration from a JSON file
func (c *Config) LoadFromJSON(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("解析JSON失败: %w", err)
	}

	c.data = jsonData
	return nil
}

// LoadFromJSONString 从JSON字符串加载配置
//
// 参数 / Parameters:
//   - jsonStr: JSON字符串 / JSON string
//
// 返回值 / Returns:
//   - error: 如果加载失败则返回错误 / error if loading fails
//
// 示例 / Example:
//   err := config.LoadFromJSONString(`{"key":"value"}`)
//
// LoadFromJSONString loads configuration from a JSON string
func (c *Config) LoadFromJSONString(jsonStr string) error {
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		return fmt.Errorf("解析JSON失败: %w", err)
	}

	c.data = jsonData
	return nil
}

// LoadFromEnv 从环境变量加载配置
//
// 参数 / Parameters:
//   - prefix: 环境变量前缀，如果为空则加载所有环境变量 / environment variable prefix, empty to load all
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   config.LoadFromEnv("APP_")
//
// LoadFromEnv loads configuration from environment variables
func (c *Config) LoadFromEnv(prefix string) {
	envVars := os.Environ()
	for _, env := range envVars {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		// 如果指定了前缀，只加载匹配的环境变量
		// If prefix is specified, only load matching environment variables
		if prefix != "" && !strings.HasPrefix(key, prefix) {
			continue
		}

		// 移除前缀
		// Remove prefix
		if prefix != "" {
			key = strings.TrimPrefix(key, prefix)
		}

		// 将键名转换为小写，并用点号分隔层级
		// Convert key to lowercase and use dots to separate levels
		key = strings.ToLower(key)
		key = strings.ReplaceAll(key, "_", ".")

		// 尝试解析为数字或布尔值
		// Try to parse as number or boolean
		parsedValue := parseValue(value)
		c.Set(key, parsedValue)
	}
}

// parseValue 尝试将字符串值解析为适当的类型
// parseValue attempts to parse a string value into an appropriate type
func parseValue(value string) interface{} {
	// 尝试解析为布尔值
	// Try to parse as boolean
	if value == "true" || value == "TRUE" {
		return true
	}
	if value == "false" || value == "FALSE" {
		return false
	}

	// 尝试解析为整数
	// Try to parse as integer
	if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
		return intVal
	}

	// 尝试解析为浮点数
	// Try to parse as float
	if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
		return floatVal
	}

	// 返回原始字符串
	// Return original string
	return value
}

// Set 设置配置值
//
// 参数 / Parameters:
//   - key: 配置键，支持点号分隔的嵌套键（如 "database.host"） / config key, supports dot-separated nested keys
//   - value: 配置值 / config value
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   config.Set("database.host", "localhost")
//
// Set sets a configuration value
func (c *Config) Set(key string, value interface{}) {
	keys := strings.Split(key, ".")
	c.setNested(keys, value, c.data)
}

// setNested 递归设置嵌套配置值
// setNested recursively sets nested configuration values
func (c *Config) setNested(keys []string, value interface{}, data map[string]interface{}) {
	if len(keys) == 1 {
		data[keys[0]] = value
		return
	}

	key := keys[0]
	if _, exists := data[key]; !exists {
		data[key] = make(map[string]interface{})
	}

	if nested, ok := data[key].(map[string]interface{}); ok {
		c.setNested(keys[1:], value, nested)
	} else {
		// 如果已存在但不是map，则覆盖
		// If exists but not a map, overwrite
		data[key] = make(map[string]interface{})
		c.setNested(keys[1:], value, data[key].(map[string]interface{}))
	}
}

// Get 获取配置值
//
// 参数 / Parameters:
//   - key: 配置键，支持点号分隔的嵌套键 / config key, supports dot-separated nested keys
//
// 返回值 / Returns:
//   - interface{}: 配置值，如果不存在则返回nil / config value, nil if not exists
//   - bool: 是否存在 / whether the key exists
//
// 示例 / Example:
//   value, exists := config.Get("database.host")
//
// Get gets a configuration value
func (c *Config) Get(key string) (interface{}, bool) {
	keys := strings.Split(key, ".")
	return c.getNested(keys, c.data)
}

// getNested 递归获取嵌套配置值
// getNested recursively gets nested configuration values
func (c *Config) getNested(keys []string, data map[string]interface{}) (interface{}, bool) {
	if len(keys) == 0 {
		return nil, false
	}

	key := keys[0]
	value, exists := data[key]
	if !exists {
		return nil, false
	}

	if len(keys) == 1 {
		return value, true
	}

	if nested, ok := value.(map[string]interface{}); ok {
		return c.getNested(keys[1:], nested)
	}

	return nil, false
}

// GetString 获取字符串配置值
//
// 参数 / Parameters:
//   - key: 配置键 / config key
//   - defaultValue: 默认值，如果不存在则返回此值 / default value if key doesn't exist
//
// 返回值 / Returns:
//   - string: 配置值或默认值 / config value or default value
//
// 示例 / Example:
//   host := config.GetString("database.host", "localhost")
//
// GetString gets a string configuration value
func (c *Config) GetString(key string, defaultValue string) string {
	value, exists := c.Get(key)
	if !exists {
		return defaultValue
	}

	if str, ok := value.(string); ok {
		return str
	}

	return fmt.Sprintf("%v", value)
}

// GetInt 获取整数配置值
//
// 参数 / Parameters:
//   - key: 配置键 / config key
//   - defaultValue: 默认值，如果不存在或类型不匹配则返回此值 / default value if key doesn't exist or type mismatch
//
// 返回值 / Returns:
//   - int64: 配置值或默认值 / config value or default value
//
// 示例 / Example:
//   port := config.GetInt("database.port", 3306)
//
// GetInt gets an integer configuration value
func (c *Config) GetInt(key string, defaultValue int64) int64 {
	value, exists := c.Get(key)
	if !exists {
		return defaultValue
	}

	switch v := value.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		if intVal, err := strconv.ParseInt(v, 10, 64); err == nil {
			return intVal
		}
	}

	return defaultValue
}

// GetFloat 获取浮点数配置值
//
// 参数 / Parameters:
//   - key: 配置键 / config key
//   - defaultValue: 默认值，如果不存在或类型不匹配则返回此值 / default value if key doesn't exist or type mismatch
//
// 返回值 / Returns:
//   - float64: 配置值或默认值 / config value or default value
//
// 示例 / Example:
//   ratio := config.GetFloat("app.ratio", 0.5)
//
// GetFloat gets a float configuration value
func (c *Config) GetFloat(key string, defaultValue float64) float64 {
	value, exists := c.Get(key)
	if !exists {
		return defaultValue
	}

	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int64:
		return float64(v)
	case int:
		return float64(v)
	case string:
		if floatVal, err := strconv.ParseFloat(v, 64); err == nil {
			return floatVal
		}
	}

	return defaultValue
}

// GetBool 获取布尔配置值
//
// 参数 / Parameters:
//   - key: 配置键 / config key
//   - defaultValue: 默认值，如果不存在或类型不匹配则返回此值 / default value if key doesn't exist or type mismatch
//
// 返回值 / Returns:
//   - bool: 配置值或默认值 / config value or default value
//
// 示例 / Example:
//   debug := config.GetBool("app.debug", false)
//
// GetBool gets a boolean configuration value
func (c *Config) GetBool(key string, defaultValue bool) bool {
	value, exists := c.Get(key)
	if !exists {
		return defaultValue
	}

	switch v := value.(type) {
	case bool:
		return v
	case string:
		if boolVal, err := strconv.ParseBool(v); err == nil {
			return boolVal
		}
	}

	return defaultValue
}

// GetStringSlice 获取字符串切片配置值
//
// 参数 / Parameters:
//   - key: 配置键 / config key
//   - defaultValue: 默认值，如果不存在或类型不匹配则返回此值 / default value if key doesn't exist or type mismatch
//
// 返回值 / Returns:
//   - []string: 配置值或默认值 / config value or default value
//
// 示例 / Example:
//   hosts := config.GetStringSlice("database.hosts", []string{"localhost"})
//
// GetStringSlice gets a string slice configuration value
func (c *Config) GetStringSlice(key string, defaultValue []string) []string {
	value, exists := c.Get(key)
	if !exists {
		return defaultValue
	}

	// 尝试转换为字符串切片
	// Try to convert to string slice
	if slice, ok := value.([]interface{}); ok {
		result := make([]string, 0, len(slice))
		for _, item := range slice {
			if str, ok := item.(string); ok {
				result = append(result, str)
			} else {
				result = append(result, fmt.Sprintf("%v", item))
			}
		}
		return result
	}

	// 如果是字符串，尝试按逗号分割
	// If it's a string, try to split by comma
	if str, ok := value.(string); ok {
		if str == "" {
			return defaultValue
		}
		parts := strings.Split(str, ",")
		result := make([]string, len(parts))
		for i, part := range parts {
			result[i] = strings.TrimSpace(part)
		}
		return result
	}

	return defaultValue
}

// Has 检查配置键是否存在
//
// 参数 / Parameters:
//   - key: 配置键 / config key
//
// 返回值 / Returns:
//   - bool: 如果存在则返回true / true if key exists
//
// 示例 / Example:
//   if config.Has("database.host") { ... }
//
// Has checks if a configuration key exists
func (c *Config) Has(key string) bool {
	_, exists := c.Get(key)
	return exists
}

// Unmarshal 将配置解析到结构体
//
// 参数 / Parameters:
//   - v: 目标结构体指针 / target struct pointer
//
// 返回值 / Returns:
//   - error: 如果解析失败则返回错误 / error if unmarshaling fails
//
// 示例 / Example:
//   type AppConfig struct {
//       Database struct {
//           Host string `json:"host"`
//           Port int    `json:"port"`
//       } `json:"database"`
//   }
//   var cfg AppConfig
//   err := config.Unmarshal(&cfg)
//
// Unmarshal unmarshals configuration into a struct
func (c *Config) Unmarshal(v interface{}) error {
	// 将配置数据转换为JSON，然后解析到结构体
	// Convert config data to JSON, then unmarshal into struct
	jsonData, err := json.Marshal(c.data)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %w", err)
	}

	if err := json.Unmarshal(jsonData, v); err != nil {
		return fmt.Errorf("解析配置到结构体失败: %w", err)
	}

	return nil
}

// Merge 合并另一个配置对象
//
// 参数 / Parameters:
//   - other: 要合并的配置对象 / config object to merge
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   config.Merge(otherConfig)
//
// Merge merges another config object
func (c *Config) Merge(other *Config) {
	c.mergeMaps(c.data, other.data)
}

// mergeMaps 递归合并两个map
// mergeMaps recursively merges two maps
func (c *Config) mergeMaps(dest, src map[string]interface{}) {
	for key, value := range src {
		if existing, exists := dest[key]; exists {
			// 如果两个值都是map，递归合并
			// If both values are maps, merge recursively
			if destMap, ok := existing.(map[string]interface{}); ok {
				if srcMap, ok := value.(map[string]interface{}); ok {
					c.mergeMaps(destMap, srcMap)
					continue
				}
			}
		}
		// 否则直接覆盖
		// Otherwise, overwrite directly
		dest[key] = value
	}
}

// SetDefaults 设置默认值
//
// 参数 / Parameters:
//   - defaults: 默认配置的map / map of default configuration
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   config.SetDefaults(map[string]interface{}{
//       "database.host": "localhost",
//       "database.port": 3306,
//   })
//
// SetDefaults sets default values for configuration
func (c *Config) SetDefaults(defaults map[string]interface{}) {
	for key, value := range defaults {
		if !c.Has(key) {
			c.Set(key, value)
		}
	}
}

// Validate 验证配置值
//
// 参数 / Parameters:
//   - key: 配置键 / config key
//   - validator: 验证函数，返回true表示验证通过 / validator function, returns true if valid
//
// 返回值 / Returns:
//   - error: 如果验证失败则返回错误 / error if validation fails
//
// 示例 / Example:
//   err := config.Validate("database.port", func(v interface{}) bool {
//       port, ok := v.(float64)
//       return ok && port > 0 && port < 65536
//   })
//
// Validate validates a configuration value
func (c *Config) Validate(key string, validator func(interface{}) bool) error {
	value, exists := c.Get(key)
	if !exists {
		return fmt.Errorf("配置键不存在: %s", key)
	}

	if !validator(value) {
		return fmt.Errorf("配置值验证失败: %s", key)
	}

	return nil
}

// All 获取所有配置数据
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - map[string]interface{}: 所有配置数据的副本 / copy of all configuration data
//
// 示例 / Example:
//   all := config.All()
//
// All returns all configuration data
func (c *Config) All() map[string]interface{} {
	return c.deepCopy(c.data).(map[string]interface{})
}

// deepCopy 深拷贝配置数据
// deepCopy performs deep copy of configuration data
func (c *Config) deepCopy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	switch v := src.(type) {
	case map[string]interface{}:
		dst := make(map[string]interface{})
		for k, val := range v {
			dst[k] = c.deepCopy(val)
		}
		return dst
	case []interface{}:
		dst := make([]interface{}, len(v))
		for i, val := range v {
			dst[i] = c.deepCopy(val)
		}
		return dst
	default:
		return v
	}
}

// Clear 清空所有配置
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - 无 / none
//
// 示例 / Example:
//   config.Clear()
//
// Clear clears all configuration
func (c *Config) Clear() {
	c.data = make(map[string]interface{})
}

// Keys 获取所有配置键
//
// 参数 / Parameters:
//   - 无 / none
//
// 返回值 / Returns:
//   - []string: 所有配置键的列表 / list of all configuration keys
//
// 示例 / Example:
//   keys := config.Keys()
//
// Keys returns all configuration keys
func (c *Config) Keys() []string {
	return c.getKeys("", c.data)
}

// getKeys 递归获取所有键
// getKeys recursively gets all keys
func (c *Config) getKeys(prefix string, data map[string]interface{}) []string {
	var keys []string
	for key, value := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		if nested, ok := value.(map[string]interface{}); ok {
			keys = append(keys, c.getKeys(fullKey, nested)...)
		} else {
			keys = append(keys, fullKey)
		}
	}
	return keys
}

// LoadConfigFromJSON 从JSON文件创建配置对象（便捷函数）
//
// 参数 / Parameters:
//   - filepath: JSON文件路径 / JSON file path
//
// 返回值 / Returns:
//   - *Config: 配置对象 / config object
//   - error: 如果加载失败则返回错误 / error if loading fails
//
// 示例 / Example:
//   config, err := LoadConfigFromJSON("config.json")
//
// LoadConfigFromJSON creates a config object from a JSON file (convenience function)
func LoadConfigFromJSON(filepath string) (*Config, error) {
	config := NewConfig()
	if err := config.LoadFromJSON(filepath); err != nil {
		return nil, err
	}
	return config, nil
}

// LoadConfigFromEnv 从环境变量创建配置对象（便捷函数）
//
// 参数 / Parameters:
//   - prefix: 环境变量前缀，如果为空则加载所有环境变量 / environment variable prefix, empty to load all
//
// 返回值 / Returns:
//   - *Config: 配置对象 / config object
//
// 示例 / Example:
//   config := LoadConfigFromEnv("APP_")
//
// LoadConfigFromEnv creates a config object from environment variables (convenience function)
func LoadConfigFromEnv(prefix string) *Config {
	config := NewConfig()
	config.LoadFromEnv(prefix)
	return config
}


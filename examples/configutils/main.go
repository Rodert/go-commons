package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Rodert/go-commons/configutils"
)

func main() {
	fmt.Println("=== Config Utils Examples ===\n")

	// 示例1: 创建配置对象并设置值
	// Example 1: Create config object and set values
	fmt.Println("1. Basic Configuration:")
	config := configutils.NewConfig()
	config.Set("app.name", "MyApp")
	config.Set("app.version", "1.0.0")
	config.Set("database.host", "localhost")
	config.Set("database.port", 3306)
	config.Set("database.debug", true)

	fmt.Printf("   App Name: %s\n", config.GetString("app.name", ""))
	fmt.Printf("   App Version: %s\n", config.GetString("app.version", ""))
	fmt.Printf("   Database Host: %s\n", config.GetString("database.host", ""))
	fmt.Printf("   Database Port: %d\n", config.GetInt("database.port", 0))
	fmt.Printf("   Database Debug: %v\n\n", config.GetBool("database.debug", false))

	// 示例2: 从JSON字符串加载配置
	// Example 2: Load configuration from JSON string
	fmt.Println("2. Load from JSON String:")
	jsonStr := `{
		"server": {
			"host": "0.0.0.0",
			"port": 8080,
			"timeout": 30
		},
		"features": {
			"cache": true,
			"logging": true
		},
		"allowed_hosts": ["localhost", "127.0.0.1", "example.com"]
	}`
	
	jsonConfig := configutils.NewConfig()
	if err := jsonConfig.LoadFromJSONString(jsonStr); err != nil {
		fmt.Printf("   Error: %v\n\n", err)
	} else {
		fmt.Printf("   Server Host: %s\n", jsonConfig.GetString("server.host", ""))
		fmt.Printf("   Server Port: %d\n", jsonConfig.GetInt("server.port", 0))
		fmt.Printf("   Server Timeout: %d\n", jsonConfig.GetInt("server.timeout", 0))
		fmt.Printf("   Cache Enabled: %v\n", jsonConfig.GetBool("features.cache", false))
		hosts := jsonConfig.GetStringSlice("allowed_hosts", []string{})
		fmt.Printf("   Allowed Hosts: %v\n\n", hosts)
	}

	// 示例3: 从JSON文件加载配置
	// Example 3: Load configuration from JSON file
	fmt.Println("3. Load from JSON File:")
	tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("config_%d.json", time.Now().Unix()))
	jsonContent := `{
		"app": {
			"name": "FileConfigApp",
			"env": "development"
		},
		"database": {
			"host": "db.example.com",
			"port": 5432
		}
	}`
	
	if err := os.WriteFile(tmpFile, []byte(jsonContent), 0644); err != nil {
		fmt.Printf("   Error creating temp file: %v\n\n", err)
	} else {
		defer os.Remove(tmpFile)
		
		fileConfig, err := configutils.LoadConfigFromJSON(tmpFile)
		if err != nil {
			fmt.Printf("   Error: %v\n\n", err)
		} else {
			fmt.Printf("   App Name: %s\n", fileConfig.GetString("app.name", ""))
			fmt.Printf("   App Env: %s\n", fileConfig.GetString("app.env", ""))
			fmt.Printf("   Database Host: %s\n", fileConfig.GetString("database.host", ""))
			fmt.Printf("   Database Port: %d\n\n", fileConfig.GetInt("database.port", 0))
		}
	}

	// 示例4: 从环境变量加载配置
	// Example 4: Load configuration from environment variables
	fmt.Println("4. Load from Environment Variables:")
	// 设置一些测试环境变量
	// Set some test environment variables
	os.Setenv("APP_CONFIG_NAME", "EnvApp")
	os.Setenv("APP_CONFIG_PORT", "9090")
	os.Setenv("APP_CONFIG_DEBUG", "true")
	os.Setenv("APP_CONFIG_RATIO", "0.85")
	defer func() {
		os.Unsetenv("APP_CONFIG_NAME")
		os.Unsetenv("APP_CONFIG_PORT")
		os.Unsetenv("APP_CONFIG_DEBUG")
		os.Unsetenv("APP_CONFIG_RATIO")
	}()

	envConfig := configutils.LoadConfigFromEnv("APP_CONFIG_")
	fmt.Printf("   Name: %s\n", envConfig.GetString("name", ""))
	fmt.Printf("   Port: %d\n", envConfig.GetInt("port", 0))
	fmt.Printf("   Debug: %v\n", envConfig.GetBool("debug", false))
	fmt.Printf("   Ratio: %.2f\n\n", envConfig.GetFloat("ratio", 0.0))

	// 示例5: 设置默认值
	// Example 5: Set default values
	fmt.Println("5. Set Default Values:")
	defaultConfig := configutils.NewConfig()
	defaultConfig.Set("app.name", "CustomApp") // 这个值不会被默认值覆盖
	// This value won't be overwritten by defaults
	
	defaults := map[string]interface{}{
		"app.name":    "DefaultApp",
		"app.version": "1.0.0",
		"app.port":    8080,
		"app.debug":   false,
	}
	defaultConfig.SetDefaults(defaults)
	
	fmt.Printf("   App Name (custom, not overwritten): %s\n", defaultConfig.GetString("app.name", ""))
	fmt.Printf("   App Version (default): %s\n", defaultConfig.GetString("app.version", ""))
	fmt.Printf("   App Port (default): %d\n", defaultConfig.GetInt("app.port", 0))
	fmt.Printf("   App Debug (default): %v\n\n", defaultConfig.GetBool("app.debug", true))

	// 示例6: 配置验证
	// Example 6: Configuration validation
	fmt.Println("6. Configuration Validation:")
	validateConfig := configutils.NewConfig()
	validateConfig.Set("server.port", 8080)
	validateConfig.Set("server.timeout", 30)
	
	// 验证端口范围
	// Validate port range
	err := validateConfig.Validate("server.port", func(v interface{}) bool {
		port, ok := v.(int64)
		return ok && port > 0 && port < 65536
	})
	if err != nil {
		fmt.Printf("   Port validation failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Port validation passed\n")
	}
	
	// 验证超时值
	// Validate timeout value
	err = validateConfig.Validate("server.timeout", func(v interface{}) bool {
		timeout, ok := v.(int64)
		return ok && timeout > 0 && timeout <= 300
	})
	if err != nil {
		fmt.Printf("   Timeout validation failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Timeout validation passed\n")
	}
	fmt.Println()

	// 示例7: 配置合并
	// Example 7: Configuration merging
	fmt.Println("7. Configuration Merging:")
	config1 := configutils.NewConfig()
	config1.Set("app.name", "App1")
	config1.Set("database.host", "host1")
	config1.Set("shared.key", "value1")
	
	config2 := configutils.NewConfig()
	config2.Set("app.version", "2.0.0")
	config2.Set("database.port", 3306)
	config2.Set("shared.key", "value2") // 这个值会覆盖config1的值
	// This value will overwrite config1's value
	
	config1.Merge(config2)
	fmt.Printf("   App Name: %s\n", config1.GetString("app.name", ""))
	fmt.Printf("   App Version: %s\n", config1.GetString("app.version", ""))
	fmt.Printf("   Database Host: %s\n", config1.GetString("database.host", ""))
	fmt.Printf("   Database Port: %d\n", config1.GetInt("database.port", 0))
	fmt.Printf("   Shared Key (merged): %s\n\n", config1.GetString("shared.key", ""))

	// 示例8: 解析到结构体
	// Example 8: Unmarshal to struct
	fmt.Println("8. Unmarshal to Struct:")
	type ServerConfig struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		Timeout int    `json:"timeout"`
	}
	
	type AppConfig struct {
		Name   string        `json:"name"`
		Server ServerConfig  `json:"server"`
		Debug  bool          `json:"debug"`
	}
	
	structConfig := configutils.NewConfig()
	structConfig.LoadFromJSONString(`{
		"name": "StructApp",
		"server": {
			"host": "0.0.0.0",
			"port": 8080,
			"timeout": 30
		},
		"debug": true
	}`)
	
	var appCfg AppConfig
	if err := structConfig.Unmarshal(&appCfg); err != nil {
		fmt.Printf("   Error: %v\n\n", err)
	} else {
		fmt.Printf("   App Name: %s\n", appCfg.Name)
		fmt.Printf("   Server Host: %s\n", appCfg.Server.Host)
		fmt.Printf("   Server Port: %d\n", appCfg.Server.Port)
		fmt.Printf("   Server Timeout: %d\n", appCfg.Server.Timeout)
		fmt.Printf("   Debug: %v\n\n", appCfg.Debug)
	}

	// 示例9: 获取所有配置键
	// Example 9: Get all configuration keys
	fmt.Println("9. Get All Configuration Keys:")
	keysConfig := configutils.NewConfig()
	keysConfig.Set("app.name", "TestApp")
	keysConfig.Set("app.version", "1.0.0")
	keysConfig.Set("database.host", "localhost")
	keysConfig.Set("database.port", 3306)
	
	keys := keysConfig.Keys()
	fmt.Printf("   Total keys: %d\n", len(keys))
	for _, key := range keys {
		fmt.Printf("   - %s\n", key)
	}
	fmt.Println()

	// 示例10: 检查配置是否存在
	// Example 10: Check if configuration exists
	fmt.Println("10. Check Configuration Existence:")
	checkConfig := configutils.NewConfig()
	checkConfig.Set("existing.key", "value")
	
	if checkConfig.Has("existing.key") {
		fmt.Printf("   ✓ 'existing.key' exists\n")
	}
	
	if !checkConfig.Has("nonexistent.key") {
		fmt.Printf("   ✓ 'nonexistent.key' does not exist\n")
	}
}


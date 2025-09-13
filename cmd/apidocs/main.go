//go:build !noswagger
// +build !noswagger

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/Rodert/go-commons/docs" // 导入生成的docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Commons API
// @version 1.0
// @description Go Commons 是一个精简的Go实用工具库，专注于字符串处理和系统工具
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/Rodert/go-commons
// @contact.email support@example.com

// @license.name Unlicense
// @license.url https://github.com/Rodert/go-commons/blob/main/LICENSE

// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/static", "./docs/assets")

	// API路由组
	api := r.Group("/api/v1")
	{
		// 字符串工具API
		stringAPI := api.Group("/string")
		{
			stringAPI.GET("/isEmpty/:str", IsEmptyHandler)
			stringAPI.GET("/isNotEmpty/:str", IsNotEmptyHandler)
			stringAPI.GET("/reverse/:str", ReverseHandler)
			stringAPI.GET("/swapCase/:str", SwapCaseHandler)
			stringAPI.POST("/padCenter", PadCenterHandler)
		}

		// 系统工具API
		systemAPI := api.Group("/system")
		{
			systemAPI.GET("/cpu", GetCPUInfoHandler)
			systemAPI.GET("/memory", GetMemInfoHandler)
			systemAPI.GET("/disk", GetDiskInfoHandler)
		}
	}

	// 使用gin-swagger中间件提供API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 首页重定向到Swagger UI
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}

// IsEmptyHandler godoc
// @Summary 检查字符串是否为空
// @Description 检查提供的字符串是否为空
// @Tags 字符串工具
// @Accept json
// @Produce json
// @Param str path string true "要检查的字符串"
// @Success 200 {object} map[string]interface{}
// @Router /string/isEmpty/{str} [get]
func IsEmptyHandler(c *gin.Context) {
	str := c.Param("str")
	c.JSON(200, gin.H{
		"string":  str,
		"isEmpty": str == "",
	})
}

// IsNotEmptyHandler godoc
// @Summary 检查字符串是否非空
// @Description 检查提供的字符串是否非空
// @Tags 字符串工具
// @Accept json
// @Produce json
// @Param str path string true "要检查的字符串"
// @Success 200 {object} map[string]interface{}
// @Router /string/isNotEmpty/{str} [get]
func IsNotEmptyHandler(c *gin.Context) {
	str := c.Param("str")
	c.JSON(200, gin.H{
		"string":     str,
		"isNotEmpty": str != "",
	})
}

// ReverseHandler godoc
// @Summary 反转字符串
// @Description 反转提供的字符串
// @Tags 字符串工具
// @Accept json
// @Produce json
// @Param str path string true "要反转的字符串"
// @Success 200 {object} map[string]interface{}
// @Router /string/reverse/{str} [get]
func ReverseHandler(c *gin.Context) {
	str := c.Param("str")
	reversed := reverseString(str)
	c.JSON(200, gin.H{
		"original": str,
		"reversed": reversed,
	})
}

// SwapCaseHandler godoc
// @Summary 交换字符串大小写
// @Description 交换提供的字符串中字母的大小写
// @Tags 字符串工具
// @Accept json
// @Produce json
// @Param str path string true "要处理的字符串"
// @Success 200 {object} map[string]interface{}
// @Router /string/swapCase/{str} [get]
func SwapCaseHandler(c *gin.Context) {
	str := c.Param("str")
	swapped := swapCase(str)
	c.JSON(200, gin.H{
		"original": str,
		"swapped":  swapped,
	})
}

// PadCenterRequest 表示填充请求
type PadCenterRequest struct {
	Str    string `json:"str" binding:"required"`
	Size   int    `json:"size" binding:"required"`
	PadStr string `json:"padStr" binding:"required"`
}

// PadCenterHandler godoc
// @Summary 居中填充字符串
// @Description 在字符串两侧填充字符，使其居中
// @Tags 字符串工具
// @Accept json
// @Produce json
// @Param request body PadCenterRequest true "填充请求"
// @Success 200 {object} map[string]interface{}
// @Router /string/padCenter [post]
func PadCenterHandler(c *gin.Context) {
	var req PadCenterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	padded := padCenter(req.Str, req.Size, req.PadStr)
	c.JSON(200, gin.H{
		"original": req.Str,
		"padded":   padded,
		"size":     req.Size,
		"padStr":   req.PadStr,
	})
}

// GetCPUInfoHandler godoc
// @Summary 获取CPU信息
// @Description 获取CPU核心数、使用率百分比和负载平均值
// @Tags 系统工具
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /system/cpu [get]
func GetCPUInfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"cores":   4,
		"usage":   "25%",
		"load":    1.5,
		"message": "这是一个示例API，实际实现需要调用systemutils/cpuutils包",
	})
}

// GetMemInfoHandler godoc
// @Summary 获取内存信息
// @Description 获取总内存、可用内存和已用内存
// @Tags 系统工具
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /system/memory [get]
func GetMemInfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"total":     "16GB",
		"available": "8GB",
		"used":      "8GB",
		"message":   "这是一个示例API，实际实现需要调用systemutils/memutils包",
	})
}

// GetDiskInfoHandler godoc
// @Summary 获取磁盘信息
// @Description 获取磁盘空间信息，包括总空间、可用空间、已用空间和使用率
// @Tags 系统工具
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /system/disk [get]
func GetDiskInfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"total":     "500GB",
		"available": "300GB",
		"used":      "200GB",
		"usage":     "40%",
		"message":   "这是一个示例API，实际实现需要调用systemutils/diskutils包",
	})
}

// 辅助函数
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func swapCase(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = r - 'a' + 'A'
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = r - 'A' + 'a'
		}
	}
	return string(runes)
}

func padCenter(str string, size int, padStr string) string {
	strLen := len(str)
	if strLen >= size {
		return str
	}

	padsNeeded := size - strLen
	padLeft := padsNeeded / 2
	padRight := padsNeeded - padLeft

	result := ""
	for i := 0; i < padLeft; i++ {
		result += padStr
	}
	result += str
	for i := 0; i < padRight; i++ {
		result += padStr
	}

	return result
}

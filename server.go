package main

import (
	"gin_hello/auth"
	"gin_hello/config"
	"gin_hello/database"
	"gin_hello/kimi"
	"gin_hello/middle_ware"
	"gin_hello/models"
	"gin_hello/openai"
	"gin_hello/wechat"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func GinServer() *gin.Engine {
	ginServer := gin.Default()
	ginServer.Use(middle_ware.ErrorHandlingMiddleware())
	//ginServer.Use(middle_ware.AuthMiddleware())
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.NewSuccessResponse(
			map[string]interface{}{
				"message": "这是首页",
			},
		))
	})
	ginServer.POST("/send_wechat_msg/:to/:send_type/:msg", func(c *gin.Context) {
		wechat.SendWechatMsg(c)
	})
	ginServer.POST("/send_wechat_msg2", func(c *gin.Context) {
		wechat.SendWechatMsg2(c)
	})
	ginServer.POST("/received_wechat_msg", func(c *gin.Context) {
		wechat.ReceivedWechatMsg(c)
	})
	ginServer.POST("/kimi/single_chat", func(c *gin.Context) {
		kimi.Chat(c)
	})
	ginServer.POST("/chatgpt/single_chat", func(c *gin.Context) {
		openai.Chat(c)
	})

	ginServer.GET("/users", func(c *gin.Context) {
		database.GetUsers(c)

	})
	ginServer.GET("/recording-service-config", func(c *gin.Context) {

		c.JSON(http.StatusOK, models.NewSuccessResponse(
			map[string]interface{}{
				"configUrl":            config.ReadConf("recording.configUrl").(string),
				"endpoint":             config.ReadConf("recording.endpoint").(string),
				"region":               config.ReadConf("recording.region").(string),
				"accessKey":            config.ReadConf("recording.accessKey").(string),
				"secretKey":            config.ReadConf("recording.secretKey").(string),
				"bucketName":           config.ReadConf("recording.bucketName").(string),
				"maxConcurrentTaskNum": config.ReadConf("recording.maxConcurrentTaskNum").(int),
				"audioSliceSize":       config.ReadConf("recording.audioSliceSize").(int),
				"skipSilence":          config.ReadConf("recording.skipSilence").(bool),
				"latestVersion":        config.ReadConf("recording.latestVersion").(int),
				"apkDownloadUrl":       "https://117.50.199.110:8081/download/latest.apk",
			},
		))
	})

	// 定义文件下载接口
	ginServer.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := filepath.Join("uploads", filename) // 假设文件存放在uploads目录下

		// 检查文件是否存在
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		// 设置响应头，告诉浏览器这是一个文件下载
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")

		// 发送文件
		c.File(filePath)
	})

	// 为 multipart forms 设置更低的内存限制 (默认是 32 MiB)
	ginServer.MaxMultipartMemory = 8 << 20 // 8 MiB

	// 定义文件上传接口
	ginServer.POST("/upload", func(c *gin.Context) {
		// 获取上传的文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 构建文件保存路径
		dst := filepath.Join("uploads", file.Filename)

		// 保存文件到指定路径
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, models.NewSuccessResponse(
			map[string]interface{}{
				"message":  "File uploaded successfully",
				"filename": file.Filename,
			},
		))
	})

	ginServer.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	ginServer.POST("/register", func(c *gin.Context) {
		auth.Register(c)
	})
	ginServer.POST("/login", func(c *gin.Context) {
		auth.Login(c)
	})
	return ginServer
}

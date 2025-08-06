package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 靜態檔服務
	r.Static("/doc", "./doc")
	r.Static("/static", "./static")

	// 載入模板，註冊安全 HTML 函式
	r.SetFuncMap(template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	r.LoadHTMLGlob("templates/*")

	// Markdown 頁面路由
	r.GET("/apiGuide", markdownHandler)

	// 測試 API
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	return r
}

func markdownHandler(c *gin.Context) {
	// 讀取md 用套件轉換
	content, err := os.ReadFile("README.md")
	if err != nil {
		log.Printf("Failed to read README.md: %v", err)
		c.String(http.StatusInternalServerError, "Failed to read file")
		return
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		log.Printf("Markdown conversion error: %v", err)
		c.String(http.StatusInternalServerError, "Markdown conversion error")
		return
	}

	// 注入模板
	c.HTML(http.StatusOK, "markdown.html", gin.H{
		"Content": buf.String(),
		// "ImageURL": "/static/1.png", // 如有需要可動態處理
	})
}

func main() {
	r := setupRouter()
	port := ":5001"
	log.Printf("Server running at http://localhost%s/", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

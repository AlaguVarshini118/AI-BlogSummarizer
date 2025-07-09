package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ai-blog-summarizer/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("❌ .env file not loaded")
	} else {
		fmt.Println("✅ .env loaded")
	}

	apiKey := os.Getenv("COHERE_API_KEY")
	if apiKey == "" {
		fmt.Println("🔴 Cohere API Key is EMPTY! Check your .env file")
	} else {
		fmt.Println("🔑 Cohere API Key is:", apiKey[:8]+"...") 
	}
}

func main() {
	r := gin.Default()

	
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Next()
	})

	
	r.GET("/api/articles", routes.GetSummarizedArticles)

	fmt.Println("🚀 Server running at http://localhost:8080")
	r.Run(":8080")
}
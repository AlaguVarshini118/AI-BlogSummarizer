package routes

import (
	"fmt"
	"net/http"
	"ai-blog-summarizer/services"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

type Article struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Link    string `json:"link"`
}

func GetSummarizedArticles(c *gin.Context) {
	fmt.Println("📰 Request received at /api/articles")

	feedURL := "https://techcrunch.com/feed/"
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		fmt.Println(" Failed to fetch RSS feed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch RSS feed"})
		return
	}

	var summarized []Article

	for i, item := range feed.Items[:5] {
		fmt.Printf("📨 Summarizing article %d: %s\n", i+1, item.Title)

		
		textToSummarize := item.Content
		if textToSummarize == "" {
			textToSummarize = item.Description
		}

		if len(textToSummarize) < 250 {
			fmt.Printf(" Skipping article %d (too short: %d chars)\n", i+1, len(textToSummarize))
			summarized = append(summarized, Article{
				Title:   item.Title,
				Summary: " Summary unavailable (Too short for Cohere)",
				Link:    item.Link,
			})
			continue
		}

		summary, err := services.SummarizeText(textToSummarize)
		if err != nil {
			fmt.Println(" Error summarizing article:", err)
			summary = " Summary unavailable (API error)"
		}

		summarized = append(summarized, Article{
			Title:   item.Title,
			Summary: summary,
			Link:    item.Link,
		})
	}

	fmt.Printf(" Returning %d summarized articles\n", len(summarized))
	c.JSON(http.StatusOK, summarized)
}

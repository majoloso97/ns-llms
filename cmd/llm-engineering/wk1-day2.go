package main

import (
	"fmt"

	"github.com/majoloso97/llms/llm-engineering/scrapper"
	"github.com/majoloso97/llms/llm-engineering/summarizer"
	"github.com/majoloso97/llms/shared/env"
)

func main() {
	env.LoadEnv()

	pageContent := new(scrapper.WebPageContent)
	fmt.Println("Welcome to the scrap-summarize tool!")
	fmt.Print("Enter the URL: ")

	var url string
	fmt.Scanln(&url)
	fmt.Println("Summarizing...")

	scrapper.ScrapURL(url, pageContent)
	summarizer.SendOllamaMessage(pageContent)
}

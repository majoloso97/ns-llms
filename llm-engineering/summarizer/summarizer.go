package summarizer

import (
	"fmt"

	"github.com/majoloso97/llms/llm-engineering/scrapper"
	"github.com/majoloso97/llms/shared/ollama"
)

func chatCallback(resp ollama.ChatResponse) error {
	fmt.Println("RUNNING CHAT CALLBACK")
	fmt.Println(resp.Message.Content)
	return nil
}

func PrepareWebpageContentPrompt(content *scrapper.WebPageContent) string {
	user_prompt := `You're looking at a webpage titled %s. The contents of this 
	webpage are as follows. Please provide a short summary of this website in markdown.
	If it includes news or announcements, please include them in the summary.\n%s\n`
	return fmt.Sprintf(user_prompt, content.Title, content.Content)
}

func RequestSummarization(content *scrapper.WebPageContent) {
	client := ollama.GetClient()
	model := ollama.GetModelFromEnv()

	var stream bool = false
	var messages []ollama.Message

	systemPrompt := ollama.Message{
		Role: "system",
		Content: `You're an assistant that summarizes the content of a webpage and
		provides a short summary, ignoring text that might be navigation related. You
		will get the content of the page passed as the body tag of the html document.
		Ignore script, style, image and input tags. Respond in markdown format.`,
	}
	userPrompt := ollama.Message{
		Role:    "user",
		Content: PrepareWebpageContentPrompt(content),
	}

	ollama.AddMessages(&messages, systemPrompt, userPrompt)
	ollama.Chat(client, model, &messages, &stream, chatCallback)
}

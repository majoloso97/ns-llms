package summarizer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/majoloso97/llms/llm-engineering/scrapper"
	"github.com/majoloso97/llms/shared/env"

	ollama "github.com/ollama/ollama/api"
	"github.com/openai/openai-go"
	openaiOption "github.com/openai/openai-go/option"
	openaiShared "github.com/openai/openai-go/shared"
)

func chatCallback(resp ollama.ChatResponse) error {
	fmt.Println("RUNNING CHAT CALLBACK")
	fmt.Println(resp.Message.Content)
	return nil
}

func PrepareWebpageContent(content *scrapper.WebPageContent) string {
	user_prompt := `You're looking at a webpage titled %s. The contents of this 
	webpage are as follows. Please provide a short summary of this website in markdown.
	If it includes news or announcements, please include them in the summary.\n%s\n`
	return fmt.Sprintf(user_prompt, content.Title, content.Content)
}

func SendOllamaMessage(content *scrapper.WebPageContent) {
	ollama_host := env.GetEnv("OLLAMA_HOST")
	model := env.GetEnv("OLLAMA_MODEL")

	host, err := url.Parse(ollama_host)
	if err != nil {
		fmt.Errorf("Error parsing ollama host: %s", err)
	}
	client := ollama.NewClient(host, http.DefaultClient)

	var stream bool = false
	var messages []ollama.Message
	messages = append(messages, ollama.Message{
		Role: "system",
		Content: `You're an assistant that summarizes the content of a webpage and
		provides a short summary, ignoring text that might be navigation related. You
		will get the content of the page passed as the body tag of the html document.
		Ignore script, style, image and input tags. Respond in markdown format.`,
	},
		ollama.Message{
			Role:    "user",
			Content: PrepareWebpageContent(content),
		})

	req := ollama.ChatRequest{
		Model:    model,
		Messages: messages,
		Stream:   &stream,
	}
	ctx := context.TODO()
	if err := client.Chat(ctx, &req, chatCallback); err != nil {
		fmt.Errorf("Ollama chat failed: %v", err)
	}
}

func SendOpenAIMessage() {
	ollama_host := env.GetEnv("OLLAMA_HOST")
	openai_api_key := env.GetEnv("OPENAI_API_KEY")

	client := openai.NewClient(
		openaiOption.WithBaseURL(ollama_host+"/v1"),
		openaiOption.WithAPIKey(openai_api_key),
	)
	model_str := env.GetEnv("OLLAMA_MODEL")
	model := openaiShared.ChatModel(model_str)

	completionReq := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Hello, how are you?"),
		},
		Model: model,
	}

	ctx := context.TODO()
	chatCompletion, err := client.Chat.Completions.New(ctx, completionReq)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(chatCompletion.Choices[0].Message.Content)
}

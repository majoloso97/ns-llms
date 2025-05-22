package ollama

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/majoloso97/llms/shared/env"

	"github.com/ollama/ollama/api"
)

type Message = api.Message
type ChatResponse = api.ChatResponse

func GetClient() *api.Client {
	ollamaHost := env.GetEnv("OLLAMA_HOST")

	host, err := url.Parse(ollamaHost)
	if err != nil {
		fmt.Errorf("Error parsing ollama host: %s", err)
	}
	client := api.NewClient(host, http.DefaultClient)
	return client
}

func GetModelFromEnv() string {
	return env.GetEnv("OLLAMA_MODEL")
}

func AddMessages(messages *[]api.Message, newMessages ...api.Message) {
	*messages = append(*messages, newMessages...)
}

func Chat(client *api.Client, model string, messages *[]api.Message, stream *bool, callback func(resp api.ChatResponse) error) {
	req := api.ChatRequest{
		Model:    model,
		Messages: *messages,
		Stream:   stream,
	}
	ctx := context.TODO()
	if err := client.Chat(ctx, &req, callback); err != nil {
		fmt.Errorf("Ollama chat failed: %v", err)
	}
}

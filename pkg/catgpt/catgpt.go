package catgpt

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/sashabaranov/go-openai"
)

type Options struct {
	APIKey  string `usage:"OpenAI API Key" env:"OPENAI_API_KEY" name:"api-key"`
	BaseURL string `usage:"OpenAI Base URL" env:"OPENAI_API_BASE" name:"base-url" default:"https://api.openai.com/v1"`
	Model   string `usage:"The openai model to use" default:"gpt-3.5-turbo" short:"m"`
}

func Run(ctx context.Context, prompt []string, input io.Reader, opts Options) error {
	data, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	var part openai.ChatMessagePart
	if utf8.Valid(data) {
		part.Text = string(data)
		part.Type = openai.ChatMessagePartTypeText
	} else {
		log.Fatal("I don't support binary content yet, but would be swell to do \"cat image.png | catgpt what is this\"")
	}

	cfg := openai.DefaultConfig(opts.APIKey)
	if opts.BaseURL != "" {
		cfg.BaseURL = opts.BaseURL
	}
	c := openai.NewClientWithConfig(cfg)

	request := openai.ChatCompletionRequest{
		Model:  opts.Model,
		Stream: true,
	}

	if len(prompt) > 0 {
		request.Messages = append(request.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: strings.Join(prompt, " "),
		})
	}

	request.Messages = append(request.Messages, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleUser,
		MultiContent: []openai.ChatMessagePart{
			part,
		},
	})

	response, err := c.CreateChatCompletionStream(ctx, request)
	if err != nil {
		return err
	}
	defer response.Close()

	for {
		resp, err := response.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		fmt.Print(resp.Choices[0].Delta.Content)
	}

	return nil
}

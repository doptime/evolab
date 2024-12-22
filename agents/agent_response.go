package agents

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/doptime/evolab/utils"
	openai "github.com/sashabaranov/go-openai"
	"golang.design/x/clipboard"
)

func (a *Agent) GetResponse(req openai.ChatCompletionRequest) (resp openai.ChatCompletionResponse, err error) {

	// Send the request to the OpenAI API
	if a.msgDeFile != "" {
		resp, err = utils.FileToResponse(a.msgDeFile)
		return resp, err
	}
	if a.msgDeCliboard {
		textbytes := clipboard.Read(clipboard.FmtText)
		if len(textbytes) == 0 {
			return resp, fmt.Errorf("no data in clipboard")
		}
		msg := openai.ChatCompletionMessage{
			Role:    "assistant",
			Content: string(textbytes),
		}
		resp = openai.ChatCompletionResponse{
			Choices: []openai.ChatCompletionChoice{
				{
					Message: msg,
				},
			},
		}
		return resp, nil
	}
	ctx := context.Background()
	//not load from file yet, then send request to openai
	if len(req.Messages) > 0 {
		resp, err = a.Model.Client.CreateChatCompletion(ctx, req)
	}

	if a.msgToFile != "" && len(resp.Choices) > 0 {
		jsonbytes, err := json.Marshal(resp)
		if err == nil {
			saveToFile(&SaveToFile{Filename: a.msgToFile, Content: string(jsonbytes)})
		}
	}
	return resp, err
}
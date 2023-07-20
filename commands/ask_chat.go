package commands

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
	openai "github.com/sashabaranov/go-openai"
)

func handleAskChatInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	questionOption := i.ApplicationCommandData().Options[0]
	question := questionOption.StringValue()

	response, err := chatWithGPT(question)
	if err != nil {
		log.Printf("Failed to get chat response: %s", err)
		return
	}

	reply := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	}

	err = s.InteractionRespond(i.Interaction, reply)
	if err != nil {
		log.Printf("Failed to send response: %s", err)
	}
}

func chatWithGPT(question string) (string, error) {
	client := openai.NewClient(GuiderToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

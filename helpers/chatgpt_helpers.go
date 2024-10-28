package helpers

/*
 * Copyright (C) 2024  Nicholas Popovic
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <https://www.gnu.org/licenses/>.
 */

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sashabaranov/go-openai"
)

type FetchedDataMsg struct {
	Data string
	Err  error
}

func GetChatCompletion(txt string) tea.Cmd {
	return func() tea.Msg {
		apiKey := os.Getenv("OPENAI_API_KEY")
		client := openai.NewClient(apiKey)
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4oLatest,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: txt,
					},
				},
			},
		)

		if err != nil {
			return FetchedDataMsg{Data: "", Err: fmt.Errorf("ChatCompletion error: %v", err)}
		}

		return FetchedDataMsg{Data: resp.Choices[0].Message.Content, Err: nil}
	}
}

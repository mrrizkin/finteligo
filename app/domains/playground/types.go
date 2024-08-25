package playground

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Llm struct {
}

func (*Llm) Prompt(prompt string) {
	url := "https://api.anthropic.com/v1/messages"
	jsonStr := []byte(`{"message": "` + prompt + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set(
		"x-api-key",
		"you-api-key",
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status: %v", resp.Status)
	fmt.Println("response Headers: %v", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body: %v", string(body))
}

type Service struct {
	llm *Llm
}

type Prompt struct {
	Role    string `json:"role"`
	Content string `json:"content"`

	Model       string  `json:"model"       validate:"required"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
	TopK        int     `json:"top_k"`

	Message string `json:"message" validate:"required"`
}

type PromptResponse struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

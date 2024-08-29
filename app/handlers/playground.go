package handlers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mrrizkin/finteligo/app/domains/playground"
	"github.com/mrrizkin/finteligo/system/types"
	lcTypes "github.com/mrrizkin/finteligo/third_party/langchain/types"
	"github.com/valyala/fasthttp"
)

type StreamPromptResponse struct {
	Message string `json:"message"`
}

type StreamResponse struct {
	ID    string
	Event string
	Data  interface{}
}

func generateStreamResponse(w *bufio.Writer, response *StreamResponse) error {
	data, err := json.Marshal(response.Data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "id: %s\nevent: %s\ndata: %s\n\n", response.ID, response.Event, data)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handlers) Prompt(c *fiber.Ctx) error {
	payload := new(playground.PromptPayload)
	err := c.BodyParser(payload)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "payload not valid",
		}
	}

	validationError := h.System.Validator.MustValidate(payload)
	if validationError != nil {
		return validationError
	}

	if !payload.Stream {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan string)
		defer close(ch)

		streamFunc := func(cx context.Context, chunk []byte) error {
			return nil
		}

		promptPayload := lcTypes.PromptPayload{
			Role:        payload.Role,
			Content:     payload.Content,
			Model:       payload.Model,
			Temperature: payload.Temperature,
			TopP:        payload.TopP,
			TopK:        payload.TopK,
			Message:     payload.Message,
			StreamFunc:  &streamFunc,
			Channel:     ch,
		}

		go func() {
			err := h.playgroundService.Prompt(payload.Token, promptPayload)
			if err != nil {
				h.System.Logger.Error().Err(err).Msg("failed to prompt")
				ch <- "error: " + err.Error()
			}
			cancel()
		}()

		for {
			select {
			case <-ctx.Done():
				return h.SendJson(c, types.Response{
					Title:   "Prompt",
					Message: "Prompt interrupted",
					Status:  "error",
				})
			case promptResponse := <-ch:
				if promptResponse != "" {
					if strings.HasPrefix(promptResponse, "error:") {
						return h.SendJson(c, types.Response{
							Title:   "Prompt",
							Message: "Prompt error",
							Status:  "error",
							Data:    promptResponse[6:],
						})
					} else {
						return h.SendJson(c, types.Response{
							Title:   "Prompt",
							Message: "Prompt success",
							Status:  "success",
							Data:    promptResponse,
						})
					}
				}
			}
		}
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")
	c.Status(fiber.StatusOK).
		Context().
		SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := make(chan string)
			defer close(ch)
			streamFunc := func(cx context.Context, chunk []byte) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					err := generateStreamResponse(w, &StreamResponse{
						ID:    "1",
						Event: "message",
						Data:  string(chunk),
					})
					if err != nil {
						h.System.Logger.Error().Err(err).Msg("failed to write chunk")
						return err
					}
					return w.Flush()
				}
			}

			promptPayload := lcTypes.PromptPayload{
				Role:        payload.Role,
				Content:     payload.Content,
				Model:       payload.Model,
				Temperature: payload.Temperature,
				TopP:        payload.TopP,
				TopK:        payload.TopK,
				Message:     payload.Message,
				StreamFunc:  &streamFunc,
				Channel:     ch,
			}

			go func() {
				err := h.playgroundService.Prompt(payload.Token, promptPayload)
				if err != nil {
					h.System.Logger.Error().Err(err).Msg("failed to prompt")
					ch <- "error: " + err.Error()
				}
				cancel()
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case promptResponse := <-ch:
					if promptResponse != "" {
						if strings.HasPrefix(promptResponse, "error:") {
							_, err := fmt.Fprintf(
								w,
								"event: error\ndata: %s\n\n",
								promptResponse[6:],
							)
							if err != nil {
								h.System.Logger.Error().Err(err).Msg("failed to write error")
							}
							return
						}
					}
				}
			}
		}))

	return nil
}

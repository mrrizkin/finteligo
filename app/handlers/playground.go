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
	"github.com/mrrizkin/finteligo/third_party/logger"
	"github.com/valyala/fasthttp"
)

func playgroundPrompting(
	w *bufio.Writer,
	payload *playground.PromptPayload,
	service *playground.Service,
	log *logger.Logger,
	sendStream func(w *bufio.Writer, response *StreamResponse) error,
) *types.Response {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan string)
	defer close(ch)

	payload.Channel = ch
	payload.StreamFunc = nil

	if !payload.Stream {
		go func() {
			err := service.Prompt(payload)
			if err != nil {
				log.Error().Err(err).Msg("failed to prompt")
				payload.Channel <- "error: " + err.Error()
			}
			cancel()
		}()
	}

	if payload.Stream && w != nil {
		streamFunc := func(cx context.Context, chunk []byte) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				message, err := json.Marshal(string(chunk))
				if err != nil {
					log.Error().Err(err).Msg("failed to marshal message")
					return err
				}

				err = sendStream(w, &StreamResponse{
					ID:    "1",
					Event: "message",
					Data:  message,
				})
				if err != nil {
					log.Error().Err(err).Msg("failed to write chunk")
					return err
				}
				return w.Flush()
			}

		}
		payload.StreamFunc = &streamFunc
		go func() {
			err := service.Prompt(payload)
			if err != nil {
				log.Error().Err(err).Msg("failed to prompt")
				payload.Channel <- "error: " + err.Error()
			}
			cancel()
		}()

	}

	for {
		select {
		case <-ctx.Done():
			if w != nil {
				return &types.Response{
					Title:   "Prompt",
					Message: "Prompt interrupted",
					Status:  "error",
				}
			} else {
				return nil
			}
		case promptResponse := <-ch:
			if promptResponse != "" {
				if strings.HasPrefix(promptResponse, "error:") {
					if w != nil {
						_, err := fmt.Fprintf(
							w,
							"event: error\ndata: %s\n\n",
							promptResponse[6:],
						)
						if err != nil {
							log.Error().Err(err).Msg("failed to write error")
						}
						return nil
					} else {
						return &types.Response{
							Title:   "Prompt",
							Message: "Prompt error",
							Status:  "error",
							Data:    promptResponse[6:],
						}
					}
				} else {
					if w != nil {
						return nil
					} else {
						return &types.Response{
							Title:   "Prompt",
							Message: "Prompt success",
							Status:  "success",
							Data:    promptResponse,
						}
					}
				}
			}
		}
	}
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
		response := playgroundPrompting(
			nil,
			payload,
			h.playgroundService,
			h.System.Logger,
			h.SendStream,
		)
		if response != nil {
			return h.SendJson(c, *response)
		}
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")
	c.Status(fiber.StatusOK).
		Context().
		SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			response := playgroundPrompting(
				w,
				payload,
				h.playgroundService,
				h.System.Logger,
				h.SendStream,
			)
			if response != nil {
				h.SendJson(c, *response)
			}
		}))

	return nil
}

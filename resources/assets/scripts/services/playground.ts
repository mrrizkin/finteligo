import { request, validate } from "@lib/request";

import { PlaygroundPromptPayload, PlaygroundPromptResponse, playgroundPromptResponseSchema } from "@schemas/playground";

const API_PLAYGROUND_PROMPT = "/api/v1/playground/prompt";

export async function prompt(payload: PlaygroundPromptPayload) {
  return validate(
    await request.post<PlaygroundPromptResponse>(API_PLAYGROUND_PROMPT, payload),
    playgroundPromptResponseSchema,
  );
}

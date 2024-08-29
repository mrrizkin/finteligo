import { PromptRequestOptions, SSEResponse } from "@schemas/playground";

const API_PLAYGROUND_PROMPT = "/api/v1/playground/prompt";

function parseSSEResponse(value: string): SSEResponse {
  const data = value.split("\n").reduce(
    (acc, line) => {
      const [key, value] = line.split(": ");
      return { ...acc, [key]: value };
    },
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    {} as Record<string, any>,
  );

  return {
    data: data.data,
    id: data.id,
    event: data.event,
  };
}

export async function prompt(options: PromptRequestOptions = {}) {
  try {
    const apiResponse = await fetch(API_PLAYGROUND_PROMPT, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(options.payload || {}),
    });

    if (!apiResponse.ok) {
      throw new Error(`HTTP error! status: ${apiResponse.status}`);
    }

    if (!apiResponse.body) {
      throw new Error("Response body is null");
    }

    const reader = apiResponse.body.pipeThrough(new TextDecoderStream()).getReader();
    while (true) {
      const { value, done } = await reader.read();
      if (done) {
        options.done?.();
        break;
      }
      if (value) {
        options.stream?.(parseSSEResponse(value));
      }
    }
  } catch (error) {
    options.error?.(error instanceof Error ? error.message : String(error));
  }
}

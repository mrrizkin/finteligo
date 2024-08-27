import { PromptRequestOptions } from "@schemas/playground";

const API_PLAYGROUND_PROMPT = "/api/v1/playground/prompt";

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
    let buffer = "";

    while (true) {
      const { value, done } = await reader.read();

      if (done) {
        if (buffer.trim()) {
          processChunk(buffer);
        }
        options.done?.();
        break;
      }

      buffer += value;
      let newlineIndex;
      while ((newlineIndex = buffer.indexOf("\n\n")) >= 0) {
        const chunk = buffer.slice(0, newlineIndex);
        buffer = buffer.slice(newlineIndex + 2);
        processChunk(chunk);
      }
    }
  } catch (error) {
    options.error?.(error instanceof Error ? error.message : String(error));
  }

  function processChunk(chunk: string) {
    const lines = chunk.split("\n");
    let data = "";
    for (const line of lines) {
      if (line.startsWith("data: ")) {
        data += line.slice(6);
      }
    }
    if (data) {
      options.stream?.(data);
    }
  }
}

import { z } from "zod";

import { createResponseSchema } from "@schemas";

export const playgroundPromptPayloadSchema = z.object({
  name: z.string(),
  age: z.number(),
});

export const playgroundPromptResponseSchema = createResponseSchema(
  z.object({
    message: z.string(),
  }),
);

export type PlaygroundPromptPayload = z.infer<typeof playgroundPromptPayloadSchema>;
export type PlaygroundPromptResponse = z.infer<typeof playgroundPromptResponseSchema>;

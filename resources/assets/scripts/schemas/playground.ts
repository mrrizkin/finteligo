import { z } from "zod";

import { createResponseSchema } from "@schemas";

/** ================================ MODEL ================================= **/

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const playgroundPromptPayloadSchema = z.object({
  model: z.string(),
  message: z.string(),
  temperature: z.number().nullable().optional(),
  topP: z.number().nullable().optional(),
  topK: z.number().nullable().optional(),
  role: z.string().nullable().optional(),
  content: z.string().nullable().optional(),
  token: z.string().nullable().optional(),
});

export const playgroundPromptResponseSchema = createResponseSchema(
  z.object({
    answer: z.string(),
  }),
);

export type PlaygroundPromptPayload = z.infer<typeof playgroundPromptPayloadSchema>;
export type PlaygroundPromptResponse = z.infer<typeof playgroundPromptResponseSchema>;
export type PromptRequestOptions = {
  payload?: PlaygroundPromptPayload;
  stream?: (data: string) => void;
  done?: () => void;
  error?: (error: string) => void;
};

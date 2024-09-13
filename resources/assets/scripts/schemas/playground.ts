import { z } from "zod";

import { createResponseSchema } from "@schemas";

/** ================================ MODEL ================================= **/

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const playgroundPromptPayloadSchema = z.object({
  model: z.string(),
  message: z.string(),
  chat_history: z.object({
    role: z.string(),
    content: z.array(z.string()),
  }).array().optional(),
  temperature: z.number().nullable().optional(),
  topP: z.number().nullable().optional(),
  topK: z.number().nullable().optional(),
  role: z.string().nullable().optional(),
  content: z.string().nullable().optional(),
  token: z.string().nullable().optional(),
  stream: z.boolean().nullable().optional(),
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
  stream?: (data: SSEResponse) => void;
  done?: () => void;
  error?: (error: string) => void;
};

export type SSEResponse = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  data: any;
  id: string;
  event: string;
};

import { z } from "zod";

import { createResponseSchema, createResponsesSchema, generalModelSchema } from "@schemas";

/** ================================ MODEL ================================= **/

export const modelsSchema = z
  .object({
    user_id: z.number(),
    token: z.string(),
    model: z.string(),
    provider: z.string(),
    url: z.string(),
    api_key: z.string(),
    status: z.string(),
    enabled: z.boolean(),
    error: z.string().nullable(),
  })
  .merge(generalModelSchema);

export type Models = z.infer<typeof modelsSchema>;

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const modelsResponseSchema = createResponseSchema(modelsSchema);
export const modelsResponsesSchema = createResponsesSchema(modelsSchema);

export type ModelsResponse = z.infer<typeof modelsResponseSchema>;
export type ModelsResponses = z.infer<typeof modelsResponsesSchema>;

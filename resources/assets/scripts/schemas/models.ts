import { z } from "zod";

import { createResponseSchema, createResponsesSchema, generalModelSchema } from "@schemas";

/** ================================ MODEL ================================= **/

export const modelsSchema = z
  .object({
    user_id: z.number().optional(),
    token: z.string().optional(),
    model: z.string().optional(),
    provider: z.string().optional(),
    url: z.string().optional(),
    api_key: z.string().optional(),
    status: z.string().optional(),
    enabled: z.boolean().optional(),
    error: z.string().nullable(),
  })
  .merge(generalModelSchema);

export type Models = z.infer<typeof modelsSchema>;

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const modelsResponseSchema = createResponseSchema(modelsSchema);
export const modelsResponsesSchema = createResponsesSchema(modelsSchema);

export type ModelsResponse = z.infer<typeof modelsResponseSchema>;
export type ModelsResponses = z.infer<typeof modelsResponsesSchema>;

import { z } from "zod";

import { createResponseSchema, createResponsesSchema, generalModelSchema } from "@schemas";

import { userSchema } from "./user";

/** ================================ MODEL ================================= **/

export const apiTokenSchema = z
  .object({
    key: z.string().optional(),
    slug: z.string().optional(),
    token: z.string().optional(),
    user_id: z.number().optional(),
    expiry_date: z.string().optional(),
    expired: z.boolean().optional(),
    enabled: z.boolean().optional(),
    user: userSchema.optional().nullable(),
  })
  .merge(generalModelSchema);

export type ApiToken = z.infer<typeof apiTokenSchema>;

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const apiTokenResponseSchema = createResponseSchema(apiTokenSchema);
export const apiTokenResponsesSchema = createResponsesSchema(apiTokenSchema);

export type ApiTokenResponse = z.infer<typeof apiTokenResponseSchema>;
export type ApiTokenResponses = z.infer<typeof apiTokenResponsesSchema>;

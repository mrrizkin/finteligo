import { z } from "zod";

import { createResponseSchema, createResponsesSchema, generalModelSchema } from "@schemas";
import { roleSchema } from "@schemas/role";

/** ================================ MODEL ================================= **/

export const userSchema = z
  .object({
    username: z.string().optional(),
    password: z.string().optional(),
    name: z.string().optional(),
    email: z.string().optional(),
    role_id: z.number().optional(),
    role: roleSchema.optional().nullable(),
  })
  .merge(generalModelSchema);

export type User = z.infer<typeof userSchema>;

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const userResponseSchema = createResponseSchema(userSchema);
export const userResponsesSchema = createResponsesSchema(userSchema);

export type UserResponse = z.infer<typeof userResponseSchema>;
export type UserResponses = z.infer<typeof userResponsesSchema>;

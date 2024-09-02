import { z } from "zod";

import { createResponseSchema, createResponsesSchema, generalModelSchema } from "@schemas";
import { rolePermissionSchema } from "@schemas/role_permission";

/** ================================ MODEL ================================= **/

export const roleSchema = z
  .object({
    slug: z.string().optional(),
    name: z.string().optional(),
    role_permissions: z.array(rolePermissionSchema).optional().nullable(),
  })
  .merge(generalModelSchema);

export type Role = z.infer<typeof roleSchema>;

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const roleResponseSchema = createResponseSchema(roleSchema);
export const roleResponsesSchema = createResponsesSchema(roleSchema);

export type RoleResponse = z.infer<typeof roleResponseSchema>;
export type RoleResponses = z.infer<typeof roleResponsesSchema>;

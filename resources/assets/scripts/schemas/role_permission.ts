import { z } from "zod";

import { createResponseSchema, createResponsesSchema, generalModelSchema } from "@schemas";

/** ================================ MODEL ================================= **/

export const rolePermissionSchema = z
  .object({
    role_id: z.number().optional(),
    permission_id: z.number().optional(),
  })
  .merge(generalModelSchema);

export type RolePermission = z.infer<typeof rolePermissionSchema>;

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const rolePermissionResponseSchema = createResponseSchema(rolePermissionSchema);
export const rolePermissionResponsesSchema = createResponsesSchema(rolePermissionSchema);

export type RolePermissionResponse = z.infer<typeof rolePermissionResponseSchema>;
export type RolePermissionResponses = z.infer<typeof rolePermissionResponsesSchema>;

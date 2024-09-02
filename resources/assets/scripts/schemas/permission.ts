import { z } from "zod";

import { createResponseSchema, createResponsesSchema, generalModelSchema } from "@schemas";

/** ================================ MODEL ================================= **/

export const permissionSchema = z
  .object({
    group: z.string().optional(),
    name: z.string().optional(),
    slug: z.string().optional(),
  })
  .merge(generalModelSchema);

export type Permission = z.infer<typeof permissionSchema>;

/** ========================== PAYLOAD & RESPONSE ========================== **/

export const permissionResponseSchema = createResponseSchema(permissionSchema);
export const permissionResponsesSchema = createResponsesSchema(permissionSchema);

export type PermissionResponse = z.infer<typeof permissionResponseSchema>;
export type PermissionResponses = z.infer<typeof permissionResponsesSchema>;

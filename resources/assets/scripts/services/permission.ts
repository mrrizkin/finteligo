import { request, validate } from "@lib/request";

import { AnyObject, StatusResponseSchema, statusResponseSchema } from "@schemas";
import {
  PermissionResponse,
  PermissionResponses,
  permissionResponseSchema,
  permissionResponsesSchema,
} from "@schemas/permission";

const API_PERMISSION_ALL = "/api/v1/permission";
const API_PERMISSION_FIND = "/api/v1/permission/:id";
const API_PERMISSION_CREATE = "/api/v1/permission";
const API_PERMISSION_UPDATE = "/api/v1/permission/:id";
const API_PERMISSION_DELETE = "/api/v1/permission/:id";

export async function get_all(params: AnyObject = {}) {
  return validate(await request.get<PermissionResponses>(API_PERMISSION_ALL, { params }), permissionResponsesSchema);
}

export async function find(id: number, params: AnyObject = {}) {
  return validate(
    await request.get<PermissionResponse>(API_PERMISSION_FIND.replace(":id", id.toString()), { params }),
    permissionResponseSchema,
  );
}

export async function create(payload: AnyObject) {
  return validate(await request.post<PermissionResponse>(API_PERMISSION_CREATE, payload), permissionResponseSchema);
}

export async function update(id: number, payload: AnyObject) {
  return validate(
    await request.put<PermissionResponse>(API_PERMISSION_UPDATE.replace(":id", id.toString()), payload),
    permissionResponseSchema,
  );
}

export async function remove(id: number, payload: AnyObject = {}) {
  return validate(
    await request.delete<StatusResponseSchema>(API_PERMISSION_DELETE.replace(":id", id.toString()), payload),
    statusResponseSchema,
  );
}

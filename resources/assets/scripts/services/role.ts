import { request, validate } from "@lib/request";

import { AnyObject, StatusResponseSchema, statusResponseSchema } from "@schemas";
import { RoleResponse, RoleResponses, roleResponseSchema, roleResponsesSchema } from "@schemas/role";

const API_ROLE_ALL = "/api/v1/role";
const API_ROLE_FIND = "/api/v1/role/:id";
const API_ROLE_CREATE = "/api/v1/role";
const API_ROLE_UPDATE = "/api/v1/role/:id";
const API_ROLE_DELETE = "/api/v1/role/:id";

export async function get_all(params: AnyObject = {}) {
  return validate(await request.get<RoleResponses>(API_ROLE_ALL, { params }), roleResponsesSchema);
}

export async function find(id: number, params: AnyObject = {}) {
  return validate(
    await request.get<RoleResponse>(API_ROLE_FIND.replace(":id", id.toString()), { params }),
    roleResponseSchema,
  );
}

export async function create(payload: AnyObject) {
  return validate(await request.post<RoleResponse>(API_ROLE_CREATE, payload), roleResponseSchema);
}

export async function update(id: number, payload: AnyObject) {
  return validate(
    await request.put<RoleResponse>(API_ROLE_UPDATE.replace(":id", id.toString()), payload),
    roleResponseSchema,
  );
}

export async function remove(id: number, payload: AnyObject = {}) {
  return validate(
    await request.delete<StatusResponseSchema>(API_ROLE_DELETE.replace(":id", id.toString()), payload),
    statusResponseSchema,
  );
}

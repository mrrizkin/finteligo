import { request, validate } from "@lib/request";

import { AnyObject, StatusResponseSchema, statusResponseSchema } from "@schemas";
import { UserResponse, UserResponses, userResponseSchema, userResponsesSchema } from "@schemas/user";

const API_USER_ALL = "/api/v1/user";
const API_USER_FIND = "/api/v1/user/:id";
const API_USER_CREATE = "/api/v1/user";
const API_USER_UPDATE = "/api/v1/user/:id";
const API_USER_DELETE = "/api/v1/user/:id";

export async function get_all(params: AnyObject = {}) {
  return validate(await request.get<UserResponses>(API_USER_ALL, { params }), userResponsesSchema);
}

export async function find(id: number, params: AnyObject = {}) {
  return validate(
    await request.get<UserResponse>(API_USER_FIND.replace(":id", id.toString()), { params }),
    userResponseSchema,
  );
}

export async function create(payload: AnyObject) {
  return validate(await request.post<UserResponse>(API_USER_CREATE, payload), userResponseSchema);
}

export async function update(id: number, payload: AnyObject) {
  return validate(
    await request.put<UserResponse>(API_USER_UPDATE.replace(":id", id.toString()), payload),
    userResponseSchema,
  );
}

export async function remove(id: number, payload: AnyObject = {}) {
  return validate(
    await request.delete<StatusResponseSchema>(API_USER_DELETE.replace(":id", id.toString()), payload),
    statusResponseSchema,
  );
}

import { request, validate } from "@lib/request";

import { AnyObject, StatusResponseSchema, statusResponseSchema } from "@schemas";
import {
  ApiTokenResponse,
  ApiTokenResponses,
  apiTokenResponseSchema,
  apiTokenResponsesSchema,
} from "@schemas/api-token";

const API_API_TOKEN_ALL = "/api/v1/api-tokens";
const API_API_TOKEN_FIND = "/api/v1/api-tokens/:id";
const API_API_TOKEN_CREATE = "/api/v1/api-tokens";
const API_API_TOKEN_ENABLE = "/api/v1/api-tokens/:id/enable";
const API_API_TOKEN_DISABLE = "/api/v1/api-tokens/:id/disable";
const API_API_TOKEN_UPDATE = "/api/v1/api-tokens/:id";
const API_API_TOKEN_DELETE = "/api/v1/api-tokens/:id";

export async function get_all(params: AnyObject = {}) {
  return validate(await request.get<ApiTokenResponses>(API_API_TOKEN_ALL, { params }), apiTokenResponsesSchema);
}

export async function find(id: number, params: AnyObject = {}) {
  return validate(
    await request.get<ApiTokenResponse>(API_API_TOKEN_FIND.replace(":id", id.toString()), { params }),
    apiTokenResponseSchema,
  );
}

export async function create(payload: AnyObject) {
  return validate(await request.post<ApiTokenResponse>(API_API_TOKEN_CREATE, payload), apiTokenResponseSchema);
}

export async function update(id: number, payload: AnyObject) {
  return validate(
    await request.put<ApiTokenResponse>(API_API_TOKEN_UPDATE.replace(":id", id.toString()), payload),
    apiTokenResponseSchema,
  );
}

export async function enable(id: number, payload: AnyObject = {}) {
  return validate(
    await request.post<StatusResponseSchema>(API_API_TOKEN_ENABLE.replace(":id", id.toString()), payload),
    statusResponseSchema,
  );
}

export async function disable(id: number, payload: AnyObject = {}) {
  return validate(
    await request.post<StatusResponseSchema>(API_API_TOKEN_DISABLE.replace(":id", id.toString()), payload),
    statusResponseSchema,
  );
}

export async function remove(id: number, payload: AnyObject = {}) {
  return validate(
    await request.delete<StatusResponseSchema>(API_API_TOKEN_DELETE.replace(":id", id.toString()), payload),
    statusResponseSchema,
  );
}

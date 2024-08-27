import { request, validate } from "@lib/request";

import { AnyObject, StatusResponseSchema, statusResponseSchema } from "@schemas";
import { ModelsResponse, ModelsResponses, modelsResponseSchema, modelsResponsesSchema } from "@schemas/models";

const API_MODELS_ALL = "/api/v1/models";
const API_MODELS_FIND = "/api/v1/models/:id";
const API_MODELS_CREATE = "/api/v1/models";
const API_MODELS_UPDATE = "/api/v1/models/:id";
const API_MODELS_DELETE = "/api/v1/models/:id";

export async function get_all(params: AnyObject = {}) {
  return validate(await request.get<ModelsResponses>(API_MODELS_ALL, { params }), modelsResponsesSchema);
}

export async function find(id: number, params: AnyObject = {}) {
  return validate(
    await request.get<ModelsResponse>(API_MODELS_FIND.replace(":id", id.toString()), { params }),
    modelsResponseSchema,
  );
}

export async function create(payload: AnyObject) {
  return validate(await request.post<ModelsResponse>(API_MODELS_CREATE, payload), modelsResponseSchema);
}

export async function update(id: number, payload: AnyObject) {
  return validate(
    await request.put<ModelsResponse>(API_MODELS_UPDATE.replace(":id", id.toString()), payload),
    modelsResponseSchema,
  );
}

export async function remove(id: number, payload: AnyObject = {}) {
  return validate(
    await request.delete<StatusResponseSchema>(API_MODELS_DELETE.replace(":id", id.toString()), payload),
    statusResponseSchema,
  );
}

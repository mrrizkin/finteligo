import { request, validate } from "@lib/request";

import { AnyObject, StatusResponseSchema, statusResponseSchema } from "@schemas";

const API_AUTH_IDENTITY = "/api/identity";
const API_AUTH_LOGIN = "/api/login";
const API_AUTH_LOGOUT = "/api/logout";

export async function identity() {
  return validate(await request.get<StatusResponseSchema>(API_AUTH_IDENTITY), statusResponseSchema);
}

export async function login(payload: AnyObject = {}) {
  return validate(await request.post<StatusResponseSchema>(API_AUTH_LOGIN, payload), statusResponseSchema);
}

export async function logout(payload: AnyObject = {}) {
  return validate(await request.post<StatusResponseSchema>(API_AUTH_LOGOUT, payload), statusResponseSchema);
}

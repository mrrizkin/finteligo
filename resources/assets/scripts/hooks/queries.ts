import { useQuery } from "react-query";

import { PaginationFilter } from "@schemas";

import * as apiTokensService from "@services/api-token";
import * as modelsService from "@services/models";
import * as permissionService from "@services/permission";
import * as roleService from "@services/role";
import * as userService from "@services/user";

export function useApiTokens(params: PaginationFilter = {}) {
  return useQuery({
    queryKey: ["list-api-tokens", params.page, params.per_page],
    queryFn: () =>
      apiTokensService.get_all({
        page: params.page || 1,
        per_page: params.per_page || 10,
      }),
  });
}

export function useApiToken(id: number) {
  return useQuery({
    queryKey: ["find-api-token", id],
    queryFn: () => apiTokensService.find(id),
  });
}

export function useModels(params: PaginationFilter = {}) {
  return useQuery({
    queryKey: ["list-models", params.page, params.per_page],
    queryFn: () =>
      modelsService.get_all({
        page: params.page || 1,
        per_page: params.per_page || 10,
      }),
  });
}

export function useModel(id: number) {
  return useQuery({
    queryKey: ["find-model", id],
    queryFn: () => modelsService.find(id),
  });
}

export function usePermissions() {
  return useQuery({
    queryKey: ["list-permission"],
    queryFn: () => permissionService.get_all(),
  });
}

export function useRoles(params: PaginationFilter = {}) {
  return useQuery({
    queryKey: ["list-roles", params.page, params.per_page],
    queryFn: () =>
      roleService.get_all({
        page: params.page || 1,
        per_page: params.per_page || 10,
      }),
  });
}

export function useRole(id: number) {
  return useQuery({
    queryKey: ["find-role", id],
    queryFn: () => roleService.find(id),
  });
}

export function useUsers(params: PaginationFilter = {}) {
  return useQuery({
    queryKey: ["list-users", params.page, params.per_page],
    queryFn: () =>
      userService.get_all({
        page: params.page || 1,
        per_page: params.per_page || 10,
      }),
  });
}

export function useUser(id: number) {
  return useQuery({
    queryKey: ["find-user", id],
    queryFn: () => userService.find(id),
  });
}

import { useQuery } from "react-query";

import { PaginationFilter } from "@schemas";

import * as apiTokensService from "@services/api-token";
import * as modelsService from "@services/models";

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

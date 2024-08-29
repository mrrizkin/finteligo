import { useQuery } from "react-query";

import * as apiTokensService from "@services/api-token";
import * as modelsService from "@services/models";

export function useApiTokens() {
  return useQuery({
    queryKey: ["list-api-tokens"],
    queryFn: () => apiTokensService.get_all(),
  });
}

export function useApiToken(id: number) {
  return useQuery({
    queryKey: ["find-api-token", id],
    queryFn: () => apiTokensService.find(id),
  });
}

export function useModels() {
  return useQuery({
    queryKey: ["list-models"],
    queryFn: () => modelsService.get_all(),
  });
}

export function useModel(id: number) {
  return useQuery({
    queryKey: ["find-model", id],
    queryFn: () => modelsService.find(id),
  });
}

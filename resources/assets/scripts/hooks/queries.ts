import { useQuery } from "react-query";

import * as modelsService from "@services/models";

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

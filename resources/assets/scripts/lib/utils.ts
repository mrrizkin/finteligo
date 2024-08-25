import { type ClassValue, clsx } from "clsx";
import { toast } from "sonner";
import { twMerge } from "tailwind-merge";

import { ErrorResponseSchema, ToastValidationParams, ValidatedResponse } from "@schemas";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function toastValidation<T, D>(promise: Promise<ValidatedResponse<T, D>>, params?: ToastValidationParams<T>) {
  promise
    .then((data) => {
      if (data.validation.success) {
        params?.success?.(data.validation.data as T);
      } else {
        toast.error(data.validation.error?.name || "Oops something when wrong", {
          description: data.validation.error?.message || "An error occurred",
        });
      }
    })
    .catch((error) => {
      error
        .json()
        .then((errorData: ErrorResponseSchema) => {
          toast.error(errorData.title || "Oops something when wrong", {
            description: errorData.message || "An error occurred",
          });
          params?.error?.(errorData);
        })
        .catch(() => params?.error?.(error));
    })
    .finally(() => {
      params?.finally?.();
    });
}

export function formatBytes(bytes: number, decimals = 2) {
  if (bytes === 0) return "0 Bytes";

  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

  const i = Math.floor(Math.log(bytes) / Math.log(k));
  const value = bytes / Math.pow(k, i);

  const formattedValue = new Intl.NumberFormat("en-US", {
    minimumFractionDigits: dm,
    maximumFractionDigits: dm,
  }).format(value);

  return `${formattedValue} ${sizes[i]}`;
}

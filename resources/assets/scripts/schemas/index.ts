import { AxiosResponse } from "axios";
import { z } from "zod";

export type AnyObject = { [key: string]: unknown };
export type Validatation<T = unknown> = {
  success: boolean;
  error?: z.ZodError | null;
  data?: T;
};

export type ValidatedResponse<T, D> = {
  validation: Validatation<T>;
} & AxiosResponse<T, D>;

export type SafeParseAsync = (data: unknown, params?: Partial<z.ParseParams>) => Promise<Validatation>;

export type ToastValidationParams<T> = {
  success?(data: T): void;
  error?(error: unknown): void;
  finally?(): void;
};

export type Schema = {
  safeParseAsync: SafeParseAsync;
};

export type ErrorResponseSchema = z.infer<typeof errorResponseSchema>;

export type StatusResponseSchema = z.infer<typeof statusResponseSchema>;

export const errorResponseSchema = z.object({ detail: z.string().optional() }).merge(createStatusResponseSchema());

export const generalModelSchema = z.object({
  id: z.number(),
  created_at: z.string(),
  updated_at: z.string(),
  deleted_at: z.string().nullable(),
});

export const statusResponseSchema = createStatusResponseSchema();

function createStatusResponseSchema() {
  return z.object({
    title: z.string().optional(),
    message: z.string().optional(),
    status: z.enum(["success", "error", "warning", "info"]).optional(),
  });
}

function createDataResponseSchema<T>(schema: z.ZodType<T>) {
  return z.object({
    data: schema.optional(),
  });
}

function createPaginationResponseSchema() {
  return z.object({
    meta: z.object({
      page: z.number(),
      per_page: z.number(),
      total: z.number(),
      page_count: z.number(),
    }),
  });
}

export function createResponseSchema<T>(schema: z.ZodType<T>) {
  return createStatusResponseSchema().merge(createDataResponseSchema(schema));
}

export function createResponsesSchema<T>(schema: z.ZodType<T>) {
  return createStatusResponseSchema()
    .merge(createDataResponseSchema(z.array(schema)))
    .merge(createPaginationResponseSchema());
}

import axios, { AxiosResponse } from "axios";
import { json } from "react-router-dom";

import { ErrorResponseSchema, Schema, ValidatedResponse } from "@schemas";

const request = axios.create({
  withCredentials: true,
  timeout: 30000,
});

request.interceptors.response.use(null, requestErrorHandler, {
  synchronous: true,
});

async function requestErrorHandler(error: unknown) {
  if (axios.isAxiosError(error)) {
    let statusCode = error.response?.status;
    const data: ErrorResponseSchema = {
      title: "Oops, something went wrong!",
      message: error.message,
    };

    if (error.response) {
      if (error.response.data) {
        if (typeof error.response.data.text === "function") {
          let response = await error.response.data.text();
          try {
            response = JSON.parse(response);
          } catch (error) {
            console.error(error);
          }
          if (response.title) {
            data.title = response.title;
          }
          if (response.message) {
            data.message = response.message;
          }
          if (response.detail) {
            data.detail = response.detail;
          }
          if (response.status) {
            data.status = response.status;
          }
          if (response.statusCode) {
            statusCode = response.statusCode;
          }
        } else if (typeof error.response.data === "string") {
          if (error.response.data === "missing or malformed JWT") {
            console.error("You are unauthoraized to use application please re-login");
            data.title = "Unauthoraized";
            data.message = "You are not authorized to access this page";
            statusCode = 401;
          }
        } else {
          const response = error.response.data;
          if (response.title) {
            data.title = response.title;
          }
          if (response.message) {
            data.message = response.message;
          }
          if (response.detail) {
            data.detail = response.detail;
          }
          if (response.status) {
            data.status = response.status;
          }
        }
      }
    } else if (error.code === "ERR_NETWORK") {
      data.title = "Service unavailable";
      data.message = "Service is unavailable at the moment, please try again later";
      statusCode = 503;
    } else if (error.code === "ECONNABORTED") {
      data.title = "Request timeout";
      data.message = "Request timeout, please check your internet connection";
      statusCode = 408;
    }

    throw json(data, statusCode);
  }

  throw error;
}

async function validate<T, D>(response: AxiosResponse<T, D>, schema: Schema): Promise<ValidatedResponse<T, D>> {
  if (response.status < 200 || response.status >= 300) {
    return {
      ...response,
      validation: {
        success: false,
        error: null,
        data: undefined,
      },
    };
  }

  const result = await schema.safeParseAsync(response.data);
  return {
    ...response,
    validation: {
      success: result.success,
      error: result.error,
      data: result.data as T,
    },
  };
}

export { request, validate };

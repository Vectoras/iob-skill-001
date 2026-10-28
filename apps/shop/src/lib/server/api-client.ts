// The core of the private APIs' client: one apiFetch per API, made from that API's
// address and secret, then a one-line function per endpoint. Server-only: the
// browser never calls a private API. The same file lives in an app's
// src/lib/server/api-client.ts until a second frontend app calls a private API; then
// it moves into the api-client package (docs/guide.md §6).
import type { z } from "zod";

export type ApiErrorKind = "network" | "http" | "validation";

export class ApiError extends Error {
  readonly kind: ApiErrorKind;
  readonly status: number | undefined;

  constructor(
    kind: ApiErrorKind,
    message: string,
    status?: number,
    options?: ErrorOptions,
  ) {
    super(message, options);
    this.name = "ApiError";
    this.kind = kind;
    this.status = status;
  }
}

export interface ApiConfig {
  // The API's name, for error messages.
  name: string;
  // Both come from the app's src/env.ts: required when the app runs, possibly
  // missing while it builds (a prerendered page makes no call).
  url: string | undefined;
  secret: string | undefined;
}

export function createApiFetch(api: ApiConfig) {
  return async function apiFetch<Schema extends z.ZodType>(
    path: string,
    schema: Schema,
    init?: RequestInit,
  ): Promise<z.output<Schema>> {
    const headers = new Headers(init?.headers);
    headers.set("X-Api-Secret", api.secret as string);

    let response: Response;
    try {
      response = await fetch(new URL(path, api.url as string), {
        ...init,
        headers,
      });
    } catch (cause) {
      throw new ApiError(
        "network",
        `The ${api.name} API is unreachable (${path})`,
        undefined,
        { cause },
      );
    }

    if (!response.ok) {
      throw new ApiError(
        "http",
        `The ${api.name} API answered ${response.status} (${path})`,
        response.status,
      );
    }

    const parsed = schema.safeParse(await response.json());
    if (!parsed.success) {
      throw new ApiError(
        "validation",
        `Unexpected response from the ${api.name} API (${path})`,
        response.status,
        { cause: parsed.error },
      );
    }
    return parsed.data;
  };
}

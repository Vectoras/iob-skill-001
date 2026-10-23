// The API client's core. Server-only (src/lib/server): the browser never calls the API.
// One generic apiFetch here, then a one-line function per endpoint.
import { API_SECRET, API_URL } from "$app/env/private";
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

export async function apiFetch<Schema extends z.ZodType>(
  path: string,
  schema: Schema,
  init?: RequestInit,
): Promise<z.output<Schema>> {
  // Both are required when the app runs (src/env.ts); only a build may lack them.
  const headers = new Headers(init?.headers);
  headers.set("X-Api-Secret", API_SECRET as string);

  let response: Response;
  try {
    response = await fetch(new URL(path, API_URL as string), {
      ...init,
      headers,
    });
  } catch (cause) {
    throw new ApiError(
      "network",
      `The API is unreachable (${path})`,
      undefined,
      {
        cause,
      },
    );
  }

  if (!response.ok) {
    throw new ApiError(
      "http",
      `The API answered ${response.status} (${path})`,
      response.status,
    );
  }

  const parsed = schema.safeParse(await response.json());
  if (!parsed.success) {
    throw new ApiError(
      "validation",
      `Unexpected response from the API (${path})`,
      response.status,
      { cause: parsed.error },
    );
  }
  return parsed.data;
}

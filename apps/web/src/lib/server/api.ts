// The api API's client: its address and secret (src/env.ts) and the shared core.
// Endpoints import apiFetch from here, one line each.
import { API_SECRET, API_URL } from "$app/env/private";
import { createApiFetch } from "./api-client";

export const apiFetch = createApiFetch({
  name: "api",
  url: API_URL,
  secret: API_SECRET,
});

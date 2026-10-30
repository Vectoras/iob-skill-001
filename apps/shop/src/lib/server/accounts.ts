// The accounts API's client: its address and secret (src/env.ts) and the shared core.
// Endpoints import apiFetch from here, one line each.
import { ACCOUNTS_SECRET, ACCOUNTS_URL } from "$app/env/private";
import { createApiFetch } from "api-client";

export const apiFetch = createApiFetch({
  name: "accounts",
  url: ACCOUNTS_URL,
  secret: ACCOUNTS_SECRET,
});

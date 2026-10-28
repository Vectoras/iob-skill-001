// The orders API's client: its address and secret (src/env.ts) and the shared core.
// Endpoints import apiFetch from here, one line each.
import { ORDERS_SECRET, ORDERS_URL } from "$app/env/private";
import { createApiFetch } from "api-client";

export const apiFetch = createApiFetch({
  name: "orders",
  url: ORDERS_URL,
  secret: ORDERS_SECRET,
});

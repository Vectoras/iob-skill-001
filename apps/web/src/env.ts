// Every environment variable the app reads. Private (server-only) and read when the
// app starts, by default. Values come from Compose (API_URL) and the app's .env
// (API_SECRET); a missing or invalid one stops the app at startup. While building they
// may be missing: a prerendered page makes no call to the API.
import { building } from "$app/env";
import { defineEnvVars } from "@sveltejs/kit/env";
import { z } from "zod";

const url = z.url();
const secret = z.string().min(1);

export const variables = defineEnvVars({
  API_URL: {
    schema: building ? url.optional() : url,
    description: "The API's address (in Compose, its service name).",
  },
  API_SECRET: {
    schema: building ? secret.optional() : secret,
    description: "The shared secret sent to the API in X-Api-Secret.",
  },
});

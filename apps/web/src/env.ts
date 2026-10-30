// Every environment variable the app reads. Private (server-only) and read when the
// app starts, by default. Values come from Compose (addresses) and the app's .env
// (secrets); a missing or invalid one stops the app at startup. While building they
// may be missing: a prerendered page makes no call to an API.
import { building } from "$app/env";
import { defineEnvVars } from "@sveltejs/kit/env";
import { z } from "zod";

const url = z.url();
const secret = z.string().min(1);

export const variables = defineEnvVars({
  // One pair per private API the app calls.
  API_URL: {
    schema: building ? url.optional() : url,
    description: "The api API's address (in Compose, its service name).",
  },
  API_SECRET: {
    schema: building ? secret.optional() : secret,
    description: "The shared secret sent to the api API in X-Api-Secret.",
  },
});

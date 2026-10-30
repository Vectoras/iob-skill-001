// Hello world: page data from the private APIs, through the frontend's server (a
// server load), rendered on the server. Delete it when the real work starts.
import { error } from "@sveltejs/kit";
import { ApiError } from "../lib/server/api-client";
import { getHelloWorlds } from "../lib/server/hello-world";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
  try {
    return { helloWorlds: await getHelloWorlds() };
  } catch (err) {
    if (err instanceof ApiError) error(502, err.message);
    throw err;
  }
};

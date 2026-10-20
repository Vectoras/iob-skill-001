// Hello world: page data from the API, through the frontend's server (a server load),
// rendered on the server. Delete it when the real work starts.
import { error } from "@sveltejs/kit";
import { ApiError } from "../lib/server/api";
import { getHelloWorld } from "../lib/server/hello-world";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
  try {
    return { helloWorld: await getHelloWorld() };
  } catch (err) {
    if (err instanceof ApiError) error(502, err.message);
    throw err;
  }
};

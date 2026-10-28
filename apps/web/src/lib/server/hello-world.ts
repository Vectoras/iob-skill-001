// Hello world: an endpoint is one line, here one per private API the app calls.
// Delete it when the real work starts.
import { HelloWorldSchema } from "shared-types";
import { apiFetch as api } from "./api";

export const getHelloWorlds = () =>
  Promise.all(
    [api].map((apiFetch) =>
      apiFetch("/hello-world/shared-types", HelloWorldSchema),
    ),
  );

// Hello world: an endpoint is one line, here one per private API the app calls.
// Delete it when the real work starts.
import { HelloWorldSchema } from "shared-types";
import { apiFetch as accounts } from "./accounts";
import { apiFetch as orders } from "./orders";

export const getHelloWorlds = () =>
  Promise.all(
    [accounts, orders].map((apiFetch) =>
      apiFetch("/hello-world/shared-types", HelloWorldSchema),
    ),
  );

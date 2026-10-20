// Hello world: an endpoint is one line. Delete it when the real work starts.
import { HelloWorldSchema } from "shared";
import { apiFetch } from "./api";

export const getHelloWorld = () => apiFetch("/hello-world", HelloWorldSchema);

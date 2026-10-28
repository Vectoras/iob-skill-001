// Hello world: proves a Zod schema is checked against a type generated from Go
// (`satisfies` fails to compile when the two drift apart). Delete it, with its test
// and hello_world.go, when the real work starts.
import { z } from "zod";
import type { HelloWorld } from "../generated/index";

export const HelloWorldSchema = z.object({
  message: z.string(),
}) satisfies z.ZodType<HelloWorld>;

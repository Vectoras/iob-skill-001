// Hello world: proves a Zod schema is checked against a type generated from Go
// (`satisfies` fails to compile when the two drift apart). Delete it, with its test,
// when the real work starts.
import { z } from "zod";
import type { HelloWorld } from "types";

export const HelloWorldSchema = z.object({
  message: z.string(),
}) satisfies z.ZodType<HelloWorld>;

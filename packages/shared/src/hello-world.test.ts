// Hello world: proves the unit tests run. Delete it with hello-world.ts.
import { describe, expect, it } from "vitest";
import { HelloWorldSchema } from "./hello-world";

describe("HelloWorldSchema", () => {
  it("parses what the API sends", () => {
    expect(HelloWorldSchema.parse({ message: "Hello world!" })).toEqual({
      message: "Hello world!",
    });
  });

  it("refuses anything else", () => {
    expect(() => HelloWorldSchema.parse({ text: "Hello world!" })).toThrow();
  });
});

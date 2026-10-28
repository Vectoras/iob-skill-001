// Placeholder: Vitest fails on a package with no test file. It proves only that the
// entry point loads; delete it with the package's first real test.
import { expect, test } from "vitest";
import * as entry from "./index";

test("the entry point loads", () => {
  expect(entry).toBeDefined();
});

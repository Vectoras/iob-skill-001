// The package's entry point: the TS types generated from Go (generated/, gitignored,
// written by `generate`) and the hand-written code beside them, such as the Zod
// schemas checked against those types.
export type * from "../generated/index";
export { HelloWorldSchema } from "./hello-world";

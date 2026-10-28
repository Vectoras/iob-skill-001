import { defineConfig, devices } from "@playwright/test";

// The tests run against a stack Compose already started: no `webServer` here. The
// address comes from the environment; the editing service sets it to the service name.
const baseURL = process.env.ADMIN_URL;
if (!baseURL) {
  throw new Error("ADMIN_URL is not set: the address of the running admin app");
}

export default defineConfig({
  testDir: "./src",
  forbidOnly: !!process.env.CI,
  reporter: "list",
  use: { baseURL, trace: "retain-on-failure" },
  projects: [{ name: "chromium", use: { ...devices["Desktop Chrome"] } }],
});

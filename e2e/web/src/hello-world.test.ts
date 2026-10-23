// Hello world: proves the whole chain in a browser (page, frontend's server, API, the
// shared component and its stylesheet) and the prerendered page. Delete it when the
// real work starts.
import { expect, test } from "@playwright/test";

test("the home page shows the API's message, styled", async ({ page }) => {
  await page.goto("/");
  const message = page.getByText("Hello world!");
  await expect(message).toBeVisible();
  await expect(message).toHaveCSS("color", "rgb(102, 51, 153)");
});

test("the about page is served", async ({ page }) => {
  await page.goto("/about");
  await expect(page.getByRole("heading", { name: "About" })).toBeVisible();
});

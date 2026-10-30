// Hello world: proves the whole chain in a browser (page, frontend's server, the
// private APIs if the app calls any, the component and its stylesheet) and the
// prerendered page. Delete it when the real work starts.
import { expect, test } from "@playwright/test";

test("the home page shows the message, styled", async ({ page }) => {
  await page.goto("/");
  // One message per private API the app calls, or one fixed message.
  const message = page.getByText("Hello world!").first();
  await expect(message).toBeVisible();
  await expect(message).toHaveCSS("color", "rgb(102, 51, 153)");
});

test("the about page is served", async ({ page }) => {
  await page.goto("/about");
  await expect(page.getByRole("heading", { name: "About" })).toBeVisible();
});

import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import Settings from "./../src/components/Settings";

test("Settings logic", async () => {
  // Arrange component
  render(<Settings />);

  // Button should be disabled on render
  expect(await screen.findByRole("button", { name: /pay/i })).toBeDisabled();

  // Act => user event
  userEvent.type(screen.getByPlaceholderText(/amount/i), "50");
  userEvent.type(screen.getByPlaceholderText(/add a note/i), "dinner");

  // Assertion
  expect(await screen.findByRole("button", { name: /pay/i })).toBeEnabled();
});

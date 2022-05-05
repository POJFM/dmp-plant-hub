// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom'
// import { render, screen } from "@testing-library/react";
// import userEvent from "@testing-library/user-event";
// import InitForm from "./components/InitForm";

// test("Init logic", () => {
//   // Arrange component
//   render(<InitForm />);

//   expect(screen.findByRole("button", { name: /initSave/i })).toBeDisabled();
//   expect(screen.findByRole("checkbox", { name: /limitsTrigger/i })).toBeEnabled();
//   expect(screen.findByRole("checkbox", { name: /scheduledTrigger/i })).toBeDisabled();

//   // limits on, scheduled off => everything apart from hour range is enabled
//   // limits off, scheduled on => moist limit, water amount limit and water level limit are disabled
//   // limits off, scheduled off => everything is disabled
//   // limits on, scheduled on => everything is enabled

//   // Button should be disabled on render

//   // Act => user event
//   userEvent.type(screen.getByPlaceholderText(/amount/i), "50");
//   userEvent.type(screen.getByPlaceholderText(/add a note/i), "dinner");

//   // Assertion
//   expect(await screen.findByRole("button", { name: /pay/i })).toBeEnabled();
// });

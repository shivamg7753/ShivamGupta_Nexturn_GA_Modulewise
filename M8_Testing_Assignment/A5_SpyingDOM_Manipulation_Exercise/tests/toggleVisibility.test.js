const toggleVisibility = require("../src/toggleVisibility");

describe("toggleVisibility", () => {
  let element;

  beforeEach(() => {
    // Create a mock element with a mock style object
    element = {
      style: {
        display: "block", // Initial display state
      },
    };
  });

  test("should change display to 'none' when element is visible", () => {
    toggleVisibility(element);
    expect(element.style.display).toBe("none");
  });

  test("should change display to 'block' when element is hidden", () => {
    element.style.display = "none"; // Set initial state to hidden
    toggleVisibility(element);
    expect(element.style.display).toBe("block");
  });
});

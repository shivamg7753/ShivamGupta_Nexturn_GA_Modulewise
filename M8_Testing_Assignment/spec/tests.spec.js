const { capitalize, reverseString } = require("../src/app");

describe("String Utilities", () => {
  it("should capitalize a word", () => {
    expect(capitalize("hello")).toBe("Hello");
    expect(capitalize("")).toBe("");
  });

  it("should reverse a string", () => {
    expect(reverseString("hello")).toBe("olleh");
    expect(reverseString("")).toBe("");
  });
});

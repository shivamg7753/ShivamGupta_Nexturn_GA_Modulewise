const { capitalize, reverseString } = require("../src/app");

describe("String Utilities", () => {
  describe("capitalize", () => {
    it("should capitalize the first letter of a word", () => {
      expect(capitalize("hello")).toBe("Hello");
    });

    it("should handle an empty string", () => {
      expect(capitalize("")).toBe("");
    });

    it("should handle single-character words", () => {
      expect(capitalize("a")).toBe("A");
    });

    it("should leave the rest of the string unchanged", () => {
      expect(capitalize("javaScript")).toBe("JavaScript");
    });
  });

  describe("reverseString", () => {
    it("should reverse a string", () => {
      expect(reverseString("hello")).toBe("olleh");
    });

    it("should handle empty strings", () => {
      expect(reverseString("")).toBe("");
    });

    it("should handle single-character strings", () => {
      expect(reverseString("a")).toBe("a");
    });

    it("should correctly reverse a palindrome", () => {
      expect(reverseString("madam")).toBe("madam");
    });
  });
});

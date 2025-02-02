const { getElement } = require("../src/app");

describe("Error Handling: Array Index", () => {
  describe("getElement", () => {
    it("should return the element at a valid index", () => {
      const arr = [10, 20, 30, 40];
      expect(getElement(arr, 0)).toBe(10);
      expect(getElement(arr, 2)).toBe(30);
    });

    it("should throw an error for a negative index", () => {
      const arr = [10, 20, 30, 40];
      expect(() => getElement(arr, -1)).toThrowError("Index out of bounds");
    });

    it("should throw an error for an out-of-range index", () => {
      const arr = [10, 20, 30, 40];
      expect(() => getElement(arr, 4)).toThrowError("Index out of bounds");
    });
  });
});

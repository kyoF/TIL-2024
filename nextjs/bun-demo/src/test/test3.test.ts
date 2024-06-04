// https://bun.sh/guides/test/run-tests
//
import { describe, test, expect } from "bun:test";

describe("math", () => {
  test("2 + 2 は、4 なのだ。", () => {
    expect(2 + 2).toEqual(4);
  });

  test("2 * 2 は、4 である。", () => {
    expect(2 * 2).toEqual(4);
  });
});

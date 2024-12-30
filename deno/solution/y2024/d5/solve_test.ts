import { assertEquals } from "@std/assert";
import { dirname, join } from "@std/path";
import { solution } from "./solve.ts";

const __dirname = dirname(import.meta.url);
const filePath = new URL(join(__dirname, "test1.txt"));

Deno.test("part-1", () => {
  const input = Deno.readTextFileSync(filePath).trim();
  const s1 = new solution(input);
  s1.part1();
  assertEquals(s1.res(), 0);
});

Deno.test("part-2", () => {
  const input = Deno.readTextFileSync(filePath).trim();
  const s2 = new solution(input);
  s2.part2();
  assertEquals(s2.res(), 0);
});

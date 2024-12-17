import { dirname, join } from "@std/path";

export class solve {
  input: string;
  ans: number;
  constructor(input: string, ans = 0) {
    this.input = input;
    this.ans = ans;
  }
  part1() {}
  part2() {}
  res() {
    return this.ans;
  }
}

export default function run() {
  const __dirname = dirname(import.meta.url);
  const filePath = new URL(join(__dirname, "input.txt"));
  const input = Deno.readTextFileSync(filePath).trim();
  const s1 = new solve(input);
  s1.part1();
  console.log("Part1 result ->", s1.res());
  const s2 = new solve(input);
  s2.part2();
  console.log("Part2 result ->", s2.res());
}

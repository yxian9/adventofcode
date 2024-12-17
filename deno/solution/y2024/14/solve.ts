import { dirname, join } from "@std/path";
import { IntsFromString } from "../../../lib/number.ts";

type Pt = {
  x: number;
  y: number;
};
type robot = {
  Pt: Pt;
  vx: number;
  vy: number;
};
const ncol = 101, nrow = 103, seconds = 100;

export class solver {
  input: string;
  ans = 0;
  strings: string[];
  robots: robot[] = [];

  constructor(input: string) {
    this.input = input;
    this.strings = this.input.split("\n");
    for (const line of this.strings) {
      const [x, y, vx, vy] = IntsFromString(line);
      const Pt = { x, y };
      this.robots.push({ Pt, vx, vy });
    }
  }
  part1() {
    for (let i = 0; i < this.robots.length; i++) {
      const r = this.robots[i];
      r.Pt.x = (r.Pt.x + r.vx * seconds) % ncol;
      r.Pt.y = (r.Pt.y + r.vy * seconds) % nrow;
      if (r.Pt.x < 0) {
        r.Pt.x += ncol;
      }
      if (r.Pt.y < 0) {
        r.Pt.y += nrow;
      }
    }
  }
  part2() {
    this.ans = 0;
  }
  res(): number {
    const quadrants = new Array(4).fill(0);
    for (let i = 0; i < this.robots.length; i++) {
      const r = this.robots[i].Pt;
      if (r.x < ncol >> 1 && r.y < nrow >> 1) {
        quadrants[0]++;
      } else if (r.x < ncol >> 1 && r.y > nrow >> 1) {
        quadrants[1]++;
      } else if (r.x > ncol >> 1 && r.y > nrow >> 1) {
        quadrants[2]++;
      } else if (r.x > ncol >> 1 && r.y < nrow >> 1) {
        quadrants[3]++;
      }
    }
    this.ans = 1;
    for (const i of quadrants) {
      this.ans *= i;
    }
    return this.ans;
  }
}

export default function run() {
  const __dirname = dirname(import.meta.url);
  const filePath = new URL(join(__dirname, "input.txt"));
  const input = Deno.readTextFileSync(filePath).trim();
  const s1 = new solver(input);
  s1.part1();
  console.log("Part1 result ->", s1.res());
  const s2 = new solver(input);
  s2.part2();
  console.log("Part2 result ->", s2.res());
}

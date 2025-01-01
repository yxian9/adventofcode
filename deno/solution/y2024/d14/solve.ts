import { IntsFromLine } from "../../../lib/number.ts";
import { exec } from "../../../lib/cli.ts";

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

export class solution {
  input: string;
  ans = 0;
  lines: string[];
  robots: robot[] = [];

  constructor(input: string) {
    this.input = input;
    this.lines = this.input.split("\n");
    for (const line of this.lines) {
      const [x, y, vx, vy] = IntsFromLine(line);
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

if (import.meta.main) {
  exec(solution, "input.txt");
}

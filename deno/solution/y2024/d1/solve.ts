import { exec } from "../../../lib/cli.ts";

export class solution {
  input: string;
  lines: string[];
  ans = 0;

  array1: number[] = [];
  array2: number[] = [];
  arr2Freq = new Map<number, number>();

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    for (const line of this.lines) {
      const [s1, s2] = line.split("   ");
      this.array1.push(Number(s1));
      this.array2.push(Number(s2));
    }
  }
  res(): number {
    return this.ans;
  }
  part1() {
    this.array1.sort((a, b) => a - b);
    this.array2.sort((a, b) => a - b);
    for (let index = 0; index < this.array1.length; index++) {
      const diff = Math.abs(this.array1[index] - this.array2[index]);
      this.ans += diff;
    }
  }
  part2() {
    this.buildFreq();
    for (const v1 of this.array1) {
      if (this.arr2Freq.has(v1)) {
        this.ans += v1 * this.arr2Freq.get(v1)!;
      }
    }
  }
  buildFreq() {
    for (const v of this.array2) {
      // init default value
      if (!this.arr2Freq.has(v)) {
        this.arr2Freq.set(v, 0);
      }
      const prev = this.arr2Freq.get(v)!;
      this.arr2Freq.set(v, prev + 1);
    }
  }
}

if (import.meta.main) {
  exec(solution, "input.txt");
}

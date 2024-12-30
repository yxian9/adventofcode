import { process, type SolutionMethods } from "../../../lib/utils.ts";

export class solution implements SolutionMethods {
  input: string;
  lines: string[];
  ans = 0;
  array1: number[] = [];
  array2: number[] = [];

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    for (const line of this.lines) {
      const [s1, s2] = line.split("   ");
      const n1 = Number(s1);
      const n2 = Number(s2);
      this.array1.push(n1);
      this.array2.push(n2);
    }
  }
  part1() {
    this.array1.sort((a, b) => a - b);
    this.array2.sort((a, b) => a - b);
    for (let index = 0; index < this.array1.length; index++) {
      const element = Math.abs(this.array1[index] - this.array2[index]);
      this.ans += element;
    }
  }
  part2() {
  }
  res(): number {
    return this.ans;
  }
}

if (import.meta.main) {
  process(solution, "./input.txt");
}

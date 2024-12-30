import { process, type SolutionMethods } from "../lib/utils.ts";

export class solution implements SolutionMethods {
  input: string;
  lines: string[];
  ans = 0;

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
  }
  part1() {
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

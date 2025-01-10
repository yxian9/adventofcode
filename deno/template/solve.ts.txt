import { exec } from "../../../lib/cli.ts";

export class solution {
  input: string;
  lines: string[];
  ans = 0;

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
  }
  res(): number {
    return this.ans;
  }
  res2(): number {
    return this.ans;
  }
  part1() {
  }
  part2() {
  }
}

if (import.meta.main) {
  console.time("Execution time for test1.txt");
  exec(solution, "test1.txt");
  console.timeEnd("Execution time for test1.txt");

  console.time("Execution time for input.txt");
  exec(solution, "input.txt");
  console.timeEnd("Execution time for input.txt");
}

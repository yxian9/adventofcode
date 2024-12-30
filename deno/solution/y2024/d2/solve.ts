import { exec } from "../../../lib/cli.ts";

export class solution {
  input: string;
  lines: string[];
  nums: number[][] = [];
  ans = 0;

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    for (const line of this.lines) {
      this.nums.push(this.stringToInts(line));
    }
  }
  stringToInts(line: string): number[] {
    const linesplit = line.split(" ");
    return linesplit.map((v) => Number(v));
  }
  res(): number {
    return this.ans;
  }
  isSave(arr: number[]): boolean {
    const startIdx = 1;
    const initDiff = arr[1] - arr[0]; // magic 1 0
    for (let idx = startIdx; idx < arr.length; idx++) { // magic 1
      const diff = arr[idx] - arr[idx - 1]; // magic 1
      if (diff * initDiff <= 0 || Math.abs(diff) < 1 || Math.abs(diff) > 3) {
        return false;
      }
    }
    return true;
  }
  isSave2(arr: number[], skip: number): boolean {
    let startIdx = 1;
    let initDiff = arr[1] - arr[0];

    if (skip === 0) {
      startIdx = 2;
      initDiff = arr[2] - arr[1];
    }
    if (skip === 1) {
      startIdx = 2;
      initDiff = arr[2] - arr[0];
    }

    for (let idx = startIdx; idx < arr.length; idx++) {
      if (idx === skip) {
        continue;
      }
      let diff = arr[idx] - arr[idx - 1];
      if (idx === skip + 1) {
        diff = arr[idx] - arr[idx - 2];
      }
      if (diff * initDiff <= 0 || Math.abs(diff) < 1 || Math.abs(diff) > 3) {
        return false;
      }
    }
    return true;
  }
  part1() {
    for (const arr of this.nums) {
      if (this.isSave(arr)) {
        this.ans++;
      }
    }
  }
  part2() {
    for (const arr of this.nums) {
      if (this.isSave(arr)) {
        this.ans++;
      } else {
        for (let skip = 0; skip < arr.length; skip++) {
          // const newArr: number[] = [];
          // for (let idx = 0; idx < arr.length; idx++) {
          //   if (idx === skip) {
          //     continue;
          //   }
          //   const v = arr[idx];
          //   newArr.push(v);
          // }
          const newArr = arr.filter((_, idx) => idx !== skip);

          if (this.isSave(newArr)) {
            this.ans++;
            break;
          }
          // if (this.isSave2(arr, skip)) {
          //   this.ans++;
          //   break;
          // }
        }
      }
    }
  }
}

if (import.meta.main) {
  exec(solution, "input.txt");
}

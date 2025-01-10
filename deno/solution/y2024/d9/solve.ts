import { exec } from "../../../lib/cli.ts";

export class solution {
  input: string;
  lines: string[];
  ans = 0;
  arr: number[] = [];

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    let fileID = 0;
    for (let idx = 0; idx < input.length; idx++) {
      const v = input[idx];
      let count = v.charCodeAt(0) - 48;
      if (idx % 2 == 0) {
        while (count--) {
          this.arr.push(fileID);
        }
        fileID++;
      } else {
        while (count--) {
          this.arr.push(-1);
        }
      }
    }
  }
  res(): number {
    for (const [idx, v] of this.arr.entries()) {
      if (v == -1) break;
      this.ans += idx * v;
    }
    return this.ans;
  }
  res2(): number {
    for (const [idx, v] of this.arr.entries()) {
      if (v == -1) continue;
      this.ans += idx * v;
    }
    return this.ans;
  }
  part1() {
    let l = 0, r = this.arr.length - 1;
    while (l < r) {
      if (this.arr[l] !== -1) {
        l++;
        continue;
      }
      if (this.arr[r] === -1) {
        r--;
        continue;
      }
      this.swap(l, r);
      l++;
      r--;
    }
  }
  swap(l: number, r: number) {
    const temp = this.arr[l];
    this.arr[l] = this.arr[r];
    this.arr[r] = temp;
  }
  swap2(l: number, r: number, len: number, fileID: number) {
    for (let idx = l; idx < l + len; idx++) {
      this.arr[idx] = -1;
    }
    for (let idx = r; idx < r + len; idx++) {
      this.arr[idx] = fileID;
    }
  }
  findFile(fileID: number): [number, number] {
    let left = 0, len = 0;
    for (const [idx, v] of this.arr.entries()) {
      if (v == fileID) {
        left = idx;
        break;
      }
    }
    for (let idx = left; idx < this.arr.length; idx++) {
      if (this.arr[idx] == fileID) {
        len++;
        continue;
      }
      break;
    }
    return [left, len];
  }
  findSpace(required: number, right: number): [number, boolean] {
    let leftIdx = 0, len = 0;
    for (let idx = 0; idx < right; idx++) {
      if (this.arr[idx] === -1) {
        if (len === 0) {
          leftIdx = idx;
        }
        len++;
        if (len === required) {
          return [leftIdx, true];
        }
        continue;
      }
      len = 0;
    }
    return [leftIdx, false];
  }
  part2() {
    let fileID = this.arr[this.arr.length - 1];
    for (; fileID >= 0; fileID--) {
      const [fileL, len] = this.findFile(fileID);
      const [freeL, found] = this.findSpace(len, fileL);
      if (found) {
        this.swap2(fileL, freeL, len, fileID);
      }
    }
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

import { exec } from "../../../lib/cli.ts";

export class solution {
  input: string;
  lines = [];
  ans = 0;
  start = "mul(";

  constructor(input: string) {
    this.input = input;
    // this.lines = input.split("\n");
  }
  res(): number {
    return this.ans;
  }
  getNumber(idx: number, line: string) {
    let num = 0, length = 0;
    for (; idx < line.length; idx++) {
      const char = line[idx];
      if (char >= "0" && char <= "9") {
        num *= 10;
        num += char.charCodeAt(0) - "0".charCodeAt(0);
        length++;
      } else {
        return { num, length };
      }
    }
    return { num, length };
  }
  part1() {
    for (let idx = 0; idx < this.input.length - 4; idx++) {
      if (this.input.slice(idx, idx + this.start.length) !== "mul(") {
        continue;
      }
      idx += this.start.length;

      const n1 = this.getNumber(idx, this.input);
      idx += n1.length;
      if (this.input[idx] !== ",") {
        continue;
      }
      idx++;

      const n2 = this.getNumber(idx, this.input);
      idx += n2.length;
      if (this.input[idx] !== ")") {
        continue;
      }
      // idx++;
      this.ans += n1.num * n2.num;
    }
  }
  part2() {
    let enable = true;
    for (let idx = 0; idx < this.input.length - 4; idx++) {
      if (
        idx + 7 < this.input.length &&
        this.input.slice(idx, idx + 7) === "don't()"
      ) {
        enable = false;
        // idx += 6;
        continue;
      }
      if (this.input.slice(idx, idx + 4) === "do()") {
        enable = true;
        // idx += 3;
        continue;
      }
      if (this.input.slice(idx, idx + this.start.length) !== "mul(") {
        continue;
      }
      idx += this.start.length;

      const n1 = this.getNumber(idx, this.input);
      idx += n1.length;
      if (this.input[idx] !== ",") {
        continue;
      }
      idx++;

      const n2 = this.getNumber(idx, this.input);
      idx += n2.length;
      if (this.input[idx] !== ")") {
        continue;
      }
      // idx++;
      if (enable) {
        this.ans += n1.num * n2.num;
      }
    }
  }
}

if (import.meta.main) {
  exec(solution, "input.txt");
  // part1 164730528
}

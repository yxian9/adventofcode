import { exec } from "../../../lib/cli.ts";
import { IntsFromLine } from "../../../lib/number.ts";

export class solution {
  input: string;
  lines: string[];
  ans = 0;
  connects: number[][] = [];
  records: number[][] = [];
  deps = new Map<number, number[]>();
  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    let firstPart = true;
    for (const line of this.lines) {
      if (line.length === 0) {
        firstPart = false;
        continue;
      }
      if (firstPart) {
        this.connects.push(IntsFromLine(line));
      } else {
        this.records.push(IntsFromLine(line));
      }
    }
  }
  res(): number {
    return this.ans;
  }
  isSorted(record: number[]) {
    const parentSet = new Set<number>();

    for (const parent of record) {
      if (!this.deps.has(parent)) {
        continue;
      }
      for (const child of this.deps.get(parent)!) {
        if (parentSet.has(child)) {
          return false;
        }
      }
      parentSet.add(parent);
    }
    return true;
  }
  part1() {
    for (const [parent, child] of this.connects) {
      if (!this.deps.has(parent)) {
        this.deps.set(parent, []);
      }
      this.deps.get(parent)!.push(child);
    }
    for (const record of this.records) {
      if (this.isSorted(record)) {
        const mid = record[record.length >> 1];
        this.ans += mid;
      }
    }
  }
  buildDep(record: number[]) {
    const res = new Map<number, number[]>();
    for (const parent of record) {
      if (this.deps.has(parent)) {
        const cur = [];
        for (const item of this.deps.get(parent)!) {
          if (record.includes(item)) {
            cur.push(item);
          }
        }
        res.set(parent, cur);
      }
    }

    return res;
  }
  buildInDegree(deps: Map<number, number[]>): Map<number, number> {
    const res = new Map<number, number>();
    for (const [_, childs] of deps) {
      for (const child of childs) {
        if (!res.has(child)) {
          res.set(child, 0);
        }
        const cur = res.get(child)!;
        res.set(child, cur + 1);
      }
    }
    return res;
  }
  topSort(record: number[]): number[] {
    const deps = this.buildDep(record);
    const inDegree = this.buildInDegree(deps);
    console.table(deps);
    console.table(inDegree);
    const res: number[] = [];
    const queue: number[] = [];
    for (const [key, value] of inDegree) {
      if (value === 0) {
        queue.push(key);
      }
    }
    while (queue.length > 0) {
      const leftHead = queue.shift()!;
      for (const child of deps.get(leftHead)!) {
        const cur = inDegree.get(child)!;
        inDegree.set(child, cur - 1);
        if (cur - 1 === 0) {
          queue.push(cur);
        }
      }
    }

    return res;
  }
  part2() {
    for (const [parent, child] of this.connects) {
      if (!this.deps.has(parent)) {
        this.deps.set(parent, []);
      }
      this.deps.get(parent)!.push(child);
    }
    for (const record of this.records) {
      if (!this.isSorted(record)) {
        console.log("check", record);
        const newSorted = this.topSort(record);
        const mid = newSorted[newSorted.length >> 1];
        this.ans += mid;
      }
    }
  }
}

if (import.meta.main) {
  // exec(solution, "input.txt");
  exec(solution, "test1.txt");
}

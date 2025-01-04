import { exec } from "../../../lib/cli.ts";
import { Dirs, Grid, Pt } from "../../../lib/dataStructure.ts";

export class solution {
  input: string;
  lines: string[];
  ans = 0;
  grid: Grid<string>;
  seen = new Set<string>();
  start = new Pt(0, 0);

  constructor(input: string) {
    this.input = input;
    this.lines = input.split("\n");
    const grid: string[][] = [];
    for (const [r, line] of this.lines.entries()) {
      grid.push(new Array(line.length).fill(""));
      for (let c = 0; c < line.length; c++) {
        const char = line[c];
        if (char === "^") {
          this.start = new Pt(r, c);
        }
        grid[r][c] = char;
      }
    }
    // console.table(grid);
    this.grid = new Grid(grid);
  }
  res(): number {
    return this.seen.size;
  }
  res2(): number {
    return this.ans;
  }
  dfs(p: Pt, angle: number) {
    this.seen.add(p.id);
    const nextP = p.pMove(Dirs[angle]);
    if (!this.grid.isInside(nextP)) {
      return;
    }
    if (this.grid.PGet(nextP) === "#") {
      this.dfs(p, (angle + 1) % 4);
    } else {
      this.dfs(nextP, angle);
    }
  }
  part1() {
    this.dfs(this.start, 0);
  }
  dfs2(p: Pt, angle: number, seen: Set<string>) {
    const nextP = p.pMove(Dirs[angle]);
    const id = nextP.id + ";" + String(angle);
    if (!this.grid.isInside(nextP)) {
      return;
    }
    if (seen.has(id)) {
      this.ans++;
      return;
    }
    if (this.grid.PGet(nextP) === "#") {
      seen.add(id);
      this.dfs2(p, (angle + 1) % 4, seen);
    } else {
      this.dfs2(nextP, angle, seen);
    }
  }
  part2() {
    for (let r = 0; r < this.grid.nrow; r++) {
      for (let c = 0; c < this.grid.ncol; c++) {
        const char = this.grid.get(r, c);
        if (char === ".") {
          const seen = new Set<string>();
          this.grid.set(r, c, "#");
          this.dfs2(this.start, 0, seen);
          this.grid.set(r, c, ".");
        }
      }
    }
  }
}

if (import.meta.main) {
  exec(solution, "test1.txt");
  exec(solution, "input.txt");
}

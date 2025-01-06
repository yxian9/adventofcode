export class Pt {
  r: number;
  c: number;
  constructor(r: number, c: number) {
    this.r = r;
    this.c = c;
  }
  pMove(p: Pt) {
    return new Pt(this.r + p.r, this.c + p.c);
  }
  Move(r: number, c: number) {
    return new Pt(this.r + r, this.c + c);
  }
  get id(): string {
    return `${this.r}:${this.c}`;
  }
}

export const Dirs = [
  new Pt(-1, 0),
  new Pt(0, 1),
  new Pt(1, 0),
  new Pt(0, -1),
];

export class Grid<T> {
  array: T[][];
  nrow: number;
  ncol: number;
  constructor(array: T[][]) {
    this.array = array;
    this.nrow = array.length;
    this.ncol = array[0].length;
  }
  isInside(p: Pt): boolean {
    return p.r >= 0 && p.r < this.nrow && p.c >= 0 && p.c < this.ncol;
  }
  PGet(p: Pt): T {
    return this.array[p.r][p.c];
  }
  get(r: number, c: number): T {
    return this.array[r][c];
  }
  PSet(p: Pt, v: T) {
    this.array[p.r][p.c] = v;
  }
  set(r: number, c: number, v: T) {
    return this.array[r][c] = v;
  }
}

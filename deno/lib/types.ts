export interface SolutionMethods {
  input: string;
  lines: string[];
  ans: number;
  part1(): void;
  part2(): void;
  res(): number;
  res2?(): number;
}

export type SolutionConstructor = {
  new (input: string): SolutionMethods;
};

export interface SolutionMethods {
  input: string;
  lines: string[];
  ans: number;
  part1(): void;
  part2(): void;
  res(): number;
}

export type SolutionConstructor = {
  new (input: string): SolutionMethods;
};

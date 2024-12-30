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
export function process(solution: SolutionConstructor, inputPath: string) {
  const input = Deno.readTextFileSync(inputPath).trim();
  const s1 = new solution(input);
  s1.part1();
  console.log("Part1 result ->", s1.res());
  const s2 = new solution(input);
  s2.part2();
  console.log("Part2 result ->", s2.res());
}

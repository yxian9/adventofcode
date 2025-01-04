import "@std/dotenv/load";
import { ensureDirSync, existsSync } from "@std/fs";
import { SolutionConstructor } from "./types.ts";
type Day = {
  day: string;
  year: string;
};

export function getCurrentDayAndYear(): Day {
  const now = new Date();
  const day = String(now.getDate());
  const year = String(now.getFullYear()); // Convert to string

  return { day, year };
}

export function templPaths({ year, day }: Day) {
  const id = `y${year}/d${day}`;
  const solve = `./solution/${id}/solve.ts`;
  const test = `./solution/${id}/solve_test.ts`;
  const inputPath = `./solution/${id}/input.txt`;
  const testInput = `./solution/${id}/test1.txt`;
  return { solve, inputPath, testInput, test, id };
}

export async function run(date: Day) {
  const { solve, inputPath } = templPaths(date);
  const { solution } = await import("." + solve) as {
    solution: SolutionConstructor;
  };
  exec(solution, inputPath);
}

export function exec(solution: SolutionConstructor, inputPath: string) {
  const input = Deno.readTextFileSync(inputPath).trim();
  const s1 = new solution(input);
  s1.part1();
  console.log("Part1 result ->", s1.res());
  const s2 = new solution(input);
  s2.part2();
  console.log("Part2 result ->", s2.res2?.() ?? s2.res());
}
export async function test(date: Day) {
  // const absoluteTestPath = resolve(
  //   fromFileUrl(Deno.mainModule),
  //   getPaths(date).test.replace(/^\.\/src/, ".."),
  // );
  const { test } = templPaths(date);
  const command = new Deno.Command(Deno.execPath(), {
    args: ["test", "-R", test],
  });
  const { stdout, stderr } = await command.output();
  console.log(new TextDecoder().decode(stdout).trim());
  if (stderr.length) console.error(new TextDecoder().decode(stderr).trim());
}

const AOC_SESSION = Deno.env.get("AOC_SESSION") ?? "";
const BASE_URL = "https://adventofcode.com";

export async function init(date: Day) {
  const { solve, inputPath, testInput, test, id } = templPaths(date);
  ensureDirSync(`solution/${id}`);
  if (existsSync(inputPath)) {
    console.log("AOC ", date, "skip download input~");
  } else {
    await fetchInput(date, inputPath);
  }
  if (!(existsSync(solve))) {
    await Deno.copyFile("./template/solve.ts.txt", solve);
  }
  if (!(existsSync(test))) {
    await Deno.copyFile("./template/solve_test.ts", test);
  }
  if (!(existsSync(testInput))) {
    await Deno.copyFile(
      "./template/test1.txt",
      testInput,
    );
  }
  console.log("AOC ", date, "init!");
}

async function fetchInput({ year, day }: Day, path: string) {
  if (!AOC_SESSION) throw new Error("AOC_SESSION is not set");
  const input = await fetch(`${BASE_URL}/${year}/day/${day}/input`, {
    headers: { Cookie: `session=${AOC_SESSION}` },
  }).then((x) => x.text());
  if (input.length == 0) throw new Error("No input found");
  return Deno.writeTextFile(path, input.trim());
}

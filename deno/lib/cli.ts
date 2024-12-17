import { ensureDirSync, existsSync } from "@std/fs";
import "@std/dotenv/load";
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
  const id = `y${year}/${day}`;
  const solve = `./solution/${id}/solve.ts`;
  const test = `./solution/${id}/solve_test.ts`;
  const input = `./solution/${id}/input.txt`;
  const testInput = `./solution/${id}/test1.txt`;
  return { solve, input, testInput, test, id };
}

export async function run(date: Day) {
  const { solve } = templPaths(date);
  const fn = await import("." + solve); //relative import path
  fn.default();
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
  const { solve, input, testInput, test, id } = templPaths(date);
  ensureDirSync(`solution/${id}`);
  if (!(existsSync(input))) await fetchInput(date, input);
  if (!(existsSync(solve))) {
    await Deno.copyFile(solve.replace(`./solution/${id}`, "template"), solve);
  }
  if (!(existsSync(test))) {
    await Deno.copyFile(test.replace(`./solution/${id}`, "template"), test);
  }
  if (!(existsSync(testInput))) {
    await Deno.copyFile(
      testInput.replace(`./solution/${id}`, "template"),
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
  return Deno.writeTextFile(path, input.trim());
}

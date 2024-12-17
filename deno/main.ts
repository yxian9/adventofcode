import { parseArgs } from "@std/cli/parse-args";
import { init, run, test, getCurrentDayAndYear } from "./lib/cli.ts";
const flags = parseArgs(Deno.args, {
  string: ["d", "day", "y", "year"],
  boolean: ["h", "help"],
});

async function main() {
  const verb = flags._[0]?.toString()?.toLowerCase();
  const day = flags.d || flags.day || getCurrentDayAndYear().day;
  const year = flags.y || flags.year || getCurrentDayAndYear().year;
  const help = flags.h || flags.help;

  if (help || +day < 1 || +day > 25) {
    return printUsage();
  }

  switch (verb) {
    case "init":
      await init({ day, year });
      break;
    case "test":
      await test({ day, year });
      break;
    // return await solve(year, day, part);
    // return await init(year, day);
    case "run":
      await run({ day, year });
      break;
    // return await solve(year, day, part);
    default:
      return printUsage();
  }
}

function printUsage() {
  console.log("Usage:");
  console.log(
    "    deno run solve [-d, --day <day>] [-p, --part <1 | 2>] [-y, --year <year>]"
  );
  console.log("    deno run init [-d, --day <day>] [-y, --year <year>]");
}

if (import.meta.main) {
  main()
    .then(() => Deno.exit(0))
    .catch((e) => {
      console.error(e);
      Deno.exit(1);
    });
}

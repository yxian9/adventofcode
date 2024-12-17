export function printNumber(n: number): number {
  n++;
  return 2;
}

export function IntsFromString(str: string): number[] {
  const result: number[] = [];
  const codeof0 = "0".charCodeAt(0);
  const codeof9 = "9".charCodeAt(0);

  let curDigitCount = 0, n = 0, sign = 1;

  for (const r of str) {
    const rCode = r.charCodeAt(0);
    if (r == "-" && curDigitCount == 0) {
      sign = -1;
      continue;
    }

    if (rCode >= codeof0 && rCode <= codeof9) {
      n *= 10;
      n += rCode - codeof0;
      curDigitCount++;
      continue;
    }

    if (curDigitCount > 0) {
      result.push(n * sign);
      curDigitCount = 0, n = 0, sign = 1;
    }
    sign = 1;
  }

  if (curDigitCount > 0) {
    result.push(n * sign);
  }

  return result;
}

export function IntsFromLine(str: string): number[] {
  const result: number[] = [];
  const codeof0 = "0".charCodeAt(0);
  const codeof9 = "9".charCodeAt(0);

  let digitLen = 0, num = 0, sign = 1;

  for (const r of str) {
    const rCode = r.charCodeAt(0);
    if (r == "-" && digitLen == 0) {
      sign = -1;
      continue;
    }

    if (rCode >= codeof0 && rCode <= codeof9) {
      num = num * 10 + (rCode - codeof0);
      digitLen++;
      continue;
    }

    if (digitLen > 0) {
      result.push(num * sign);
      digitLen = 0, num = 0, sign = 1;
    }
    sign = 1;
  }

  if (digitLen > 0) {
    result.push(num * sign);
  }

  return result;
}

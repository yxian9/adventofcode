from __future__ import annotations

import time
from pathlib import Path

from python.h import parse


def localParse(s: str, trim: bool = False):
    if trim:
        return int(s.split(":")[1].replace(" ", ""))
    return [int(v) for v in s.split(":")[1].split()]


def allPossible(time, dist):
    res = 0
    for i in range(time):
        if i * (time - i) > dist:
            res += 1
    return res


class Solution:
    def __init__(self, text: str):
        self.input = parse(text)
        self.time1 = localParse(self.input[0])
        self.dist1 = localParse(self.input[1])
        self.time2 = localParse(self.input[0], True)
        self.dist2 = localParse(self.input[1], True)
        self.ans1 = 1
        self.ans2 = 0

    def run1(self):
        for i, v in enumerate(self.time1):
            self.ans1 *= allPossible(v, self.dist1[i])

    def run2(self):
        self.ans2 = allPossible(self.time2, self.dist2)


def part1(text: str) -> int:
    s = Solution(text)
    s.run1()
    return s.ans1


def part2(text: str) -> int:
    s = Solution(text)
    s.run2()
    return s.ans2


def main():
    text = (Path(__file__).parent / "input.txt").read_text()

    start = time.perf_counter()
    result = part1(text)
    elapsed = time.perf_counter() - start
    print(f"p1 res -> {result} (Time taken: {elapsed:.4f}s)")

    start = time.perf_counter()
    result = part2(text)
    elapsed = time.perf_counter() - start
    print(f"p2 res -> {result} (Time taken: {elapsed:.4f}s)")


if __name__ == "__main__":
    main()

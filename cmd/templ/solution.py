from __future__ import annotations

import time
from pathlib import Path

from h import parse


class Solution:
    def __init__(self, text: str):
        self.input = parse(text)
        self.ans1 = 0
        self.ans2 = 0

    def run1(self):
        pass

    def run2(self):
        pass


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

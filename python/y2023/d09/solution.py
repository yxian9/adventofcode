from __future__ import annotations

import time
from collections import Counter, defaultdict
from pathlib import Path
from pprint import pprint as pp

import python.h as h


class Solution:
    def __init__(self, text: str):
        self.input = h.parse(text)
        self.ans1 = 0
        self.ans2 = 0

    def one(self, s: str):
        nums = [int(i) for i in s.split()]
        ret = nums[-1]
        while not all(i == 0 for i in nums):
            cur = []
            for i in range(1, len(nums)):
                cur.append(nums[i] - nums[i - 1])
            ret += cur[-1]
            nums = cur
        self.ans1 += ret

    def two(self, s: str):
        nums = [int(i) for i in s.split()]
        ret = nums[0]
        c = 1
        while not all(i == 0 for i in nums):
            cur = []
            for i in range(1, len(nums)):
                cur.append(nums[i] - nums[i - 1])
            c *= -1
            ret += c * cur[0]
            nums = cur
        self.ans2 += ret

    def run1(self):
        for s in self.input:
            self.one(s)

    def run2(self):
        for s in self.input:
            self.two(s)


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

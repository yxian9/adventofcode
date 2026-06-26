from __future__ import annotations

import re
import time
from collections import Counter, defaultdict
from dataclasses import dataclass
from math import gcd
from pathlib import Path
from pprint import pprint as pp

import python.h as h


@dataclass
class node:
    id: str
    l: str
    r: str

    def next(self, s):
        if s == "L":
            return self.l
        return self.r


def lcmm(xs):
    result = 1
    for n in xs:
        result = result * n // gcd(result, n)
    return result


class Solution:
    def __init__(self, text: str):
        self.input = h.parse(text)
        self.nodes: dict[str, node] = {}
        self.dir = ""
        self.ans1 = 0
        self.ans2 = 0

    def parse(self):
        ret = {}
        pattern = re.compile(r"(\w+) = \((\w+), (\w+)\)")
        for i, s in enumerate(self.input):
            if i == 0:
                self.dir = s
                continue
            if s == "":
                continue
            m = pattern.match(s)
            if not m:
                raise ValueError("not found")
            ret[m.group(1)] = node(m.group(1), m.group(2), m.group(3))
        self.nodes = ret

    def getdir(self, i: int) -> str:
        return self.dir[i % len(self.dir)]

    def run1(self):
        self.parse()
        start = "AAA"
        c = 0
        while start != "ZZZ":
            start = self.nodes[start].next(self.getdir(c))
            c += 1
        self.ans1 = c

    def run2(self):
        self.parse()
        starts = [self.nodes[n] for n in self.nodes if n.endswith("A")]
        pp(starts)
        ret = []

        for start in starts:
            c = 0
            while not start.id.endswith("Z"):
                dir = self.getdir(c)
                start = self.nodes[start.next(dir)]
                c += 1
                # pp(starts)
            ret.append(c)
        self.ans2 = lcmm(ret)


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

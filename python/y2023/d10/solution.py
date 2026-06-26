from __future__ import annotations

import time
from collections import Counter, defaultdict
from pathlib import Path
from pprint import pprint as pp

import python.h as h
from python.h import DIR4, Pt

pipe_dir_in = {
    "|": (0, 2),
    "-": (3, 1),
    "L": (2, 3),
    "J": (2, 1),
    "7": (0, 1),
    "F": (0, 3),
}

pipe_dir_out = {
    "|": (0, 2),
    "-": (3, 1),
    "L": (0, 1),
    "J": (0, 3),
    "7": (3, 2),
    "F": (1, 2),
}


def valid(c, i) -> bool:
    if c in pipe_dir_in:
        return i in pipe_dir_in[c]
    return False


class Solution:
    def __init__(self, text):
        self.input = h.parse(text)
        self.ans1 = 0
        self.ans2 = 0
        self.m = self.parse()
        self.start = Pt(1, 1)

    def parse(self) -> dict[Pt, str]:
        m = defaultdict(str)
        for i, row in enumerate(self.input):
            for j, c in enumerate(row):
                cur = Pt(i, j)
                match c:
                    case "S":
                        self.start = cur
                    case ".":
                        continue
                    case _:
                        m[cur] = c
        return m

    def find_start_end(self) -> list[Pt]:
        ret = []
        for i, dir in enumerate(DIR4):
            next = self.start.move(dir)
            if valid(self.m[next], i):
                ret.append(next)
        pp(ret)
        return ret

    def run1(self):
        self.parse()
        s, e = self.find_start_end()
        seen = {}
        seen[self.start] = True
        track = [[0] * len(self.input[0]) for _ in range(len(self.input))]
        total = 1
        cur = s
        seen[cur] = True
        while cur != e:
            for i in pipe_dir_out[self.m[cur]]:
                next = cur.move(DIR4[i])
                if valid(self.m[next], i) and not seen.get(next):
                    track[next.r][next.c] = total
                    total += 1
                    # pp(track)
                    cur = next
                    seen[cur] = True
                    break
        print(total)
        # pp(track)
        self.ans1 = (total + 1) // 2

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

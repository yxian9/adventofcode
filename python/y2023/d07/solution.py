from __future__ import annotations

import time
from collections import Counter
from curses import mouseinterval
from pathlib import Path
from pprint import pprint as pp

import python.h as h


def get_card(s: list[str]) -> dict[str, int]:
    rst = {}
    for l in s:
        p1, p2 = l.split()
        rst[p1] = int(p2)
    return rst


FRQ = [5, 4, 3, 2, 1]
m2 = {
    "A": -1,
    "K": -2,
    "Q": -3,
    "T": -5,
    "9": -6,
    "8": -7,
    "7": -8,
    "6": -9,
    "5": -10,
    "4": -11,
    "3": -12,
    "2": -13,
    "J": -14,
}
m = {
    "A": -1,
    "K": -2,
    "Q": -3,
    "J": -4,
    "T": -5,
    "9": -6,
    "8": -7,
    "7": -8,
    "6": -9,
    "5": -10,
    "4": -11,
    "3": -12,
    "2": -13,
}


def score(s, m=m):
    return tuple(m[i] for i in s)


class Solution:
    def __init__(self, text: str):
        self.input = h.parse(text)
        self.ans1 = 0
        self.ans2 = 0
        self.cards = get_card(self.input)
        self.rank_card = {}

    def get_rank(self):
        for card in self.cards:
            c = Counter(card)
            frq_str = [0 for _ in range(5)]
            cc = Counter(c.values())
            for k, v in cc.items():
                frq_str[k - 1] = v
            self.rank_card[card] = (
                "".join([str(c) for c in reversed(frq_str)]),
                score(card),
            )

    def get_rank2(self):
        for card in self.cards:
            j_c = 0
            c = Counter[str]()
            for char in card:
                if char == "J":
                    j_c += 1
                else:
                    c.update(char)
            most_comm = c.most_common(1)
            if most_comm:
                c[most_comm[0][0]] += j_c
            else:
                c["J"] = 5

            frq_str = [0 for _ in range(5)]
            cc = Counter(c.values())
            for k, v in cc.items():
                frq_str[k - 1] = v
            self.rank_card[card] = (
                "".join([str(c) for c in reversed(frq_str)]),
                score(card, m2),
            )

    def run1(self):
        self.get_rank()
        keys = [k for k in self.cards]
        keys.sort(key=lambda k: self.rank_card[k])
        for i, v in enumerate(keys, 1):
            # print(v, self.rank_card[v], i)
            self.ans1 += i * self.cards[v]

    def run2(self):
        self.get_rank2()
        keys = [k for k in self.cards]
        keys.sort(key=lambda k: self.rank_card[k])
        for i, v in enumerate(keys, 1):
            # print(v, self.rank_card[v], i)
            self.ans2 += i * self.cards[v]


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

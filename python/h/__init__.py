from __future__ import annotations

from typing import NamedTuple


class Pt(NamedTuple):
    r: int
    c: int

    def move(self, p: Pt) -> Pt:
        return Pt(self.r + p.r, self.c + p.c)


# up, right, down, left
DIR4 = (
    Pt(-1, 0),
    Pt(0, 1),
    Pt(1, 0),
    Pt(0, -1),
)

DIR8 = [
    Pt(-1, -1),
    Pt(-1, 0),
    Pt(-1, 1),
    Pt(0, -1),
    Pt(0, 1),
    Pt(1, -1),
    Pt(1, 0),
    Pt(1, 1),
]


def parse(text: str) -> list[str]:
    return text.strip().splitlines()

import unittest
from pathlib import Path

from .solution import part1, part2

DIR = Path(__file__).parent
WANT1 = 6440
WANT2 = 5905


class TestSolution(unittest.TestCase):
    def test_part1(self):
        text = (DIR / "test1.txt").read_text()
        self.assertEqual(part1(text), WANT1)

    def test_part2(self):
        text = (DIR / "test1.txt").read_text()
        self.assertEqual(part2(text), WANT2)


if __name__ == "__main__":
    unittest.main()

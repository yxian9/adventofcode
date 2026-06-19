from pathlib import Path
import unittest

from solution import part1, part2

TEST_INPUT = (Path(__file__).parent / "test1.txt").read_text()

WANT1 = 4361
WANT2 = 467835


class TestSolution(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(part1(TEST_INPUT), WANT1)

    def test_part2(self):
        self.assertEqual(part2(TEST_INPUT), WANT2)


if __name__ == "__main__":
    unittest.main()

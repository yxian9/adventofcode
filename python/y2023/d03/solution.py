import time
from pathlib import Path

import python.h as h


def is_symbol(ch: str) -> bool:
    return ch != "." and not ch.isdigit()


def part1(text: str) -> int:
    grid = h.parse(text)
    nrow, ncol = len(grid), len(grid[0])
    total = 0

    for i in range(nrow):
        j = 0
        while j < ncol:
            if grid[i][j].isdigit():
                num, valid = 0, False
                while j < ncol and grid[i][j].isdigit():
                    num = num * 10 + int(grid[i][j])
                    for dr, dc in h.DIR8:
                        nr, nc = i + dr, j + dc
                        if (
                            0 <= nr < nrow
                            and 0 <= nc < ncol
                            and is_symbol(grid[nr][nc])
                        ):
                            valid = True
                    j += 1
                if valid:
                    total += num
            else:
                j += 1
    return total


def part2(text: str) -> int:
    grid = h.parse(text)
    nrow, ncol = len(grid), len(grid[0])
    gears: dict[tuple[int, int], list[int]] = {}

    for i in range(nrow):
        j = 0
        while j < ncol:
            if grid[i][j].isdigit():
                num = 0
                adj_gears: set[tuple[int, int]] = set()
                while j < ncol and grid[i][j].isdigit():
                    num = num * 10 + int(grid[i][j])
                    for dr, dc in h.DIR8:
                        nr, nc = i + dr, j + dc
                        if 0 <= nr < nrow and 0 <= nc < ncol and grid[nr][nc] == "*":
                            adj_gears.add((nr, nc))
                    j += 1
                for g in adj_gears:
                    gears.setdefault(g, []).append(num)
            else:
                j += 1

    return sum(nums[0] * nums[1] for nums in gears.values() if len(nums) == 2)


def main():
    input_path = Path(__file__).parent / "input.txt"
    text = input_path.read_text()

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

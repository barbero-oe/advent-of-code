import re


def part_one(lines: list[str]) -> int:
    pattern = re.compile(r" (?=([0-9]+)\b(?!:).+\|.+\b\1\b)")
    winners = [winner for line in lines
               if (winner := pattern.findall(line))]
    points = [2 ** (len(winner) - 1) for winner in winners]
    return sum(points)



import re


def part_one(lines: list[str]) -> int:
    pattern = re.compile(r" (?=([0-9]+)\b(?!:).+\|.+\b\1\b)")
    winners = [winner for line in lines
               if (winner := pattern.findall(line))]
    points = [2 ** (len(winner) - 1) for winner in winners]
    return sum(points)

def part_two(lines: list[str]) -> int:
    cards = [line.split("|") for line in lines]
    winners = [len({*winnings.split(" ")} & {*numbers.split(" ")} - {""})
               for winnings, numbers in cards]
    card_count = [1 for _ in range(len(lines))]
    for index, points in enumerate(winners):
        start = index + 1
        end = points + start
        copies = card_count[index]
        for i in range(start, end):
            card_count[i] += copies
    return sum(card_count)


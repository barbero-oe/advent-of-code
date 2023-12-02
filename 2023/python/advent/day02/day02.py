from functools import reduce
import operator


def part_one(lines: list[str]) -> int:
    games = (parse_line(line) for line in lines)
    return sum(id for id, reveals in games if is_possible(reveals))

def is_possible(reveals: list[tuple[int, int, int]]) -> bool:
    maximum_reveal = (max(x) for x in zip(*reveals))
    possible = not any(operator.gt(*comp)
                       for comp in zip(maximum_reveal, (12, 13, 14)))
    return possible

def parse_line(line: str) -> tuple[int, list[tuple[int, int, int]]]:
    parts = line.split(":")
    game_id = int(parts[0].split(" ")[1])
    reveals = [parse_reveal(part)
               for part in parts[1].split(";")]
    return game_id, reveals

def parse_reveal(reveal_str: str) -> tuple[int, int, int]:
    mapping = {"red": 0, "green": 0, "blue": 0}
    reveals = reveal_str.split(",")
    for reveal in reveals:
        parts = reveal.strip().split(" ")
        ammount = int(parts[0])
        color = parts[1]
        mapping[color] += ammount
    return mapping["red"], mapping["green"], mapping["blue"]


def part_two(lines: list[str]) -> int:
    games = (parse_line(line) for line in lines)
    return sum(game_power(reveals) for _, reveals in games)

def game_power(reveals: list[tuple[int, int, int]]) -> int:
    minimum_dices = (max(x) for x in zip(*reveals))
    without_zero_dices = (x for x in minimum_dices if x > 0)
    return reduce(operator.mul, without_zero_dices, 1)
    

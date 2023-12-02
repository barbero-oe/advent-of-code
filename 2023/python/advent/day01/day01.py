import re

def part_one(lines: list[str]) -> int:
    return sum(extract_number(line) for line in lines)

def extract_number(line: str) -> int:
    numbers = [c for c in line if c.isdigit()]
    return int(numbers[0] + numbers[-1])

NUMBERS = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
STR_NUMS = "|".join(NUMBERS)
PATTERN = re.compile(f"(?=([1-9]|{STR_NUMS}))")

def part_two(lines: list[str]) -> int:
    return sum(extract_number_from_line(line) for line in lines)

def extract_number_from_line(line: str) -> int:
    numbers = PATTERN.findall(line)
    first = get_num_str(numbers[0])
    last = get_num_str(numbers[-1])
    return int(first + last)

def get_num_str(string_number: str) -> str:
    if string_number.isdigit():
        return string_number
    return str(NUMBERS.index(string_number) + 1)

import math
from itertools import chain, product
from typing import Iterator

def part_one(lines: list[str]) -> int:
    outer = range(len(lines))
    inner = range(len(lines[0]) if lines else 0)
    is_symbol = lambda i, j: "." != lines[i][j] and not lines[i][j].isdigit()
    symbols = [(i, j) for i, j
               in product(outer, inner)
               if is_symbol(i, j)]
    numbers = [surroundings(lines, i, j) for i, j in symbols]
    return sum(chain.from_iterable(numbers))

def surroundings(lines: list[str], i: int, j: int) -> list[int]:
    top = find_in_line(lines[i - 1], j) if i != 0 else []
    center = find_in_line(lines[i], j)
    bottom = find_in_line(lines[i + 1], j) if i + 1 != len(lines) else []
    return [*top] + [*center] + [*bottom]

def find_in_line(line: str, index: int) -> Iterator[int]:
    if line[index].isdigit():
        yield extract_number(line, index)
    else:
        if index != 0 and line[index - 1].isdigit():
            yield extract_number(line, index - 1)
        if index + 1 != len(line) and line[index + 1].isdigit():
            yield extract_number(line, index + 1)

def extract_number(line: str, index: int) -> int:
    start = index
    end = index
    for i in range(index, -1, -1):
        if not line[i].isdigit():
            break
        start = i
    for i in range(index, len(line)):
        if not line[i].isdigit():
            break
        end = i
    return int(line[start : end + 1])

def part_two(lines: list[str]) -> int:
    outer = range(len(lines))
    inner = range(len(lines[0]) if lines else 0)
    maybe_gear = lambda i, j: "*" == lines[i][j]
    symbols = [(i, j) for i, j
               in product(outer, inner)
               if maybe_gear(i, j)]
    numbers = [surroundings(lines, i, j) for i, j in symbols]
    ratios = [math.prod(surrounding) 
              for surrounding in numbers 
              if len(surrounding) == 2]
    return sum(ratios)

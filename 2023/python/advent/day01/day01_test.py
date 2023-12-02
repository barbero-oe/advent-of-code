from .day01 import part_one, part_two
from pathlib import Path

DIR = Path(__file__).parent

def read(path: Path) -> list[str]:
    with open(path) as f:
        return f.read().splitlines()

def test_sample_part_one():
    lines = read(DIR / "sample01.txt")
    result = part_one(lines)
    assert result == 142

def test_part_one():
    lines = read(DIR / "input")
    result = part_one(lines)
    assert result == 55607

def test_sample_part_two():
    lines = read(DIR / "sample02.txt")
    result = part_two(lines)
    assert result == 281

def test_part_two():
    lines = read(DIR / "input")
    result = part_two(lines)
    assert result == 55291

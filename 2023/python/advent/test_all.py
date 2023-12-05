import importlib
import pytest
from pathlib import Path
from typing import Any

DIR = Path(__file__).parent

def read(path: Path) -> list[str]:
    with open(path) as f:
        return f.read().splitlines()

@pytest.mark.parametrize("module_name,part,file,expect", [
    ("day01", "part_one", "sample01", 142),
    ("day01", "part_one", "input", 55607),
    ("day01", "part_two", "sample02", 281),
    ("day01", "part_two", "input", 55291),
    ("day02", "part_one", "sample01", 8),
    ("day02", "part_one", "input", 3059),
    ("day02", "part_two", "sample01", 2286),
    ("day02", "part_two", "input", 65371),
    ("day03", "part_one", "sample01", 4361),
    ("day03", "part_two", "sample01", 467835),
    ("day03", "part_two", "input", 80694070),
    ("day04", "part_one", "sample01", 13),
    ("day04", "part_one", "input", 25004),
])
def test_parts(module_name: str, part: str, file: str, expect: Any):
    module = importlib.import_module(f"advent.{module_name}.{module_name}")
    function = getattr(module, part)
    lines = read(DIR / module_name / file)
    result = function(lines)
    assert result == expect

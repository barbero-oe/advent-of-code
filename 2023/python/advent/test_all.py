import importlib
import pytest
from pathlib import Path
from typing import Any

DIR = Path(__file__).parent

def read(path: Path) -> list[str]:
    with open(path) as f:
        return f.read().splitlines()

@pytest.mark.parametrize("module_name,part,file,expect", [
    ("day01", "part_one", "sample01.txt", 142),
    ("day01", "part_one", "input.txt", 55607),
    ("day01", "part_two", "sample02.txt", 281),
    ("day01", "part_two", "input.txt", 55291),
])
def test_parts(module_name: str, part: str, file: str, expect: Any):
    module = importlib.import_module(f"advent.{module_name}.{module_name}")
    function = getattr(module, part)
    lines = read(DIR / module_name / file)
    result = function(lines)
    assert result == expect

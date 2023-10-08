package day06

func PartOne(lines []string) int {
	return findStart(lines[0], 4)
}

func PartTwo(lines []string) int {
	return findStart(lines[0], 14)
}

func findStart(line string, uniqueCount int) int {
	last := []rune(line[:uniqueCount])
	for i := uniqueCount; i < len(line); i++ {
		copy(last, last[1:uniqueCount])
		last[uniqueCount-1] = rune(line[i])
		set := make(map[rune]struct{})
		for _, value := range last {
			set[value] = struct{}{}
		}
		if len(set) == uniqueCount {
			return i + 1
		}
	}
	panic("Not found")
}

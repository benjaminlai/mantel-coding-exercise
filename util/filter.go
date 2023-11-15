package util

import "sort"

type Frequency struct {
	Input     string
	Frequency int
}

func FilterByTop(inputs []string, count int) []string {
	itemCount := make(map[string]int)

	for _, item := range inputs {
		itemCount[item]++
	}

	frequencies := make([]Frequency, len(itemCount))

	i := 0
	for item, count := range itemCount {
		frequencies[i] = Frequency{item, count}
		i++
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Frequency > frequencies[j].Frequency
	})

	topItems := []string{}
	for i := 0; i < count && i < len(frequencies); i++ {
		topItems = append(topItems, frequencies[i].Input)
	}

	return topItems
}

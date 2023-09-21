package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	items := string.Split(str, "\n")

	// Create a map to store the item counts.
	summary := make(map[string]int)

	// Iterate over the items and count their occurrences.
	for _, item := range items {
		if item[0] == '-' {
			// Get the item name.
			itemName := item[1:]

			// Increment the count for the item.
			summary[itemName]++
		}
	}

	// Return the summary map.
	return summary
}

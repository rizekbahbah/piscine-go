package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	// Split the string into a slice of items.
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

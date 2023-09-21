package piscine

// import "fmt"

func splitItems(s string) []string {
	var items []string
	var currentItem string

	if s == " " {
		items = append(items, "")
		items = append(items, "")
		return items
	}
	if s == "" {
		items = append(items, "")
		return items
	}
	for i, ch := range s {
		if ch == ' ' && i == len(s)-1 {
			items = append(items, "")
		}
		if ch != ' ' {
			currentItem += string(ch)
		} else if currentItem != "" {
			items = append(items, currentItem)
			currentItem = ""
		} else {
			// This condition handles consecutive spaces and treats them as a separate item.
			items = append(items, "")
		}
	}

	if currentItem != "" {
		items = append(items, currentItem)
	}

	return items
}

func ShoppingSummaryCounter(str string) map[string]int {
	items := splitItems(str)
	summary := make(map[string]int)

	for _, item := range items {
		summary[item]++
	}

	return summary
}

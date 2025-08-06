package utils

// GetNextOrder calculates the next order number based on a slice of items with order field
func GetNextOrder[T any](items []T, getOrder func(T) int) int {
	maxOrder := 0
	for _, item := range items {
		if order := getOrder(item); order > maxOrder {
			maxOrder = order
		}
	}
	return maxOrder + 1
}

// ReorderItemsWithTarget reorders items when a specific item is moved to a target position
// This is a simplified version that directly updates order values
func ReorderItemsWithTarget[T any](items []T, getOrder func(T) int, setOrder func(*T, int),
	getID func(T) interface{}, targetID interface{}, newOrder int) []T {

	if len(items) == 0 {
		return items
	}

	// Create a copy of items to avoid modifying the original slice
	result := make([]T, len(items))
	copy(result, items)

	// Find the target item and its current order
	var currentOrder int
	found := false

	for _, item := range result {
		if getID(item) == targetID {
			currentOrder = getOrder(item)
			found = true
			break
		}
	}

	if !found {
		return result
	}

	// Validate new order
	if newOrder < 1 {
		newOrder = 1
	}
	if newOrder > len(result) {
		newOrder = len(result)
	}

	// If the new order is the same as current order, no need to reorder
	if newOrder == currentOrder {
		return result
	}

	// Update order values based on the move direction
	for i := range result {
		itemOrder := getOrder(result[i])

		if getID(result[i]) == targetID {
			// Set target item to new order
			setOrder(&result[i], newOrder)
		} else if newOrder > currentOrder {
			// Moving down: shift items between current and new position up by 1
			if itemOrder > currentOrder && itemOrder <= newOrder {
				setOrder(&result[i], itemOrder-1)
			}
		} else {
			// Moving up: shift items between new and current position down by 1
			if itemOrder >= newOrder && itemOrder < currentOrder {
				setOrder(&result[i], itemOrder+1)
			}
		}
	}

	return result
}

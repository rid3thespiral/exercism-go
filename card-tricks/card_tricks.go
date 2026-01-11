package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if len(slice) <= index || index < 0 {
		return -1
	} else {
		return slice[index]
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if GetItem(slice, index) != -1 {
		slice[index] = value
		return slice
	} else {
		slice := append(slice, value)
		return slice
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	newValues := values[:]
	oldValues := slice[:]
	prependedItems := append(newValues, oldValues...)
	return prependedItems
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if GetItem(slice, index) != -1 {
		num := slice[index]
		i := 0
		for idx, item := range slice {
			if item != num {
				slice[i] = slice[idx]
				i++
			}
		}
		return slice[:i]
	} else {
		return slice
	}
}

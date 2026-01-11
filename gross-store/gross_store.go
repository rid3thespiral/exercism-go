package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["gross"] = 144
	units["small_gross"] = 120
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if units[unit] != 0 {

		bill[item] = bill[item] + units[unit]
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if units[unit] != 0 {
		if bill[item] != 0 {
			newQty := bill[item] - units[unit]
			if newQty < 0 {
				return false
			}
			if newQty == 0 {
				delete(bill, item)
			} else {
				bill[item] = newQty
			}
			return true
		}

		return false
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if bill[item] != 0 {
		return bill[item], true
	}
	return 0, false
}

package gross

var units = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	// why did
	// var something map[string]int
	// return something
	// not work here?
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	n, ok := units[unit]

	if !ok {
		return false
	}

	bill[item] = n
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	number, itemExists := bill[item]
	amount, unitExists := units[unit]

	if !itemExists {
		return false
	}

	if !unitExists {
		return false
	}

	if number-amount < 0 {
		return false
	}

	if number-amount == 0 {
		delete(bill, item)
	} else {
		bill[item] = bill[item] - amount
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	amount, ok := bill[item]

	if !ok {
		return 0, false
	}

	return amount, true
}

package purchase

// NeedsLicence determines whether a licence is need to drive a type of vehicle. Only "car" and "truck" require a licence.
func NeedsLicence(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	phrase := " is clearly the better choice."
	if option1 < option2 {
		return option1 + phrase
	}
	return option2 + phrase
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3:
		return originalPrice * 0.8
	case age < 10:
		return originalPrice * 0.7
	default:
		return originalPrice * 0.5
	}
}

package lasagna

// OvenTime how long a single layer needs to cook for in minutes
const OvenTime = 40

// RemainingOvenTime calculate how long your lasagne still needs to cook for
func RemainingOvenTime(timeElapsed int) int {
	return OvenTime - timeElapsed
}

// PreparationTime calculate how long it will take to prepare your layers
func PreparationTime(numberOfLayers int) int {
	return 2 * numberOfLayers
}

// ElapsedTime calculate how long you have been making lasagne for
func ElapsedTime(numberOfLayers int, ovenTime int) int {
	return PreparationTime(numberOfLayers) + ovenTime
}

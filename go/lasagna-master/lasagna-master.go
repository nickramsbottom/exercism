package lasagna

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		return len(layers) * 2
	}

	return len(layers) * prepTime
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, name := range layers {
		if name == "noodles" {
			noodles += 50
		}
		if name == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendList, layers []string) []string {
	return append(layers, friendList[len(friendList)-1])
}

func ScaleRecipe(quantites []float64, portions int) []float64 {
	var newAmounts []float64

	for _, amount := range quantites {
		newAmounts = append(newAmounts, (amount/2)*float64(portions))
	}

	return newAmounts
}

const ALPHABET = [..."abcdefghijklmnopqrstuvwxyz"];

export const isPangram = input => {
	const inputLower = input.toLowerCase();
	return ALPHABET.every(char => inputLower.includes(char));
};

function translateCodon(codon) {
	switch(codon) {
		case "AUG":
			return "Methionine";
		case "UUU":
		case "UUC":
			return "Phenylalanine";
		case "UUA":
		case "UUG":
			return "Leucine";
		case "UCU":
		case "UCC":
		case "UCA":
		case "UCG":
			return "Serine";
		case "UAU":
		case "UAC":
			return "Tyrosine";
		case "UGU":
		case "UGC":
			return "Cysteine";
		case "UGG":
			return "Tryptophan";
		case "UAA":
		case "UAG":
		case "UGA":
			return "STOP";
		default:
			throw Error("Invalid codon");
	}
}

export function translate(rna) {
	if (rna === undefined || rna.length === 0) {
		return [];
	}

	const firstThree = rna.slice(0, 3);
	const rest = rna.slice(3);

	const protein = translateCodon(firstThree);

	if (protein === "STOP") {
		return [];
	}

	return [protein].concat(translate(rest));
}

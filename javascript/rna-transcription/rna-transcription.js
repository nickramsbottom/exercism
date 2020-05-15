const mapping = {
	G: 'C',
	C: 'G',
	T: 'A',
	A: 'U',
};

export const toRna = dnas => [...dnas].map(dna => mapping[dna]).join('');
export class Matrix {
	constructor(matrixRows) {
		this._matrix = this.parse(matrixRows);
	}

	get rows() {
		return copy(this._matrix);
	}

	get columns() {
		return transpose(this._matrix);
	}

	parse(string) {
		return string.split('\n').map(row => row.split(' ').map(Number))
	}
}

function copy(matrix) {
	return matrix.map(row => row.map(number => number));
}

function transpose(matrix) {
	const height = matrix.length;
	const width = matrix[0].length;

	const transposed = [];

	for (let i = 0; i < width; i++) {
		const col = [];
		for (let j = 0; j < height; j++) {
			col.push(matrix[j][i]);
		}
		transposed.push(col);
	}

	return transposed;
}

export class Matrix {
	constructor(string) {
		this.matrix = string.split('\n').map(row => row.split(' ').map(Number))
	}

	get rows() {
		return this.matrix;
	}

	get columns() {
		const height = this.matrix.length;
		const width = this.matrix[0].length;

		const transposed = [];

		for (let i = 0; i < width; i++) {
			const col = [];
			for (let j = 0; j < height; j++) {
				col.push(this.matrix[j][i]);
			}
			transposed.push(col);
		}

		return transposed;
	}
}

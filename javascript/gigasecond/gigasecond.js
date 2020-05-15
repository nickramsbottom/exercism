const GIGASECONDS = Math.pow(10, 12);

export const gigasecond = (startDate) => new Date(startDate.getTime() + GIGASECONDS);

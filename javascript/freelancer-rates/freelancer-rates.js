// @ts-check

const HOURS_WORKED_PER_DAY = 8;
const BILLABLE_DAYS_PER_MONTH = 22;

/**
 * The day rate, given a rate per hour
 *
 * @param {number} ratePerHour
 * @returns {number} the rate per day
 */
export function dayRate(ratePerHour) {
  return ratePerHour * HOURS_WORKED_PER_DAY;
}

/**
 * Calculates the rate per month
 *
 * @param {number} ratePerHour
 * @param {number} discount for example 20% written as 0.2
 * @returns {number} the rounded up monthly rate
 */
export function monthRate(ratePerHour, discount) {
  const rate = applyDiscount(dayRate(ratePerHour) * BILLABLE_DAYS_PER_MONTH, discount);
  return Math.ceil(rate);
}

/**
 * Calculates the number of days in a budget, rounded down
 *
 * @param {number} budget the total budget
 * @param {number} ratePerHour the rate per hour
 * @param {number} discount to apply, example 20% written as 0.2
 * @returns {number} the number of days
 */
export function daysInBudget(budget, ratePerHour, discount) {
  const discountedDayRate = applyDiscount(dayRate(ratePerHour),discount);
  const daysLeftDecimal = budget / discountedDayRate;
  return Math.floor(daysLeftDecimal);
}

/**
 * Applies a discount to the value
 *
 * @param {number} value
 * @param {number} percentage for example 20% written as 0.2
 * @returns {number} the discounted value
 */
function applyDiscount(value, percentage) {
  return value * (1 - percentage);
}

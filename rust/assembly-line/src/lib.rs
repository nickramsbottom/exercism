#![allow(unused)]

const PRODUCTION_RATE_DEFAULT: f64 = 221.0;

pub fn production_rate_per_hour(speed: u8) -> f64 {
    if speed == 0 {
        return 0.0;
    } else if speed <= 4 {
        return PRODUCTION_RATE_DEFAULT * speed as f64;
    } else if speed <= 8 {
        return PRODUCTION_RATE_DEFAULT * 0.9 * speed as f64;
    } else {
        return PRODUCTION_RATE_DEFAULT * 0.77 * speed as f64;
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}

#![allow(unused)]

const PRODUCTION_RATE_DEFAULT: f64 = 221.0;

pub fn production_rate_per_hour(speed: u8) -> f64 {
    match speed {
        0 => 0.0,
        1..=4 => PRODUCTION_RATE_DEFAULT * speed as f64,
        5..=8 => PRODUCTION_RATE_DEFAULT * 0.9 * speed as f64,
        _ => PRODUCTION_RATE_DEFAULT * 0.77 * speed as f64,
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}

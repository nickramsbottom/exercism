#![allow(unused)]

pub fn production_rate_per_hour(speed: u8) -> f64 {
    if speed == 0 {
        return 0.0;
    } else if speed <= 4 {
        return 221.0 * speed as f64;
    } else if speed <= 8 {
        return 0.9 * 221.0 * speed as f64;
    } else {
        return 0.77 * 221.0 * speed as f64;
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}

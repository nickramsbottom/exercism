const PRODUCTION_RATE_DEFAULT: f64 = 221.0;

pub fn production_rate_per_hour(speed: u8) -> f64 {
    let s = speed as f64;
    match speed {
        0 => 0.0,
        1..=4 => PRODUCTION_RATE_DEFAULT * s,
        5..=8 => PRODUCTION_RATE_DEFAULT * 0.9 * s,
        _ => PRODUCTION_RATE_DEFAULT * 0.77 * s,
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}

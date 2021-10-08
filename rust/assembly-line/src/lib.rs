pub fn production_rate_per_hour(speed: u8) -> f64 {
    let speed_rate: f64 = 221.0 * speed as f64;
    let ok_rate = match speed {
        0 => 0.0,
        1..=4 => 1.0,
        5..=8 => 0.9,
        _ => 0.77,
    };
    speed_rate * ok_rate
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}

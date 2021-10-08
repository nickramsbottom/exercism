pub fn production_rate_per_hour(speed: u8) -> f64 {
    let speed_rate: f64 = 221.0 * speed as f64;
    match speed {
        0 => 0.0,
        1..=4 => speed_rate,
        5..=8 => speed_rate * 0.9,
        _ => speed_rate * 0.77,
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}

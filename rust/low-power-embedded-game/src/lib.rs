pub fn divmod(dividend: i16, divisor: i16) -> (i16, i16) {
    let quot = dividend / divisor;
    let rem = dividend % divisor;
    (quot, rem)
}

pub fn evens<T>(iter: impl Iterator<Item = T>) -> impl Iterator<Item = T> {
    iter.enumerate().filter(|(i, _)| i % 2 == 0).map(|(_, v)| v)
}

pub struct Position(pub i16, pub i16);
impl Position {
    pub fn manhattan(&self) -> i16 {
        self.0.abs() + self.1.abs()
    }
}

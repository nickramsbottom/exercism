/// various log levels
#[derive(Clone, PartialEq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
    Debug,
}
/// primary function for emitting logs
pub fn log(level: LogLevel, message: &str) -> String {
    match level {
        LogLevel::Info => info(message),
        LogLevel::Warning => warn(message),
        LogLevel::Error => error(message),
        LogLevel::Debug => debug(message),
    }
}
pub fn info(message: &str) -> String {
    return "[INFO]: ".to_owned() + message;
}
pub fn warn(message: &str) -> String {
    return "[WARNING]: ".to_owned() + message;
}
pub fn error(message: &str) -> String {
    return "[ERROR]: ".to_owned() + message;
}
pub fn debug(message: &str) -> String {
    return "[DEBUG]: ".to_owned() + message;
}

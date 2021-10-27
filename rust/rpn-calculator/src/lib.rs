use std::ops::{Add, Div, Mul, Sub};

pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

fn calculate_and_push<T: Fn(i32, i32) -> i32>(vec: &mut Vec<i32>, func: T) -> Option<()> {
    let rhs = vec.pop()?;
    let lhs = vec.pop()?;
    let tmp = func(lhs, rhs);
    vec.push(tmp);
    Some(())
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut vec = Vec::new();

    for input in inputs {
        match input {
            CalculatorInput::Add => calculate_and_push(&mut vec, i32::add)?,
            CalculatorInput::Subtract => calculate_and_push(&mut vec, i32::sub)?,
            CalculatorInput::Multiply => calculate_and_push(&mut vec, i32::mul)?,
            CalculatorInput::Divide => calculate_and_push(&mut vec, i32::div)?,
            CalculatorInput::Value(val) => vec.push(*val),
        };
    }

    if vec.len() != 1 {
        return None;
    }

    vec.pop()
}

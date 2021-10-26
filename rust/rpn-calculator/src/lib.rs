#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut vec = Vec::new();

    // there's a lot of repetition in here, can we group
    // everything except value into a closure?
    for input in inputs {
        match input {
            CalculatorInput::Add => {
                let first = vec.pop()?;
                let second = vec.pop()?;
                vec.push(first + second);
            }
            CalculatorInput::Subtract => {
                let first = vec.pop()?;
                let second = vec.pop()?;
                vec.push(second - first);
            }
            CalculatorInput::Multiply => {
                let first = vec.pop()?;
                let second = vec.pop()?;
                vec.push(first * second);
            }
            CalculatorInput::Divide => {
                let first = vec.pop()?;
                let second = vec.pop()?;
                vec.push(second / first);
            }
            CalculatorInput::Value(val) => vec.push(*val),
        }
    }

    if vec.len() != 1 {
        return None;
    }

    vec.pop()
}

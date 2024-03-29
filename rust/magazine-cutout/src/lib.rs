use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut available_words = HashMap::new();

    for word in magazine {
        available_words
            .entry(word)
            .and_modify(|value| *value += 1)
            .or_insert(1);
    }

    for word in note {
        match available_words.get_mut(word) {
            None | Some(0) => {
                return false;
            }
            Some(count) => *count -= 1,
        }
    }

    true
}

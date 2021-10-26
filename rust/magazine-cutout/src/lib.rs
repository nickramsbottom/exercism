use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut available_words = HashMap::new();

    for word in magazine {
        *available_words.entry(word).or_insert(0) += 1;
    }

    for word in note {
        if !available_words.contains_key(word) || available_words[word] == 0 {
            return false;
        }

        *available_words.get_mut(word).unwrap() -= 1;
    }

    true
}

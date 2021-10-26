pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        (self.health == 0).then(|| Player {
            health: 100,
            mana: if self.level < 10 { None } else { Some(100) },
            level: self.level,
        })
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        match self.mana {
            None => {
                if self.health < mana_cost {
                    self.health = 0
                } else {
                    self.health = self.health - mana_cost;
                }
                return 0
            }
            Some(mana) => {
                if mana <= mana_cost {
                    return 0
                } else {
                    self.mana = Some(mana - mana_cost);
                    return mana_cost * 2
                }
            }
        }
    }
}

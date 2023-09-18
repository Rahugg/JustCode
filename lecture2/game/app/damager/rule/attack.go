package rule

import (
	"game/app/damager/characters/types"
	"math/rand"
)

type Enemy struct {
	Health uint64
	Attack uint64
}

type CharacterData struct {
	Health uint64
	Attack uint64
	Weapon types.AttackType
}

func (enemy *Enemy) Fight(hero CharacterData, damageMultiplier int, firstAttack bool) bool {
	if firstAttack {
		for enemy.Health > 0 && hero.Health > 0 {
			// Generate a random integer in the range from 1 to 5
			randNum := rand.Intn(5) + 1

			//Hero attacks
			//fmt.Println("Enemy:", enemy.Health)
			damage := (uint64(randNum) + hero.Attack) * uint64(damageMultiplier)
			if damage > enemy.Health {
				damage = enemy.Health // Cap the damage to not exceed the remaining health
			}
			enemy.Health -= damage
			//fmt.Println("Hero:", hero.Health)

			//enemy attacks
			randNum = rand.Intn(5-1+1) + 1
			damage = uint64(randNum) + enemy.Attack
			if damage > hero.Health {
				damage = hero.Health // Cap the damage to not exceed the remaining health
			}
			hero.Health -= damage
		}
	} else {
		for enemy.Health > 0 && hero.Health > 0 {
			//rand number in range from 1 to 5
			randNum := rand.Intn(5-1+1) + 1

			//enemy attacks
			//fmt.Println("Enemy:", enemy.Health)
			damage := uint64(randNum) + enemy.Attack
			//fmt.Println("Hero:", hero.Health)
			if damage > hero.Health {
				damage = hero.Health
			}
			hero.Health -= damage

			//Hero attacks
			randNum = rand.Intn(5-1+1) + 1
			damage = (uint64(randNum) + hero.Attack) * uint64(damageMultiplier)
			if damage > enemy.Health {
				damage = enemy.Health
			}
			enemy.Health -= damage
		}
	}
	return hero.Health > 0
}

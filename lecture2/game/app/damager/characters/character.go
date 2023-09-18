package characters

import (
	"fmt"
	"game/app/damager/characters/types"
)

var ErrInputIsNotMatch = fmt.Sprintf("input type is not match")

type Character struct{}

// the attack will vary with RNG

func (c *Character) CanAttackFirst(attackType types.AttackType) bool {
	switch attackType {
	case types.ArcherWeaponType:
		return true
	case types.WarriorWeaponType:
		return false
	case types.WizardWeaponType:
		return true
	default:
		return false
	}

}

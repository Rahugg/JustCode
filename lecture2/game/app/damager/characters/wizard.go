package characters

import (
	"game/app/damager/characters/types"
	"game/app/damager/rule"
)

type Wizard struct {
	*Character
}

// the attack will vary with RNG

func (w *Wizard) IsMatch(characterType types.CharacterType) bool {
	return types.WizardType == characterType
}

//if he uses his own weapon, he can attack first, because he is a master of his own weapon

func (w *Wizard) OwnWeapon(attackType types.AttackType) bool {
	return types.WizardWeaponType == attackType
}

func (w *Wizard) Attack(input interface{}, damageMultiplier int) string {
	data, ok := input.(rule.CharacterData)
	if !ok {
		return ErrInputIsNotMatch
	}
	RaidBoss := rule.Enemy{Health: 1500, Attack: 50}

	FirstAttack := w.CanAttackFirst(data.Weapon)

	alive := RaidBoss.Fight(data, damageMultiplier, FirstAttack)
	if !alive {
		return "something happened!"
	}

	return "alive"
}

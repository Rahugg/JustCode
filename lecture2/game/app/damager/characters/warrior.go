package characters

import (
	"game/app/damager/characters/types"
	"game/app/damager/rule"
)

type Warrior struct {
	*Character
}

func (wr *Warrior) IsMatch(characterType types.CharacterType) bool {
	return types.WarriorType == characterType
}

//if he uses his own weapon, he can attack first, because he is a master of his own weapon

func (wr *Warrior) OwnWeapon(attackType types.AttackType) bool {
	return types.WarriorWeaponType == attackType
}

func (wr *Warrior) Attack(input interface{}, damageMultiplier int) string {
	data, ok := input.(rule.CharacterData)
	if !ok {
		return ErrInputIsNotMatch
	}
	RaidBoss := rule.Enemy{Health: 1500, Attack: 50}

	FirstAttack := wr.CanAttackFirst(data.Weapon)

	alive := RaidBoss.Fight(data, damageMultiplier, FirstAttack)

	if !alive {
		return "good luck next time!"
	}

	return "alive"
}

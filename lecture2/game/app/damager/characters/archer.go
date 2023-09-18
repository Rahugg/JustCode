package characters

import (
	"game/app/damager/characters/types"
	"game/app/damager/rule"
)

type Archer struct {
	*Character
}

func (a *Archer) IsMatch(characterType types.CharacterType) bool {
	//fmt.Println("a")
	return types.ArcherType == characterType
}

//if he uses his own weapon, he can attack first, because he is a master of his own weapon

func (a *Archer) OwnWeapon(attackType types.AttackType) bool {
	return types.ArcherWeaponType == attackType
}

func (a *Archer) Attack(input interface{}, damageMultiplier int) string {

	data, ok := input.(rule.CharacterData)
	if !ok {
		return ErrInputIsNotMatch
	}
	RaidBoss := rule.Enemy{Health: 1500, Attack: 50}

	FirstAttack := a.CanAttackFirst(data.Weapon)

	alive := RaidBoss.Fight(data, damageMultiplier, FirstAttack)

	if !alive {
		return "good luck next time!"
	}

	return "alive"
}

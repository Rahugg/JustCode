package damager

import (
	"fmt"
	"game/app/damager/characters"
	"game/app/damager/characters/types"
)

type Validator interface {
	Attack(data interface{}, damageMultiplier int) string
	OwnWeapon(attackType types.AttackType) bool
	IsMatch(characterType types.CharacterType) bool
}

type Manager struct {
	validators []Validator
}

func NewManager(validators ...Validator) *Manager {
	return &Manager{validators: validators}
}

func (m *Manager) Execute(data interface{}, characterType types.CharacterType, attackType types.AttackType) error {

	for _, validator := range m.validators {
		if !validator.IsMatch(characterType) {
			continue
		}
		var damageMultiplier = 1
		if result := validator.OwnWeapon(attackType); result {
			damageMultiplier *= 2
		}
		err := validator.Attack(data, damageMultiplier)
		if err != "alive" || err == characters.ErrInputIsNotMatch {
			return fmt.Errorf("unlucky: %s", err)
		}
		return nil
	}

	return fmt.Errorf("cannot validate the character")
}

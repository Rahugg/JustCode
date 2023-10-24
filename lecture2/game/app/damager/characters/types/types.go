package types

type AttackType string
type CharacterType string

const (
	WizardWeaponType  AttackType = "Staff"
	WarriorWeaponType AttackType = "Sword"
	ArcherWeaponType  AttackType = "Bow"
)
const (
	WizardType  CharacterType = "Wizard"
	WarriorType CharacterType = "Warrior"
	ArcherType  CharacterType = "Archer"
	GoblinType  CharacterType = "Goblin"
)

package main

import (
	"fmt"
	"game/app/damager"
	validator2 "game/app/damager/characters"
	"game/app/damager/characters/types"
	"game/app/damager/rule"
)

func main() {
	var (
		character        = &validator2.Character{}
		wizardValidator  = &validator2.Wizard{Character: character}
		archerValidator  = &validator2.Archer{Character: character}
		warriorValidator = &validator2.Warrior{Character: character}

		validationManager = damager.NewManager(archerValidator, wizardValidator, warriorValidator)
	)
	archerValidatorData := rule.CharacterData{
		Health: 750,
		Attack: 65,
		Weapon: types.ArcherWeaponType,
	}

	// archer fight RANDOM
	fmt.Println("1")
	err := validationManager.Execute(archerValidatorData, types.ArcherType, archerValidatorData.Weapon)
	if err != nil {
		//the result: unlucky, character died: "good luck next time!"
		fmt.Println("the result:", err)

	} else {
		fmt.Println("Congratulations!🥳🎉🎉 Archer won")
	}

	fmt.Println("2")
	//validation error
	err = validationManager.Execute(archerValidatorData, types.WarriorType, archerValidatorData.Weapon)
	if err != nil {
		//failed to execute err: validation failed err: input type is not match
		fmt.Println("the result:", err)
	} else {
		fmt.Println("Congratulations!🥳🎉🎉 Archer won")
	}

	warriorValidatorData := rule.CharacterData{
		Health: 1300,
		Attack: 45,
		Weapon: types.WarriorWeaponType,
	}

	fmt.Println("3")
	//Warrior fight RANDOM
	err = validationManager.Execute(warriorValidatorData, types.WarriorType, warriorValidatorData.Weapon)
	if err != nil {
		//the result: unlucky, character died: "good luck next time!"
		fmt.Println("the result:", err)
	} else {
		fmt.Println("Congratulations!🥳🎉🎉 Warrior won")
	}

	fmt.Println("4")
	//validation error
	err = validationManager.Execute(warriorValidatorData, types.WizardType, warriorValidatorData.Weapon)
	if err != nil {
		//failed to execute err: validation failed err: input type is not match
		fmt.Println("the result:", err)
	} else {
		fmt.Println("Congratulations!🥳🎉🎉 Warrior won!!")
	}

	wizardValidatorData := rule.CharacterData{
		Health: 800,
		Attack: 70,
		Weapon: types.WizardWeaponType,
	}

	fmt.Println("5")
	//Wizard fight RANDOM
	err = validationManager.Execute(wizardValidatorData, types.WizardType, wizardValidatorData.Weapon)
	if err != nil {
		//the result: unlucky, character died: "good luck next time!"
		fmt.Println("the result:", err)
	} else {
		fmt.Println("Congratulations!🥳🎉🎉 Wizard won")

	}

	fmt.Println("6")
	//validation error
	err = validationManager.Execute(wizardValidatorData, types.ArcherType, wizardValidatorData.Weapon)
	if err != nil {
		//failed to execute err: validation failed err: input type is not match
		fmt.Println("the result:", err)
	} else {
		fmt.Println("Congratulations!🥳🎉🎉 Wizard won!")
	}

	fmt.Println("7")
	err = validationManager.Execute(wizardValidatorData, types.GoblinType, wizardValidatorData.Weapon)
	if err != nil {
		//failed to execute err: cannot validate the character
		fmt.Println("the result:", err)
	}

}

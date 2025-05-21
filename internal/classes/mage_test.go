package classes

import (
	"go_study_rpg_service/internal/skills"
	"testing"
)

type FakeWeapon struct{}

func (f *FakeWeapon) Name() string          { return "fake" }
func (f *FakeWeapon) DamageBonus() int      { return 7 }
func (f *FakeWeapon) SpecialEffect() string { return "none" }
func (f *FakeWeapon) Attack() string        { return "attacking" }

func TestMageAttack(t *testing.T) {
	mage := &Mage{Manapool: 100, Equipment: &FakeWeapon{}}
	if mage.Equipment.Attack() != 27 {
		t.Errorf("Esperava 27, recebeu %d", mage.Equipment.Attack())
	}
}

func TestMageCastSpel(t *testing.T) {
	mage := &Mage{Manapool: 50, Weapon: &FakeWeapon{}}
	spell := skills.Spell{Name: "IceSpell", CostMana: 30, Element: "Ice"}
	result := mage.UseSkill(spell)
	if mage.Manapool != 20 {
		t.Errorf("Mana deveria ser 20, recebeu %d", mage.Manapool)
	}
	if result == "Not enough mana!" {
		t.Errorf("Mage deveria conseguir lançar magia")
	}
}

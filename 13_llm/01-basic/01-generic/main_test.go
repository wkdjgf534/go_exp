package main

import "testing"

// BenchmarkInventoryAdd измеряет производительность добавления предметов в инвентарь
func BenchmarkInventoryAdd(b *testing.B) {
	inv := NewInventory[Weapon]("Test Inventory", 1000)
	sword := Weapon{Name: "Sword", BaseDmg: 4, DiceSize: 8}

	for b.Loop() {
		inv.Add(sword)
		// Сбрасывается инвентарь, чтобы не превысить вместимость
		if inv.Count() >= 100 {
			inv = NewInventory[Weapon]("Test Inventory", 1000)
		}
	}
}

// BenchmarkRollDice измеряет производительность броска кубика
func BenchmarkRollDice(b *testing.B) {
	for b.Loop() {
		_ = RollDice(20)
	}
}

// BenchmarkRollDamage измеряет производительность расчета урона
func BenchmarkRollDamage(b *testing.B) {
	for b.Loop() {
		_ = RollDamage(2, 6)
	}
}

// BenchmarkMax измеряет производительность функции Max
func BenchmarkMax(b *testing.B) {
	for b.Loop() {
		_ = Max(10, 20)
	}
}

// BenchmarkMaxString измеряет производительность для строк
func BenchmarkMaxString(b *testing.B) {
	for b.Loop() {
		_ = Max("abc", "xyz")
	}
}

package main

import "testing"

// BenchmarkDirectCall - прямой вызов метода конкретного типа.
// Компилятор инлайт вызов, накладных расходов почти нет.
func BenchmarkDirectCall(b *testing.B) {
	p := Paladin{name: "P", damage: 12, armor: 18, healPower: 15}
	b.ResetTimer()
	for b.Loop() {
		_ = p.Attack
	}
}

// BenchmarTypeAssertion - стоимость извлечения конкретного типа из интерфейса.
// Это операция сравнения типа в рантайме.
func BenchmarTypeAssertion(b *testing.B) {
	var c Combatant = Paladin{name: "P", damage: 12, armor: 18, healPower: 15}
	b.ResetTimer()
	for b.Loop() {
		if _, ok := c.(Paladin); ok {
			// Успешное извлечение.
		}
	}
}

// BenchmarkTypeSwitch — стоимость ветвления по типу.
// Сравниваем с простым вызовом метода.
func BenchmarkTypeSwitch(b *testing.B) {
	var c Combatant = Paladin{name: "P", damage: 12, armor: 18, healPower: 15}

	for b.Loop() {
		switch c.(type) {
		case Paladin:
			// Ветка паладина.
		case Fighter:
			// Ветка воина.
		case Skeleton:
			// Ветка скелета.
		}
	}
}

// BenchmarkNilCheck — стоимость проверки интерфейса на nil.
func BenchmarkNilCheck(b *testing.B) {
	var c Combatant = Paladin{name: "P", damage: 12, armor: 18, healPower: 15}

	for b.Loop() {
		_ = c != nil
	}
}

package main

import "testing"

// BenchmarkSumDamage_Int измеряет производительность для базового типа int.
// Компилятор использует stenciling (трафарет), вызовы инлайнятся.
func BenchmarkSumDamage_Int(b *testing.B) {
	damages := make([]int, 100)
	for i := range damages {
		damages[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumDamage(damages)
	}
}

// BenchmarkSumDamage_CustomInt измеряет производительность для кастомного типа.
// Stenciling работает так же эффективно, как и для int.
func BenchmarkSumDamage_CustomInt(b *testing.B) {
	damages := make([]CustomInt, 100)
	for i := range damages {
		damages[i] = CustomInt(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumDamage(damages)
	}
}

// BenchmarkCastFireball_Interface измеряет вызов методов через интерфейс.
// Здесь компилятор может использовать словари (dictionaries) или itab,
// что добавляет накладные расходы на динамическую диспетчеризацию.
func BenchmarkCastFireball_Interface(b *testing.B) {
	targets := []Targetable{
		&Goblin{name: "G1", hp: 100},
		&Dragon{name: "D1", hp: 100},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Восстанавливаем HP, чтобы цель не "умерла" и цикл отработал полностью
		for _, t := range targets {
			if g, ok := t.(*Goblin); ok {
				g.hp = 100
			}
			if d, ok := t.(*Dragon); ok {
				d.hp = 100
			}
		}
		CastFireball[Targetable, int](targets, 1)
	}
}

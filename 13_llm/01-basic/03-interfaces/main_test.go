package main

import "testing"

// BenchmarkDirectCall измеряет прямой вызов метода конкретной структуры.
// Компилятор может инлайнить этот вызов.
func BenchmarkDirectCall(b *testing.B) {
	r := Rogue{name: "Test", hp: 20}

	for b.Loop() {
		_ = r.Act()
	}
}

// BenchmarkInterfaceAllocation измеряет стоймость упаковки значения в интерфейс.
// Если значение не помещается в указатель (или требует аллокации для itab),
// это может вызвать выделение памяти в куче.
func BenchmarkInterfaceAllocation(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		var a Actor = Rogue{name: "Test", hp: 20}
		_ = a
	}
}

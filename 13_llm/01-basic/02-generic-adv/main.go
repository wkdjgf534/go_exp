package main

import (
	"fmt"
	"strings"
)

// Ограничения и множества типов

// Numeric - множество типов, ограничение Т числовыми типами.
// Оператор - означает "любой тип, чьим базовым типом является..."
// Это позволяет использовать математические операторы (+, -, *, /) внутри дженериков.

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Targetable - интерфейсное ограничение.
// Требуется, чтобы тип Т имел конкретные методы.
// В отличие от Numeric, здесь компилятор может использовать "словари" (dictionaries)
// для диспетчеризации методов, если Т является интерфейсом.

type Targetable interface {
	Name() string
	TakeDamage(amount int)
	IsAlive() bool
}

// CustomInt - пример кастомного типа.
// Благадаря оператору - в ограничении Numeric, это тип подойдёт для математических функций.
type CustomInt int

// Сущеости игры

// Goblin
type Goblin struct {
	name string
	hp   int
}

func (g *Goblin) Name() string       { return g.name }
func (g *Goblin) TakeDamage(dmg int) { g.hp -= dmg }
func (g *Goblin) IsAlive() bool      { return g.hp > 0 }

// Dragon
type Dragon struct {
	name string
	hp   int
}

func (d *Dragon) Name() string       { return d.name }
func (d *Dragon) TakeDamage(dmg int) { d.hp -= dmg }
func (d *Dragon) IsAlive() bool      { return d.hp > 0 }

// Обобщённые функции с ограничениями

// SumDamage суммирует урон из слайса.
// Ограничение [D Numeric] гарантирует, что опервтор += доступен для типа D.
func SumDamage[D Numeric](damages []D) D {
	var total D // нулевое значение для чисел - 0
	for _, d := range damages {
		total += d
	}

	return total
}

// CastFireball принимает урон по площади ко всем живым целям.
// Здесь два параметра типа:
// [T Targetable] - Цели должна уметь ролучать урон.
// [D Numeric] - урон должен быть числом.
func CastFireball[T Targetable, D Numeric](targets []T, baseDamage D) {
	fmt.Printf("\n Волшебник кастует Огненный шар (базовый урон: %v)!\n", baseDamage)

	for _, target := range targets {
		if !target.IsAlive() {
			continue
		}

		// Приведение D к int, так как TakeDamage требует int.
		// В реальном коде лучше проектировать интерфейсы так, чтобы приведения не требовались,
		// но здесь это демонстрирует работу с разными числовыми типами.
		dmg := int(baseDamage)

		fmt.Printf("   -> %s получает %d урона. \n", target.Name(), dmg)
		target.TakeDamage(dmg)

		if !target.IsAlive() {
			fmt.Printf("   %s повержен!\n", target.Name())
		}
	}
}

// Map - классический паттерн функционального программирования.
// Преобразует слайс типа Т в слайс типа U с помощью функции f.
func Map[T any, U any](items []T, f func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = f(item)
	}

	return result
}

// Демонстрация
func main() {
	fmt.Println("=== AD&D СИМУЛЯЦИЯ: ОГРАНИЧЕНИЯ И ИНСТАНЦИРОВАНИЕ ===")

	// 1. Использование множества типов (Numeric)
	fmt.Println("\n1. Расчёт общего урона (SumDamage)")

	// Неявное инстанцирование: компилятор сам выводит D = int
	physicalDmg := []int{4, 5, 3, 8}
	totalPhys := SumDamage(physicalDmg)
	fmt.Printf("   Физический урон (int): %d\n", totalPhys)

	// Неявное инстанцирование с кастомным типом (благодаря ~)
	magicDmg := []CustomInt{10, 15, 20}
	totalMagic := SumDamage(magicDmg)
	fmt.Printf("   Магический урон (CustomInt): %d\n", totalMagic)

	// Явное инстанцирование: указываем тип напрямую
	floatDmg := []float64{1.5, 2.5, 3.0}
	totalFloat := SumDamage(floatDmg)
	fmt.Printf("   Урон ядом (float64): %.1f\n", totalFloat)

	// 2. Использование интерфейсного ограничения (Targetable)
	fmt.Println("\n2. Боевое столкновение (CastFireball):")

	// Создаём поле боя
	g1 := &Goblin{name: "Гоблин-лучник", hp: 12}
	g2 := &Goblin{name: "Гоблин-воин", hp: 15}
	dragon := &Dragon{name: "Красный Дракон", hp: 45}

	// Cлайс интерфейсов Targetable
	// Обратите внимание: T здесь выводится как Targetable (интерфейс),
	// а не как конкретный тип *Goblin или *Dragon.
	battlefield := []Targetable{g1, g2, dragon}

	// Явное инстанцирование: CastFireball[Targetable, int]
	CastFireball(battlefield, 14)

	// 3. Паттерн Map (Преобразование данных)
	fmt.Println("\n3. Сброс трофеев (Паттерн Map):")

	// Извлекаем имена всех выживших
	survivors := []Targetable{}
	for _, t := range battlefield {
		if t.IsAlive() {
			survivors = append(survivors, t)
		}
	}

	// Map преобразует []Targetable в []string
	// T = Targetable, U = string
	names := Map(survivors, func(t Targetable) string {
		return strings.ToUpper(t.Name())
	})

	fmt.Printf("   Выживание после боя: %v\n", names)
	fmt.Println("\n=== СУМУЛЯЦИЯ ЗАВЕРШЕНА ===")
}

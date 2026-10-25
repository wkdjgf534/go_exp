package main

import (
	"fmt"
)

// Контракты поведения (интерфейсы)

// Actor - базовый интерфейс для любой сущности, которая может действовать в раунде
// Это контракт: "Ты должен уметь называть себя и совершать действие".
type Actor interface {
	Name() string
	Act() string
}

// Healer - узкий интерфейс для тех, кто умееет лечить.
// Демонстрация принципа разделения интерфейсов (ISP).
type Healer interface {
	Heal(target Actor) string
}

// Lootable - пустой интерфейс (any).
// Означает "любой тип". Используется для лута, потому что лут может быть чем угодно.
type Lootable any

// Реализации (Конкретные типы)

// Rogue - Плут. Реализует Actor неявно.
type Rogue struct {
	name string
	hp   int
}

func (r Rogue) Name() string { return r.name }
func (r Rogue) Act() string {
	return fmt.Sprintf("%s наносит удар из тени (Sneak Attack)!", r.name)
}

// Wizard - волшебник. Реализует Actor и может лечить (Healer).
type Wizard struct {
	name  string
	hp    int
	spell string
}

func (w Wizard) Name() string { return w.name }
func (w Wizard) Act() string  { return fmt.Sprintf("%s кастует %s!", w.name, w.spell) }
func (w Wizard) Heal(target Actor) string {
	return fmt.Sprintf("%s ракладывает 'Лечение легких ран' на %s.", w.name, target.Name())
}

// Goblin - Гоблин. Тоже реализует Actor!
// Компилятору всё равно, что это монстр. Если есть методы Name() и Act(), это Actor.
type Goblin struct {
	name string
	hp   int
}

func (g Goblin) Name() string { return g.name }
func (g Goblin) Act() string {
	return fmt.Sprintf("%s бросается с кривым кинжалом!", g.name)
}

// Функции, работающие с интерфейсами

// TakeTurn выполняет ход сущности.
// Фкнкция принимает интерфейс Actor. Ей не важно, кто перед ней: Плут, Волшебник или Гоблин.
func TakeTurn(a Actor) {
	fmt.Printf("[%s] %s\n", a.Name(), a.Act())
}

// DistributeLoot раздаёт лут.
// Использует пусто интерфейс (any), потому что лут может быть строкой, числом или структурой.
func DistributeLoot(recipient Actor, item Lootable) {
	fmt.Printf("   -> %s получает предмет: %v (тип: %T)\n", recipient.Name(), item, item)
}

// TryHeal пытается вылечить цель, если целитель умее это делать.
// Здесь используется type assertion (проверка типа), чтобы узнать, реализует ли Actor интерфейс Healer.
func TryHeal(healer Actor, target Actor) {
	// Попытка привести Actor к Healer
	if h, ok := healer.(Healer); ok {
		fmt.Printf("  %s\n", h.Heal(target))
	} else {
		fmt.Printf("  %s не умеет лечить.\n", healer.Name())
	}
}

// Демонстрация

func main() {
	fmt.Println("=== AD&D СИМУЛЯЦИЯ: ИНТЕРФЕЙСЫ И КОНТРАКТЫ ===n")

	// Создаём сущности
	rogue := Rogue{name: "Торин", hp: 24}
	wizard := Wizard{name: "Элина", hp: 18, spell: "Магическая стрела"}
	goblin := Goblin{name: "Снит", hp: 8}

	// 1. Полиморфизм("1. Боевой раунд (полиморфизм):")

	// Слайс интерфейсов Actor.
	// Сюда можно положить любой тип, который реализует Name() и Act().
	turnOrder := []Actor{rogue, goblin, wizard, goblin}

	for _, actor := range turnOrder {
		TakeTurn(actor)
	}
	fmt.Println()

	// 2. Работа с пустым интерфейсом (any)
	fmt.Println("2. Раздача лута (пустой интерфейс any):")

	// Лут может быть чем угодно: строкой, числом, структурой.
	DistributeLoot(rogue, "Кинжал +1")
	DistributeLoot(wizard, 50) // Золотые монеты
	DistributeLoot(goblin, struct{ Name string }{Name: "Гнилое яблоко"})
	fmt.Println()

	// 3. Проверка возможностей (Type Assertion)
	fmt.Println("3. Попытки лечения (проверка интерфейса):")

	TryHeal(wizard, rogue) // Волшебник умеет лечить
	TryHeal(rogue, wizard) // Плут не умеет лечить
	fmt.Println()

	var nilActor Actor
	fmt.Printf("   nilActor == nil: %v\n", nilActor == nil)

	// Типизированный nil (указатель на Rogue, который равен nil, но упакован в интерфейс)
	var nilRogue *Rogue
	var wrappedActor Actor = nilRogue

	// Внимание! wrappedActor Не равен nil, потому что внутри него есть информация о типе (*Rogue).
	fmt.Printf("   wrappedActor == nil: %v (тип внутри: %T)\n", wrappedActor == nil, wrappedActor)

	fmt.Println("\n=== СИМУЛЯЦИЯ ЗАВЕРШЕНА ===")
}

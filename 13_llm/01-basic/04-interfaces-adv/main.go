package main

import "fmt"

// Композиция интерфейсов

// Маленькие, сфокумированные интерфейсы.
// Каждый описывает ровно одно умение.

// Attacker - контракт того, кто умеет атаковать.
type Attacker interface {
	Attack() int
}

// Defender - контракт того, кто умеет защищаться.
type Defender interface {
	Defend() int
}

// Healer - контракт того, кто умеет лечить.
type Healer interface {
	Heal(target string) int
}

// Combatant - композиция; тот, кто умеет и атакоывть, и защищаться.
// Мы не переопределяем методы, а встраиваем существующие интерфейсы.
type Combatant interface {
	Attacker
	Defender
}

// DivineWarrior - композиция трех интерфейсов.
// Пример: Паладин или Жрец-воин.
type DivineWarrior interface {
	Combatant
	Healer
}

// Конкретные типы (реализации)

// Fighter - воин. Умеет атаковать и защищаться, но не лечить.
type Fighter struct {
	name   string
	damage int
	armor  int
}

func (f Fighter) Attack() int { return f.damage }
func (f Fighter) Defend() int { return f.armor }

// Paladin - паладин. Реализзует все три интерфейса, значит удовлетворяет DivineWarrior.
type Paladin struct {
	name      string
	damage    int
	armor     int
	healPower int
}

func (p Paladin) Attack() int            { return p.damage }
func (p Paladin) Defend() int            { return p.armor }
func (p Paladin) Heal(target string) int { return p.healPower }

// Skeleton - скелет (нежить). Атакует и защищается, имеет особое свойство.
// Это свойство не входит в общий интерфейс, и мы покажим, как ео извлечь.
type Skeleton struct {
	name        string
	damage      int
	armor       int
	weakToBlunt bool // уязвим к дробящему оружию
}

func (s Skeleton) Attack() int { return s.damage }
func (s Skeleton) Defend() int { return s.armor }

// Динамическая диспетчерезация

// ResolveCombat демонстрирует полиморфный вызов.
// Функция не знает конкретных типов - только контракт Combatant.
func ResolveCombat(c Combatant) {
	// Здесь происходит динамическая диспетчерезация:
	// рантайм смотрит в itab и вызывает метод реального типа.
	atk := c.Attack()
	def := c.Defend()
	fmt.Printf("   Атака: %2d, Защита: %2d (тип: %T)\n", atk, def, c)
}

// Type switch: Истинное зрение

// IdentifyCreature использует type switch для определения тактики против врага.
// Это аналог заклинания Истинное Зрение: раскрываем природу существа.
func IdentifyCreature(c Combatant) {
	// switch по типу: переменная v получает конкретное значение.
	// В каждой ветке v имеет соответствующий тип.
	switch v := c.(type) {
	case Paladin:
		fmt.Printf("  [Истинное Зрение] %s - Паладин! Может лечить. Фокус на контроль.\n", v.name)
	case Skeleton:
		// Здесь мы получаем доступ к полю, которого нет в интерфейсе.
		if v.weakToBlunt {
			fmt.Printf("  [Истинное Зрение] %s - Скелет! Уязвим к дробящему оружию. Берём булаву.\n", v.name)
		} else {
			fmt.Printf("  [Истинное Зрение] %s - Скелет. Обычная нежить.\n", v.name)
		}
	case Fighter:
		fmt.Printf("  [Истинное Зрение] %s - Воин. Прямой бой.\n", v.name)
	default:
		// Ветка для всех остальных типов, реализующих Combatant.
		fmt.Printf("  [Истинное Зрение] Неизвестное существо типа %T.\n", v)
	}
}

// Type assertion: опознание

// TryHealParty проверяет, может ли боец лечить, и если да - лечит.
// Используем type assertion с проверкой, чтобы избежать паники.
func TryHealParty(c Combatant) {
	// Безопасная форма: если тип не совпадает, ok будет false, паники не будет.
	if healer, ok := c.(Healer); ok {
		healed := healer.Heal("Раненый союзник")
		fmt.Printf("  %s лечит союзника на %d HP.\n", fmt.Sprintf("%T", c), healed)
	} else {
		fmt.Printf("  %T не умеет лечить. Ищем жреца.\n", c)
	}
}

// Nil - ловушка

// GetPatrol возвращает отряд патруля.
// Это демонстрация классической ошибки: возврат типизированного nil как интерфейса.
func GetPatrol(hasPatrol bool) Combatant {
	if !hasPatrol {
		// ПЛОХО: так делать не стоит.
		// Мы возвращаем типизированный указатель, равный nil.
		// Интерфейс внутри будет хранить тип *Fighter, поэтому он НЕ равен nil
		var f *Fighter
		return f
	}

	return &Fighter{name: "Патрульный", damage: 8, armor: 12}
}

// GetPatrolCorrect - правильный вариант.
func GetPatrolCorrect(hasPatrol bool) Combatant {
	if !hasPatrol {
		// ХОРОШО: возвращаем чистый нетипизированный nil.
		return nil
	}

	return &Fighter{name: "Патрульный", damage: 8, armor: 12}
}

// Демонстрация

func main() {
	fmt.Println("=== AD&D Симуляция: Композиция, диспетчеризация, type switch, nil ===")

	// Создаём отряд и врагов.
	paladin := Paladin{name: "Сэр Роланд", damage: 12, armor: 18, healPower: 15}
	fighter := Fighter{name: "Брен", damage: 10}
	skeleton := Skeleton{name: "Костяной страж", damage: 7, armor: 13, weakToBlunt: true}

	// Слайс интерфейсов: сюда можно положить любой тип, реализующий Combatant.
	party := []Combatant{paladin, fighter, skeleton}

	// 1. Динамическая диспетчеризация.
	fmt.Println("1. Разрешение боя (динамическая диспетчеризация):")
	for _, member := range party {
		ResolveCombat(member)
	}

	fmt.Println()

	// 2. Type switch: определяем тактику.
	fmt.Println("2. Индентификация целей (type switch):")
	for _, member := range party {
		IdentifyCreature(member)
	}

	// 3. Type assertion: пытаемся лечить.
	fmt.Println(" 3. Попытка лечения (type assertion):")
	for _, member := range party {
		TryHealParty(member)
	}

	fmt.Println()

	// 4. Nil-ловушка.
	fmt.Println("4. Демонстрация nil-ловушки:")

	// Неправильный вариант: функция вернула типизированный nil.
	wrongPatrol := GetPatrol(false)
	fmt.Printf("  wrongPatrol == nil: %v\n", wrongPatrol == nil)
	// Паники нет, но значение НЕ nil. Это ловушка.

	// Правильный вариант: функция вернула чистый nil.
	correctPatrol := GetPatrolCorrect(false)
	fmt.Printf("  correctPatrol == nil: %v\n", correctPatrol == nil)

	// Покажим, что внутри "неправильного патруля".
	if wrongPatrol != nil {
		fmt.Printf("  wrongPatrol имеет тип внутри: %T\n", wrongPatrol)
		// Если бы мы попытались вызвать метод, была бы паника (разыменование nil-указателя).
		// Поэтому такие проверки критичны в реальном коде.
	}

	fmt.Println("\n=== Симуляция завершена ===")
}

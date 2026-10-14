package main

import (
	"fmt"
	"math/rand"
)

// Базовые типы

// Weapon представляет оружие с уроном
type Weapon struct {
	Name     string
	BaseDmg  int
	DiceSize int // размер кубика урона (d4, d6, d8)
}

// Armor представялет броню с защитой
type Armor struct {
	Name       string
	Protection int
}

// Potion предсатвляет зельк с эффектом
type Potion struct {
	Name   string
	Heal   int
	Effect string
}

// Обобщённая структура: инвертарь

// Inventory - обобщёный инвертарь, которые может хранить предметы любого типа Т.
// Параметр типа [T any] означает: "Т может бвть любым типом".
// Это как магический мешок - он не знает, что внутри, но умеет с этим работать.
type Inventory[T any] struct {
	items    []T
	capacity int
	name     string
}

// NewInventory создаёт новый инвертарь с заданной вместимостью.
// Обратите внимание: функция - конструктор тоже параметризована [T any].
func NewInventory[T any](name string, capacity int) *Inventory[T] {
	return &Inventory[T]{
		items:    make([]T, 0, capacity),
		capacity: capacity,
		name:     name,
	}
}

// Add добавляет предмет в инвертарь.
// Возвращает true, если предмеи добавлен, false - если инвентарь полон.
func (inv *Inventory[T]) Add(item T) bool {
	if len(inv.items) >= inv.capacity {
		return false
	}
	inv.items = append(inv.items, item)
	return true
}

// Get возврашает предмет по индексу.
// Второе возвращаемое значение показывает, существует ли предмет.
func (inv *Inventory[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(inv.items) {
		var zero T // нулевое значение типа T
		return zero, false
	}
	return inv.items[index], true
}

// Count возвращает количество предметов в инвентаре.
func (inv *Inventory[T]) Count() int {
	return len(inv.items)
}

// IsFull проверяет, заполнен ли инвентарь.
func (inv *Inventory[T]) IsFull() bool {
	return len(inv.items) >= inv.capacity
}

// RollDice бросает кубик заданного размера.
// Параметр типа [T constraints.Integer] ограничивает T целыми числами.
// В Go 1.27.1 мы используем собсветнное ограничение, так как пакет constraints
// может исеть изменения. Для простоты используем any и проверяем на лету.
func RollDice[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}](diceSize T) T {
	// Бросок кубика: случайное число от 1 до diceSize
	return T(1 + rand.Int63n(int64(diceSize)))
}

// RollDamage бросает несколько кубиков урона и суммирует результат.
// Это обобщённая функция для расчёта урона оружия.
func RollDamage[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}](numDice T, diceSize T) T {
	var totalDamage T
	for i := T(0); i < numDice; i++ {
		totalDamage += RollDice(diceSize)
	}

	return totalDamage
}

// Max возвращает большее из двух значений.
// Параметр [T comparable] требует, чтобы тип поддерживал сравнение.
func Max[T interface{ ~int | ~float64 | ~string }](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// Структура для отряда приключенцев

// Adventurer представляет персонажа с инвентарём
type Adventurer struct {
	Name      string
	Class     string
	Hitpoints int
	Inventory *Inventory[any] // инвентарь для любых предметов
}

// NewAdventurer создаёт нового приключенца
func NewAdventurer(name, class string, hp int) *Adventurer {
	return &Adventurer{
		Name:      name,
		Class:     class,
		Hitpoints: hp,
		Inventory: NewInventory[any](name+"'s inventory", 10),
	}
}

// главная функция

func main() {
	fmt.Println("== AD&D симуляция: дженерики в действии ==")
	fmt.Println()

	// Пример 1. Инвентарь для оружия.

	fmt.Println("1. Создание инвентаря для оружия:")
	weaponInv := NewInventory[Weapon]("Weapon Pouch", 5)

	sword := Weapon{Name: "Длинный меч", BaseDmg: 4, DiceSize: 8}
	axe := Weapon{Name: "Боевой топор", BaseDmg: 6, DiceSize: 10}
	dagger := Weapon{Name: "Кинжал", BaseDmg: 2, DiceSize: 4}

	// Добавляем оружие
	weaponInv.Add(sword)
	weaponInv.Add(axe)
	weaponInv.Add(dagger)

	fmt.Printf("   Инвентарь '%s': %d/%d предметов\n", weaponInv.name, weaponInv.Count(), weaponInv.capacity)

	// Получаем первый предмет
	if w, ok := weaponInv.Get(0); ok {
		fmt.Printf("    Первое оружие: %s (урон: %dd%d)\n", w.Name, w.BaseDmg, w.DiceSize)
	}
	fmt.Println()

	// Пример 2. Инвентарь зелий.

	fmt.Println("2. Создание инвентаря для зелий:")
	potionInv := NewInventory[Potion]("Potion Belt", 3)

	healPotion := Potion{Name: "Зелье лечения", Heal: 10, Effect: "Восстанавливает HP"}
	magicPotion := Potion{Name: "Зелье маны", Heal: 0, Effect: "Восстанавливает заклинания"}

	potionInv.Add(healPotion)
	potionInv.Add(magicPotion)

	fmt.Printf("   Инвентарь '%s': %d/%d предметов\n", potionInv.name, potionInv.Count(), potionInv.capacity)

	if p, ok := potionInv.Get(0); ok {
		fmt.Printf("   Первое зелье: %s (лечение: %d HP, эффект: %s)\n", p.Name, p.Heal, p.Effect)
	}

	// Пример 3. Инвентарь брони.

	fmt.Println("3. Создание инвентаря для брони:")
	armorInv := NewInventory[Armor]("Armor Rack", 2)

	chainmail := Armor{Name: "Кольчуга", Protection: 5}
	plate := Armor{Name: "Латы", Protection: 8}

	armorInv.Add(chainmail)
	armorInv.Add(plate)

	fmt.Printf("   Инвентарь '%s': %d/%d предметов\n", armorInv.name, armorInv.Count(), armorInv.capacity)

	fmt.Printf(" Инвентарь полон: %v\n", armorInv.IsFull())
	fmt.Println()

	// Пример 4. Броски кубиков c дженериками

	fmt.Println("4. Броски кубиков урона (дженерик-функции):")

	// Бросок d20 для атаки
	attackRoll := RollDice(20)
	fmt.Printf("   Бросок атаки (d20): %d\n", attackRoll)

	// Бросок урона длинным мечом (1d8+4)
	swordDamage := RollDamage(1, 8) + 4
	fmt.Printf("   Урон длинным мечом (1d8+4): %d\n", swordDamage)

	// Бросок урона боевым топором (1d10+6)
	axeDamage := RollDamage(1, 10) + 6
	fmt.Printf("   Урон боевым топором (1d10+6): %d\n", axeDamage)

	// Выбираем максимальный урон
	maxDamage := Max(swordDamage, axeDamage)
	fmt.Printf("   Максимальный урон: %d\n", maxDamage)
	fmt.Println()

	// Пример 5. Вывод типов (Type Inference).

	fmt.Println("5. Вывод типов (Type Inference):")

	// Компилятор сам выводит тип Т из аргументов
	inv1 := NewInventory[string]("Auto-typed Inventory 1", 5) // = string
	inv1.Add("Стрела")
	inv1.Add("Факел")

	inv2 := NewInventory[int]("Auto-typed Inventory 2", 3) // = int
	inv2.Add(100)
	inv2.Add(50)

	fmt.Printf("   inv1 содержит %d предметов типа string\n", inv1.Count())
	fmt.Printf("   inv2 содержит %d предметов типа int\n", inv2.Count())
	fmt.Println()

	// Пример 6. Создание приключенцев.

	fmt.Println("6. Создание отряда приключенцев:")

	fighter := NewAdventurer("Торин", "Воин", 46)
	wizard := NewAdventurer("Элрина", "Волшебник", 33)
	rogue := NewAdventurer("Шандо", "Вор", 35)

	// Добавляем предметы в инвентарь (используем any)
	fighter.Inventory.Add(sword)
	fighter.Inventory.Add(plate)
	fighter.Inventory.Add(healPotion)

	wizard.Inventory.Add(magicPotion)
	wizard.Inventory.Add(dagger)

	rogue.Inventory.Add(dagger)
	rogue.Inventory.Add(healPotion)

	fmt.Printf("   %s (%s): %d HP, %d предметов\n", fighter.Name, fighter.Class, fighter.Hitpoints, fighter.Inventory.Count())
	fmt.Printf("   %s (%s): %d HP, %d предметов\n", wizard.Name, wizard.Class, wizard.Hitpoints, wizard.Inventory.Count())
	fmt.Printf("   %s (%s): %d HP, %d предметов\n", rogue.Name, rogue.Class, rogue.Hitpoints, rogue.Inventory.Count())
	fmt.Println()

	// Пример 7. Безопасность типов

	fmt.Println("7. Демонстрация безопасности типов:")

	// Этот код не скомпилируется
	// weaponInv.Add(healPotion) // ОШИБКА: нельзя добавить Potion в inventory[Weapon]

	// Но мы можем проверить типы через type assertion
	if item, ok := fighter.Inventory.Get(0); ok {
		if w, isWeapon := item.(Weapon); isWeapon {
			fmt.Printf("   Торин держит: %s (это оружие)\n", w.Name)
		} else if p, isPotion := item.(Potion); isPotion {
			fmt.Printf(" торин держит: %s (это зельк)\n", p.Name)
		} else if a, isArmor := item.(Armor); isArmor {
			fmt.Printf("   Торин держит: %s (это броня)\n", a.Name)
		}
	}

	fmt.Println()
	fmt.Println("=== СИМУЛЯЦИЯ ЗАВЕРШЕНА ===")
}

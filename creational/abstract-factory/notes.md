# 🏭 Abstract Factory

The **Abstract Factory** is a *creational design pattern* that provides an interface for creating **families of related objects** without specifying their concrete classes.  
It helps ensure that related products (like buttons and checkboxes) are created together, maintaining consistency across a family.

---

## 💡 Intent

Provide an interface to create **groups of related or dependent objects**, while keeping their concrete implementations hidden from the client.  
This pattern lets you switch between **entire product families** (e.g., Windows UI or Mac UI) without changing your core application logic.

---

## 🧠 Problem

Suppose you are building a **cross-platform UI** library that supports **Windows** and **Mac**.

Each platform has its own style for UI elements:
- Buttons  
- Checkboxes  

If you scatter conditional logic like `if os == "windows"` across your codebase, you’ll end up with:
- Tight coupling between the UI and platform-specific code  
- Repeated logic for every platform  
- Difficulty adding new UI themes or OS variants  

You need a way to **create related UI components together** while keeping the code clean and platform-agnostic.

---

## ⚙️ Solution

The Abstract Factory pattern suggests creating a **common interface** (e.g., `GUIFactory`) with methods for all related products (like `CreateButton()` and `CreateCheckbox()`).

Each **concrete factory** (e.g., `WinFactory` or `MacFactory`) implements this interface to produce its own variant of UI elements (`WinButton`, `WinCheckbox`, `MacButton`, `MacCheckbox`).

This ensures:
- You can switch product families by changing only the factory instance.  
- All created objects are consistent within the chosen family.  

In Go, we use **interfaces and structs** to model factories and products — no inheritance required.

---

## 🧩 Example (Go Implementation)

```go
package main

import "fmt"

// Product Interfaces
type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

// Concrete Products (Windows)
type WinButton struct{}
func (WinButton) Paint() { fmt.Println("Render a Windows button.") }

type WinCheckbox struct{}
func (WinCheckbox) Paint() { fmt.Println("Render a Windows checkbox.") }

// Concrete Products (Mac)
type MacButton struct{}
func (MacButton) Paint() { fmt.Println("Render a Mac button.") }

type MacCheckbox struct{}
func (MacCheckbox) Paint() { fmt.Println("Render a Mac checkbox.") }

// Abstract Factory
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Concrete Factories
type WinFactory struct{}
func (WinFactory) CreateButton() Button { return WinButton{} }
func (WinFactory) CreateCheckbox() Checkbox { return WinCheckbox{} }

type MacFactory struct{}
func (MacFactory) CreateButton() Button { return MacButton{} }
func (MacFactory) CreateCheckbox() Checkbox { return MacCheckbox{} }

// Client Code
func BuildUI(factory GUIFactory) {
	btn := factory.CreateButton()
	chk := factory.CreateCheckbox()

	btn.Paint()
	chk.Paint()
}

func main() {
	var factory GUIFactory

	fmt.Println("Building Windows UI:")
	factory = WinFactory{}
	BuildUI(factory)

	fmt.Println("\nBuilding Mac UI:")
	factory = MacFactory{}
	BuildUI(factory)
}

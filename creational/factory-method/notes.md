# 🏭 Factory Method

The **Factory Method** is a *creational design pattern* that defines a method for creating objects through a shared interface.  
In Go, it allows different structs (acting as creators) to provide their own logic for instantiating concrete types that implement a common interface.

---

## 💡 Intent

Provide an interface for creating objects, but let the implementing struct decide which concrete type to instantiate.  
This helps you write code that depends only on abstractions, not concrete implementations.

---

## 🧠 Problem

Imagine you are building a **dialogue box** for an application that should run across multiple operating systems — say **Windows** and **Mac**.

Each operating system has its own way of rendering UI elements like buttons.  
If you hardcode button creation logic for each OS throughout your codebase, you’ll end up with:

- Multiple `if` or `switch` statements checking for OS type.  
- Tight coupling between your app logic and platform-specific implementations.  
- Code that’s harder to maintain or extend when adding a new OS.

---

## ⚙️ Solution

The Factory Method pattern suggests defining a **common interface** for the product (e.g., `Button`).  
Then, each OS-specific struct (like `WindowsDialog` and `MacDialog`) implements its **own factory method** that decides which concrete type to instantiate (`WindowsButton` or `MacButton`).

This allows:

- Your main app logic to depend only on the `Button` interface.  
- Each OS to handle its own creation logic without breaking abstraction.  

In Go, since we don’t have inheritance, this pattern is implemented using **interfaces** and **methods returning interfaces** — a more composition-oriented approach.

---

## 🧩 Example (Go Implementation)

```go
package main

import "fmt"

// Product interface
type Button interface {
	Render()
	OnClick()
}

// Concrete Products
type WindowsButton struct{}

func (b *WindowsButton) Render()  { fmt.Println("Rendering a Windows button.") }
func (b *WindowsButton) OnClick() { fmt.Println("Windows button clicked!") }

type MacButton struct{}

func (b *MacButton) Render()  { fmt.Println("Rendering a Mac button.") }
func (b *MacButton) OnClick() { fmt.Println("Mac button clicked!") }

// Creator interface
type Dialog interface {
	CreateButton() Button
	RenderDialog()
}

// Concrete Creators
type WindowsDialog struct{}

func (d *WindowsDialog) CreateButton() Button {
	return &WindowsButton{}
}

func (d *WindowsDialog) RenderDialog() {
	btn := d.CreateButton()
	btn.Render()
	btn.OnClick()
}

type MacDialog struct{}

func (d *MacDialog) CreateButton() Button {
	return &MacButton{}
}

func (d *MacDialog) RenderDialog() {
	btn := d.CreateButton()
	btn.Render()
	btn.OnClick()
}

func main() {
	var dialog Dialog

	dialog = &WindowsDialog{}
	dialog.RenderDialog()

	dialog = &MacDialog{}
	dialog.RenderDialog()
}
```

## 🔗 References

- [Refactoring.Guru – Factory Method](https://refactoring.guru/design-patterns/factory-method)

package main

import (
	"fmt"
)

// Define a struct for Account
type Account struct {
	Name    string
	Balance float64
}

// Deposit function
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	fmt.Printf("✅ Deposited %.2f. New Balance: %.2f\n", amount, a.Balance)
}

// Withdraw function
func (a *Account) Withdraw(amount float64) {
	if amount > a.Balance {
		fmt.Println("❌ Insufficient balance")
		return
	}
	a.Balance -= amount
	fmt.Printf("✅ Withdrawn %.2f. New Balance: %.2f\n", amount, a.Balance)
}

// Show balance
func (a *Account) ShowBalance() {
	fmt.Printf("📊 Current Balance for %s: %.2f\n", a.Name, a.Balance)
}

func main() {
	var account Account
	var choice int
	var amount float64

	fmt.Print("Enter your name to create an account: ")
	fmt.Scan(&account.Name)

	for {
		fmt.Println("\n--- Go Bank Menu ---")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.ShowBalance()
		case 4:
			fmt.Println("👋 Exiting. Thank you for using Go Bank.")
			return
		default:
			fmt.Println("❗ Invalid choice. Please try again.")
		}
	}
}

package main

import "fmt"

//Options:
const (
	DepositOption  = 1
	WithdrawOption = 2
	ViewBalanceOption = 3
	ViewHistoryOption = 4
	ExitOption     = 5
)

//Required data to be stored in accounts
type Account struct {
	ID          int
	Name        string
	Balance     float64
	Transactions []string
}

//To store all the data
var accounts []Account

//Creating new account
func createAccount(id int, name string, initialBalance float64) {
	account := Account{
		ID:          id,
		Name:        name,
		Balance:     initialBalance,
		Transactions: []string{fmt.Sprintf("Account created with balance: %.2f", initialBalance)},
	}
	accounts = append(accounts, account)
}

//To deposit money
func deposit(id int, amount float64) string {
	for i, account := range accounts {
		if account.ID == id {
			if amount <= 0 {
				return "Deposit amount must be greater than zero."
			}
			accounts[i].Balance += amount
			accounts[i].Transactions = append(accounts[i].Transactions, fmt.Sprintf("Deposited: %f", amount))
			return fmt.Sprintf("Deposited %f to account %d. New balance: %f", amount, id, accounts[i].Balance)
		}
	}
	return "Account not found."
}

//To withdraw money
func withdraw(id int, amount float64) string {
	for i, account := range accounts {
		if account.ID == id {
			if amount <= 0 {
				return "Withdrawal amount must be greater than zero."
					}
			if amount > account.Balance {
				return "Insufficient balance."
			}
			accounts[i].Balance -= amount
			accounts[i].Transactions = append(accounts[i].Transactions, fmt.Sprintf("Withdrew: %f", amount))
			return fmt.Sprintf("Withdrew %f from account %d. New balance: %f", amount, id, accounts[i].Balance)
		}
	}
	return "Account not found."
}

//Viewing account balance
func viewBalance(id int) string {
	for _, account := range accounts {
		if account.ID == id {
			return fmt.Sprintf("Account balance: %f", account.Balance)
		}
	}
	return "Account not found."
}

// To view transaction history
func viewHistory(id int) string {
	for _, account := range accounts {
		if account.ID == id {
			if len(account.Transactions) == 0 {
				return "No transactions found."
			}
			history := "Transaction History:\n"
			for _, transaction := range account.Transactions {
				history += transaction + "\n"
			}
			return history
		}
	}
	return "Account not found."
}

func displayMenu() {
	fmt.Println("Menu:")
	fmt.Println("1. Deposit Money")
	fmt.Println("2. Withdraw Money")
	fmt.Println("3. View Balance")
	fmt.Println("4. View Transaction History")
	fmt.Println("5. Exit")
}

func main() {
	createAccount(1, "Alice", 1000.00)
	createAccount(2, "Bob", 500.00)

	var choice int
	var accountID int
	var amount float64
	var exit bool

	for !exit {
		displayMenu()
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case DepositOption:
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			fmt.Println(deposit(accountID, amount))

		case WithdrawOption:
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			fmt.Println(withdraw(accountID, amount))

		case ViewBalanceOption:
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			fmt.Println(viewBalance(accountID))

		case ViewHistoryOption:
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accountID)
			fmt.Println(viewHistory(accountID))

		case ExitOption:
			fmt.Println("Exiting the program.")
			exit = true 

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

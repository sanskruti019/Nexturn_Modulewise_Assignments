package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

// Menu constants
const (
    DEPOSIT          = 1
    WITHDRAW         = 2
    CHECK_BALANCE    = 3
    VIEW_HISTORY     = 4
    EXIT            = 5
)

// Transaction types
const (
    DEPOSIT_TYPE    = "DEPOSIT"
    WITHDRAW_TYPE   = "WITHDRAW"
)

// Account represents a bank account
type Account struct {
    ID              int
    Name            string
    Balance         float64
    Transactions    []string
}

// BankSystem manages all bank operations
type BankSystem struct {
    accounts        []*Account
    scanner         *bufio.Scanner
}

// NewBankSystem creates a new instance of BankSystem
func NewBankSystem() *BankSystem {
    return &BankSystem{
        accounts: make([]*Account, 0),
        scanner:  bufio.NewScanner(os.Stdin),
    }
}

// CreateAccount creates a new bank account
func (bs *BankSystem) CreateAccount(id int, name string) (*Account, error) {
    // Check for duplicate ID
    for _, acc := range bs.accounts {
        if acc.ID == id {
            return nil, fmt.Errorf("account with ID %d already exists", id)
        }
    }

    account := &Account{
        ID:           id,
        Name:         name,
        Balance:      0,
        Transactions: make([]string, 0),
    }
    
    bs.accounts = append(bs.accounts, account)
    return account, nil
}

// FindAccount finds an account by ID
func (bs *BankSystem) FindAccount(id int) (*Account, error) {
    for _, acc := range bs.accounts {
        if acc.ID == id {
            return acc, nil
        }
    }
    return nil, fmt.Errorf("account with ID %d not found", id)
}

// Deposit adds money to an account
func (bs *BankSystem) Deposit(id int, amount float64) error {
    if amount <= 0 {
        return errors.New("deposit amount must be greater than zero")
    }

    account, err := bs.FindAccount(id)
    if err != nil {
        return err
    }

    account.Balance += amount
    transaction := fmt.Sprintf("%s: + Rs. %.2f (Balance: Rs.%.2f) - %s",
        DEPOSIT_TYPE, amount, account.Balance, time.Now().Format("2006-01-02 15:04:05"))
    account.Transactions = append(account.Transactions, transaction)
    
    return nil
}

// Withdraw removes money from an account
func (bs *BankSystem) Withdraw(id int, amount float64) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be greater than zero")
    }

    account, err := bs.FindAccount(id)
    if err != nil {
        return err
    }

    if account.Balance < amount {
        return fmt.Errorf("insufficient balance. Current balance: Rs.%.2f", account.Balance)
    }

    account.Balance -= amount
    transaction := fmt.Sprintf("%s: -Rs.%.2f (Balance: Rs.%.2f) - %s",
        WITHDRAW_TYPE, amount, account.Balance, time.Now().Format("2006-01-02 15:04:05"))
    account.Transactions = append(account.Transactions, transaction)
    
    return nil
}

// DisplayTransactionHistory shows all transactions for an account
func (bs *BankSystem) DisplayTransactionHistory(id int) error {
    account, err := bs.FindAccount(id)
    if err != nil {
        return err
    }

    if len(account.Transactions) == 0 {
        fmt.Println("No transactions found.")
        return nil
    }

    fmt.Printf("\nTransaction History for Account %d (%s):\n", account.ID, account.Name)
    fmt.Println("----------------------------------------")
    for _, transaction := range account.Transactions {
        fmt.Println(transaction)
    }
    return nil
}

// readInput reads a line from standard input
func (bs *BankSystem) readInput() string {
    bs.scanner.Scan()
    return strings.TrimSpace(bs.scanner.Text())
}

// RunMenu starts the interactive menu system
func (bs *BankSystem) RunMenu() {
    fmt.Println("Welcome to the Bank Transaction System!")
    
    // Creating a sample account for testing
    account, err := bs.CreateAccount(1, "Amit kumar")
    if err != nil {
        fmt.Printf("Error creating account: %v\n", err)
        return
    }
    fmt.Printf("Created account for %s (ID: %d)\n\n", account.Name, account.ID)

    for {
        fmt.Println("\nPlease select an option:")
        fmt.Printf("%d. Deposit\n", DEPOSIT)
        fmt.Printf("%d. Withdraw\n", WITHDRAW)
        fmt.Printf("%d. Check Balance\n", CHECK_BALANCE)
        fmt.Printf("%d. View Transaction History\n", VIEW_HISTORY)
        fmt.Printf("%d. Exit\n", EXIT)
        
        choice, err := strconv.Atoi(bs.readInput())
        if err != nil {
            fmt.Println("Invalid input. Please enter a number.")
            continue
        }

        switch choice {
        case DEPOSIT:
            fmt.Print("Enter amount to deposit: Rs.")
            amount, err := strconv.ParseFloat(bs.readInput(), 64)
            if err != nil {
                fmt.Println("Invalid amount.")
                continue
            }
            
            if err := bs.Deposit(1, amount); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Successfully deposited Rs.%.2f\n", amount)
            }

        case WITHDRAW:
            fmt.Print("Enter amount to withdraw: Rs.")
            amount, err := strconv.ParseFloat(bs.readInput(), 64)
            if err != nil {
                fmt.Println("Invalid amount.")
                continue
            }
            
            if err := bs.Withdraw(1, amount); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Successfully withdrew Rs. %.2f\n", amount)
            }

        case CHECK_BALANCE:
            account, err := bs.FindAccount(1)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Current balance: Rs. %.2f\n", account.Balance)
            }

        case VIEW_HISTORY:
            if err := bs.DisplayTransactionHistory(1); err != nil {
                fmt.Printf("Error: %v\n", err)
            }

        case EXIT:
            fmt.Println("Thank you for using the Bank Transaction System!")
            return

        default:
            fmt.Println("Invalid option. Please try again.")
        }
    }
}

func main() {
    bankSystem := NewBankSystem()
    bankSystem.RunMenu()
}
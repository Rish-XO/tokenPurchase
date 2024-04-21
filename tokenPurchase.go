package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	Timestamp   time.Time
	WalletAddress string
	IsNewWallet bool
}

func analyzeTokenPurchases(transactions []Transaction, projectName string, days int) map[string]int {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -days)

	wallets := make(map[string]struct {
		isNew     bool
		purchases int
	})

	var (
		newBuyerPurchases     int
		returningBuyerPurchases int
		newWalletFirstPurchases int
		newWalletRepeatPurchases int
	)

	for _, transaction := range transactions {
		if transaction.Timestamp.After(startDate) && transaction.Timestamp.Before(endDate) {
			wallet, ok := wallets[transaction.WalletAddress]
			if !ok {
				wallet = struct {
					isNew     bool
					purchases int
				}{isNew: true}
			}

			if wallet.isNew {
				newBuyerPurchases++
				if transaction.IsNewWallet {
					newWalletFirstPurchases++
				}
				wallet.isNew = false
			} else {
				returningBuyerPurchases++
				if transaction.IsNewWallet {
					newWalletRepeatPurchases++
				}
			}

			wallet.purchases++
			wallets[transaction.WalletAddress] = wallet
		}
	}

	totalPurchases := 0
	for _, wallet := range wallets {
		totalPurchases += wallet.purchases
	}

	return map[string]int{
		"projectName":             len(projectName),
		"totalPurchases":          totalPurchases,
		"newBuyerPurchases":       newBuyerPurchases,
		"returningBuyerPurchases": returningBuyerPurchases,
		"newWalletFirstPurchases": newWalletFirstPurchases,
		"newWalletRepeatPurchases": newWalletRepeatPurchases,
	}
}

func main() {
	// Example usage
	transactions := []Transaction{
		{Timestamp: time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC), WalletAddress: "0x123", IsNewWallet: true},
		{Timestamp: time.Date(2024, 4, 16, 0, 0, 0, 0, time.UTC), WalletAddress: "0x456", IsNewWallet: true},
		{Timestamp: time.Date(2024, 4, 17, 0, 0, 0, 0, time.UTC), WalletAddress: "0x123", IsNewWallet: false},
		// ... more transactions
	}

	projectName := "MyProject"
	analysis := analyzeTokenPurchases(transactions, projectName, 7)
    fmt.Println(analysis)
}
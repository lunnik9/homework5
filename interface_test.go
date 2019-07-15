package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPayByWallet(t *testing.T) {
	amount := 50
	payment := 30
	w := &Wallet{amount}
	w.Pay(payment)

	assert.Equal(t, w.funds, amount-payment)
}

func TestPayByWalletGivesError(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}
	err := w.Pay(payment)
	assert.Error(t, err)
}

func TestBuyByWallet(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}
	_ = Buy(w, payment)
	assert.Equal(t, w.funds, amount-payment)
}

func TestBuyByCreditCard(t *testing.T) {
	amount := 50
	payment := 30
	c := &CreditCard{funds: amount}
	_ = Buy(c, payment)
	assert.Equal(t, c.funds, amount-payment)
}
func TestBuyByBitcoin(t *testing.T) {
	amount := 50
	payment1 := 228
	payment := 30
	b := &Bitcoin{funds: amount, owner: "arnur"}
	_ = Buy(b, payment)
	fmt.Println(b)

	_ = Buy(b, payment1)
	fmt.Println(len(b.transactions))
	assert.Equal(t, b.funds, amount-payment)
}

func TestCheckAndBuyWallet(t *testing.T) {
	amount := 50
	w := &Wallet{amount}
	checkAndBuy(w, amount)
	assert.Equal(t, w.funds, amount)
}

func TestCheckPaymentType(t *testing.T) {
	amount := 50
	c := &CreditCard{amount, "arur", time.Time{}, nil}
	fmt.Println(CheckPaymentType(c))
}

func TestHashTranaction(t *testing.T) {
	test := "hello"
	fmt.Println(HashTransaction(test))
}

func TestBitcoin(t *testing.T) {
	b := &Bitcoin{funds: 20, owner: "arnur"}
	fmt.Println(b.transactions[0].previousTransaction)
}

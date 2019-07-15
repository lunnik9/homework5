package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	newTransaction Transaction
)

type Wallet struct {
	funds int
}
type CreditCard struct {
	funds      int
	owner      string
	expireTime time.Time
	bonuses    []Bonus
}
type Bitcoin struct {
	funds        int
	owner        string
	transactions []Transaction
}
type Transaction struct {
	sumOfTransaction    int
	transactionTime     time.Time
	previousTransaction string
}

func HashTransaction(previousTransactionToString string) string {
	h := sha1.New()
	h.Write([]byte(previousTransactionToString))
	previousTransaction := h.Sum(nil)
	return hex.EncodeToString(previousTransaction)
}

type Bonus struct {
	bonusName   string
	description string
}

type Payer interface {
	Pay(int) error
}

type Funder interface {
	GetFunds() int
}

type PayFunder interface {
	Payer
	Funder
}

func (w *Wallet) Pay(amount int) error {
	if w.funds < amount {
		return errors.New("not enough founds")
	}
	w.funds -= amount
	return nil
}
func (c *CreditCard) Pay(amount int) error {
	if c.funds < amount {
		return errors.New("not enough founds")
	}
	c.funds -= amount
	return nil
}
func (b *Bitcoin) Pay(amount int) error {
	beforeTransaction := b.funds
	if b.funds < amount {
		return errors.New("not enough founds")
	}
	b.funds -= amount
	if len(b.transactions) != 0 {
		newTransaction = Transaction{
			b.funds - beforeTransaction,
			time.Now(),
			convertToNextTransaction(b.transactions[len(b.transactions)-1].sumOfTransaction, b.transactions[len(b.transactions)-1].transactionTime, b.transactions[len(b.transactions)-1].previousTransaction),
		}

		b.transactions = append(b.transactions, newTransaction)
	} else {
		newTransaction = Transaction{
			b.funds - beforeTransaction,
			time.Now(),
			"",
		}
		b.transactions = append(b.transactions, newTransaction)
	}

	return nil
}

func Buy(p Payer, amount int) error {
	err := p.Pay(amount)
	if err != nil {
		return err
	}
	return nil
}
func checkAndBuy(p PayFunder, amount int) error {
	if p.GetFunds() <= 0 {
		fmt.Println("пополните счет")
	}
	err := p.Pay(amount)
	if err != nil {
		return err
	}
	return nil
}

func (w *Wallet) GetFunds() int {
	return w.funds
}

func CheckPaymentType(p Payer) interface{} {
	switch p.(type) {
	case *Wallet:
		fmt.Println("ты пользуешься кошельком")
		return p.(*Wallet).funds
	case *CreditCard:
		fmt.Println("ты пользуешься кредиткой ")
		fmt.Println(p.(*CreditCard).owner)
		return p.(*CreditCard).funds
	case *Bitcoin:
		fmt.Println("ты КУЛХАЦКЕР И ЮЗАЕШЬ БИТКОИН ")
		fmt.Println(p.(*Bitcoin).owner)
		return p.(*Bitcoin).funds

	default:
		return nil
	}
}

func convertToNextTransaction(sumOfTransaction int, transactionTime time.Time, previousTransaction string) string {
	return HashTransaction(fmt.Sprintf(strconv.Itoa(sumOfTransaction), transactionTime, previousTransaction))
}

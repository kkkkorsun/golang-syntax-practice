package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	secondTask()
}

func firstTask() {
	a := 10000
	b := 2

	result, err := SumNumbers(a, b)
	if err != nil {
		fmt.Println("Переданные числа a: ", a, "b: ", b)
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func MultiplyNumbers(a int, b int) (int, error) {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return 0, errors.New("[MultiplyNumbers] переданные числа вышли за диапазон от -1000 до 1000")
	}

	return a * b, nil
}

func DivideNumbers(a int, b int) (float64, error) {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return 0, errors.New("[DivideNumbers] переданные числа вышли за диапазон от -1000 до 1000")
	}
	if b == 0 {
		return 0, errors.New("[DivideNumbers] деление не может быть на 0")
	}
	return float64(a) / float64(b), nil
}

func SubtractNumbers(a int, b int) (int, error) {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return 0, errors.New("[SubtractNumbers] переданные числа вышли за диапазон от -1000 до 1000")
	}

	return a - b, nil
}

func SumNumbers(a int, b int) (int, error) {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return 0, errors.New("[SumNumbers] переданные числа вышли за диапазон от -1000 до 1000")
	}

	return a + b, nil
}

func secondTask() {

	bankAcc := CreateBankAccount(10000)

	for {
		showMenu()

		var userOperationChoose int
		_, _ = fmt.Scanln(&userOperationChoose)

		isWorking, bankErr := checkBankAvailability()

		if !isWorking {
			fmt.Println("\nБанк не работает, по причине: ", bankErr)
			continue
		}

		switch userOperationChoose {
		case 1:
			var amount int
			fmt.Println("\nВведите количество денег для снятия")
			_, _ = fmt.Scanln(&amount)
			ok, err := bankAcc.CashOut(amount)

			if !ok {
				fmt.Println("\nОшибка при снятии наличных по причине: ", err)
				continue
			}

			fmt.Println("\nУспешно сняли с банковского счета: ", amount)

		case 2:
			var amount int
			fmt.Println("\nВведите количество денег для оплаты")
			_, _ = fmt.Scanln(&amount)
			ok, err := bankAcc.Pay(amount)

			if !ok {
				fmt.Println("\nОшибка при оплате по причине: ", err)
				continue
			}

			fmt.Println("\nУспешно оплатили с банковского счета: ", amount)
		case 3:
			fmt.Println("\nВаш текущий баланс: ", bankAcc.ShowBalance())
		}
	}
}

type BankAccount struct {
	Balance int
}

func CreateBankAccount(balance int) *BankAccount {
	return &BankAccount{Balance: balance}
}

func (b *BankAccount) CashOut(amount int) (bool, error) {
	if b.Balance < amount {
		return false, errors.New("[CashOut] недостаточно средств на балансе для снятия")
	}

	fmt.Println("Производим снятие наличных с банковского счета")

	b.Balance = b.Balance - amount

	return true, nil
}

func (b *BankAccount) ShowBalance() int {
	return b.Balance
}

func (b *BankAccount) Pay(amount int) (bool, error) {
	if b.Balance < amount {
		return false, errors.New("[Pay] недостаточно средств на балансе для оплаты")
	}

	fmt.Println("Производим оплату с банковского счета")

	b.Balance = b.Balance - amount

	return true, nil
}

func checkBankAvailability() (bool, error) {
	ok := rand.Intn(100) < 30
	if ok {
		return false, errors.New("в банке ведутся технические работы")
	}

	return true, nil
}

func showMenu() {
	fmt.Println("\n\nДобро пожаловать в банк, выберите операцию: ")
	fmt.Println("\n1. Снять наличные")
	fmt.Println("\n2. Провести онлайн оплату")
	fmt.Println("\n3. Показать баланс")
}

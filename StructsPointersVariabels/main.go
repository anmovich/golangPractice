package main

import ("SPV/BankAccount"
				"fmt")

func main(){
	NewAccount := BankAccount.CreateAccount("Anme", 0)

	NewAccount.Deposit(1000)
	fmt.Println("Баланс аккаунта", NewAccount.GetOwner(), "=", NewAccount.GetBalance())
	err := NewAccount.Widthdraw(500)
	if err!= nil{
		fmt.Println("Чета деняк у вас нет")
	} else{
		fmt.Println("Баланс аккаунта", NewAccount.GetOwner(), "=", NewAccount.GetBalance())
	}
	err = NewAccount.Widthdraw(600)
	if err!= nil{
		fmt.Println("Чета деняк у вас нет")
	} else{
		fmt.Println("Баланс аккаунта", NewAccount.GetOwner(), "=", NewAccount.GetBalance())
	}

	BankAccount.ChangeOwner(&NewAccount, "Hui s gory")
	fmt.Println("Новое имя аккаунта:", NewAccount.GetOwner())
}

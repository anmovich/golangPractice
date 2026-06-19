package BankAccount

import "errors"

type Account struct{
	id int
	owner string
	balance float64
}

func CreateAccount( owner string, balance float64) Account{
	return Account{
		id: idGenerator(owner),
		owner: owner,
		balance: balance,
	}
}

func (a Account) GetId() int{
	return a.id
}

func (a Account) GetOwner() string{
	return a.owner
}

func (a Account)GetBalance() float64{
	return a.balance
}

func (a *Account) Deposit(amount float64){
	a.balance+=amount
}

func (a *Account) Widthdraw(amount float64) error{
	if a.balance - amount < 0{
		return errors.New("Недостаточно средств")
	} else{
		a.balance -= amount
		return nil
	}
}

func ChangeOwner(a *Account, newName string){
	a.owner=newName
}

func idGenerator(name string) int{
	return len(name) * 4
}

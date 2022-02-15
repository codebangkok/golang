package commands

type OpenAccountCommand struct {
	AccountHolder  string
	AccountType    int
	OpeningBalance float64
}

type DepositFundCommand struct {
	ID     string
	Amount float64
}

type WithdrawFundCommand struct {
	ID     string
	Amount float64
}

type CloseAccountCommand struct {
	ID string
}

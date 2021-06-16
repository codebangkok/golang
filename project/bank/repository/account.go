package repository

type Account struct {
	AccountID   int     `db:"account_id"`
	CustomerID  int     `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      int     `db:"status"`
}

//go:generate mockgen -destination=../mock/mock_repository/mock_account_repository.go bank/repository AccountRepository
type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAll(int) ([]Account, error)
}

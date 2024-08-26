package usecase

type ITransactionRepository interface {
	Begin() error
	Commit() error
	Rollback() error
}

type TransactionUsecase struct {
	TransactionRepository ITransactionRepository
}

func (u *TransactionUsecase) Begin() error {
	return u.TransactionRepository.Begin()
}

func (u *TransactionUsecase) Commit() error {
	return u.TransactionRepository.Commit()
}

func (u *TransactionUsecase) Rollback() error {
	return u.TransactionRepository.Rollback()
}

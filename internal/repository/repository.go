package repository

type Repo struct {
	Base
}

type Repository interface {
	Init(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error
}

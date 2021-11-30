package service

type Service interface {
	Run() (int, error)
}

package service

type Authorization interface {
	
}

type Post interface {

}

type Service struct {
	Authorization
	Post
}

func New() *Service {
	return &Service{
		
	}
}
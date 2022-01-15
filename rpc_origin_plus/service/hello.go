package service

type HelloService struct {}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = "hello," + req
	return nil
}
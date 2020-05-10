package HelloService

type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "Hello:" + request.GetValue()
	return nil
}

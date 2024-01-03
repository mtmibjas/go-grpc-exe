package alerts

import context "context"

func New() *Service {
	return &Service{}
}

type Service struct {
	Message string
	UnimplementedAlertServiceServer
}

func (s *Service) Send(ctx context.Context, req *AlertRequest) (*AlertResponse, error) {
	s.Message = req.Body
	return &AlertResponse{Response: s.Message}, nil
}

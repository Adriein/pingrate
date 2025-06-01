package session

type CreateSessionService struct{}

func NewCreateSessionService() *CreateSessionService {
	return &CreateSessionService{}
}

func (c *CreateSessionService) Execute() (string, error) {
	return "Create Session", nil
}

package access_token

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{}
}

func GetById(string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}

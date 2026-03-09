package base

import "KimJin/src/internal/repository/base"

type PublicRepoService struct {
	repo *base.PublicRepo
}

func (s *PublicRepoService) Login(username string, password string) (string, error) {
	return s.repo.Login(username, password)
}

// NewPublicRepoService 创建公共仓库服务实例
func NewPublicRepoService() *PublicRepoService {
	return &PublicRepoService{
		repo: &base.PublicRepo{},
	}
}

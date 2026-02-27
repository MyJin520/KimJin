package service

import (
	"KimJin/src/internal/model"
	"KimJin/src/internal/repository"
)

// FormService 表单业务服务
type FormService struct {
	repo *repository.FormRepo
}

// NewFormService 创建表单服务实例
func NewFormService() *FormService {
	return &FormService{
		repo: &repository.FormRepo{},
	}
}

// GetFormConfig 获取表单配置
func (s *FormService) GetFormConfig(formID string) (*model.FormConfig, error) {
	return s.repo.GetFormConfigByFormID(formID)
}

// SubmitForm 提交表单数据
func (s *FormService) SubmitForm(formID, data string) (uint, error) {
	submission := &model.FormSubmission{
		FormID: formID,
		Data:   data,
	}
	if err := s.repo.CreateFormSubmission(submission); err != nil {
		return 0, err
	}
	return submission.ID, nil
}

// 初始化默认配置 --> todo 后续需要删除
func (s *FormService) InitDefaultConfig() error {
	return s.repo.InitDefaultFormConfig()
}

// GetFormSubmissions 查询指定表单的所有提交记录
func (s *FormService) GetFormSubmissions(formID string) ([]model.FormSubmission, error) {
	return s.repo.GetFormSubmissionsByFormID(formID)
}

// GetFormSubmissionByID 根据ID查询单条提交记录
func (s *FormService) GetFormSubmissionByID(id uint) (*model.FormSubmission, error) {
	return s.repo.GetFormSubmissionByID(id)
}

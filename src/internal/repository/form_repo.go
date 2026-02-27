package repository

import (
	"KimJin/src/internal/model"
	"KimJin/src/pkg/db"
)

type FormRepo struct{}

// GetFormConfigByFormID 根据FormID获取表单配置
func (r *FormRepo) GetFormConfigByFormID(formID string) (*model.FormConfig, error) {
	var config model.FormConfig
	err := db.DB.Where("form_id = ?", formID).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// CreateFormSubmission 保存表单提交结果
func (r *FormRepo) CreateFormSubmission(submission *model.FormSubmission) error {
	return db.DB.Create(submission).Error
}

// InitDefaultFormConfig 初始化默认表单配置
func (r *FormRepo) InitDefaultFormConfig() error {
	var count int64
	db.DB.Model(&model.FormConfig{}).Where("form_id = ?", "payment_form").Count(&count)
	if count > 0 {
		return nil // 已存在则跳过
	}

	defaultConfig := model.FormConfig{
		FormID:   "payment_form",
		FormName: "订单支付表单",
		Fields: `[
			{"type":"select","name":"payType","label":"支付方式","options":["creditCard","alipay"],"optionLabels":["信用卡","支付宝"],"required":true},
			{"type":"text","name":"cardNo","label":"信用卡号","showIf":{"payType":"creditCard"},"required":true},
			{"type":"text","name":"validDate","label":"有效期","showIf":{"payType":"creditCard"},"required":true},
			{"type":"text","name":"alipayAccount","label":"支付宝账号","showIf":{"payType":"alipay"},"required":true}
		]`,
	}
	return db.DB.Create(&defaultConfig).Error
}

// GetFormSubmissionsByFormID 根据FormID查询所有提交记录
func (r *FormRepo) GetFormSubmissionsByFormID(formID string) ([]model.FormSubmission, error) {
	var submissions []model.FormSubmission
	err := db.DB.Where("form_id = ?", formID).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

// GetFormSubmissionByID 根据ID查询单条提交记录
func (r *FormRepo) GetFormSubmissionByID(id uint) (*model.FormSubmission, error) {
	var submission model.FormSubmission
	err := db.DB.Where("id = ?", id).First(&submission).Error
	if err != nil {
		return nil, err
	}
	return &submission, nil
}

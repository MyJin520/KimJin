package api

import "KimJin/src/internal/api/base"

var (
	FormAPI   = NewFormController()
	PublicAPI = base.NewPublicController()
)

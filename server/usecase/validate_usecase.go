package usecase

import (
	"jobfair2024/model"
	"jobfair2024/pkg"
)

func validateAdminRole(userInfo *UserInfo) error {
	if userInfo.Role != model.Admin {
		return pkg.NotHaveRight
	}
	return nil
}

func validateCompanyRole(userInfo *UserInfo) error {
	if userInfo.Role != model.Company {
		return pkg.NotHaveRight
	}
	return nil
}

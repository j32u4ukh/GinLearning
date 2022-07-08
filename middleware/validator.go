package middleware

import (
	"GinLearning/structs"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsValidPassword(filed validator.FieldLevel) bool {
	if match, _ := regexp.MatchString(`^[A-Z]\w{4,10}$`, filed.Field().String()); match {
		return true
	}

	return false
}

func UserList(filed validator.StructLevel) {
	users := filed.Current().Interface().(structs.Users)

	if users.UserListSize == len(users.UserList) {

	} else {
		filed.ReportError(users.UserListSize, "Size of user list", "UserListSize", "UserListSizeMustEqualsUserList", "")
	}
}

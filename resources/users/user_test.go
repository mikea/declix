package users_test

import (
	"mikea/declix/interfaces"
	"mikea/declix/resources/users"
)

var _ interfaces.Resource = &users.UserImpl{}

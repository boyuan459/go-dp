package user

import (
	"dp/component"
	"dp/httpd/handler/user"
)

func Routes() {
	uRoute := component.Router.Group("/user")
	{
		uRoute.POST("/", user.Create())
	}
}

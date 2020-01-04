package misc

import (
	"dp/component"
	"dp/httpd/handler/user"
)

func Routes() {
	component.Router.POST("/login", user.Login)
}

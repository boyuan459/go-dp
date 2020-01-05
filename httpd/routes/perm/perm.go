package perm

import (
	"dp/component"
	"dp/httpd/handler/perm"
)

func Routes() {
	pRoute := component.Router.Group("/permission")
	{
		pRoute.POST("/", perm.AddPerm)
	}
}

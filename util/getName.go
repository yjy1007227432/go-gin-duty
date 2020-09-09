package util

import "github.com/gin-gonic/gin"

type GetName struct {
	C gin.Context
}

func (g *GetName) GetName() string {
	name, _ := g.C.Get("name")
	return name.(string)
}
func (g *GetName) GetGroup() string {
	group, _ := g.C.Get("group")
	return group.(string)
}

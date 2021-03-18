package controllers

import "github.com/gin-gonic/gin"
import "rdo/models"
import "net/http"

func UsersIndex(c *gin.Context) {
	if !BeforeAll("user", c) {
		return
	}
	users, err := models.SelectUsers(Db)

	c.HTML(http.StatusOK, "users__index.tmpl", gin.H{
		"flash": err,
		"users": users,
		"name":  "name",
	})

}
func UsersCreate(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
func UsersNew(c *gin.Context) {
	c.HTML(http.StatusOK, "users__new.tmpl", gin.H{
		"flash":  "",
		"user":   map[string]string{},
		"flavor": c.Query("flavor"),
	})

}

func UsersShow(c *gin.Context) {
	c.HTML(http.StatusOK, "users__show.tmpl", gin.H{
		"flash": "",
		"user":  map[string]string{},
		"name":  "name",
	})

}

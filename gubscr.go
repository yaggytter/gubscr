package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "fmt"
    "os"
	"os/exec"
//    "html/template"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })
    router.Static("/public", "./public")
    router.Static("/gitrepos", "./gitrepos")

	/*
	   router.POST("/submit", func(c *gin.Context) {
	       c.String(http.StatusUnauthorized, "not authorized")
	   })
	   router.PUT("/error", func(c *gin.Context) {
	       c.String(http.StatusInternalServerError, "and error happened :(")
	   })
	*/
	router.GET("/view/:user/:reponame/:branch", func(c *gin.Context) {

		branch := c.Params.ByName("branch")
		user := c.Params.ByName("user")
		reponame := c.Params.ByName("reponame")

		gitcheckout("./gitrepos", branch, user, reponame)

/*
    t := template.Must(template.ParseFiles("template/index.html"))

    err := t.Execute(c.Writer, pv)
    if err != nil {
        c.String(500, "b"+err.Error())
        return
    }
*/
		c.String(http.StatusOK, branch+" "+user+" "+reponame)
	})
	router.Run(":8088")
}

func gitcheckout(localpath string, branch string, user string, reponame string) {
	cmd := exec.Command("cmd/gitcheckout.sh", localpath, branch, user, reponame)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

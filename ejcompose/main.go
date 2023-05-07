package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	//json tag to serialize json body
	Number int `json:"number"`
}

func main() {

	db, err := sql.Open("mysql", "root:pass@tcp(dbmysql:3306)/nums")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtOut, err := db.Prepare("SELECT squareNumber FROM sqarenum WHERE number = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	engine := gin.New()
	engine.GET("/test", func(context *gin.Context) {

		squareNum := 0
		body := Body{}
		err = stmtOut.QueryRow(13).Scan(&squareNum)
		if err != nil {
			panic(err.Error())
		}
		body.Number = squareNum
		context.JSON(http.StatusAccepted, &body)
	})
	engine.Run(":3000")

}

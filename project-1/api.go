package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
  Name string `json:"name"`
}


/*
  user: {
  key : value
  name : "krish"

}
*/

type Response struct {
  Message string `json:"message"`

} 

func SimpleHandler(ctx echo.Context) error {
  var req Request

  if err := ctx.Bind(&req); err != nil {
    return ctx.JSON(http.StatusBadRequest, Response{Message: "empty name, please provide name"})

  }
  message := fmt.Sprintf("Hello %s", req.Name)

    return ctx.JSON(http.StatusOK, Response{Message: message})
}

func main(){
  e := echo.New()

  e.POST("/hello",SimpleHandler)
  e.Start(":8080")
}	
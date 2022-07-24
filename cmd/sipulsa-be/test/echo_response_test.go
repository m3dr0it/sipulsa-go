package test

import (
	"log"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestEchoResponse(t *testing.T) {
	e := echo.ErrBadGateway
	log.Println(e)

	t.Log(e.Code)
	t.Log(e.Internal)
	t.Log(e.Error())
	t.Fail()
}

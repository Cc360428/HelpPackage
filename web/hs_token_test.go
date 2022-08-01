package web

import (
	"fmt"
	"testing"

	"github.com/Cc360428/HelpPackage/token"
)

func TestCreateToken(t *testing.T) {
	tokenStruct := new(token.Token)
	tokenStruct.Name = "lcc"
	tokenString := token.NewToken(tokenStruct)
	fmt.Printf("%s \n", tokenString)
}
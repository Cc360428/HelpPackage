package utils

import (
	"fmt"
	"github.com/Cc360428/HelpPackage/utils/token"
	"testing"
)

//
func TestCreateToken(t *testing.T) {
	tokenStruct := new(token.Token)
	tokenStruct.Name = "lcc"
	tokenString := token.NewToken(tokenStruct)
	fmt.Printf("%s \n", tokenString)
}

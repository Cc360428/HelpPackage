package swagger

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"testing"
)

func TestGetSwagger(t *testing.T) {
	swagger := GetSwagger("swagger.json")
	var requestAll []Swagger
	for key, value := range swagger {
		details := GetDetails(value)
		details.Url = key
		requestAll = append(requestAll, details)
	}
	logs.Info(requestAll[0])
}

package generator

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/reivaj05/GoConfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GeneratorTestSuite struct {
	suite.Suite
	assert          *assert.Assertions
	mockServiceName string
}

func (suite *GeneratorTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.mockServiceName = "mockService"

	GoConfig.Init(&GoConfig.ConfigOptions{
		ConfigType: "json",
		ConfigFile: "config",
		ConfigPath: "..",
	})
	GoConfig.SetConfigValue("templatesPath", "templates/")
}

func (suite *GeneratorTestSuite) TearDownSuite() {
	rollback(suite.mockServiceName)
}

func (suite *GeneratorTestSuite) TestJoinPath() {
	path := joinPath()
	suite.assert.True(strings.HasPrefix(path, os.Getenv("GOPATH")))
	suite.assert.True(strings.HasSuffix(path, "/src/github.com/reivaj05/apigateway"))
}

// TODO: Add unsuccessful tests
func (suite *GeneratorTestSuite) TestGenerateAPIFileSuccessful() {
	path := joinPath()
	suite.assert.Nil(generateAPIFile(path, suite.mockServiceName))
	_, err := os.Stat(fmt.Sprintf("../api/%s/%s.go",
		suite.mockServiceName, suite.mockServiceName))
	suite.assert.False(os.IsNotExist(err))
}

func (suite *GeneratorTestSuite) TestGenerateServiceFileSuccessful() {
	path := joinPath()
	suite.assert.Nil(generateServiceFile(path, suite.mockServiceName))
	_, err := os.Stat(fmt.Sprintf("../services/%s/%s.go",
		suite.mockServiceName, suite.mockServiceName))
	suite.assert.False(os.IsNotExist(err))
}

func (suite *GeneratorTestSuite) TestGenerateProtoFilesSuccessful() {
	path := joinPath()
	suite.assert.Nil(generateProtoFiles(path, suite.mockServiceName))
	_, err := os.Stat(fmt.Sprintf("../protos/api/%s.proto", suite.mockServiceName))
	suite.assert.False(os.IsNotExist(err))
	_, err = os.Stat(fmt.Sprintf("../protos/services/%s.proto", suite.mockServiceName))
	suite.assert.False(os.IsNotExist(err))
}

func TestGenerator(t *testing.T) {
	suite.Run(t, new(GeneratorTestSuite))
}

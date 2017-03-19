package generator

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GeneratorTestSuite struct {
	suite.Suite
	assert *assert.Assertions
}

func (suite *GeneratorTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	templatesPath = "templates/"
}

func (suite *GeneratorTestSuite) TearDownSuite() {
}

func (suite *GeneratorTestSuite) TestJoinPath() {
	path := joinPath()
	suite.assert.True(strings.HasPrefix(path, os.Getenv("GOPATH")))
	suite.assert.True(strings.HasSuffix(path, "/src/github.com/reivaj05/apigateway"))
}

// TODO: Add unsuccessful tests
func (suite *GeneratorTestSuite) TestGenerateAPIFileSuccessful() {
	path := joinPath()
	suite.assert.Nil(generateAPIFile(path, "mockAPIServiceName"))
	_, err := os.Stat("../api/mockAPIServiceName/mockAPIServiceName.go")
	suite.assert.False(os.IsNotExist(err))
}

func (suite *GeneratorTestSuite) TestGenerateServiceFileSuccessful() {
	path := joinPath()
	suite.assert.Nil(generateServiceFile(path, "mockServiceName"))
	_, err := os.Stat("../services/mockServiceName/mockServiceName.go")
	suite.assert.False(os.IsNotExist(err))
}

func (suite *GeneratorTestSuite) TestGenerateProtoFilesSuccessful() {
	path := joinPath()
	suite.assert.Nil(generateProtoFiles(path, "mockProtoName"))
	_, err := os.Stat("../protos/api/mockProtoName.proto")
	suite.assert.False(os.IsNotExist(err))
	// _, err = os.Stat("../protos/services/mockProtoName.proto")
	// suite.assert.False(os.IsNotExist(err))
}

func TestGenerator(t *testing.T) {
	suite.Run(t, new(GeneratorTestSuite))
}

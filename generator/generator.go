package generator

import (
	"fmt"
	"os"
)

func Generate(args ...interface{}) error {
	fmt.Println("TODO: Implement generator")
	path := joinPath()
	serviceName := "changeServiceName"
	if err := generateAPIFile(path, serviceName); err != nil {
		fmt.Println("api", err)
		return err
	}
	if err := generateServiceFile(path, serviceName); err != nil {
		fmt.Println("services", err)
		return err
	}
	if err := generateProtoFiles(path, serviceName); err != nil {
		fmt.Println("protos", err)
		return err
	}
	// TODO: Change to service name
	// TODO: Implement template files
	// TODO: Add tests

	return nil
}

func joinPath() string {
	const relativePath = "/src/github.com/reivaj05/apigateway"
	goPath := os.Getenv("GOPATH")
	return goPath + relativePath
}

func generateAPIFile(path, serviceName string) error {
	path += "/api/" + serviceName + "/"
	return _generateFile(path, serviceName, ".go")

}

func generateServiceFile(path, serviceName string) error {
	path += "/services/" + serviceName + "/"
	return _generateFile(path, serviceName, ".go")
}

func generateProtoFiles(path, serviceName string) error {
	pathProtosAPI := path + "/protos/api/"
	if err := _generateFile(pathProtosAPI, serviceName, ".proto"); err != nil {
		return err
	}
	pathProtosServices := path + "/protos/services/"
	return _generateFile(pathProtosServices, serviceName, ".proto")
}

func _generateFile(path, serviceName, extension string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(path + serviceName + extension)
	if err != nil {
		return err
	}
	defer file.Close()
	// tmpl := template.Must(template.ParseFiles(path))
	return nil
}

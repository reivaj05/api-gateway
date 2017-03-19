package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/chuckpreslar/inflect"
	"github.com/fatih/camelcase"
)

type generateOptions struct {
	path          string
	serviceName   string
	fileExtension string
	fileTemplate  string
	data          interface{}
}

func (op *generateOptions) getFilePath() string {
	return op.path + op.serviceName + op.fileExtension
}

type goAPITemplateData struct {
	ServiceName      string
	UpperServiceName string
}

type protoAPITemplateData struct {
	ServiceName      string
	ResourcePath     string
	UpperServiceName string
}

type EndpointsData struct {
	Services []string
}

// TODO: Add tests
// TODO: Refactor code when done

// TODO: Move to Config
var templatesPath = "generator/templates/"
var protosPath = "protos/api/"

func Generate(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Service name not provided")
	}
	createServices(args...)
	updateServerEndpoints()
	return runProtoGenScript()
}

func createServices(args ...string) {
	// TODO: Create properly template files (protos and go files)
	// TODO: Implement rest of template files
	basePath := joinPath()
	for _, serviceName := range args {
		if err := generateFiles(basePath, serviceName); err != nil {
			// TODO: Change msg
			fmt.Errorf("Service not created rolling back: " + err.Error())
			// TODO: Implement rollback when creating a service fails
			rollback(basePath, serviceName)
		}
	}

}

func joinPath() string {
	const relativePath = "/src/github.com/reivaj05/apigateway"
	goPath := os.Getenv("GOPATH")
	return goPath + relativePath
}

func generateFiles(path, serviceName string) error {
	fmt.Println(path, serviceName)
	if err := generateAPIFile(path, serviceName); err != nil {
		fmt.Println("api", err)
		return err
	}
	// TODO: Dont forget to uncomment :D
	// if err := generateServiceFile(path, serviceName); err != nil {
	// 	fmt.Println("services", err)
	// 	return err
	// }
	if err := generateProtoFiles(path, serviceName); err != nil {
		fmt.Println("protos", err)
		return err
	}
	return nil
}

func generateAPIFile(path, serviceName string) error {
	path += "/api/" + serviceName + "/"
	return _generateFile(&generateOptions{
		path:          path,
		serviceName:   serviceName,
		fileExtension: ".go",
		fileTemplate:  "goAPI.txt",
		data: &goAPITemplateData{
			ServiceName:      serviceName,
			UpperServiceName: inflect.Titleize(serviceName),
		},
	})

}

func generateServiceFile(path, serviceName string) error {
	path += "/services/" + serviceName + "/"
	return _generateFile(&generateOptions{
		path:          path,
		serviceName:   serviceName,
		fileExtension: ".go",
		fileTemplate:  "goAPI.txt",
	})
}

func generateProtoFiles(path, serviceName string) error {
	sp := camelcase.Split(serviceName)[0]
	if err := _generateFile(&generateOptions{
		path:          path + "/protos/api/",
		serviceName:   serviceName,
		fileExtension: ".proto",
		fileTemplate:  "protoAPI.txt",
		data: &protoAPITemplateData{
			ServiceName:      serviceName,
			ResourcePath:     sp,
			UpperServiceName: inflect.Titleize(serviceName),
		},
	}); err != nil {
		return err
	}
	// TODO: Dont forget to uncomment :D
	// return _generateFile(&generateOptions{
	// 	path:          path + "/protos/services/",
	// 	serviceName:   serviceName,
	// 	fileExtension: ".proto",
	// 	fileTemplate:  "protoAPI.txt",
	// })
	return nil
}

func _generateFile(options *generateOptions) error {
	file, err := _createFile(options)
	if err != nil {
		return err
	}
	return _writeTemplateContent(file, options)
}

func _createFile(options *generateOptions) (*os.File, error) {
	err := os.MkdirAll(options.path, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(options.getFilePath())
}

func _writeTemplateContent(file *os.File, options *generateOptions) error {
	defer file.Close()
	tmpl := template.Must(template.ParseFiles(
		templatesPath + options.fileTemplate),
	)
	return tmpl.Execute(file, options.data)
}

func rollback(path, serviceName string) {
	// TODO: Rollback to previous stable point if something goes wrong
}

func updateServerEndpoints() {
	services := getServicesNames()
	_updateServerFiles(&EndpointsData{Services: services})

}

func getServicesNames() (services []string) {
	files, _ := ioutil.ReadDir(protosPath)
	for _, file := range files {
		services = append(services, strings.TrimSuffix(file.Name(), ".proto"))
	}
	return services
}

func _updateServerFiles(data *EndpointsData) {
	_updateFile(data, "registeredGRPCEndpoints.go", "grpcEndpoints.txt")
	_updateFile(data, "registeredHTTPEndpoints.go", "httpEndpoints.txt")
}

func _updateFile(data *EndpointsData, endpointsFile, fileTemplate string) {
	file, err := os.OpenFile("server/"+endpointsFile, os.O_RDWR, 0660)
	if err != nil {
		// rollback(path, serviceName)
	}
	if err = _writeTemplateContent(file, &generateOptions{
		data: data, fileTemplate: fileTemplate}); err != nil {
		// rollback(path, serviceName)
	}
}

func runProtoGenScript() error {
	cmd := exec.Command("/bin/sh", "proto-gen.sh")
	return cmd.Run()
}

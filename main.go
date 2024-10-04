package main

import (
	"encoding/json"
	"github.com/docker/docker/client"
	"github.com/docker/go-plugins-helpers/secrets"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"text/template"
)

const (
	// socket address
	socketAddress = "/run/docker/plugins/vault.sock"
)

var (
	log                      = logrus.New()
	policyTemplateExpression string
	policyTemplate           *template.Template
	// secretZero               string
)

type vaultSecretsDriver struct {
	dockerClient *client.Client
}

func (d vaultSecretsDriver) Get(req secrets.Request) secrets.Response {
	//errorResponse := func(s string, err error) secrets.Response {
	//	log.Errorf("Error getting secret %q: %s: %v", req.SecretName, s, err)
	//	return secrets.Response{
	//		Value: []byte("-"),
	//		Err:   fmt.Sprintf("%s: %v", s, err),
	//	}
	//}
	valueResponse := func(s string) secrets.Response {
		return secrets.Response{
			Value:      []byte(s),
			DoNotReuse: true,
		}
	}

	resultBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Failed to json.Marshal: %v", err)
	}

	return valueResponse(string(resultBytes))
}

// Read "secret zero" from the file system of a helper service task container, then serve the plugin.
func main() {
	// Create Docker client
	var httpClient *http.Client
	dockerAPIVersion := os.Getenv("DOCKER_API_VERSION")
	cli, err := client.NewClient("unix:///var/run/docker.sock", dockerAPIVersion, httpClient, nil)
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	// Create the driver
	d := vaultSecretsDriver{
		dockerClient: cli,
	}
	h := secrets.NewHandler(d)

	// Serve plugin
	if err := h.ServeUnix("plugin", 0); err != nil {
		log.Errorf("Error serving plugin: %v", err)
	}

}

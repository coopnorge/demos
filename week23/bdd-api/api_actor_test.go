package actor

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/xeipuuv/gojsonschema"
)

const errActorNotCreated = "API Actor context not created"

var (
	opts = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "progress",
	}
	actorContext *actorApi
)

type actorApi struct {
	client   *http.Client
	lastResp *response
	headers  map[string]string
}

type response struct {
	StatusCode  int
	Body        string
	ResponseObj *http.Response
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                "TestingAPIs",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	actorContext = &actorApi{
		client:   new(http.Client),
		lastResp: nil,
		headers:  make(map[string]string),
	}

	ctx.Step(`^I set header "([^"]*)" with value "([^"]*)"$`, iSetHeaderWithValue)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, iSendRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response should be a valid json$`, theResponseShouldBeValidJSON)
	ctx.Step(`^the response should match json schema "([^"]*)"$`, theResponseShouldMatchJsonSchema)
}

func iSetHeaderWithValue(name string, value string) error {
	if actorContext != nil {
		actorContext.headers[name] = value
	}

	return nil
}

func iSendRequestTo(method, endpoint string) error {
	if actorContext == nil {
		return fmt.Errorf(errActorNotCreated)
	}

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}

	// Add headers to request
	for name, value := range actorContext.headers {
		req.Header.Set(name, value)
	}
	resp, err := actorContext.client.Do(req)
	if err != nil {
		return err
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return err2
	}

	actorContext.lastResp = &response{
		StatusCode:  resp.StatusCode,
		ResponseObj: resp,
		Body:        string(body),
	}

	return nil
}

func theResponseCodeShouldBe(code int) error {
	if actorContext != nil && code != actorContext.lastResp.StatusCode {
		return fmt.Errorf("expected response code: %d, but it is: %d", code, actorContext.lastResp.StatusCode)
	}

	return nil
}

func theResponseShouldBeValidJSON() error {
	if actorContext == nil {
		return fmt.Errorf(errActorNotCreated)
	}

	var data interface{}
	return json.Unmarshal([]byte(actorContext.lastResp.Body), &data)
}

func theResponseShouldMatchJsonSchema(path string) error {

	path = strings.Trim(path, "/")
	absPath, absErr := filepath.Abs(path)
	if absErr != nil {
		return fmt.Errorf("unalbe to resolve file path: %s", absPath)
	}
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", absPath)
	}

	schemaContents, err := ioutil.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}

	schemaLoader := gojsonschema.NewBytesLoader(schemaContents)
	documentLoader := gojsonschema.NewStringLoader(actorContext.lastResp.Body)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		return err
	}

	if !result.Valid() {
		var schemaErrors []string
		for _, schemaErr := range result.Errors() {
			schemaErrors = append(schemaErrors, schemaErr.String())
		}

		return fmt.Errorf("The response has other Json scheam %s\n %v", path, schemaErrors)
	}

	return nil
}

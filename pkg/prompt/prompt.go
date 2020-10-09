package prompt

import (
	"context"
	"fmt"
	"regexp"

	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	sectionLength   = 80
	morePrompt      = "Add another?"
	appNameDefault  = "my-app"
	httpPortDefault = 8080
	grpcPortDefault = 50505
)

// Start starts the wizard
func Start(ctx context.Context) (app *model.App, err error) {
	app = &model.App{}

	Header("Application")

	// name
	app.Meta.Name = ForString("Name: ", appNameDefault)

	// app type
	app.Meta.Type = ForOption("App type: ", model.AppTypeCLI, model.AppTypeGRPC, model.AppTypeHTTP)
	switch app.Meta.Type {
	case model.AppTypeGRPC:
		app.Meta.Port = grpcPortDefault
	case model.AppTypeHTTP:
		app.Meta.Port = httpPortDefault
	}

	// lang
	app.Meta.Lang = ForOption("Language: ", model.LangGo)
	switch app.Meta.Lang {
	case model.LangGo:
		app.Meta.Main = "main.go"
	}

	// port
	if app.Meta.Type != model.AppTypeCLI {
		app.Meta.Port = ForInt("App protocol port: ", app.Meta.Port)
	}

	Header("Pub/Sub")

	// pubsub
	if ForBool("Subscribe to topic?") {
		list := make([]*model.PubSub, 0)
		for {
			comp, err := ForPubSub()
			if err != nil {
				return nil, errors.Errorf("Error getting pub/sub components: %v.", err)
			}
			list = append(list, comp)
			if !ForBool(morePrompt) {
				break
			}
		}
		if len(list) > 0 {
			app.PubSubs = list
		}
	}

	Header("Binding")

	// binding
	if ForBool("Use input binding?") {
		list := make([]*model.Component, 0)
		for {
			comp, err := ForBinding()
			if err != nil {
				return nil, errors.Errorf("Error getting binding components: %v.", err)
			}
			list = append(list, comp)
			if !ForBool(morePrompt) {
				break
			}
		}
		if len(list) > 0 {
			app.Bindings = list
		}
	}

	Header("Service Invocation")

	// service
	if ForBool("Enable service invocation?") {
		list := make([]*model.Service, 0)
		for {
			comp, err := ForService()
			if err != nil {
				return nil, errors.Errorf("Error getting service: %v.", err)
			}
			list = append(list, comp)

			if !ForBool(morePrompt) {
				break
			}
		}
		if len(list) > 0 {
			app.Services = list
		}
	}

	Header("Secrets")

	// secret
	app.Components = make([]*model.Component, 0)
	if ForBool("Use secrets?") {
		secretComp, err := ForComponents(model.SecretComponentTypes(), "secret", "secretstores")
		if err != nil {
			return nil, errors.Errorf("Error parsing answer: %v.", err)
		}
		app.Components = append(app.Components, secretComp...)
	}

	Header("Dapr Client")

	// client
	app.Meta.UsesClient = ForBool("Uses Dapr client?")
	if app.Meta.UsesClient {

		// state
		if ForBool("Add state components?") {
			stateComp, err := ForComponents(model.StateComponentTypes(), "store", "state")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, stateComp...)
		}

		// pubsub
		if ForBool("Add pub/sub components?") {
			pubsubComp, err := ForComponents(model.PubsubComponentTypes(), "pubsub", "pubsub")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, pubsubComp...)
		}

		// binding
		if ForBool("Add output binding components?") {
			outBindComp, err := ForComponents(model.OutputBindingComponentTypes(), "binding", "bindings")
			if err != nil {
				return nil, errors.Errorf("Error parsing answer: %v.", err)
			}
			app.Components = append(app.Components, outBindComp...)
		}

	}

	return
}

func codeSafeString(val string) string {
	reg := regexp.MustCompile("[^a-zA-Z]+")
	return reg.ReplaceAllString(val, "")
}

// Content prints the provided object
func Content(v interface{}) {
	fmt.Println()
	fmt.Println(v)
	fmt.Println()
}

// Header prints provided title as header
func Header(title string) {
	s := fmt.Sprintf("=== %s ", title)
	for i := len(s); i < sectionLength; i++ {
		s = s + "="
	}

	fmt.Println()
	fmt.Println(s)
	fmt.Println()
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/grokify/simplego/fmt/fmtutil"
	hum "github.com/grokify/simplego/net/httputilmore"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	rcu "github.com/grokify/go-ringcentral-client/office/v1/util"
	ro "github.com/grokify/oauth2more/ringcentral"
)

func loadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

func listRules(apiClient *rc.APIClient) {
	ctx := context.Background()
	resp, err := apiClient.CallHandlingSettingsApi.LoadAnsweringRulesList(
		ctx, "~", "~", &rc.LoadAnsweringRulesListOpts{})
	if err != nil {
		panic(err)
	}

	err = hum.PrintResponse(resp, true)
	if err != nil {
		panic(err)
	}
}

type DemoConfig struct {
	ListRules bool
}

func main() {
	var listRulesFlag = flag.Int("list", 0, "Create user")

	flag.Parse()

	fmt.Printf("[%v]\n", *listRulesFlag)

	cfg := DemoConfig{}
	if *listRulesFlag != 0 {
		cfg.ListRules = true
	}

	fmtutil.PrintJSON(cfg)

	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	apiClient, err := rcu.NewApiClientPassword(
		ro.NewApplicationCredentialsEnv(),
		ro.NewPasswordCredentialsEnv(),
	)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.ListRules {
		listRules(apiClient)
	}

	fmt.Println("DONE")
}

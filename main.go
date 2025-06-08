package main

import (
	"context"
	"fmt"
	"os"

	"github.com/keenbytes/broccli/v3"

	"github.com/Cardinal-Cryptography/github-actions-validator/pkg/dotgithub"
)

func main() {
	cli := broccli.NewBroccli("github-actions-validator", "Validates GitHub Actions' .github directory", "infra-team@cardinals")
	cmdValidate := cli.Command("validate", "Runs the validation on files from a specified directory", validateHandler)
	cmdValidate.Flag("path", "p", "", "Path to .github directory", broccli.TypePathFile, broccli.IsDirectory|broccli.IsExistent|broccli.IsRequired)
	cmdValidate.Flag("vars-file", "z", "", "Check if variable names exist in this file (one per line)", broccli.TypePathFile, broccli.IsExistent)
	cmdValidate.Flag("secrets-file", "s", "", "Check if secret names exist in this file (one per line)", broccli.TypePathFile, broccli.IsExistent)
	_ = cli.Command("version", "Prints version", versionHandler)
	if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		os.Args = []string{"App", "version"}
	}

	os.Exit(cli.Run(context.Background()))
}

func versionHandler(ctx context.Context, c *broccli.Broccli) int {
	fmt.Fprintf(os.Stdout, VERSION+"\n")
	return 0
}

func validateHandler(ctx context.Context, c *broccli.Broccli) int {
	dotGithub := dotgithub.DotGithub{
		Path:        c.Flag("path"),
		VarsFile:    c.Flag("vars-file"),
		SecretsFile: c.Flag("secrets-file"),
	}
	err := dotGithub.InitFiles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "!!!! Error with initialization: %s\n", err.Error())
		return 1
	}
	validationErrors, err := dotGithub.Validate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "!!!! Error with validation: %s\n", err.Error())
		return 1
	}
	for _, verr := range validationErrors {
		fmt.Fprintf(os.Stdout, "%s\n", verr)
	}
	return 0
}

package cli

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/auth0/auth0-cli/internal/prompt"
	"github.com/spf13/cobra"
	"gopkg.in/auth0.v5"
)

type commandInput interface {
	GetName() string
	GetLabel() string
	GetHelp() string
	GetIsRequired() bool
}

func ask(cmd *cobra.Command, i commandInput, value interface{}, defaultValue *string, isUpdate bool) error {
	isRequired := !isUpdate && i.GetIsRequired()
	input := prompt.TextInput("", i.GetLabel(), i.GetHelp(), auth0.StringValue(defaultValue), isRequired)

	if err := prompt.AskOne(input, value); err != nil {
		return handleInputError(err)
	}

	return nil
}

func _select(cmd *cobra.Command, i commandInput, value interface{}, options []string, defaultValue *string, isUpdate bool) error {
	isRequired := !isUpdate && i.GetIsRequired()
	input := prompt.SelectInput("", i.GetLabel(), i.GetHelp(), options, auth0.StringValue(defaultValue), isRequired)

	if err := prompt.AskOne(input, value); err != nil {
		return handleInputError(err)
	}

	return nil
}

func inputLabel(name string) string {
	return fmt.Sprintf("%s:", name)
}

func handleInputError(err error) error {
	if err == terminal.InterruptErr {
		os.Exit(0)
	}

	return unexpectedError(err)
}

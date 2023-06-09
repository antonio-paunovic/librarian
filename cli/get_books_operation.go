// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"librarian/client/books"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBooksGetBooksCmd returns a cmd to handle operation getBooks
func makeOperationBooksGetBooksCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetBooks",
		Short: ``,
		RunE:  runOperationBooksGetBooks,
	}

	if err := registerOperationBooksGetBooksParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBooksGetBooks uses cmd flags to call endpoint api
func runOperationBooksGetBooks(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := books.NewGetBooksParams()
	if err, _ := retrieveOperationBooksGetBooksFilterFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBooksGetBooksResult(appCli.Books.GetBooks(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBooksGetBooksParamFlags registers all flags needed to fill params
func registerOperationBooksGetBooksParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBooksGetBooksFilterParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBooksGetBooksFilterParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDescription := ``

	var filterFlagName string
	if cmdPrefix == "" {
		filterFlagName = "filter"
	} else {
		filterFlagName = fmt.Sprintf("%v.filter", cmdPrefix)
	}

	var filterFlagDefault string

	_ = cmd.PersistentFlags().String(filterFlagName, filterFlagDefault, filterDescription)

	return nil
}

func retrieveOperationBooksGetBooksFilterFlag(m *books.GetBooksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter") {

		var filterFlagName string
		if cmdPrefix == "" {
			filterFlagName = "filter"
		} else {
			filterFlagName = fmt.Sprintf("%v.filter", cmdPrefix)
		}

		filterFlagValue, err := cmd.Flags().GetString(filterFlagName)
		if err != nil {
			return err, false
		}
		m.Filter = &filterFlagValue

	}
	return nil, retAdded
}

// parseOperationBooksGetBooksResult parses request result and return the string content
func parseOperationBooksGetBooksResult(resp0 *books.GetBooksOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*books.GetBooksDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*books.GetBooksOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

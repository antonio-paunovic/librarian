// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"librarian/client/collections"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationCollectionsDeleteCollectionsCmd returns a cmd to handle operation deleteCollections
func makeOperationCollectionsDeleteCollectionsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "DeleteCollections",
		Short: ``,
		RunE:  runOperationCollectionsDeleteCollections,
	}

	if err := registerOperationCollectionsDeleteCollectionsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationCollectionsDeleteCollections uses cmd flags to call endpoint api
func runOperationCollectionsDeleteCollections(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := collections.NewDeleteCollectionsParams()
	if err, _ := retrieveOperationCollectionsDeleteCollectionsNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationCollectionsDeleteCollectionsResult(appCli.Collections.DeleteCollections(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationCollectionsDeleteCollectionsParamFlags registers all flags needed to fill params
func registerOperationCollectionsDeleteCollectionsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationCollectionsDeleteCollectionsNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationCollectionsDeleteCollectionsNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Name of collection`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func retrieveOperationCollectionsDeleteCollectionsNameFlag(m *collections.DeleteCollectionsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

	}
	return nil, retAdded
}

// parseOperationCollectionsDeleteCollectionsResult parses request result and return the string content
func parseOperationCollectionsDeleteCollectionsResult(resp0 *collections.DeleteCollectionsNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*collections.DeleteCollectionsDefault)
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
		resp0, ok := iResp0.(*collections.DeleteCollectionsNoContent)
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

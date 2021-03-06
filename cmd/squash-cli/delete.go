package main

import (
	"errors"

	"fmt"

	"github.com/solo-io/squash/pkg/client/debugattachment"

	"github.com/spf13/cobra"
)

func init() {

	var rmCmd = &cobra.Command{
		Use:     "delete [id]",
		Short:   "delete a debug attachment object",
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				return errors.New("Debug attachment id required")
			}

			c, err := getClient()
			if err != nil {
				return err
			}

			for _, id := range args {

				params := debugattachment.NewDeleteDebugAttachmentParams()
				params.DebugAttachmentID = id
				_, err := c.Debugattachment.DeleteDebugAttachment(params)

				if !jsonoutput && (err != nil) {
					fmt.Print("Error deleting id: ", id, err)
				}

			}

			if jsonoutput {
				fmt.Println("{}")
			}
			return nil
		},
	}
	RootCmd.AddCommand(rmCmd)

}

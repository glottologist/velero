/*
Copyright 2018 the Heptio Ark contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package snapshotlocation

import (
	"github.com/spf13/cobra"

	"github.com/heptio/ark/pkg/client"
)

func NewCommand(f client.Factory) *cobra.Command {
	c := &cobra.Command{
		Use:   "snapshot-location",
		Short: "Work with snapshot locations",
		Long:  "Work with snapshot locations",
	}

	c.AddCommand(
		NewCreateCommand(f, "create"),
		NewGetCommand(f, "get"),
	)

	return c
}

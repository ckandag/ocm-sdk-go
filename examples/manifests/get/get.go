/*
Copyright (c) 2024 Red Hat, Inc.

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

// This example shows how to update the display name of a cluster.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/openshift-online/ocm-sdk-go/logging"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := logging.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Client(clientId, clientSecret).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of clusters:
	collection := connection.ClustersMgmt().V1().Clusters()

	// Get the client for the resource that manages the cluster that we want to update. Note
	// that this will not yet send any request to the server, so it will succeed even if the
	// cluster doesn't exist.
	resource := collection.Cluster("<cluster_id>")
	manifest, err := resource.ExternalConfiguration().Manifests().Manifest("example").Get().Parameter("fetch_live", "true").SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't retrieve manifest of cluster: %v\n", err)
		os.Exit(1)
	}
	err = v1.MarshalManifest(manifest.Body(), os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't write manifest to stdout: %v\n", err)
		os.Exit(1)
	}
}

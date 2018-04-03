// Copyright © 2017 The Kubicorn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import(
	"errors"
	"os"
	
	"github.com/kubicorn/kubicorn/pkg/cli"
	"github.com/kubicorn/kubicorn/pkg/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

func SshCmd() *cobra.Command {
	var sshOpt = &cli.SshOptions{}
	var sshCmd = &cobra.Command {
		Use:	"ssh <name>",
		Short:	"Access cluster node",
		Long:	`Use this command to run ssh into the master node`,
		Run:	func(cmd *cobra.Command, args []string){
			if(len(args) != 1) {
				logger.Critical("Incorrect number of arguments.")
				os.Exit(1)
			}
			sshOpt.Name = args[0]
			if err := runSsh(sshOpt); err != nil {
				logger.Critical(err.Error())
				os.Exit(1)
			}
		},
	}

	fs := sshCmd.Flags()

	fs.StringVarP(&sshOpt.StateStorePath, keyStateStorePath, "S", viper.GetString(keyStateStorePath), descStateStorePath)
	fs.StringVarP(&sshOpt.StateStore, keyStateStore, "s", viper.GetString(keyStateStore), descStateStore)

	return sshCmd
}

func runSsh(options *cli.SshOptions) error{

	// Name field is required
	name := options.Name
	if name == "" {
		return errors.New("Empty name. Must specify the name of the cluster to access")
	}
	
	// Expand state store path
	options.StateStorePath = cli.ExpandPath(options.StateStorePath)

	// Register state store
	stateStore, err := options.NewStateStore()
	if err != nil{
		logger.Critical("Unable to retrieve state store: " + err.Error())
	}

	// Retrieve cluster configuration
	cluster, err := stateStore.GetCluster()
	if err != nil{
		logger.Critical("Unable to retrieve clusters: " + err.Error())
	}
	logger.Info("Retrieved cluster %s", cluster.Name)

	// Make the call
	config := &ssh.ClientConfig{
		User: cluster.ProviderConfig().SSH.User,
		Auth: []ssh.AuthMethod{
			ssh.Password("yourpassword"),
		},
	}
	client, err := ssh.Dial("tcp", cluster.ProviderConfig().KubernetesAPI.Endpoint, config)

	session, err := client.NewSession()
	if err != nil {
		logger.Critical("Failed to create session: ", err.Error())
	}
	defer session.Close()

	return nil
}

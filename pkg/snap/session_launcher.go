// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snap

import (
	"time"

	"github.com/intelsdi-x/snap/mgmt/rest/client"
	"github.com/intelsdi-x/snap/scheduler/wmap"
)

// SessionConfig holds configuration for snap session
type SessionConfig struct {
	SnapteldAddress string
	Publisher       *wmap.PublishWorkflowMapNode
	Interval        time.Duration
	Plugins         []string
	TaskName        string
	Metrics         []string
	Tags            map[string]interface{}
}

// NewSessionLauncher construct snap.Session based on provied config
// It also try load plugins mentioned in config.Plugins and can fails
// if plugins are unavailable.
func NewSessionLauncher(config SessionConfig) (*Session, error) {
	snapClient, err := client.New(config.SnapteldAddress, "v1", true)
	if err != nil {
		return nil, err
	}

	err = LoadPlugins(config.Plugins...)
	if err != nil {
		return nil, err
	}

	return NewSession(
		config.TaskName,
		config.Metrics,
		config.Interval,
		snapClient,
		config.Publisher,
		config.Tags,
	), nil
}

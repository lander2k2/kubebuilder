/*
Copyright 2020 The Kubernetes Authors.

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

package scaffold

import (
	"sigs.k8s.io/kubebuilder/internal/config"
)

var _ Scaffolder = &editScaffolder{}

type editScaffolder struct {
	config     *config.Config
	multigroup bool
}

// NewEditScaffolder returns a new Scaffolder for configuration edit operations
func NewEditScaffolder(config *config.Config, multigroup bool) Scaffolder {
	return &editScaffolder{
		config:     config,
		multigroup: multigroup,
	}
}

// Scaffold implements Scaffolder
func (s *editScaffolder) Scaffold() error {
	s.config.MultiGroup = s.multigroup

	return s.config.Save()
}

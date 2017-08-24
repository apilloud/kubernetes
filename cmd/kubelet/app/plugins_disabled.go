// +build !kubernetes_plugins

/*
Copyright 2017 The Kubernetes Authors.

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

package app

// This file exists to force no plugin implementations to be linked.
import (
	// Network plugins
	"k8s.io/kubernetes/pkg/kubelet/network"
	// Volume plugins
	"k8s.io/kubernetes/pkg/volume"
)

// ProbeVolumePlugins collects all volume plugins into an easy to use list.
// PluginDir specifies the directory to search for additional third party
// volume plugins.
func ProbeVolumePlugins(pluginDir string) []volume.VolumePlugin {
	return []volume.VolumePlugin{}
}

// ProbeNetworkPlugins collects all compiled-in plugins
func ProbeNetworkPlugins(pluginDir, cniConfDir, cniBinDir string) []network.NetworkPlugin {
	return []network.NetworkPlugin{}
}

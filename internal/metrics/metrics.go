/*
Copyright 2024 IBM Corporation.

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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	AppWrapperPhaseCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "appwrapper_phase_total",
			Help: `The total number of times an appwrapper transitioned to a given phase per namespace.`,
		}, []string{"namespace", "phase"},
	)
)

func Register() {
	metrics.Registry.MustRegister(AppWrapperPhaseCounter)
}

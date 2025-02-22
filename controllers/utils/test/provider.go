// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package test

import (
	"go.uber.org/fx"

	"github.com/chaos-mesh/chaos-mesh/cmd/chaos-controller-manager/provider"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/recorder"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/test/manager"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/test/metrics"
)

var Module = fx.Provide(
	provider.NewOption,
	provider.NewClient,
	provider.NewLogger,
	provider.NewAuthCli,
	provider.NewScheme,
	provider.NewNoCacheReader,
	provider.NewGlobalCacheReader,
	provider.NewControlPlaneCacheReader,
	manager.NewTestManager,
	recorder.NewRecorderBuilder,
	metrics.NewTestChaosControllerManagerMetricsCollector,
)

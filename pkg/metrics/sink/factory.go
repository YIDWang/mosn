/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sink

import (
	"errors"
	"fmt"
	"plugin"

	"mosn.io/api"
	v2 "mosn.io/mosn/pkg/config/v2"
	"mosn.io/mosn/pkg/types"
)

// MetricsSinkCreator creates a MetricsSink according to config
type MetricsSinkCreator func(config map[string]interface{}) (types.MetricsSink, error)

var metricsSinkFactory map[string]MetricsSinkCreator

func init() {
	metricsSinkFactory = make(map[string]MetricsSinkCreator)
}

// RegisterSink registers the sinkType as MetricsSinkCreator
func RegisterSink(sinkType string, creator MetricsSinkCreator) {
	metricsSinkFactory[sinkType] = creator
}

// CreateMetricsSink creates a MetricsSink according to sinkType
func CreateMetricsSink(sinkType string, config map[string]interface{}) (types.MetricsSink, error) {
	if creator, ok := metricsSinkFactory[sinkType]; ok {
		sink, err := creator(config)
		if err != nil {
			return nil, fmt.Errorf("create metrics sink failed: %v", err)
		}
		return sink, nil
	}
	return nil, fmt.Errorf("unsupported metrics sink type: %v", sinkType)
}

func CreateMetricsSinkByPlugin(pluginConfig *v2.StreamFilterGoPluginConfig, factoryConfig map[string]interface{}) (api.MetricsSink, error) {
	if pluginConfig.SoPath == "" {
		return nil, errors.New("so file path could not be found")
	}
	p, err := plugin.Open(pluginConfig.SoPath)
	if err != nil {
		return nil, err
	}
	if pluginConfig.FactoryMethod == "" {
		pluginConfig.FactoryMethod = "CreateMetricsSink"
	}
	f, err := p.Lookup(pluginConfig.FactoryMethod)
	if err != nil {
		return nil, err
	}
	function, ok := f.(func(map[string]interface{}) (api.MetricsSink, error))
	if !ok {
		return nil, errors.New("failed to get correct factory method")
	}
	return function(factoryConfig)
}

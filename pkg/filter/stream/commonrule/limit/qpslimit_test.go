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

package limit

import (
	"testing"
	"time"
	"math"
	"github.com/alipay/sofa-mosn/pkg/log"
	"github.com/alipay/sofa-mosn/pkg/filter/stream/commonrule/metrix"
)

func TestQpsLimiter_TryAcquire(t *testing.T) {
	limiter, err := NewQpsLimiter(0, 1000)
	if err != nil {
		t.Errorf("%v",err)
	}
	res := limiter.TryAcquire()
	if res {
		t.Errorf("false")
	} else {
		t.Log("ok")
	}
}

func TestQpsLimiter_TryAcquire1(t *testing.T) {
	limiter, err := NewQpsLimiter(1, 1000)
	if err != nil {
		t.Errorf("%v",err)
	}

	total := 0
	success := 0
	ticker := metrix.NewTicker(func() {
		total ++
		res := limiter.TryAcquire();
		if res {
			success ++
		}
	})
	ticker.Start(time.Millisecond * 100)
	time.Sleep(5* time.Second)
	ticker.Stop()
	if math.Abs(float64(success - 5)) > 1 {
		t.Errorf("false, success=%d", success)
	} else {
		t.Log("total = ", total)
		t.Log("success = ", success)
	}
}

func TestQpsLimiter_TryAcquire2(t *testing.T) {
	maxAllowsList := []int64{1,5,10,100,500}
	periodMsList := []int64{1000, 5000}
	intervals := []time.Duration{1000, 100, 10}
	sleepDuration := 5 * time.Second

	log.DefaultLogger.Infof("start")
	for _, maxAllow := range maxAllowsList {
		for _, periodMs := range periodMsList {
			for _, interval := range intervals {
				log.DefaultLogger.Infof("maxAllow=%d, periodMs=%d, interval=%d", maxAllow, periodMs, interval)
				limiter, _ := NewQpsLimiter(maxAllow, periodMs)

				total := 0
				success := 0
				ticker := metrix.NewTicker(func() {
					total ++
					res := limiter.TryAcquire();
					if res {
						success ++
					}
				})
				ticker.Start(time.Millisecond * interval)
				time.Sleep(sleepDuration)
				ticker.Stop()
				threshold := math.Min(float64(maxAllow * 1000 * 5) / float64(periodMs), float64(5 * 1000/interval))
				log.DefaultLogger.Infof("total=%d, success=%d, threshold=%f", total, success, threshold)
				if math.Abs(float64(success) - threshold) > 1 {
					t.Errorf("false, success=%d", success)
				}
			}
		}
	}
	t.Log("end")
}
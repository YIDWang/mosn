package rules

import (
	"context"
	"sync"

	"mosn.io/mosn/pkg/types"
)

var tranferInstance *TransferRule

func init() {
	tranferInstance = NewTransferRule()
}

type TransferRule struct {
	rules sync.Map
}

func NewTransferRule() *TransferRule {
	return &TransferRule{}
}

func GetInstanceTransferRule() *TransferRule {
	return tranferInstance
}

type TransferRuleInfo struct {
	UpstreamProtocol    string
	UpstreamSubProtocol string
	Config              map[string]interface{}
	Description         string
}

// type TransferRule struct {
// Matchers []TransferRuleMatcher `json:"transfer_matcher"`
// RuleInfo *TransferRuleInfo `json:"transfer_rule_info"`
// }

func (l *TransferRule) AddOrUpdateTransferRule(value string) error {
	/*
		if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
			// log.DefaultLogger.Errorf("[drm_transfer_rule] [OnDRMPush] dataId:%v,value:%v", dataId, value)
		}
		transferRule := &TransferRule{}
		if err := json.Unmarshal([]byte("value"), transferRule); err != nil {
			// log.DefaultLogger.Errorf("[drm_transfer_rule] [OnDRMPush] json Unmarshal failed,err:%v", err)
			return
		}

		// md5Value := common.CalculateMd5(value)
		if oldRule, ok := l.rules.LoadOrStore("xx", transferRule); ok {
			if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
				// log.DefaultLogger.Errorf("[drm_transfer_rule] [OnDRMPush] dataId:%v,oldvalue:%v", dataId, oldRule)
			}
		}
	*/
}

func (tf *TransferRule) Match(ctx context.Context, headers types.HeaderMap) (ruleInfo *TransferRuleInfo, result bool) {
	/*
		// TODO listener
		tf.rules.Range(func(key, value interface{}) bool {
			tranRule := value.(*TransferRule)
			for _, match := range tranRule.Matchers {
				if result = match.Match(ctx, headers); !result {
					continue
				}
				ruleInfo = tranRule.RuleInfo
				break
			}
			return !result
		})

		if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
			log.DefaultLogger.Errorf("[drm_transfer_rule] [Match] info:%v,result:%v", ruleInfo, result)
		}
	*/
	return ruleInfo, result
}

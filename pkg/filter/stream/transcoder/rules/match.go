package rules

import (
	"context"
	"encoding/json"

	"mosn.io/mosn/pkg/router"
	"mosn.io/mosn/pkg/types"
)

type TransferMatcher interface {
	Match(ctx context.Context, headers types.HeaderMap) bool
}

type transferMatcher struct {
	// VariableMatchers []router        `json:"variables"`
	HeaderMatchers []headerMatcher `json:"headers"`
}

func (tfm *transferMatcher) Match(ctx context.Context, headers types.HeaderMap) bool {
	result := false
	for _, matcher := range tfm.HeaderMatchers {
		result = matcher.Match(ctx, headers)
		if result {
			return result
		}
	}
	for _, matcher := range tfm.VariableMatchers {
		result = matcher.Match(ctx, headers)
		if result {
			return result
		}
	}
	return false
}

type headerMatcher struct {
}

func (hm *headerMatcher) UnmarshalJSON(v []byte) error {
	// router.ParseToVariableMatchItem()
	if err := json.Unmarshal(v, &hm); err != nil {
		return err
	}
	// hm.matcher = requirement.BuildRequirement(requirement.REQUIREMENT_INFO_HEADER, hm.Name, hm.Operator, hm.Values)
	return nil
}

func (hm *headerMatcher) Match(ctx context.Context, headers types.HeaderMap) bool {
	return hm.matcher.Match(headers, ctx)
}

type variableMatcher struct {
	Variables []*router.VariableMatchItem
}

func (vm *variableMatcher) UnmarshalJSON(v []byte) error {
}

func (vm *variableMatcher) Match(ctx context.Context, headers types.HeaderMap) bool {
	return vm.matcher.Match(headers, ctx)
}

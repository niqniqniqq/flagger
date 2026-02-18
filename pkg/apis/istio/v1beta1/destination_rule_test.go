package v1beta1

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLocalityLbSettingEnabledFalseIsMarshaled(t *testing.T) {
	enabled := false
	spec := DestinationRuleSpec{
		Host: "podinfo",
		TrafficPolicy: &TrafficPolicy{
			LoadBalancer: &LoadBalancerSettings{
				Simple: SimpleLBRandom,
				LocalityLbSetting: &LocalityLbSetting{
					Enabled: &enabled,
				},
			},
		},
	}

	b, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("failed to marshal DestinationRuleSpec: %v", err)
	}

	if !strings.Contains(string(b), `"enabled":false`) {
		t.Fatalf("expected JSON to contain enabled=false, got: %s", string(b))
	}
}

func TestLocalityLbSettingEnabledNilIsOmitted(t *testing.T) {
	spec := DestinationRuleSpec{
		Host: "podinfo",
		TrafficPolicy: &TrafficPolicy{
			LoadBalancer: &LoadBalancerSettings{
				Simple:            SimpleLBRandom,
				LocalityLbSetting: &LocalityLbSetting{},
			},
		},
	}

	b, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("failed to marshal DestinationRuleSpec: %v", err)
	}

	if strings.Contains(string(b), `"enabled":`) {
		t.Fatalf("expected JSON not to contain enabled when nil, got: %s", string(b))
	}
}

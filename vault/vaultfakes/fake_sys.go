// Code generated by counterfeiter. DO NOT EDIT.
package vaultfakes

import (
	"sync"

	"github.com/dkyanakiev/vaulty/vault"
	"github.com/hashicorp/vault/api"
)

type FakeSys struct {
	GetPolicyStub        func(string) (string, error)
	getPolicyMutex       sync.RWMutex
	getPolicyArgsForCall []struct {
		arg1 string
	}
	getPolicyReturns struct {
		result1 string
		result2 error
	}
	getPolicyReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	HealthStub        func() (*api.HealthResponse, error)
	healthMutex       sync.RWMutex
	healthArgsForCall []struct {
	}
	healthReturns struct {
		result1 *api.HealthResponse
		result2 error
	}
	healthReturnsOnCall map[int]struct {
		result1 *api.HealthResponse
		result2 error
	}
	ListMountsStub        func() (map[string]*api.MountOutput, error)
	listMountsMutex       sync.RWMutex
	listMountsArgsForCall []struct {
	}
	listMountsReturns struct {
		result1 map[string]*api.MountOutput
		result2 error
	}
	listMountsReturnsOnCall map[int]struct {
		result1 map[string]*api.MountOutput
		result2 error
	}
	ListPoliciesStub        func() ([]string, error)
	listPoliciesMutex       sync.RWMutex
	listPoliciesArgsForCall []struct {
	}
	listPoliciesReturns struct {
		result1 []string
		result2 error
	}
	listPoliciesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSys) GetPolicy(arg1 string) (string, error) {
	fake.getPolicyMutex.Lock()
	ret, specificReturn := fake.getPolicyReturnsOnCall[len(fake.getPolicyArgsForCall)]
	fake.getPolicyArgsForCall = append(fake.getPolicyArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetPolicyStub
	fakeReturns := fake.getPolicyReturns
	fake.recordInvocation("GetPolicy", []interface{}{arg1})
	fake.getPolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSys) GetPolicyCallCount() int {
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	return len(fake.getPolicyArgsForCall)
}

func (fake *FakeSys) GetPolicyCalls(stub func(string) (string, error)) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = stub
}

func (fake *FakeSys) GetPolicyArgsForCall(i int) string {
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	argsForCall := fake.getPolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSys) GetPolicyReturns(result1 string, result2 error) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = nil
	fake.getPolicyReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSys) GetPolicyReturnsOnCall(i int, result1 string, result2 error) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = nil
	if fake.getPolicyReturnsOnCall == nil {
		fake.getPolicyReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getPolicyReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSys) Health() (*api.HealthResponse, error) {
	fake.healthMutex.Lock()
	ret, specificReturn := fake.healthReturnsOnCall[len(fake.healthArgsForCall)]
	fake.healthArgsForCall = append(fake.healthArgsForCall, struct {
	}{})
	stub := fake.HealthStub
	fakeReturns := fake.healthReturns
	fake.recordInvocation("Health", []interface{}{})
	fake.healthMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSys) HealthCallCount() int {
	fake.healthMutex.RLock()
	defer fake.healthMutex.RUnlock()
	return len(fake.healthArgsForCall)
}

func (fake *FakeSys) HealthCalls(stub func() (*api.HealthResponse, error)) {
	fake.healthMutex.Lock()
	defer fake.healthMutex.Unlock()
	fake.HealthStub = stub
}

func (fake *FakeSys) HealthReturns(result1 *api.HealthResponse, result2 error) {
	fake.healthMutex.Lock()
	defer fake.healthMutex.Unlock()
	fake.HealthStub = nil
	fake.healthReturns = struct {
		result1 *api.HealthResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSys) HealthReturnsOnCall(i int, result1 *api.HealthResponse, result2 error) {
	fake.healthMutex.Lock()
	defer fake.healthMutex.Unlock()
	fake.HealthStub = nil
	if fake.healthReturnsOnCall == nil {
		fake.healthReturnsOnCall = make(map[int]struct {
			result1 *api.HealthResponse
			result2 error
		})
	}
	fake.healthReturnsOnCall[i] = struct {
		result1 *api.HealthResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSys) ListMounts() (map[string]*api.MountOutput, error) {
	fake.listMountsMutex.Lock()
	ret, specificReturn := fake.listMountsReturnsOnCall[len(fake.listMountsArgsForCall)]
	fake.listMountsArgsForCall = append(fake.listMountsArgsForCall, struct {
	}{})
	stub := fake.ListMountsStub
	fakeReturns := fake.listMountsReturns
	fake.recordInvocation("ListMounts", []interface{}{})
	fake.listMountsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSys) ListMountsCallCount() int {
	fake.listMountsMutex.RLock()
	defer fake.listMountsMutex.RUnlock()
	return len(fake.listMountsArgsForCall)
}

func (fake *FakeSys) ListMountsCalls(stub func() (map[string]*api.MountOutput, error)) {
	fake.listMountsMutex.Lock()
	defer fake.listMountsMutex.Unlock()
	fake.ListMountsStub = stub
}

func (fake *FakeSys) ListMountsReturns(result1 map[string]*api.MountOutput, result2 error) {
	fake.listMountsMutex.Lock()
	defer fake.listMountsMutex.Unlock()
	fake.ListMountsStub = nil
	fake.listMountsReturns = struct {
		result1 map[string]*api.MountOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeSys) ListMountsReturnsOnCall(i int, result1 map[string]*api.MountOutput, result2 error) {
	fake.listMountsMutex.Lock()
	defer fake.listMountsMutex.Unlock()
	fake.ListMountsStub = nil
	if fake.listMountsReturnsOnCall == nil {
		fake.listMountsReturnsOnCall = make(map[int]struct {
			result1 map[string]*api.MountOutput
			result2 error
		})
	}
	fake.listMountsReturnsOnCall[i] = struct {
		result1 map[string]*api.MountOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeSys) ListPolicies() ([]string, error) {
	fake.listPoliciesMutex.Lock()
	ret, specificReturn := fake.listPoliciesReturnsOnCall[len(fake.listPoliciesArgsForCall)]
	fake.listPoliciesArgsForCall = append(fake.listPoliciesArgsForCall, struct {
	}{})
	stub := fake.ListPoliciesStub
	fakeReturns := fake.listPoliciesReturns
	fake.recordInvocation("ListPolicies", []interface{}{})
	fake.listPoliciesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSys) ListPoliciesCallCount() int {
	fake.listPoliciesMutex.RLock()
	defer fake.listPoliciesMutex.RUnlock()
	return len(fake.listPoliciesArgsForCall)
}

func (fake *FakeSys) ListPoliciesCalls(stub func() ([]string, error)) {
	fake.listPoliciesMutex.Lock()
	defer fake.listPoliciesMutex.Unlock()
	fake.ListPoliciesStub = stub
}

func (fake *FakeSys) ListPoliciesReturns(result1 []string, result2 error) {
	fake.listPoliciesMutex.Lock()
	defer fake.listPoliciesMutex.Unlock()
	fake.ListPoliciesStub = nil
	fake.listPoliciesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeSys) ListPoliciesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listPoliciesMutex.Lock()
	defer fake.listPoliciesMutex.Unlock()
	fake.ListPoliciesStub = nil
	if fake.listPoliciesReturnsOnCall == nil {
		fake.listPoliciesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listPoliciesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeSys) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	fake.healthMutex.RLock()
	defer fake.healthMutex.RUnlock()
	fake.listMountsMutex.RLock()
	defer fake.listMountsMutex.RUnlock()
	fake.listPoliciesMutex.RLock()
	defer fake.listPoliciesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSys) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ vault.Sys = new(FakeSys)

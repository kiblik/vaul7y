// Code generated by counterfeiter. DO NOT EDIT.
package vaultfakes

import (
	"sync"

	"github.com/dkyanakiev/vaulty/vault"
)

type FakeClient struct {
	AddressStub        func() string
	addressMutex       sync.RWMutex
	addressArgsForCall []struct {
	}
	addressReturns struct {
		result1 string
	}
	addressReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Address() string {
	fake.addressMutex.Lock()
	ret, specificReturn := fake.addressReturnsOnCall[len(fake.addressArgsForCall)]
	fake.addressArgsForCall = append(fake.addressArgsForCall, struct {
	}{})
	stub := fake.AddressStub
	fakeReturns := fake.addressReturns
	fake.recordInvocation("Address", []interface{}{})
	fake.addressMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) AddressCallCount() int {
	fake.addressMutex.RLock()
	defer fake.addressMutex.RUnlock()
	return len(fake.addressArgsForCall)
}

func (fake *FakeClient) AddressCalls(stub func() string) {
	fake.addressMutex.Lock()
	defer fake.addressMutex.Unlock()
	fake.AddressStub = stub
}

func (fake *FakeClient) AddressReturns(result1 string) {
	fake.addressMutex.Lock()
	defer fake.addressMutex.Unlock()
	fake.AddressStub = nil
	fake.addressReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeClient) AddressReturnsOnCall(i int, result1 string) {
	fake.addressMutex.Lock()
	defer fake.addressMutex.Unlock()
	fake.AddressStub = nil
	if fake.addressReturnsOnCall == nil {
		fake.addressReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.addressReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addressMutex.RLock()
	defer fake.addressMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ vault.Client = new(FakeClient)

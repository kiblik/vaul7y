// Code generated by counterfeiter. DO NOT EDIT.
package vaultfakes

import (
	"context"
	"sync"

	"github.com/dkyanakiev/vaulty/vault"
	"github.com/hashicorp/vault/api"
)

type FakeKV2 struct {
	GetStub        func(context.Context, string) (*api.KVSecret, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getReturns struct {
		result1 *api.KVSecret
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *api.KVSecret
		result2 error
	}
	GetMetadataStub        func(context.Context, string) (*api.KVMetadata, error)
	getMetadataMutex       sync.RWMutex
	getMetadataArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getMetadataReturns struct {
		result1 *api.KVMetadata
		result2 error
	}
	getMetadataReturnsOnCall map[int]struct {
		result1 *api.KVMetadata
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKV2) Get(arg1 context.Context, arg2 string) (*api.KVSecret, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKV2) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeKV2) GetCalls(stub func(context.Context, string) (*api.KVSecret, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeKV2) GetArgsForCall(i int) (context.Context, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKV2) GetReturns(result1 *api.KVSecret, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *api.KVSecret
		result2 error
	}{result1, result2}
}

func (fake *FakeKV2) GetReturnsOnCall(i int, result1 *api.KVSecret, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *api.KVSecret
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *api.KVSecret
		result2 error
	}{result1, result2}
}

func (fake *FakeKV2) GetMetadata(arg1 context.Context, arg2 string) (*api.KVMetadata, error) {
	fake.getMetadataMutex.Lock()
	ret, specificReturn := fake.getMetadataReturnsOnCall[len(fake.getMetadataArgsForCall)]
	fake.getMetadataArgsForCall = append(fake.getMetadataArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetMetadataStub
	fakeReturns := fake.getMetadataReturns
	fake.recordInvocation("GetMetadata", []interface{}{arg1, arg2})
	fake.getMetadataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKV2) GetMetadataCallCount() int {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return len(fake.getMetadataArgsForCall)
}

func (fake *FakeKV2) GetMetadataCalls(stub func(context.Context, string) (*api.KVMetadata, error)) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = stub
}

func (fake *FakeKV2) GetMetadataArgsForCall(i int) (context.Context, string) {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	argsForCall := fake.getMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKV2) GetMetadataReturns(result1 *api.KVMetadata, result2 error) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = nil
	fake.getMetadataReturns = struct {
		result1 *api.KVMetadata
		result2 error
	}{result1, result2}
}

func (fake *FakeKV2) GetMetadataReturnsOnCall(i int, result1 *api.KVMetadata, result2 error) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = nil
	if fake.getMetadataReturnsOnCall == nil {
		fake.getMetadataReturnsOnCall = make(map[int]struct {
			result1 *api.KVMetadata
			result2 error
		})
	}
	fake.getMetadataReturnsOnCall[i] = struct {
		result1 *api.KVMetadata
		result2 error
	}{result1, result2}
}

func (fake *FakeKV2) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKV2) recordInvocation(key string, args []interface{}) {
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

var _ vault.KV2 = new(FakeKV2)

// Code generated by counterfeiter. DO NOT EDIT.
package componentfakes

import (
	"sync"

	"github.com/dkyanakiev/vaulty/component"
	"github.com/rivo/tview"
)

type FakeDropDown struct {
	PrimitiveStub        func() tview.Primitive
	primitiveMutex       sync.RWMutex
	primitiveArgsForCall []struct {
	}
	primitiveReturns struct {
		result1 tview.Primitive
	}
	primitiveReturnsOnCall map[int]struct {
		result1 tview.Primitive
	}
	SetCurrentOptionStub        func(int)
	setCurrentOptionMutex       sync.RWMutex
	setCurrentOptionArgsForCall []struct {
		arg1 int
	}
	SetOptionsStub        func([]string, func(text string, index int))
	setOptionsMutex       sync.RWMutex
	setOptionsArgsForCall []struct {
		arg1 []string
		arg2 func(text string, index int)
	}
	SetSelectedFuncStub        func(func(text string, index int))
	setSelectedFuncMutex       sync.RWMutex
	setSelectedFuncArgsForCall []struct {
		arg1 func(text string, index int)
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDropDown) Primitive() tview.Primitive {
	fake.primitiveMutex.Lock()
	ret, specificReturn := fake.primitiveReturnsOnCall[len(fake.primitiveArgsForCall)]
	fake.primitiveArgsForCall = append(fake.primitiveArgsForCall, struct {
	}{})
	stub := fake.PrimitiveStub
	fakeReturns := fake.primitiveReturns
	fake.recordInvocation("Primitive", []interface{}{})
	fake.primitiveMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDropDown) PrimitiveCallCount() int {
	fake.primitiveMutex.RLock()
	defer fake.primitiveMutex.RUnlock()
	return len(fake.primitiveArgsForCall)
}

func (fake *FakeDropDown) PrimitiveCalls(stub func() tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = stub
}

func (fake *FakeDropDown) PrimitiveReturns(result1 tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = nil
	fake.primitiveReturns = struct {
		result1 tview.Primitive
	}{result1}
}

func (fake *FakeDropDown) PrimitiveReturnsOnCall(i int, result1 tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = nil
	if fake.primitiveReturnsOnCall == nil {
		fake.primitiveReturnsOnCall = make(map[int]struct {
			result1 tview.Primitive
		})
	}
	fake.primitiveReturnsOnCall[i] = struct {
		result1 tview.Primitive
	}{result1}
}

func (fake *FakeDropDown) SetCurrentOption(arg1 int) {
	fake.setCurrentOptionMutex.Lock()
	fake.setCurrentOptionArgsForCall = append(fake.setCurrentOptionArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.SetCurrentOptionStub
	fake.recordInvocation("SetCurrentOption", []interface{}{arg1})
	fake.setCurrentOptionMutex.Unlock()
	if stub != nil {
		fake.SetCurrentOptionStub(arg1)
	}
}

func (fake *FakeDropDown) SetCurrentOptionCallCount() int {
	fake.setCurrentOptionMutex.RLock()
	defer fake.setCurrentOptionMutex.RUnlock()
	return len(fake.setCurrentOptionArgsForCall)
}

func (fake *FakeDropDown) SetCurrentOptionCalls(stub func(int)) {
	fake.setCurrentOptionMutex.Lock()
	defer fake.setCurrentOptionMutex.Unlock()
	fake.SetCurrentOptionStub = stub
}

func (fake *FakeDropDown) SetCurrentOptionArgsForCall(i int) int {
	fake.setCurrentOptionMutex.RLock()
	defer fake.setCurrentOptionMutex.RUnlock()
	argsForCall := fake.setCurrentOptionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDropDown) SetOptions(arg1 []string, arg2 func(text string, index int)) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setOptionsMutex.Lock()
	fake.setOptionsArgsForCall = append(fake.setOptionsArgsForCall, struct {
		arg1 []string
		arg2 func(text string, index int)
	}{arg1Copy, arg2})
	stub := fake.SetOptionsStub
	fake.recordInvocation("SetOptions", []interface{}{arg1Copy, arg2})
	fake.setOptionsMutex.Unlock()
	if stub != nil {
		fake.SetOptionsStub(arg1, arg2)
	}
}

func (fake *FakeDropDown) SetOptionsCallCount() int {
	fake.setOptionsMutex.RLock()
	defer fake.setOptionsMutex.RUnlock()
	return len(fake.setOptionsArgsForCall)
}

func (fake *FakeDropDown) SetOptionsCalls(stub func([]string, func(text string, index int))) {
	fake.setOptionsMutex.Lock()
	defer fake.setOptionsMutex.Unlock()
	fake.SetOptionsStub = stub
}

func (fake *FakeDropDown) SetOptionsArgsForCall(i int) ([]string, func(text string, index int)) {
	fake.setOptionsMutex.RLock()
	defer fake.setOptionsMutex.RUnlock()
	argsForCall := fake.setOptionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDropDown) SetSelectedFunc(arg1 func(text string, index int)) {
	fake.setSelectedFuncMutex.Lock()
	fake.setSelectedFuncArgsForCall = append(fake.setSelectedFuncArgsForCall, struct {
		arg1 func(text string, index int)
	}{arg1})
	stub := fake.SetSelectedFuncStub
	fake.recordInvocation("SetSelectedFunc", []interface{}{arg1})
	fake.setSelectedFuncMutex.Unlock()
	if stub != nil {
		fake.SetSelectedFuncStub(arg1)
	}
}

func (fake *FakeDropDown) SetSelectedFuncCallCount() int {
	fake.setSelectedFuncMutex.RLock()
	defer fake.setSelectedFuncMutex.RUnlock()
	return len(fake.setSelectedFuncArgsForCall)
}

func (fake *FakeDropDown) SetSelectedFuncCalls(stub func(func(text string, index int))) {
	fake.setSelectedFuncMutex.Lock()
	defer fake.setSelectedFuncMutex.Unlock()
	fake.SetSelectedFuncStub = stub
}

func (fake *FakeDropDown) SetSelectedFuncArgsForCall(i int) func(text string, index int) {
	fake.setSelectedFuncMutex.RLock()
	defer fake.setSelectedFuncMutex.RUnlock()
	argsForCall := fake.setSelectedFuncArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDropDown) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.primitiveMutex.RLock()
	defer fake.primitiveMutex.RUnlock()
	fake.setCurrentOptionMutex.RLock()
	defer fake.setCurrentOptionMutex.RUnlock()
	fake.setOptionsMutex.RLock()
	defer fake.setOptionsMutex.RUnlock()
	fake.setSelectedFuncMutex.RLock()
	defer fake.setSelectedFuncMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDropDown) recordInvocation(key string, args []interface{}) {
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

var _ component.DropDown = new(FakeDropDown)

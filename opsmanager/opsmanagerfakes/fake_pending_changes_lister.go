// Code generated by counterfeiter. DO NOT EDIT.
package opsmanagerfakes

import (
	"sync"

	"github.com/pivotal-cf/aqueduct-courier/opsmanager"
	"github.com/pivotal-cf/om/api"
)

type FakePendingChangesLister struct {
	ListStagedPendingChangesStub        func() (api.PendingChangesOutput, error)
	listStagedPendingChangesMutex       sync.RWMutex
	listStagedPendingChangesArgsForCall []struct{}
	listStagedPendingChangesReturns     struct {
		result1 api.PendingChangesOutput
		result2 error
	}
	listStagedPendingChangesReturnsOnCall map[int]struct {
		result1 api.PendingChangesOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePendingChangesLister) ListStagedPendingChanges() (api.PendingChangesOutput, error) {
	fake.listStagedPendingChangesMutex.Lock()
	ret, specificReturn := fake.listStagedPendingChangesReturnsOnCall[len(fake.listStagedPendingChangesArgsForCall)]
	fake.listStagedPendingChangesArgsForCall = append(fake.listStagedPendingChangesArgsForCall, struct{}{})
	fake.recordInvocation("ListStagedPendingChanges", []interface{}{})
	fake.listStagedPendingChangesMutex.Unlock()
	if fake.ListStagedPendingChangesStub != nil {
		return fake.ListStagedPendingChangesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listStagedPendingChangesReturns.result1, fake.listStagedPendingChangesReturns.result2
}

func (fake *FakePendingChangesLister) ListStagedPendingChangesCallCount() int {
	fake.listStagedPendingChangesMutex.RLock()
	defer fake.listStagedPendingChangesMutex.RUnlock()
	return len(fake.listStagedPendingChangesArgsForCall)
}

func (fake *FakePendingChangesLister) ListStagedPendingChangesReturns(result1 api.PendingChangesOutput, result2 error) {
	fake.ListStagedPendingChangesStub = nil
	fake.listStagedPendingChangesReturns = struct {
		result1 api.PendingChangesOutput
		result2 error
	}{result1, result2}
}

func (fake *FakePendingChangesLister) ListStagedPendingChangesReturnsOnCall(i int, result1 api.PendingChangesOutput, result2 error) {
	fake.ListStagedPendingChangesStub = nil
	if fake.listStagedPendingChangesReturnsOnCall == nil {
		fake.listStagedPendingChangesReturnsOnCall = make(map[int]struct {
			result1 api.PendingChangesOutput
			result2 error
		})
	}
	fake.listStagedPendingChangesReturnsOnCall[i] = struct {
		result1 api.PendingChangesOutput
		result2 error
	}{result1, result2}
}

func (fake *FakePendingChangesLister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listStagedPendingChangesMutex.RLock()
	defer fake.listStagedPendingChangesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePendingChangesLister) recordInvocation(key string, args []interface{}) {
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

var _ opsmanager.PendingChangesLister = new(FakePendingChangesLister)
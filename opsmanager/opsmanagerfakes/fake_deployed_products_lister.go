// Code generated by counterfeiter. DO NOT EDIT.
package opsmanagerfakes

import (
	"sync"

	"github.com/pivotal-cf/aqueduct-courier/opsmanager"
	"github.com/pivotal-cf/om/api"
)

type FakeDeployedProductsLister struct {
	ListDeployedProductsStub        func() ([]api.DeployedProductOutput, error)
	listDeployedProductsMutex       sync.RWMutex
	listDeployedProductsArgsForCall []struct{}
	listDeployedProductsReturns     struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	listDeployedProductsReturnsOnCall map[int]struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeployedProductsLister) ListDeployedProducts() ([]api.DeployedProductOutput, error) {
	fake.listDeployedProductsMutex.Lock()
	ret, specificReturn := fake.listDeployedProductsReturnsOnCall[len(fake.listDeployedProductsArgsForCall)]
	fake.listDeployedProductsArgsForCall = append(fake.listDeployedProductsArgsForCall, struct{}{})
	fake.recordInvocation("ListDeployedProducts", []interface{}{})
	fake.listDeployedProductsMutex.Unlock()
	if fake.ListDeployedProductsStub != nil {
		return fake.ListDeployedProductsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDeployedProductsReturns.result1, fake.listDeployedProductsReturns.result2
}

func (fake *FakeDeployedProductsLister) ListDeployedProductsCallCount() int {
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	return len(fake.listDeployedProductsArgsForCall)
}

func (fake *FakeDeployedProductsLister) ListDeployedProductsReturns(result1 []api.DeployedProductOutput, result2 error) {
	fake.ListDeployedProductsStub = nil
	fake.listDeployedProductsReturns = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployedProductsLister) ListDeployedProductsReturnsOnCall(i int, result1 []api.DeployedProductOutput, result2 error) {
	fake.ListDeployedProductsStub = nil
	if fake.listDeployedProductsReturnsOnCall == nil {
		fake.listDeployedProductsReturnsOnCall = make(map[int]struct {
			result1 []api.DeployedProductOutput
			result2 error
		})
	}
	fake.listDeployedProductsReturnsOnCall[i] = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployedProductsLister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeployedProductsLister) recordInvocation(key string, args []interface{}) {
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

var _ opsmanager.DeployedProductsLister = new(FakeDeployedProductsLister)
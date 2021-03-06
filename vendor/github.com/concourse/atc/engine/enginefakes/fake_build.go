// Code generated by counterfeiter. DO NOT EDIT.
package enginefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/engine"
)

type FakeBuild struct {
	MetadataStub        func() string
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct{}
	metadataReturns     struct {
		result1 string
	}
	metadataReturnsOnCall map[int]struct {
		result1 string
	}
	AbortStub        func(lager.Logger) error
	abortMutex       sync.RWMutex
	abortArgsForCall []struct {
		arg1 lager.Logger
	}
	abortReturns struct {
		result1 error
	}
	abortReturnsOnCall map[int]struct {
		result1 error
	}
	ResumeStub        func(lager.Logger)
	resumeMutex       sync.RWMutex
	resumeArgsForCall []struct {
		arg1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuild) Metadata() string {
	fake.metadataMutex.Lock()
	ret, specificReturn := fake.metadataReturnsOnCall[len(fake.metadataArgsForCall)]
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct{}{})
	fake.recordInvocation("Metadata", []interface{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.metadataReturns.result1
}

func (fake *FakeBuild) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeBuild) MetadataReturns(result1 string) {
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) MetadataReturnsOnCall(i int, result1 string) {
	fake.MetadataStub = nil
	if fake.metadataReturnsOnCall == nil {
		fake.metadataReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.metadataReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) Abort(arg1 lager.Logger) error {
	fake.abortMutex.Lock()
	ret, specificReturn := fake.abortReturnsOnCall[len(fake.abortArgsForCall)]
	fake.abortArgsForCall = append(fake.abortArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Abort", []interface{}{arg1})
	fake.abortMutex.Unlock()
	if fake.AbortStub != nil {
		return fake.AbortStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.abortReturns.result1
}

func (fake *FakeBuild) AbortCallCount() int {
	fake.abortMutex.RLock()
	defer fake.abortMutex.RUnlock()
	return len(fake.abortArgsForCall)
}

func (fake *FakeBuild) AbortArgsForCall(i int) lager.Logger {
	fake.abortMutex.RLock()
	defer fake.abortMutex.RUnlock()
	return fake.abortArgsForCall[i].arg1
}

func (fake *FakeBuild) AbortReturns(result1 error) {
	fake.AbortStub = nil
	fake.abortReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) AbortReturnsOnCall(i int, result1 error) {
	fake.AbortStub = nil
	if fake.abortReturnsOnCall == nil {
		fake.abortReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.abortReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) Resume(arg1 lager.Logger) {
	fake.resumeMutex.Lock()
	fake.resumeArgsForCall = append(fake.resumeArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Resume", []interface{}{arg1})
	fake.resumeMutex.Unlock()
	if fake.ResumeStub != nil {
		fake.ResumeStub(arg1)
	}
}

func (fake *FakeBuild) ResumeCallCount() int {
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	return len(fake.resumeArgsForCall)
}

func (fake *FakeBuild) ResumeArgsForCall(i int) lager.Logger {
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	return fake.resumeArgsForCall[i].arg1
}

func (fake *FakeBuild) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	fake.abortMutex.RLock()
	defer fake.abortMutex.RUnlock()
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuild) recordInvocation(key string, args []interface{}) {
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

var _ engine.Build = new(FakeBuild)

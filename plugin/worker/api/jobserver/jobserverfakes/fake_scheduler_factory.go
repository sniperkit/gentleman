// Code generated by counterfeiter. DO NOT EDIT.
package jobserverfakes

import (
	"sync"

	"github.com/concourse/atc/api/jobserver"
	"github.com/concourse/atc/creds"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/scheduler"
)

type FakeSchedulerFactory struct {
	BuildSchedulerStub        func(db.Pipeline, string, creds.Variables) scheduler.BuildScheduler
	buildSchedulerMutex       sync.RWMutex
	buildSchedulerArgsForCall []struct {
		arg1 db.Pipeline
		arg2 string
		arg3 creds.Variables
	}
	buildSchedulerReturns struct {
		result1 scheduler.BuildScheduler
	}
	buildSchedulerReturnsOnCall map[int]struct {
		result1 scheduler.BuildScheduler
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSchedulerFactory) BuildScheduler(arg1 db.Pipeline, arg2 string, arg3 creds.Variables) scheduler.BuildScheduler {
	fake.buildSchedulerMutex.Lock()
	ret, specificReturn := fake.buildSchedulerReturnsOnCall[len(fake.buildSchedulerArgsForCall)]
	fake.buildSchedulerArgsForCall = append(fake.buildSchedulerArgsForCall, struct {
		arg1 db.Pipeline
		arg2 string
		arg3 creds.Variables
	}{arg1, arg2, arg3})
	fake.recordInvocation("BuildScheduler", []interface{}{arg1, arg2, arg3})
	fake.buildSchedulerMutex.Unlock()
	if fake.BuildSchedulerStub != nil {
		return fake.BuildSchedulerStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.buildSchedulerReturns.result1
}

func (fake *FakeSchedulerFactory) BuildSchedulerCallCount() int {
	fake.buildSchedulerMutex.RLock()
	defer fake.buildSchedulerMutex.RUnlock()
	return len(fake.buildSchedulerArgsForCall)
}

func (fake *FakeSchedulerFactory) BuildSchedulerArgsForCall(i int) (db.Pipeline, string, creds.Variables) {
	fake.buildSchedulerMutex.RLock()
	defer fake.buildSchedulerMutex.RUnlock()
	return fake.buildSchedulerArgsForCall[i].arg1, fake.buildSchedulerArgsForCall[i].arg2, fake.buildSchedulerArgsForCall[i].arg3
}

func (fake *FakeSchedulerFactory) BuildSchedulerReturns(result1 scheduler.BuildScheduler) {
	fake.BuildSchedulerStub = nil
	fake.buildSchedulerReturns = struct {
		result1 scheduler.BuildScheduler
	}{result1}
}

func (fake *FakeSchedulerFactory) BuildSchedulerReturnsOnCall(i int, result1 scheduler.BuildScheduler) {
	fake.BuildSchedulerStub = nil
	if fake.buildSchedulerReturnsOnCall == nil {
		fake.buildSchedulerReturnsOnCall = make(map[int]struct {
			result1 scheduler.BuildScheduler
		})
	}
	fake.buildSchedulerReturnsOnCall[i] = struct {
		result1 scheduler.BuildScheduler
	}{result1}
}

func (fake *FakeSchedulerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildSchedulerMutex.RLock()
	defer fake.buildSchedulerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSchedulerFactory) recordInvocation(key string, args []interface{}) {
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

var _ jobserver.SchedulerFactory = new(FakeSchedulerFactory)

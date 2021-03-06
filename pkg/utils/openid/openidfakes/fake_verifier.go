// Code generated by counterfeiter. DO NOT EDIT.
package openidfakes

import (
	"context"
	"sync"

	"github.com/appvia/kore/pkg/utils/openid"
)

type FakeVerifier struct {
	VerifyStub        func(context.Context, string) (openid.IDToken, error)
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	verifyReturns struct {
		result1 openid.IDToken
		result2 error
	}
	verifyReturnsOnCall map[int]struct {
		result1 openid.IDToken
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVerifier) Verify(arg1 context.Context, arg2 string) (openid.IDToken, error) {
	fake.verifyMutex.Lock()
	ret, specificReturn := fake.verifyReturnsOnCall[len(fake.verifyArgsForCall)]
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Verify", []interface{}{arg1, arg2})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.verifyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVerifier) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeVerifier) VerifyCalls(stub func(context.Context, string) (openid.IDToken, error)) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = stub
}

func (fake *FakeVerifier) VerifyArgsForCall(i int) (context.Context, string) {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	argsForCall := fake.verifyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVerifier) VerifyReturns(result1 openid.IDToken, result2 error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 openid.IDToken
		result2 error
	}{result1, result2}
}

func (fake *FakeVerifier) VerifyReturnsOnCall(i int, result1 openid.IDToken, result2 error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	if fake.verifyReturnsOnCall == nil {
		fake.verifyReturnsOnCall = make(map[int]struct {
			result1 openid.IDToken
			result2 error
		})
	}
	fake.verifyReturnsOnCall[i] = struct {
		result1 openid.IDToken
		result2 error
	}{result1, result2}
}

func (fake *FakeVerifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVerifier) recordInvocation(key string, args []interface{}) {
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

var _ openid.Verifier = new(FakeVerifier)

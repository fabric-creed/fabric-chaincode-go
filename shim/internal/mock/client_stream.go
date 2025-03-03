// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/fabric-creed/fabric-protos-go/peer"
)

type ClientStream struct {
	CloseSendStub        func() error
	closeSendMutex       sync.RWMutex
	closeSendArgsForCall []struct {
	}
	closeSendReturns struct {
		result1 error
	}
	closeSendReturnsOnCall map[int]struct {
		result1 error
	}
	RecvStub        func() (*peer.ChaincodeMessage, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct {
	}
	recvReturns struct {
		result1 *peer.ChaincodeMessage
		result2 error
	}
	recvReturnsOnCall map[int]struct {
		result1 *peer.ChaincodeMessage
		result2 error
	}
	SendStub        func(*peer.ChaincodeMessage) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 *peer.ChaincodeMessage
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ClientStream) CloseSend() error {
	fake.closeSendMutex.Lock()
	ret, specificReturn := fake.closeSendReturnsOnCall[len(fake.closeSendArgsForCall)]
	fake.closeSendArgsForCall = append(fake.closeSendArgsForCall, struct {
	}{})
	fake.recordInvocation("CloseSend", []interface{}{})
	fake.closeSendMutex.Unlock()
	if fake.CloseSendStub != nil {
		return fake.CloseSendStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeSendReturns
	return fakeReturns.result1
}

func (fake *ClientStream) CloseSendCallCount() int {
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	return len(fake.closeSendArgsForCall)
}

func (fake *ClientStream) CloseSendCalls(stub func() error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = stub
}

func (fake *ClientStream) CloseSendReturns(result1 error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = nil
	fake.closeSendReturns = struct {
		result1 error
	}{result1}
}

func (fake *ClientStream) CloseSendReturnsOnCall(i int, result1 error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = nil
	if fake.closeSendReturnsOnCall == nil {
		fake.closeSendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeSendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ClientStream) Recv() (*peer.ChaincodeMessage, error) {
	fake.recvMutex.Lock()
	ret, specificReturn := fake.recvReturnsOnCall[len(fake.recvArgsForCall)]
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct {
	}{})
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if fake.RecvStub != nil {
		return fake.RecvStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.recvReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ClientStream) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *ClientStream) RecvCalls(stub func() (*peer.ChaincodeMessage, error)) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = stub
}

func (fake *ClientStream) RecvReturns(result1 *peer.ChaincodeMessage, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *peer.ChaincodeMessage
		result2 error
	}{result1, result2}
}

func (fake *ClientStream) RecvReturnsOnCall(i int, result1 *peer.ChaincodeMessage, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	if fake.recvReturnsOnCall == nil {
		fake.recvReturnsOnCall = make(map[int]struct {
			result1 *peer.ChaincodeMessage
			result2 error
		})
	}
	fake.recvReturnsOnCall[i] = struct {
		result1 *peer.ChaincodeMessage
		result2 error
	}{result1, result2}
}

func (fake *ClientStream) Send(arg1 *peer.ChaincodeMessage) error {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 *peer.ChaincodeMessage
	}{arg1})
	fake.recordInvocation("Send", []interface{}{arg1})
	fake.sendMutex.Unlock()
	if fake.SendStub != nil {
		return fake.SendStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendReturns
	return fakeReturns.result1
}

func (fake *ClientStream) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *ClientStream) SendCalls(stub func(*peer.ChaincodeMessage) error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = stub
}

func (fake *ClientStream) SendArgsForCall(i int) *peer.ChaincodeMessage {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	argsForCall := fake.sendArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ClientStream) SendReturns(result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *ClientStream) SendReturnsOnCall(i int, result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ClientStream) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ClientStream) recordInvocation(key string, args []interface{}) {
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

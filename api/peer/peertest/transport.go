// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: go.uber.org/yarpc/api/peer (interfaces: Transport,Subscriber)

package peertest

import (
	gomock "github.com/golang/mock/gomock"
	peer "go.uber.org/yarpc/api/peer"
)

// Mock of Transport interface
type MockTransport struct {
	ctrl     *gomock.Controller
	recorder *_MockTransportRecorder
}

// Recorder for MockTransport (not exported)
type _MockTransportRecorder struct {
	mock *MockTransport
}

func NewMockTransport(ctrl *gomock.Controller) *MockTransport {
	mock := &MockTransport{ctrl: ctrl}
	mock.recorder = &_MockTransportRecorder{mock}
	return mock
}

func (_m *MockTransport) EXPECT() *_MockTransportRecorder {
	return _m.recorder
}

func (_m *MockTransport) ReleasePeer(_param0 peer.Identifier, _param1 peer.Subscriber) error {
	ret := _m.ctrl.Call(_m, "ReleasePeer", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTransportRecorder) ReleasePeer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReleasePeer", arg0, arg1)
}

func (_m *MockTransport) RetainPeer(_param0 peer.Identifier, _param1 peer.Subscriber) (peer.Peer, error) {
	ret := _m.ctrl.Call(_m, "RetainPeer", _param0, _param1)
	ret0, _ := ret[0].(peer.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTransportRecorder) RetainPeer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetainPeer", arg0, arg1)
}

// Mock of Subscriber interface
type MockSubscriber struct {
	ctrl     *gomock.Controller
	recorder *_MockSubscriberRecorder
}

// Recorder for MockSubscriber (not exported)
type _MockSubscriberRecorder struct {
	mock *MockSubscriber
}

func NewMockSubscriber(ctrl *gomock.Controller) *MockSubscriber {
	mock := &MockSubscriber{ctrl: ctrl}
	mock.recorder = &_MockSubscriberRecorder{mock}
	return mock
}

func (_m *MockSubscriber) EXPECT() *_MockSubscriberRecorder {
	return _m.recorder
}

func (_m *MockSubscriber) NotifyStatusChanged(_param0 peer.Identifier) {
	_m.ctrl.Call(_m, "NotifyStatusChanged", _param0)
}

func (_mr *_MockSubscriberRecorder) NotifyStatusChanged(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NotifyStatusChanged", arg0)
}

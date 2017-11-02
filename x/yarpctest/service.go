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

package yarpctest

import (
	"fmt"
	"net"

	"github.com/stretchr/testify/assert"
	"github.com/uber-go/tally"
	"go.uber.org/multierr"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/transport/grpc"
	"go.uber.org/yarpc/transport/http"
	"go.uber.org/yarpc/transport/tchannel"
)

// Lifecycles is a wrapper around a list of Lifecycle definitions.
func Lifecycles(l ...Lifecycle) Lifecycle {
	return lifecycles(l)
}

type lifecycles []Lifecycle

// Start the lifecycles. If there are any errors, stop any started lifecycles
// and fail the test.
func (ls lifecycles) Start(t TestingT) error {
	startedLifecycles := make(lifecycles, 0, len(ls))
	for _, l := range ls {
		err := l.Start(t)
		if !assert.NoError(t, err) {
			// Cleanup started lifecycles (this could fail)
			return multierr.Append(err, startedLifecycles.Stop(t))
		}
		startedLifecycles = append(startedLifecycles, l)
	}
	return nil
}

// Stop the lifecycles. Record all errors. If any lifecycle failed to stop
// fail the test.
func (ls lifecycles) Stop(t TestingT) error {
	var err error
	for _, l := range ls {
		err = multierr.Append(err, l.Stop(t))
	}
	assert.NoError(t, err)
	return err
}

// Lifecycle defines test infra that needs to be started before the actions
// and stopped afterwards.
type Lifecycle interface {
	Start(TestingT) error
	Stop(TestingT) error
}

// ServiceOpts are the configuration options for a yarpc service.
type ServiceOpts struct {
	Name       string
	Port       int
	Procedures []transport.Procedure
}

// ServiceOption is an option when creating a Service.
type ServiceOption interface {
	ApplyService(*ServiceOpts)
}

// ServiceOptionFunc converts a function into a ServiceOption.
type ServiceOptionFunc func(*ServiceOpts)

// ApplyService implements ServiceOption.
func (f ServiceOptionFunc) ApplyService(opts *ServiceOpts) { f(opts) }

// HTTPService will create a runnable HTTP service.
func HTTPService(options ...ServiceOption) Lifecycle {
	opts := ServiceOpts{}
	for _, option := range options {
		option.ApplyService(&opts)
	}
	inbound := http.NewTransport().NewInbound(fmt.Sprintf(":%d", opts.Port))
	return createService(opts.Name, inbound, opts.Procedures)
}

// TChannelService will create a runnable TChannel service.
func TChannelService(options ...ServiceOption) Lifecycle {
	opts := ServiceOpts{}
	for _, option := range options {
		option.ApplyService(&opts)
	}
	trans, err := tchannel.NewTransport(
		tchannel.ListenAddr(fmt.Sprintf(":%d", opts.Port)),
		tchannel.ServiceName(opts.Name),
	)
	if err != nil {
		panic(err)
	}
	inbound := trans.NewInbound()
	return createService(opts.Name, inbound, opts.Procedures)
}

// GRPCService will create a runnable GRPC service.
func GRPCService(options ...ServiceOption) Lifecycle {
	opts := ServiceOpts{}
	for _, option := range options {
		option.ApplyService(&opts)
	}
	trans := grpc.NewTransport()
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.Port))
	if err != nil {
		panic(err)
	}
	inbound := trans.NewInbound(grpcListener)
	return createService(opts.Name, inbound, opts.Procedures)
}

func createService(
	name string,
	inbound transport.Inbound,
	procedures []transport.Procedure,
) *wrappedDispatcher {
	d := yarpc.NewDispatcher(
		yarpc.Config{
			Name:     name,
			Inbounds: yarpc.Inbounds{inbound},
			Metrics: yarpc.MetricsConfig{
				Tally: tally.NoopScope,
			},
		},
	)
	d.Register(procedures)
	return &wrappedDispatcher{d}
}

type wrappedDispatcher struct {
	*yarpc.Dispatcher
}

func (w *wrappedDispatcher) Start(t TestingT) error {
	err := w.Dispatcher.Start()
	assert.NoError(t, err, "error starting dispatcher: %s", w.Name())
	return err
}

func (w *wrappedDispatcher) Stop(t TestingT) error {
	err := w.Dispatcher.Stop()
	assert.NoError(t, err, "error stopping dispatcher: %s", w.Name())
	return err
}

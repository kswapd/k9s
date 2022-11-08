// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/kswapd/k12s/internal/config (interfaces: KubeSettings)

package config_test

import (
	pegomock "github.com/petergtz/pegomock"
	v1 "k8s.io/api/core/v1"
	"reflect"
	"time"
)

type MockKubeSettings struct {
	fail func(message string, callerSkip ...int)
}

func NewMockKubeSettings(options ...pegomock.Option) *MockKubeSettings {
	mock := &MockKubeSettings{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockKubeSettings) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockKubeSettings) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockKubeSettings) ClusterNames() (map[string]struct{}, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKubeSettings().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ClusterNames", params, []reflect.Type{reflect.TypeOf((*map[string]struct{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]struct{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]struct{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockKubeSettings) CurrentClusterName() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKubeSettings().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentClusterName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockKubeSettings) CurrentContextName() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKubeSettings().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentContextName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockKubeSettings) CurrentNamespaceName() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKubeSettings().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentNamespaceName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockKubeSettings) NamespaceNames(_param0 []v1.Namespace) []string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKubeSettings().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NamespaceNames", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem()})
	var ret0 []string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
	}
	return ret0
}

func (mock *MockKubeSettings) VerifyWasCalledOnce() *VerifierMockKubeSettings {
	return &VerifierMockKubeSettings{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockKubeSettings) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockKubeSettings {
	return &VerifierMockKubeSettings{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockKubeSettings) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockKubeSettings {
	return &VerifierMockKubeSettings{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockKubeSettings) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockKubeSettings {
	return &VerifierMockKubeSettings{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockKubeSettings struct {
	mock                   *MockKubeSettings
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockKubeSettings) ClusterNames() *MockKubeSettings_ClusterNames_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ClusterNames", params, verifier.timeout)
	return &MockKubeSettings_ClusterNames_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockKubeSettings_ClusterNames_OngoingVerification struct {
	mock              *MockKubeSettings
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockKubeSettings_ClusterNames_OngoingVerification) GetCapturedArguments() {
}

func (c *MockKubeSettings_ClusterNames_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockKubeSettings) CurrentClusterName() *MockKubeSettings_CurrentClusterName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentClusterName", params, verifier.timeout)
	return &MockKubeSettings_CurrentClusterName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockKubeSettings_CurrentClusterName_OngoingVerification struct {
	mock              *MockKubeSettings
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockKubeSettings_CurrentClusterName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockKubeSettings_CurrentClusterName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockKubeSettings) CurrentContextName() *MockKubeSettings_CurrentContextName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentContextName", params, verifier.timeout)
	return &MockKubeSettings_CurrentContextName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockKubeSettings_CurrentContextName_OngoingVerification struct {
	mock              *MockKubeSettings
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockKubeSettings_CurrentContextName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockKubeSettings_CurrentContextName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockKubeSettings) CurrentNamespaceName() *MockKubeSettings_CurrentNamespaceName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentNamespaceName", params, verifier.timeout)
	return &MockKubeSettings_CurrentNamespaceName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockKubeSettings_CurrentNamespaceName_OngoingVerification struct {
	mock              *MockKubeSettings
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockKubeSettings_CurrentNamespaceName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockKubeSettings_CurrentNamespaceName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockKubeSettings) NamespaceNames(_param0 []v1.Namespace) *MockKubeSettings_NamespaceNames_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NamespaceNames", params, verifier.timeout)
	return &MockKubeSettings_NamespaceNames_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockKubeSettings_NamespaceNames_OngoingVerification struct {
	mock              *MockKubeSettings
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockKubeSettings_NamespaceNames_OngoingVerification) GetCapturedArguments() []v1.Namespace {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockKubeSettings_NamespaceNames_OngoingVerification) GetAllCapturedArguments() (_param0 [][]v1.Namespace) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]v1.Namespace, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.([]v1.Namespace)
		}
	}
	return
}

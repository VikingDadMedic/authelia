// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/storage/provider.go

package storage

import (
	"context"
	"reflect"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/authelia/authelia/v4/internal/models"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// AppendAuthenticationLog mocks base method.
func (m *MockProvider) AppendAuthenticationLog(arg0 context.Context, arg1 models.AuthenticationAttempt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendAuthenticationLog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendAuthenticationLog indicates an expected call of AppendAuthenticationLog.
func (mr *MockProviderMockRecorder) AppendAuthenticationLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendAuthenticationLog", reflect.TypeOf((*MockProvider)(nil).AppendAuthenticationLog), arg0, arg1)
}

// Close mocks base method.
func (m *MockProvider) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockProviderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProvider)(nil).Close))
}

// DeleteTOTPConfiguration mocks base method.
func (m *MockProvider) DeleteTOTPConfiguration(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTOTPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTOTPConfiguration indicates an expected call of DeleteTOTPConfiguration.
func (mr *MockProviderMockRecorder) DeleteTOTPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTOTPConfiguration", reflect.TypeOf((*MockProvider)(nil).DeleteTOTPConfiguration), arg0, arg1)
}

// FindIdentityVerification mocks base method.
func (m *MockProvider) FindIdentityVerification(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIdentityVerification", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIdentityVerification indicates an expected call of FindIdentityVerification.
func (mr *MockProviderMockRecorder) FindIdentityVerification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIdentityVerification", reflect.TypeOf((*MockProvider)(nil).FindIdentityVerification), arg0, arg1)
}

// LoadAuthenticationLogs mocks base method.
func (m *MockProvider) LoadAuthenticationLogs(arg0 context.Context, arg1 string, arg2 time.Time, arg3, arg4 int) ([]models.AuthenticationAttempt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAuthenticationLogs", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]models.AuthenticationAttempt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAuthenticationLogs indicates an expected call of LoadAuthenticationLogs.
func (mr *MockProviderMockRecorder) LoadAuthenticationLogs(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAuthenticationLogs", reflect.TypeOf((*MockProvider)(nil).LoadAuthenticationLogs), arg0, arg1, arg2, arg3, arg4)
}

// LoadPreferred2FAMethod mocks base method.
func (m *MockProvider) LoadPreferred2FAMethod(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPreferred2FAMethod", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadPreferred2FAMethod indicates an expected call of LoadPreferred2FAMethod.
func (mr *MockProviderMockRecorder) LoadPreferred2FAMethod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPreferred2FAMethod", reflect.TypeOf((*MockProvider)(nil).LoadPreferred2FAMethod), arg0, arg1)
}

// LoadTOTPConfiguration mocks base method.
func (m *MockProvider) LoadTOTPConfiguration(arg0 context.Context, arg1 string) (*models.TOTPConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTOTPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*models.TOTPConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTOTPConfiguration indicates an expected call of LoadTOTPConfiguration.
func (mr *MockProviderMockRecorder) LoadTOTPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTOTPConfiguration", reflect.TypeOf((*MockProvider)(nil).LoadTOTPConfiguration), arg0, arg1)
}

// LoadTOTPConfigurations mocks base method.
func (m *MockProvider) LoadTOTPConfigurations(arg0 context.Context, arg1, arg2 int) ([]models.TOTPConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTOTPConfigurations", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.TOTPConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTOTPConfigurations indicates an expected call of LoadTOTPConfigurations.
func (mr *MockProviderMockRecorder) LoadTOTPConfigurations(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTOTPConfigurations", reflect.TypeOf((*MockProvider)(nil).LoadTOTPConfigurations), arg0, arg1, arg2)
}

// LoadU2FDevice mocks base method.
func (m *MockProvider) LoadU2FDevice(arg0 context.Context, arg1 string) (*models.U2FDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadU2FDevice", arg0, arg1)
	ret0, _ := ret[0].(*models.U2FDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadU2FDevice indicates an expected call of LoadU2FDevice.
func (mr *MockProviderMockRecorder) LoadU2FDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadU2FDevice", reflect.TypeOf((*MockProvider)(nil).LoadU2FDevice), arg0, arg1)
}

// LoadUserInfo mocks base method.
func (m *MockProvider) LoadUserInfo(arg0 context.Context, arg1 string) (models.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUserInfo", arg0, arg1)
	ret0, _ := ret[0].(models.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUserInfo indicates an expected call of LoadUserInfo.
func (mr *MockProviderMockRecorder) LoadUserInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUserInfo", reflect.TypeOf((*MockProvider)(nil).LoadUserInfo), arg0, arg1)
}

// RemoveIdentityVerification mocks base method.
func (m *MockProvider) RemoveIdentityVerification(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIdentityVerification", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIdentityVerification indicates an expected call of RemoveIdentityVerification.
func (mr *MockProviderMockRecorder) RemoveIdentityVerification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIdentityVerification", reflect.TypeOf((*MockProvider)(nil).RemoveIdentityVerification), arg0, arg1)
}

// SaveIdentityVerification mocks base method.
func (m *MockProvider) SaveIdentityVerification(arg0 context.Context, arg1 models.IdentityVerification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveIdentityVerification", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveIdentityVerification indicates an expected call of SaveIdentityVerification.
func (mr *MockProviderMockRecorder) SaveIdentityVerification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveIdentityVerification", reflect.TypeOf((*MockProvider)(nil).SaveIdentityVerification), arg0, arg1)
}

// SavePreferred2FAMethod mocks base method.
func (m *MockProvider) SavePreferred2FAMethod(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePreferred2FAMethod", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePreferred2FAMethod indicates an expected call of SavePreferred2FAMethod.
func (mr *MockProviderMockRecorder) SavePreferred2FAMethod(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePreferred2FAMethod", reflect.TypeOf((*MockProvider)(nil).SavePreferred2FAMethod), arg0, arg1, arg2)
}

// SaveTOTPConfiguration mocks base method.
func (m *MockProvider) SaveTOTPConfiguration(arg0 context.Context, arg1 models.TOTPConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTOTPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTOTPConfiguration indicates an expected call of SaveTOTPConfiguration.
func (mr *MockProviderMockRecorder) SaveTOTPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTOTPConfiguration", reflect.TypeOf((*MockProvider)(nil).SaveTOTPConfiguration), arg0, arg1)
}

// SaveU2FDevice mocks base method.
func (m *MockProvider) SaveU2FDevice(arg0 context.Context, arg1 models.U2FDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveU2FDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveU2FDevice indicates an expected call of SaveU2FDevice.
func (mr *MockProviderMockRecorder) SaveU2FDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveU2FDevice", reflect.TypeOf((*MockProvider)(nil).SaveU2FDevice), arg0, arg1)
}

// SchemaEncryptionChangeKey mocks base method.
func (m *MockProvider) SchemaEncryptionChangeKey(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaEncryptionChangeKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchemaEncryptionChangeKey indicates an expected call of SchemaEncryptionChangeKey.
func (mr *MockProviderMockRecorder) SchemaEncryptionChangeKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaEncryptionChangeKey", reflect.TypeOf((*MockProvider)(nil).SchemaEncryptionChangeKey), arg0, arg1)
}

// SchemaEncryptionCheckKey mocks base method.
func (m *MockProvider) SchemaEncryptionCheckKey(arg0 context.Context, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaEncryptionCheckKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchemaEncryptionCheckKey indicates an expected call of SchemaEncryptionCheckKey.
func (mr *MockProviderMockRecorder) SchemaEncryptionCheckKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaEncryptionCheckKey", reflect.TypeOf((*MockProvider)(nil).SchemaEncryptionCheckKey), arg0, arg1)
}

// SchemaLatestVersion mocks base method.
func (m *MockProvider) SchemaLatestVersion() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaLatestVersion")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaLatestVersion indicates an expected call of SchemaLatestVersion.
func (mr *MockProviderMockRecorder) SchemaLatestVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaLatestVersion", reflect.TypeOf((*MockProvider)(nil).SchemaLatestVersion))
}

// SchemaMigrate mocks base method.
func (m *MockProvider) SchemaMigrate(arg0 context.Context, arg1 bool, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaMigrate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchemaMigrate indicates an expected call of SchemaMigrate.
func (mr *MockProviderMockRecorder) SchemaMigrate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaMigrate", reflect.TypeOf((*MockProvider)(nil).SchemaMigrate), arg0, arg1, arg2)
}

// SchemaMigrationHistory mocks base method.
func (m *MockProvider) SchemaMigrationHistory(arg0 context.Context) ([]models.Migration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaMigrationHistory", arg0)
	ret0, _ := ret[0].([]models.Migration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaMigrationHistory indicates an expected call of SchemaMigrationHistory.
func (mr *MockProviderMockRecorder) SchemaMigrationHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaMigrationHistory", reflect.TypeOf((*MockProvider)(nil).SchemaMigrationHistory), arg0)
}

// SchemaMigrationsDown mocks base method.
func (m *MockProvider) SchemaMigrationsDown(arg0 context.Context, arg1 int) ([]SchemaMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaMigrationsDown", arg0, arg1)
	ret0, _ := ret[0].([]SchemaMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaMigrationsDown indicates an expected call of SchemaMigrationsDown.
func (mr *MockProviderMockRecorder) SchemaMigrationsDown(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaMigrationsDown", reflect.TypeOf((*MockProvider)(nil).SchemaMigrationsDown), arg0, arg1)
}

// SchemaMigrationsUp mocks base method.
func (m *MockProvider) SchemaMigrationsUp(arg0 context.Context, arg1 int) ([]SchemaMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaMigrationsUp", arg0, arg1)
	ret0, _ := ret[0].([]SchemaMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaMigrationsUp indicates an expected call of SchemaMigrationsUp.
func (mr *MockProviderMockRecorder) SchemaMigrationsUp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaMigrationsUp", reflect.TypeOf((*MockProvider)(nil).SchemaMigrationsUp), arg0, arg1)
}

// SchemaTables mocks base method.
func (m *MockProvider) SchemaTables(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaTables", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaTables indicates an expected call of SchemaTables.
func (mr *MockProviderMockRecorder) SchemaTables(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaTables", reflect.TypeOf((*MockProvider)(nil).SchemaTables), arg0)
}

// SchemaVersion mocks base method.
func (m *MockProvider) SchemaVersion(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaVersion", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaVersion indicates an expected call of SchemaVersion.
func (mr *MockProviderMockRecorder) SchemaVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaVersion", reflect.TypeOf((*MockProvider)(nil).SchemaVersion), arg0)
}

// StartupCheck mocks base method.
func (m *MockProvider) StartupCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartupCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartupCheck indicates an expected call of StartupCheck.
func (mr *MockProviderMockRecorder) StartupCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartupCheck", reflect.TypeOf((*MockProvider)(nil).StartupCheck))
}

// UpdateTOTPConfigurationSecret mocks base method.
func (m *MockProvider) UpdateTOTPConfigurationSecret(arg0 context.Context, arg1 models.TOTPConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTOTPConfigurationSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTOTPConfigurationSecret indicates an expected call of UpdateTOTPConfigurationSecret.
func (mr *MockProviderMockRecorder) UpdateTOTPConfigurationSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTOTPConfigurationSecret", reflect.TypeOf((*MockProvider)(nil).UpdateTOTPConfigurationSecret), arg0, arg1)
}
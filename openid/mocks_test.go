// Code generated by mockery v1.0.0
package openid

import http "net/http"
import mock "github.com/stretchr/testify/mock"
import jose "gopkg.in/square/go-jose.v2"
import jwt "github.com/dgrijalva/jwt-go"

// mockSigningKeyGetter is an autogenerated mock type for the signingKeyGetter type
type mockSigningKeyGetter struct {
	mock.Mock
}

// flushCachedSigningKeys provides a mock function with given fields: issuer
func (_m *mockSigningKeyGetter) flushCachedSigningKeys(issuer string) error {
	ret := _m.Called(issuer)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(issuer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// getSigningKey provides a mock function with given fields: r, issuer, kid
func (_m *mockSigningKeyGetter) getSigningKey(r *http.Request, issuer string, kid string) ([]byte, error) {
	ret := _m.Called(r, issuer, kid)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*http.Request, string, string) []byte); ok {
		r0 = rf(r, issuer, kid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request, string, string) error); ok {
		r1 = rf(r, issuer, kid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockJwksGetter is an autogenerated mock type for the jwksGetter type
type mockJwksGetter struct {
	mock.Mock
}

// getJwkSet provides a mock function with given fields: r, url
func (_m *mockJwksGetter) getJwkSet(r *http.Request, url string) (jose.JSONWebKeySet, error) {
	ret := _m.Called(r, url)

	var r0 jose.JSONWebKeySet
	if rf, ok := ret.Get(0).(func(*http.Request, string) jose.JSONWebKeySet); ok {
		r0 = rf(r, url)
	} else {
		r0 = ret.Get(0).(jose.JSONWebKeySet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request, string) error); ok {
		r1 = rf(r, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockConfigurationGetter is an autogenerated mock type for the configurationGetter type
type mockConfigurationGetter struct {
	mock.Mock
}

// getConfiguration provides a mock function with given fields: r, url
func (_m *mockConfigurationGetter) getConfiguration(r *http.Request, url string) (configuration, error) {
	ret := _m.Called(r, url)

	var r0 configuration
	if rf, ok := ret.Get(0).(func(*http.Request, string) configuration); ok {
		r0 = rf(r, url)
	} else {
		r0 = ret.Get(0).(configuration)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request, string) error); ok {
		r1 = rf(r, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockJwtTokenValidator is an autogenerated mock type for the jwtTokenValidator type
type mockJwtTokenValidator struct {
	mock.Mock
}

// validate provides a mock function with given fields: r, t
func (_m *mockJwtTokenValidator) validate(r *http.Request, t string) (*jwt.Token, error) {
	ret := _m.Called(r, t)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(*http.Request, string) *jwt.Token); ok {
		r0 = rf(r, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request, string) error); ok {
		r1 = rf(r, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockJwtParser is an autogenerated mock type for the jwtParser type
type mockJwtParser struct {
	mock.Mock
}

// parse provides a mock function with given fields: _a0, _a1
func (_m *mockJwtParser) parse(_a0 string, _a1 jwt.Keyfunc) (*jwt.Token, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(string, jwt.Keyfunc) *jwt.Token); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, jwt.Keyfunc) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

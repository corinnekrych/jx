// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	auth "github.com/jenkins-x/jx/pkg/auth"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyAuthConfigService() auth.ConfigService {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(auth.ConfigService))(nil)).Elem()))
	var nullValue auth.ConfigService
	return nullValue
}

func EqAuthConfigService(value auth.ConfigService) auth.ConfigService {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue auth.ConfigService
	return nullValue
}

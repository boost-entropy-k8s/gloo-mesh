// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/networking/v1/virtual_mesh.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *VirtualMeshSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualMeshSpec)
	if !ok {
		that2, ok := that.(VirtualMeshSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetMeshes()) != len(target.GetMeshes()) {
		return false
	}
	for idx, v := range m.GetMeshes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMeshes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMeshes()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetMtlsConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMtlsConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMtlsConfig(), target.GetMtlsConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFederation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFederation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFederation(), target.GetFederation()) {
			return false
		}
	}

	if m.GetGlobalAccessPolicy() != target.GetGlobalAccessPolicy() {
		return false
	}

	return true
}

// Equal function
func (m *VirtualMeshStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualMeshStatus)
	if !ok {
		that2, ok := that.(VirtualMeshStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetMeshes()) != len(target.GetMeshes()) {
		return false
	}
	for k, v := range m.GetMeshes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMeshes()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMeshes()[k]) {
				return false
			}
		}

	}

	if len(m.GetDestinations()) != len(target.GetDestinations()) {
		return false
	}
	for k, v := range m.GetDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDestinations()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDestinations()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualMeshSpec_MTLSConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualMeshSpec_MTLSConfig)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_MTLSConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetAutoRestartPods() != target.GetAutoRestartPods() {
		return false
	}

	switch m.TrustModel.(type) {

	case *VirtualMeshSpec_MTLSConfig_Shared:
		if _, ok := target.TrustModel.(*VirtualMeshSpec_MTLSConfig_Shared); !ok {
			return false
		}

		if h, ok := interface{}(m.GetShared()).(equality.Equalizer); ok {
			if !h.Equal(target.GetShared()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetShared(), target.GetShared()) {
				return false
			}
		}

	case *VirtualMeshSpec_MTLSConfig_Limited:
		if _, ok := target.TrustModel.(*VirtualMeshSpec_MTLSConfig_Limited); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLimited()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLimited()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLimited(), target.GetLimited()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.TrustModel != target.TrustModel {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualMeshSpec_RootCertificateAuthority) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualMeshSpec_RootCertificateAuthority)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_RootCertificateAuthority)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.CaSource.(type) {

	case *VirtualMeshSpec_RootCertificateAuthority_Generated:
		if _, ok := target.CaSource.(*VirtualMeshSpec_RootCertificateAuthority_Generated); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGenerated()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGenerated()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGenerated(), target.GetGenerated()) {
				return false
			}
		}

	case *VirtualMeshSpec_RootCertificateAuthority_Secret:
		if _, ok := target.CaSource.(*VirtualMeshSpec_RootCertificateAuthority_Secret); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSecret()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSecret()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSecret(), target.GetSecret()) {
				return false
			}
		}

	case *VirtualMeshSpec_RootCertificateAuthority_Vault:
		if _, ok := target.CaSource.(*VirtualMeshSpec_RootCertificateAuthority_Vault); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVault()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVault()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVault(), target.GetVault()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CaSource != target.CaSource {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualMeshSpec_Federation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualMeshSpec_Federation)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Federation)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetFlatNetwork() != target.GetFlatNetwork() {
		return false
	}

	if strings.Compare(m.GetHostnameSuffix(), target.GetHostnameSuffix()) != 0 {
		return false
	}

	switch m.Mode.(type) {

	case *VirtualMeshSpec_Federation_Permissive:
		if _, ok := target.Mode.(*VirtualMeshSpec_Federation_Permissive); !ok {
			return false
		}

		if h, ok := interface{}(m.GetPermissive()).(equality.Equalizer); ok {
			if !h.Equal(target.GetPermissive()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetPermissive(), target.GetPermissive()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Mode != target.Mode {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualMeshSpec_MTLSConfig_SharedTrust) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualMeshSpec_MTLSConfig_SharedTrust)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_MTLSConfig_SharedTrust)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRootCertificateAuthority()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRootCertificateAuthority()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRootCertificateAuthority(), target.GetRootCertificateAuthority()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIntermediateCertOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIntermediateCertOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIntermediateCertOptions(), target.GetIntermediateCertOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualMeshSpec_MTLSConfig_LimitedTrust) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualMeshSpec_MTLSConfig_LimitedTrust)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_MTLSConfig_LimitedTrust)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

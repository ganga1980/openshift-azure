// Code generated by protoc-gen-gogo.
// source: github.com/openshift/api/servicecertsigner/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		github.com/openshift/api/servicecertsigner/v1alpha1/generated.proto

	It has these top-level messages:
		DelegatedAuthentication
		DelegatedAuthorization
		ServiceCertSignerOperatorConfig
		ServiceCertSignerOperatorConfigList
		ServiceCertSignerOperatorConfigSpec
		ServiceCertSignerOperatorConfigStatus
		ServiceServingCertSignerConfig
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *DelegatedAuthentication) Reset()                    { *m = DelegatedAuthentication{} }
func (*DelegatedAuthentication) ProtoMessage()               {}
func (*DelegatedAuthentication) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *DelegatedAuthorization) Reset()                    { *m = DelegatedAuthorization{} }
func (*DelegatedAuthorization) ProtoMessage()               {}
func (*DelegatedAuthorization) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *ServiceCertSignerOperatorConfig) Reset()      { *m = ServiceCertSignerOperatorConfig{} }
func (*ServiceCertSignerOperatorConfig) ProtoMessage() {}
func (*ServiceCertSignerOperatorConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{2}
}

func (m *ServiceCertSignerOperatorConfigList) Reset()      { *m = ServiceCertSignerOperatorConfigList{} }
func (*ServiceCertSignerOperatorConfigList) ProtoMessage() {}
func (*ServiceCertSignerOperatorConfigList) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{3}
}

func (m *ServiceCertSignerOperatorConfigSpec) Reset()      { *m = ServiceCertSignerOperatorConfigSpec{} }
func (*ServiceCertSignerOperatorConfigSpec) ProtoMessage() {}
func (*ServiceCertSignerOperatorConfigSpec) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{4}
}

func (m *ServiceCertSignerOperatorConfigStatus) Reset()      { *m = ServiceCertSignerOperatorConfigStatus{} }
func (*ServiceCertSignerOperatorConfigStatus) ProtoMessage() {}
func (*ServiceCertSignerOperatorConfigStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{5}
}

func (m *ServiceServingCertSignerConfig) Reset()      { *m = ServiceServingCertSignerConfig{} }
func (*ServiceServingCertSignerConfig) ProtoMessage() {}
func (*ServiceServingCertSignerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{6}
}

func init() {
	proto.RegisterType((*DelegatedAuthentication)(nil), "github.com.openshift.api.servicecertsigner.v1alpha1.DelegatedAuthentication")
	proto.RegisterType((*DelegatedAuthorization)(nil), "github.com.openshift.api.servicecertsigner.v1alpha1.DelegatedAuthorization")
	proto.RegisterType((*ServiceCertSignerOperatorConfig)(nil), "github.com.openshift.api.servicecertsigner.v1alpha1.ServiceCertSignerOperatorConfig")
	proto.RegisterType((*ServiceCertSignerOperatorConfigList)(nil), "github.com.openshift.api.servicecertsigner.v1alpha1.ServiceCertSignerOperatorConfigList")
	proto.RegisterType((*ServiceCertSignerOperatorConfigSpec)(nil), "github.com.openshift.api.servicecertsigner.v1alpha1.ServiceCertSignerOperatorConfigSpec")
	proto.RegisterType((*ServiceCertSignerOperatorConfigStatus)(nil), "github.com.openshift.api.servicecertsigner.v1alpha1.ServiceCertSignerOperatorConfigStatus")
	proto.RegisterType((*ServiceServingCertSignerConfig)(nil), "github.com.openshift.api.servicecertsigner.v1alpha1.ServiceServingCertSignerConfig")
}
func (m *DelegatedAuthentication) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatedAuthentication) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *DelegatedAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatedAuthorization) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *ServiceCertSignerOperatorConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceCertSignerOperatorConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *ServiceCertSignerOperatorConfigList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceCertSignerOperatorConfigList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n4, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ServiceCertSignerOperatorConfigSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceCertSignerOperatorConfigSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.OperatorSpec.Size()))
	n5, err := m.OperatorSpec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ServiceServingCertSignerConfig.Size()))
	n6, err := m.ServiceServingCertSignerConfig.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func (m *ServiceCertSignerOperatorConfigStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceCertSignerOperatorConfigStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.OperatorStatus.Size()))
	n7, err := m.OperatorStatus.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	return i, nil
}

func (m *ServiceServingCertSignerConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceServingCertSignerConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ServingInfo.Size()))
	n8, err := m.ServingInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Authentication.Size()))
	n9, err := m.Authentication.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n9
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Authorization.Size()))
	n10, err := m.Authorization.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n10
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Signer.Size()))
	n11, err := m.Signer.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n11
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DelegatedAuthentication) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}

func (m *DelegatedAuthorization) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}

func (m *ServiceCertSignerOperatorConfig) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ServiceCertSignerOperatorConfigList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ServiceCertSignerOperatorConfigSpec) Size() (n int) {
	var l int
	_ = l
	l = m.OperatorSpec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.ServiceServingCertSignerConfig.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ServiceCertSignerOperatorConfigStatus) Size() (n int) {
	var l int
	_ = l
	l = m.OperatorStatus.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ServiceServingCertSignerConfig) Size() (n int) {
	var l int
	_ = l
	l = m.ServingInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Authentication.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Authorization.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Signer.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DelegatedAuthentication) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DelegatedAuthentication{`,
		`Disabled:` + fmt.Sprintf("%v", this.Disabled) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DelegatedAuthorization) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DelegatedAuthorization{`,
		`Disabled:` + fmt.Sprintf("%v", this.Disabled) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServiceCertSignerOperatorConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServiceCertSignerOperatorConfig{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ServiceCertSignerOperatorConfigSpec", "ServiceCertSignerOperatorConfigSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ServiceCertSignerOperatorConfigStatus", "ServiceCertSignerOperatorConfigStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServiceCertSignerOperatorConfigList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServiceCertSignerOperatorConfigList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "ServiceCertSignerOperatorConfig", "ServiceCertSignerOperatorConfig", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServiceCertSignerOperatorConfigSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServiceCertSignerOperatorConfigSpec{`,
		`OperatorSpec:` + strings.Replace(strings.Replace(this.OperatorSpec.String(), "OperatorSpec", "github_com_openshift_api_operator_v1alpha1.OperatorSpec", 1), `&`, ``, 1) + `,`,
		`ServiceServingCertSignerConfig:` + strings.Replace(strings.Replace(this.ServiceServingCertSignerConfig.String(), "RawExtension", "k8s_io_apimachinery_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServiceCertSignerOperatorConfigStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServiceCertSignerOperatorConfigStatus{`,
		`OperatorStatus:` + strings.Replace(strings.Replace(this.OperatorStatus.String(), "OperatorStatus", "github_com_openshift_api_operator_v1alpha1.OperatorStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServiceServingCertSignerConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServiceServingCertSignerConfig{`,
		`ServingInfo:` + strings.Replace(strings.Replace(this.ServingInfo.String(), "HTTPServingInfo", "github_com_openshift_api_config_v1.HTTPServingInfo", 1), `&`, ``, 1) + `,`,
		`Authentication:` + strings.Replace(strings.Replace(this.Authentication.String(), "DelegatedAuthentication", "DelegatedAuthentication", 1), `&`, ``, 1) + `,`,
		`Authorization:` + strings.Replace(strings.Replace(this.Authorization.String(), "DelegatedAuthorization", "DelegatedAuthorization", 1), `&`, ``, 1) + `,`,
		`Signer:` + strings.Replace(strings.Replace(this.Signer.String(), "CertInfo", "github_com_openshift_api_config_v1.CertInfo", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DelegatedAuthentication) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DelegatedAuthentication: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatedAuthentication: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Disabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DelegatedAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DelegatedAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatedAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Disabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceCertSignerOperatorConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceCertSignerOperatorConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceCertSignerOperatorConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceCertSignerOperatorConfigList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceCertSignerOperatorConfigList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceCertSignerOperatorConfigList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, ServiceCertSignerOperatorConfig{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceCertSignerOperatorConfigSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceCertSignerOperatorConfigSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceCertSignerOperatorConfigSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OperatorSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceServingCertSignerConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServiceServingCertSignerConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceCertSignerOperatorConfigStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceCertSignerOperatorConfigStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceCertSignerOperatorConfigStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OperatorStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceServingCertSignerConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceServingCertSignerConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceServingCertSignerConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServingInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServingInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Authentication.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Signer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/openshift/api/servicecertsigner/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 744 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xee, 0x52, 0x20, 0xcd, 0x14, 0x9a, 0x5f, 0xe6, 0xa7, 0xd8, 0xf4, 0xb0, 0x90, 0x1a, 0x0d,
	0x07, 0x98, 0x15, 0x30, 0x86, 0x70, 0x73, 0xc1, 0x3f, 0x44, 0x0c, 0x66, 0xdb, 0x83, 0x21, 0x1e,
	0x9c, 0x6e, 0xa7, 0xdb, 0x01, 0xba, 0xb3, 0xd9, 0x99, 0x56, 0xe1, 0x64, 0x8c, 0x07, 0x0f, 0x1e,
	0xfc, 0x02, 0x9e, 0xfc, 0x14, 0x7e, 0x03, 0x8e, 0x1c, 0x39, 0x11, 0xa9, 0x5f, 0xc3, 0x18, 0x33,
	0xb3, 0xd3, 0x76, 0x97, 0xba, 0x96, 0xd4, 0x70, 0xeb, 0xbc, 0x7d, 0xdf, 0xe7, 0x79, 0xde, 0xf7,
	0x79, 0xf3, 0x2e, 0xd8, 0xf4, 0xa8, 0x68, 0xb6, 0x6b, 0xc8, 0x65, 0x2d, 0x8b, 0x05, 0xc4, 0xe7,
	0x4d, 0xda, 0x10, 0x16, 0x0e, 0xa8, 0xc5, 0x49, 0xd8, 0xa1, 0x2e, 0x71, 0x49, 0x28, 0x38, 0xf5,
	0x7c, 0x12, 0x5a, 0x9d, 0x15, 0x7c, 0x18, 0x34, 0xf1, 0x8a, 0xe5, 0x11, 0x9f, 0x84, 0x58, 0x90,
	0x3a, 0x0a, 0x42, 0x26, 0x18, 0x5c, 0x1b, 0x80, 0xa0, 0x3e, 0x08, 0xc2, 0x01, 0x45, 0x43, 0x20,
	0xa8, 0x07, 0x52, 0x5a, 0x8e, 0x31, 0x7b, 0xcc, 0x63, 0x96, 0xc2, 0xaa, 0xb5, 0x1b, 0xea, 0xa5,
	0x1e, 0xea, 0x57, 0xc4, 0x51, 0x5a, 0x4d, 0x15, 0xea, 0x32, 0xbf, 0x41, 0x3d, 0xab, 0x33, 0xa4,
	0xab, 0xb4, 0x91, 0x5a, 0xc3, 0x02, 0x99, 0xc8, 0xd2, 0x7b, 0x2a, 0xdd, 0x3f, 0x58, 0xe7, 0x88,
	0x32, 0x99, 0xdd, 0xc2, 0x6e, 0x93, 0xfa, 0x24, 0x3c, 0xb2, 0x82, 0x03, 0x4f, 0x06, 0xb8, 0xd5,
	0x22, 0x02, 0xff, 0x89, 0xd1, 0x4a, 0xab, 0x0a, 0xdb, 0xbe, 0xa0, 0x2d, 0x32, 0x54, 0xf0, 0x60,
	0x54, 0x01, 0x77, 0x9b, 0xa4, 0x85, 0x2f, 0xd7, 0x95, 0x9f, 0x80, 0x5b, 0x5b, 0xe4, 0x90, 0x78,
	0x32, 0xf4, 0xb0, 0x2d, 0x9a, 0xc4, 0x17, 0xd4, 0xc5, 0x82, 0x32, 0x1f, 0x2e, 0x81, 0x5c, 0x9d,
	0x72, 0x5c, 0x3b, 0x24, 0xf5, 0xa2, 0xb1, 0x60, 0x2c, 0xe6, 0xec, 0xff, 0x4e, 0xce, 0xe7, 0x33,
	0xdd, 0xf3, 0xf9, 0xdc, 0x96, 0x8e, 0x3b, 0xfd, 0x8c, 0xf2, 0x63, 0x30, 0x97, 0x00, 0x62, 0x21,
	0x3d, 0x1e, 0x07, 0xe7, 0x43, 0x16, 0xcc, 0x57, 0x22, 0xb7, 0x37, 0x49, 0x28, 0x2a, 0xca, 0xed,
	0x5d, 0x3d, 0xe6, 0x4d, 0x65, 0x11, 0x7c, 0x0d, 0x72, 0x72, 0x70, 0x75, 0x2c, 0xb0, 0x42, 0xcc,
	0xaf, 0xde, 0x43, 0x51, 0xff, 0x28, 0xde, 0x3f, 0x0a, 0x0e, 0x3c, 0x19, 0xe0, 0x48, 0x66, 0xa3,
	0xce, 0x0a, 0xda, 0xad, 0xed, 0x13, 0x57, 0x3c, 0x27, 0x02, 0xdb, 0x50, 0x6b, 0x00, 0x83, 0x98,
	0xd3, 0x47, 0x85, 0xc7, 0x60, 0x92, 0x07, 0xc4, 0x2d, 0x4e, 0x28, 0xf4, 0x97, 0x68, 0x8c, 0xc5,
	0x44, 0x23, 0xba, 0xa8, 0x04, 0xc4, 0xb5, 0x67, 0xb4, 0x8a, 0x49, 0xf9, 0x72, 0x14, 0x27, 0x7c,
	0x6f, 0x80, 0x69, 0x2e, 0xb0, 0x68, 0xf3, 0x62, 0x56, 0xd1, 0xef, 0x5d, 0x0b, 0xbd, 0x62, 0xb0,
	0x0b, 0x5a, 0xc0, 0x74, 0xf4, 0x76, 0x34, 0x73, 0xf9, 0x97, 0x01, 0x6e, 0x8f, 0x40, 0xd8, 0xa1,
	0x5c, 0xc0, 0x57, 0x43, 0x56, 0xa0, 0xab, 0x59, 0x21, 0xab, 0x95, 0x11, 0xfd, 0x65, 0xe8, 0x45,
	0x62, 0x36, 0x1c, 0x81, 0x29, 0x2a, 0x48, 0x8b, 0x17, 0x27, 0x16, 0xb2, 0x8b, 0xf9, 0xd5, 0xea,
	0x75, 0x0c, 0xc2, 0x9e, 0xd5, 0x02, 0xa6, 0xb6, 0x25, 0x95, 0x13, 0x31, 0x96, 0xbf, 0x4d, 0x8c,
	0x1c, 0x80, 0xf4, 0x0c, 0x86, 0x60, 0xa6, 0x77, 0x04, 0xe4, 0x5b, 0x0f, 0x61, 0x3d, 0x5d, 0x69,
	0x2f, 0x7b, 0x20, 0x70, 0x37, 0x56, 0x6f, 0xdf, 0xd0, 0x6a, 0x66, 0xe2, 0x51, 0x27, 0xc1, 0x01,
	0xbf, 0x18, 0xc0, 0xd4, 0x0d, 0x2b, 0x89, 0xbe, 0x37, 0x90, 0x18, 0x49, 0xd3, 0x8b, 0xbb, 0x9c,
	0xea, 0x85, 0x3e, 0x0b, 0xc8, 0xc1, 0x6f, 0x1e, 0xbd, 0x15, 0xc4, 0xe7, 0x94, 0xf9, 0xf6, 0x5d,
	0xcd, 0x6d, 0x56, 0xfe, 0x0a, 0xee, 0x8c, 0x20, 0x2f, 0x7f, 0x35, 0xc0, 0x9d, 0x2b, 0xad, 0x1f,
	0x3c, 0x06, 0x85, 0x7e, 0x67, 0xd1, 0xca, 0x47, 0xf3, 0xdb, 0x18, 0x6b, 0x7e, 0xd1, 0x4a, 0xcf,
	0xe9, 0x2e, 0x0a, 0xc9, 0xb8, 0x73, 0x89, 0xa9, 0xfc, 0x33, 0x0b, 0x46, 0x34, 0x0a, 0xf7, 0x41,
	0x9e, 0x47, 0x7f, 0x6d, 0xfb, 0x0d, 0xa6, 0xb5, 0xad, 0xa5, 0x6b, 0x8b, 0x3e, 0x21, 0x72, 0xbd,
	0x9f, 0x56, 0xab, 0x2f, 0x2a, 0x83, 0x52, 0xfb, 0x7f, 0x2d, 0x2a, 0x1f, 0x0b, 0x3a, 0x71, 0x70,
	0xf8, 0xc9, 0x00, 0x05, 0x9c, 0xb8, 0xc0, 0xda, 0xc4, 0x9d, 0xb1, 0xb6, 0x3e, 0xe5, 0xaa, 0x0f,
	0xa6, 0x93, 0x8c, 0x3b, 0x97, 0xb8, 0xe1, 0x47, 0x03, 0xcc, 0xe2, 0xf8, 0x1d, 0xd7, 0xc7, 0xe8,
	0xd9, 0xbf, 0xab, 0xe9, 0x43, 0xda, 0x37, 0xb5, 0x98, 0xd9, 0x44, 0xd8, 0x49, 0x12, 0xc3, 0x2a,
	0x98, 0x8e, 0xf0, 0x8a, 0x93, 0x4a, 0xc2, 0xd2, 0x55, 0x0c, 0x90, 0x5e, 0xaa, 0xc9, 0x0f, 0x2e,
	0x9c, 0xc2, 0x70, 0x34, 0x96, 0x8d, 0x4e, 0x2e, 0xcc, 0xcc, 0xe9, 0x85, 0x99, 0x39, 0xbb, 0x30,
	0x33, 0xef, 0xba, 0xa6, 0x71, 0xd2, 0x35, 0x8d, 0xd3, 0xae, 0x69, 0x9c, 0x75, 0x4d, 0xe3, 0x7b,
	0xd7, 0x34, 0x3e, 0xff, 0x30, 0x33, 0x7b, 0xb9, 0x5e, 0x0b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0xa4, 0xa7, 0x07, 0xe2, 0x08, 0x00, 0x00,
}

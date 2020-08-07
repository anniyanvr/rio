// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/auth/common.proto

package envoy_api_v2_auth

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TlsParameters_TlsProtocol int32

const (
	TlsParameters_TLS_AUTO TlsParameters_TlsProtocol = 0
	TlsParameters_TLSv1_0  TlsParameters_TlsProtocol = 1
	TlsParameters_TLSv1_1  TlsParameters_TlsProtocol = 2
	TlsParameters_TLSv1_2  TlsParameters_TlsProtocol = 3
	TlsParameters_TLSv1_3  TlsParameters_TlsProtocol = 4
)

var TlsParameters_TlsProtocol_name = map[int32]string{
	0: "TLS_AUTO",
	1: "TLSv1_0",
	2: "TLSv1_1",
	3: "TLSv1_2",
	4: "TLSv1_3",
}

var TlsParameters_TlsProtocol_value = map[string]int32{
	"TLS_AUTO": 0,
	"TLSv1_0":  1,
	"TLSv1_1":  2,
	"TLSv1_2":  3,
	"TLSv1_3":  4,
}

func (x TlsParameters_TlsProtocol) String() string {
	return proto.EnumName(TlsParameters_TlsProtocol_name, int32(x))
}

func (TlsParameters_TlsProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa28efa2dc487b1f, []int{0, 0}
}

type CertificateValidationContext_TrustChainVerification int32

const (
	CertificateValidationContext_VERIFY_TRUST_CHAIN CertificateValidationContext_TrustChainVerification = 0
	CertificateValidationContext_ACCEPT_UNTRUSTED   CertificateValidationContext_TrustChainVerification = 1
)

var CertificateValidationContext_TrustChainVerification_name = map[int32]string{
	0: "VERIFY_TRUST_CHAIN",
	1: "ACCEPT_UNTRUSTED",
}

var CertificateValidationContext_TrustChainVerification_value = map[string]int32{
	"VERIFY_TRUST_CHAIN": 0,
	"ACCEPT_UNTRUSTED":   1,
}

func (x CertificateValidationContext_TrustChainVerification) String() string {
	return proto.EnumName(CertificateValidationContext_TrustChainVerification_name, int32(x))
}

func (CertificateValidationContext_TrustChainVerification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa28efa2dc487b1f, []int{4, 0}
}

type TlsParameters struct {
	TlsMinimumProtocolVersion TlsParameters_TlsProtocol `protobuf:"varint,1,opt,name=tls_minimum_protocol_version,json=tlsMinimumProtocolVersion,proto3,enum=envoy.api.v2.auth.TlsParameters_TlsProtocol" json:"tls_minimum_protocol_version,omitempty"`
	TlsMaximumProtocolVersion TlsParameters_TlsProtocol `protobuf:"varint,2,opt,name=tls_maximum_protocol_version,json=tlsMaximumProtocolVersion,proto3,enum=envoy.api.v2.auth.TlsParameters_TlsProtocol" json:"tls_maximum_protocol_version,omitempty"`
	CipherSuites              []string                  `protobuf:"bytes,3,rep,name=cipher_suites,json=cipherSuites,proto3" json:"cipher_suites,omitempty"`
	EcdhCurves                []string                  `protobuf:"bytes,4,rep,name=ecdh_curves,json=ecdhCurves,proto3" json:"ecdh_curves,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                  `json:"-"`
	XXX_unrecognized          []byte                    `json:"-"`
	XXX_sizecache             int32                     `json:"-"`
}

func (m *TlsParameters) Reset()         { *m = TlsParameters{} }
func (m *TlsParameters) String() string { return proto.CompactTextString(m) }
func (*TlsParameters) ProtoMessage()    {}
func (*TlsParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa28efa2dc487b1f, []int{0}
}

func (m *TlsParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsParameters.Unmarshal(m, b)
}
func (m *TlsParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsParameters.Marshal(b, m, deterministic)
}
func (m *TlsParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsParameters.Merge(m, src)
}
func (m *TlsParameters) XXX_Size() int {
	return xxx_messageInfo_TlsParameters.Size(m)
}
func (m *TlsParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsParameters.DiscardUnknown(m)
}

var xxx_messageInfo_TlsParameters proto.InternalMessageInfo

func (m *TlsParameters) GetTlsMinimumProtocolVersion() TlsParameters_TlsProtocol {
	if m != nil {
		return m.TlsMinimumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (m *TlsParameters) GetTlsMaximumProtocolVersion() TlsParameters_TlsProtocol {
	if m != nil {
		return m.TlsMaximumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (m *TlsParameters) GetCipherSuites() []string {
	if m != nil {
		return m.CipherSuites
	}
	return nil
}

func (m *TlsParameters) GetEcdhCurves() []string {
	if m != nil {
		return m.EcdhCurves
	}
	return nil
}

type PrivateKeyProvider struct {
	ProviderName string `protobuf:"bytes,1,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*PrivateKeyProvider_Config
	//	*PrivateKeyProvider_TypedConfig
	ConfigType           isPrivateKeyProvider_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *PrivateKeyProvider) Reset()         { *m = PrivateKeyProvider{} }
func (m *PrivateKeyProvider) String() string { return proto.CompactTextString(m) }
func (*PrivateKeyProvider) ProtoMessage()    {}
func (*PrivateKeyProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa28efa2dc487b1f, []int{1}
}

func (m *PrivateKeyProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateKeyProvider.Unmarshal(m, b)
}
func (m *PrivateKeyProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateKeyProvider.Marshal(b, m, deterministic)
}
func (m *PrivateKeyProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKeyProvider.Merge(m, src)
}
func (m *PrivateKeyProvider) XXX_Size() int {
	return xxx_messageInfo_PrivateKeyProvider.Size(m)
}
func (m *PrivateKeyProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKeyProvider.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKeyProvider proto.InternalMessageInfo

func (m *PrivateKeyProvider) GetProviderName() string {
	if m != nil {
		return m.ProviderName
	}
	return ""
}

type isPrivateKeyProvider_ConfigType interface {
	isPrivateKeyProvider_ConfigType()
}

type PrivateKeyProvider_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type PrivateKeyProvider_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*PrivateKeyProvider_Config) isPrivateKeyProvider_ConfigType() {}

func (*PrivateKeyProvider_TypedConfig) isPrivateKeyProvider_ConfigType() {}

func (m *PrivateKeyProvider) GetConfigType() isPrivateKeyProvider_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *PrivateKeyProvider) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*PrivateKeyProvider_Config); ok {
		return x.Config
	}
	return nil
}

func (m *PrivateKeyProvider) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*PrivateKeyProvider_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PrivateKeyProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PrivateKeyProvider_Config)(nil),
		(*PrivateKeyProvider_TypedConfig)(nil),
	}
}

type TlsCertificate struct {
	CertificateChain           *core.DataSource    `protobuf:"bytes,1,opt,name=certificate_chain,json=certificateChain,proto3" json:"certificate_chain,omitempty"`
	PrivateKey                 *core.DataSource    `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PrivateKeyProvider         *PrivateKeyProvider `protobuf:"bytes,6,opt,name=private_key_provider,json=privateKeyProvider,proto3" json:"private_key_provider,omitempty"`
	Password                   *core.DataSource    `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	OcspStaple                 *core.DataSource    `protobuf:"bytes,4,opt,name=ocsp_staple,json=ocspStaple,proto3" json:"ocsp_staple,omitempty"`
	SignedCertificateTimestamp []*core.DataSource  `protobuf:"bytes,5,rep,name=signed_certificate_timestamp,json=signedCertificateTimestamp,proto3" json:"signed_certificate_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}            `json:"-"`
	XXX_unrecognized           []byte              `json:"-"`
	XXX_sizecache              int32               `json:"-"`
}

func (m *TlsCertificate) Reset()         { *m = TlsCertificate{} }
func (m *TlsCertificate) String() string { return proto.CompactTextString(m) }
func (*TlsCertificate) ProtoMessage()    {}
func (*TlsCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa28efa2dc487b1f, []int{2}
}

func (m *TlsCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsCertificate.Unmarshal(m, b)
}
func (m *TlsCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsCertificate.Marshal(b, m, deterministic)
}
func (m *TlsCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsCertificate.Merge(m, src)
}
func (m *TlsCertificate) XXX_Size() int {
	return xxx_messageInfo_TlsCertificate.Size(m)
}
func (m *TlsCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_TlsCertificate proto.InternalMessageInfo

func (m *TlsCertificate) GetCertificateChain() *core.DataSource {
	if m != nil {
		return m.CertificateChain
	}
	return nil
}

func (m *TlsCertificate) GetPrivateKey() *core.DataSource {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *TlsCertificate) GetPrivateKeyProvider() *PrivateKeyProvider {
	if m != nil {
		return m.PrivateKeyProvider
	}
	return nil
}

func (m *TlsCertificate) GetPassword() *core.DataSource {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *TlsCertificate) GetOcspStaple() *core.DataSource {
	if m != nil {
		return m.OcspStaple
	}
	return nil
}

func (m *TlsCertificate) GetSignedCertificateTimestamp() []*core.DataSource {
	if m != nil {
		return m.SignedCertificateTimestamp
	}
	return nil
}

type TlsSessionTicketKeys struct {
	Keys                 []*core.DataSource `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TlsSessionTicketKeys) Reset()         { *m = TlsSessionTicketKeys{} }
func (m *TlsSessionTicketKeys) String() string { return proto.CompactTextString(m) }
func (*TlsSessionTicketKeys) ProtoMessage()    {}
func (*TlsSessionTicketKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa28efa2dc487b1f, []int{3}
}

func (m *TlsSessionTicketKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsSessionTicketKeys.Unmarshal(m, b)
}
func (m *TlsSessionTicketKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsSessionTicketKeys.Marshal(b, m, deterministic)
}
func (m *TlsSessionTicketKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsSessionTicketKeys.Merge(m, src)
}
func (m *TlsSessionTicketKeys) XXX_Size() int {
	return xxx_messageInfo_TlsSessionTicketKeys.Size(m)
}
func (m *TlsSessionTicketKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsSessionTicketKeys.DiscardUnknown(m)
}

var xxx_messageInfo_TlsSessionTicketKeys proto.InternalMessageInfo

func (m *TlsSessionTicketKeys) GetKeys() []*core.DataSource {
	if m != nil {
		return m.Keys
	}
	return nil
}

type CertificateValidationContext struct {
	TrustedCa                         *core.DataSource                                    `protobuf:"bytes,1,opt,name=trusted_ca,json=trustedCa,proto3" json:"trusted_ca,omitempty"`
	VerifyCertificateSpki             []string                                            `protobuf:"bytes,3,rep,name=verify_certificate_spki,json=verifyCertificateSpki,proto3" json:"verify_certificate_spki,omitempty"`
	VerifyCertificateHash             []string                                            `protobuf:"bytes,2,rep,name=verify_certificate_hash,json=verifyCertificateHash,proto3" json:"verify_certificate_hash,omitempty"`
	VerifySubjectAltName              []string                                            `protobuf:"bytes,4,rep,name=verify_subject_alt_name,json=verifySubjectAltName,proto3" json:"verify_subject_alt_name,omitempty"` // Deprecated: Do not use.
	MatchSubjectAltNames              []*matcher.StringMatcher                            `protobuf:"bytes,9,rep,name=match_subject_alt_names,json=matchSubjectAltNames,proto3" json:"match_subject_alt_names,omitempty"`
	RequireOcspStaple                 *wrappers.BoolValue                                 `protobuf:"bytes,5,opt,name=require_ocsp_staple,json=requireOcspStaple,proto3" json:"require_ocsp_staple,omitempty"`
	RequireSignedCertificateTimestamp *wrappers.BoolValue                                 `protobuf:"bytes,6,opt,name=require_signed_certificate_timestamp,json=requireSignedCertificateTimestamp,proto3" json:"require_signed_certificate_timestamp,omitempty"`
	Crl                               *core.DataSource                                    `protobuf:"bytes,7,opt,name=crl,proto3" json:"crl,omitempty"`
	AllowExpiredCertificate           bool                                                `protobuf:"varint,8,opt,name=allow_expired_certificate,json=allowExpiredCertificate,proto3" json:"allow_expired_certificate,omitempty"`
	TrustChainVerification            CertificateValidationContext_TrustChainVerification `protobuf:"varint,10,opt,name=trust_chain_verification,json=trustChainVerification,proto3,enum=envoy.api.v2.auth.CertificateValidationContext_TrustChainVerification" json:"trust_chain_verification,omitempty"`
	XXX_NoUnkeyedLiteral              struct{}                                            `json:"-"`
	XXX_unrecognized                  []byte                                              `json:"-"`
	XXX_sizecache                     int32                                               `json:"-"`
}

func (m *CertificateValidationContext) Reset()         { *m = CertificateValidationContext{} }
func (m *CertificateValidationContext) String() string { return proto.CompactTextString(m) }
func (*CertificateValidationContext) ProtoMessage()    {}
func (*CertificateValidationContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa28efa2dc487b1f, []int{4}
}

func (m *CertificateValidationContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateValidationContext.Unmarshal(m, b)
}
func (m *CertificateValidationContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateValidationContext.Marshal(b, m, deterministic)
}
func (m *CertificateValidationContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateValidationContext.Merge(m, src)
}
func (m *CertificateValidationContext) XXX_Size() int {
	return xxx_messageInfo_CertificateValidationContext.Size(m)
}
func (m *CertificateValidationContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateValidationContext.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateValidationContext proto.InternalMessageInfo

func (m *CertificateValidationContext) GetTrustedCa() *core.DataSource {
	if m != nil {
		return m.TrustedCa
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifyCertificateSpki() []string {
	if m != nil {
		return m.VerifyCertificateSpki
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifyCertificateHash() []string {
	if m != nil {
		return m.VerifyCertificateHash
	}
	return nil
}

// Deprecated: Do not use.
func (m *CertificateValidationContext) GetVerifySubjectAltName() []string {
	if m != nil {
		return m.VerifySubjectAltName
	}
	return nil
}

func (m *CertificateValidationContext) GetMatchSubjectAltNames() []*matcher.StringMatcher {
	if m != nil {
		return m.MatchSubjectAltNames
	}
	return nil
}

func (m *CertificateValidationContext) GetRequireOcspStaple() *wrappers.BoolValue {
	if m != nil {
		return m.RequireOcspStaple
	}
	return nil
}

func (m *CertificateValidationContext) GetRequireSignedCertificateTimestamp() *wrappers.BoolValue {
	if m != nil {
		return m.RequireSignedCertificateTimestamp
	}
	return nil
}

func (m *CertificateValidationContext) GetCrl() *core.DataSource {
	if m != nil {
		return m.Crl
	}
	return nil
}

func (m *CertificateValidationContext) GetAllowExpiredCertificate() bool {
	if m != nil {
		return m.AllowExpiredCertificate
	}
	return false
}

func (m *CertificateValidationContext) GetTrustChainVerification() CertificateValidationContext_TrustChainVerification {
	if m != nil {
		return m.TrustChainVerification
	}
	return CertificateValidationContext_VERIFY_TRUST_CHAIN
}

func init() {
	proto.RegisterEnum("envoy.api.v2.auth.TlsParameters_TlsProtocol", TlsParameters_TlsProtocol_name, TlsParameters_TlsProtocol_value)
	proto.RegisterEnum("envoy.api.v2.auth.CertificateValidationContext_TrustChainVerification", CertificateValidationContext_TrustChainVerification_name, CertificateValidationContext_TrustChainVerification_value)
	proto.RegisterType((*TlsParameters)(nil), "envoy.api.v2.auth.TlsParameters")
	proto.RegisterType((*PrivateKeyProvider)(nil), "envoy.api.v2.auth.PrivateKeyProvider")
	proto.RegisterType((*TlsCertificate)(nil), "envoy.api.v2.auth.TlsCertificate")
	proto.RegisterType((*TlsSessionTicketKeys)(nil), "envoy.api.v2.auth.TlsSessionTicketKeys")
	proto.RegisterType((*CertificateValidationContext)(nil), "envoy.api.v2.auth.CertificateValidationContext")
}

func init() { proto.RegisterFile("envoy/api/v2/auth/common.proto", fileDescriptor_fa28efa2dc487b1f) }

var fileDescriptor_fa28efa2dc487b1f = []byte{
	// 1125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4d, 0x6f, 0xdb, 0x36,
	0x18, 0x8e, 0x6c, 0xc7, 0x75, 0xe8, 0xa4, 0x70, 0x39, 0xaf, 0x51, 0x83, 0xb4, 0x75, 0xbd, 0x0d,
	0xc8, 0xb0, 0x40, 0x5a, 0xdd, 0xd3, 0x3e, 0x30, 0xcc, 0x72, 0x13, 0xa4, 0xdf, 0x81, 0xac, 0x66,
	0xeb, 0x89, 0x60, 0x64, 0xc6, 0xe6, 0x2c, 0x89, 0x1a, 0x49, 0x29, 0xf1, 0x6d, 0xd8, 0x71, 0x3b,
	0x0c, 0xe8, 0xa9, 0x3f, 0x60, 0xff, 0x61, 0xc3, 0x7e, 0xc1, 0xae, 0xfd, 0x09, 0xfb, 0x0b, 0x3b,
	0x16, 0xc3, 0x30, 0x88, 0x94, 0x12, 0x3b, 0x76, 0xd3, 0x00, 0xbb, 0x89, 0x7c, 0xde, 0xe7, 0x79,
	0xc9, 0xf7, 0x8b, 0x02, 0xb7, 0x48, 0x94, 0xb2, 0x89, 0x8d, 0x63, 0x6a, 0xa7, 0x1d, 0x1b, 0x27,
	0x72, 0x64, 0xfb, 0x2c, 0x0c, 0x59, 0x64, 0xc5, 0x9c, 0x49, 0x06, 0xaf, 0x29, 0xdc, 0xc2, 0x31,
	0xb5, 0xd2, 0x8e, 0x95, 0xe1, 0x1b, 0x9b, 0x33, 0x14, 0x9f, 0x71, 0x62, 0x1f, 0x62, 0x41, 0x34,
	0x61, 0xe3, 0xb6, 0x46, 0xe5, 0x24, 0x26, 0x76, 0x88, 0xa5, 0x3f, 0x22, 0xdc, 0x16, 0x92, 0xd3,
	0x68, 0x98, 0x1b, 0xdc, 0x18, 0x32, 0x36, 0x0c, 0x88, 0xad, 0x56, 0x87, 0xc9, 0x91, 0x8d, 0xa3,
	0x49, 0x0e, 0x6d, 0x9e, 0x87, 0x84, 0xe4, 0x89, 0x2f, 0x73, 0xf4, 0xd6, 0x79, 0xf4, 0x98, 0xe3,
	0x38, 0x26, 0x5c, 0x14, 0x78, 0x32, 0x88, 0xb1, 0x8d, 0xa3, 0x88, 0x49, 0x2c, 0x29, 0x8b, 0x84,
	0x1d, 0xd2, 0x21, 0xc7, 0xb2, 0x38, 0x59, 0x6b, 0x0e, 0x17, 0x24, 0x12, 0x54, 0xd2, 0xb4, 0xb0,
	0xb8, 0x39, 0x6f, 0x21, 0xb1, 0x4c, 0x0a, 0x07, 0xeb, 0x29, 0x0e, 0xe8, 0x00, 0x4b, 0x62, 0x17,
	0x1f, 0x1a, 0x68, 0xbf, 0x2a, 0x83, 0x35, 0x2f, 0x10, 0xfb, 0x98, 0xe3, 0x90, 0x48, 0xc2, 0x05,
	0x3c, 0x06, 0x9b, 0x32, 0x10, 0x28, 0xa4, 0x11, 0x0d, 0x93, 0x10, 0x29, 0x33, 0x9f, 0x05, 0x28,
	0x25, 0x5c, 0x50, 0x16, 0x99, 0x46, 0xcb, 0xd8, 0xba, 0xda, 0xd9, 0xb6, 0xe6, 0xa2, 0x6b, 0xcd,
	0xe8, 0xa8, 0x55, 0xce, 0x75, 0x6a, 0x6f, 0x9c, 0xe5, 0x1f, 0x8d, 0x52, 0xc3, 0x70, 0x6f, 0xc8,
	0x40, 0x3c, 0xd1, 0xd2, 0x05, 0x7a, 0xa0, 0x85, 0x4f, 0x1d, 0xe3, 0x93, 0xc5, 0x8e, 0x4b, 0xff,
	0xdf, 0xb1, 0x96, 0x3e, 0xef, 0xf8, 0x03, 0xb0, 0xe6, 0xd3, 0x78, 0x44, 0x38, 0x12, 0x09, 0x95,
	0x44, 0x98, 0xe5, 0x56, 0x79, 0x6b, 0xc5, 0x5d, 0xd5, 0x9b, 0x7d, 0xb5, 0x07, 0x6f, 0x83, 0x3a,
	0xf1, 0x07, 0x23, 0xe4, 0x27, 0x3c, 0x25, 0xc2, 0xac, 0x28, 0x13, 0x90, 0x6d, 0xf5, 0xd4, 0x4e,
	0xfb, 0x19, 0xa8, 0x4f, 0x79, 0x86, 0xab, 0xa0, 0xe6, 0x3d, 0xee, 0xa3, 0xee, 0x73, 0xef, 0x59,
	0x63, 0x09, 0xd6, 0xc1, 0x15, 0xef, 0x71, 0x3f, 0xbd, 0x8b, 0x3e, 0x6d, 0x18, 0x67, 0x8b, 0xbb,
	0x8d, 0xd2, 0xd9, 0xa2, 0xd3, 0x28, 0x9f, 0x2d, 0xee, 0x35, 0x2a, 0xed, 0xd7, 0x06, 0x80, 0xfb,
	0x9c, 0xa6, 0x58, 0x92, 0x47, 0x64, 0xb2, 0xcf, 0x59, 0x4a, 0x07, 0x84, 0xc3, 0x6d, 0xb0, 0x16,
	0xe7, 0xdf, 0x28, 0xc2, 0x21, 0x51, 0x09, 0x59, 0x71, 0xae, 0xbc, 0x71, 0x2a, 0xbc, 0xd4, 0x32,
	0xdc, 0xd5, 0x02, 0x7d, 0x8a, 0x43, 0x02, 0xbf, 0x00, 0x55, 0x9f, 0x45, 0x47, 0x74, 0xa8, 0xc2,
	0x57, 0xef, 0xac, 0x5b, 0xba, 0x14, 0xad, 0xa2, 0x14, 0xad, 0xbe, 0x2a, 0x54, 0xa7, 0xf6, 0xfb,
	0x6f, 0x3f, 0xff, 0x5a, 0x32, 0x4c, 0x63, 0x6f, 0xc9, 0xcd, 0x29, 0xb0, 0x0b, 0x56, 0xb3, 0x66,
	0x18, 0xa0, 0x5c, 0xa2, 0xac, 0x24, 0x9a, 0x73, 0x12, 0xdd, 0x68, 0xe2, 0x54, 0x35, 0x7f, 0x6f,
	0xc9, 0xad, 0x2b, 0x4e, 0x4f, 0x51, 0x9c, 0x35, 0x50, 0xd7, 0x64, 0x94, 0xed, 0xb6, 0xff, 0x2a,
	0x83, 0xab, 0x5e, 0x20, 0x7a, 0x84, 0x4b, 0x7a, 0x44, 0x7d, 0x2c, 0x09, 0x7c, 0x08, 0xae, 0xf9,
	0x67, 0x4b, 0xe4, 0x8f, 0x30, 0xd5, 0x45, 0x56, 0xef, 0xdc, 0x9c, 0xcd, 0x75, 0xd6, 0xaf, 0xd6,
	0x7d, 0x2c, 0x71, 0x9f, 0x25, 0xdc, 0x27, 0x6e, 0x63, 0x8a, 0xd7, 0xcb, 0x68, 0x70, 0x17, 0xd4,
	0x63, 0x1d, 0x31, 0x34, 0x26, 0x93, 0xfc, 0xca, 0x17, 0xab, 0x14, 0x07, 0x77, 0x41, 0x7c, 0x1a,
	0x6b, 0xf8, 0x0d, 0x68, 0x4e, 0xe9, 0xa0, 0x22, 0xa2, 0x66, 0x55, 0x09, 0x7e, 0xb4, 0xa0, 0x04,
	0xe7, 0x13, 0xe5, 0xc2, 0x78, 0x3e, 0x79, 0x5d, 0x50, 0x8b, 0xb1, 0x10, 0xc7, 0x8c, 0x0f, 0xf2,
	0x68, 0x5e, 0xf2, 0x74, 0xa7, 0x34, 0xf8, 0x15, 0xa8, 0x33, 0x5f, 0xc4, 0x48, 0x48, 0x1c, 0x07,
	0xc4, 0xac, 0x5c, 0x26, 0x52, 0x20, 0x63, 0xf4, 0x15, 0x01, 0x22, 0xb0, 0x29, 0xe8, 0x30, 0xca,
	0xb2, 0x3a, 0x15, 0x76, 0x49, 0x43, 0x22, 0x24, 0x0e, 0x63, 0x73, 0xb9, 0x55, 0x7e, 0xb7, 0xe0,
	0x86, 0x96, 0x98, 0xca, 0xa3, 0x57, 0x08, 0xb4, 0x5f, 0x80, 0xa6, 0x17, 0x88, 0x3e, 0x11, 0x59,
	0x73, 0x79, 0xd4, 0x1f, 0x13, 0xf9, 0x88, 0x4c, 0x04, 0xec, 0x82, 0xca, 0x98, 0x4c, 0x84, 0x69,
	0x5c, 0xc2, 0x81, 0x73, 0xf5, 0x8d, 0xb3, 0xfc, 0xd2, 0x28, 0xd5, 0x8c, 0xfc, 0xfe, 0x8a, 0xda,
	0xfe, 0xa7, 0x0a, 0x36, 0xa7, 0x7c, 0x1e, 0xe8, 0x59, 0x46, 0x59, 0xd4, 0x63, 0x91, 0x24, 0x27,
	0x12, 0x7e, 0x09, 0x80, 0xe4, 0x89, 0x90, 0xd9, 0xed, 0xf0, 0xe5, 0xaa, 0x68, 0x25, 0x27, 0xf4,
	0x30, 0xdc, 0x05, 0xeb, 0x29, 0xe1, 0xf4, 0x68, 0x32, 0x13, 0x1a, 0x11, 0x8f, 0xa9, 0x1e, 0x09,
	0xd9, 0xa9, 0xea, 0x2f, 0x8d, 0x5a, 0xbb, 0xca, 0x2b, 0xad, 0xed, 0xad, 0x6d, 0xf7, 0x7d, 0x6d,
	0x3e, 0x75, 0xa4, 0x7e, 0x3c, 0xa6, 0x6f, 0xd1, 0x19, 0x61, 0x31, 0x32, 0x4b, 0x73, 0x3a, 0x5f,
	0x6f, 0xa1, 0x05, 0x3a, 0x7b, 0x58, 0x8c, 0xe0, 0x67, 0xa7, 0x3a, 0x22, 0x39, 0xfc, 0x8e, 0xf8,
	0x12, 0xe1, 0x40, 0xea, 0xa6, 0x57, 0xf3, 0xc7, 0x29, 0x99, 0x86, 0xdb, 0xd4, 0x26, 0x7d, 0x6d,
	0xd1, 0x0d, 0xa4, 0xea, 0xfb, 0x6f, 0xc1, 0xba, 0x7a, 0xc2, 0xe6, 0x98, 0xc2, 0x5c, 0x51, 0xf1,
	0xbf, 0x93, 0x47, 0x25, 0x6b, 0x4b, 0x2b, 0x7f, 0xed, 0xb2, 0x59, 0x40, 0xa3, 0xe1, 0x13, 0xbd,
	0x72, 0x9b, 0x6a, 0x7b, 0x56, 0x58, 0xc0, 0x87, 0xe0, 0x3d, 0x4e, 0xbe, 0x4f, 0x28, 0x27, 0x68,
	0xba, 0x0e, 0x97, 0x55, 0xac, 0x37, 0xe6, 0x66, 0x83, 0xc3, 0x58, 0x70, 0x80, 0x83, 0x84, 0xb8,
	0xd7, 0x72, 0xda, 0xb3, 0xb3, 0x5a, 0x1c, 0x83, 0x0f, 0x0b, 0xad, 0x0b, 0x6b, 0xb2, 0xfa, 0x4e,
	0xf1, 0x3b, 0xb9, 0x4e, 0xff, 0xad, 0x75, 0x09, 0x6d, 0x50, 0xf6, 0x79, 0x60, 0x5e, 0xb9, 0x4c,
	0x51, 0x64, 0x96, 0xf0, 0x73, 0x70, 0x03, 0x07, 0x01, 0x3b, 0x46, 0xe4, 0x24, 0xa6, 0x7c, 0xf6,
	0x70, 0x66, 0xad, 0x65, 0x6c, 0xd5, 0xdc, 0x75, 0x65, 0xb0, 0xa3, 0xf1, 0xe9, 0xa9, 0xf6, 0x93,
	0x01, 0x4c, 0x55, 0x58, 0x7a, 0xa0, 0x21, 0x95, 0xa4, 0x0c, 0xca, 0x5e, 0x32, 0xa0, 0x5e, 0xb2,
	0xdd, 0x05, 0x63, 0xe4, 0xa2, 0xe2, 0xb6, 0xbc, 0x4c, 0x4f, 0x0d, 0xba, 0x83, 0x29, 0xb5, 0xa9,
	0x37, 0xee, 0xba, 0x5c, 0x68, 0xd1, 0xde, 0x05, 0xd7, 0x17, 0x73, 0xe1, 0x75, 0x00, 0x0f, 0x76,
	0xdc, 0x07, 0xbb, 0x2f, 0x90, 0xe7, 0x3e, 0xef, 0x7b, 0xa8, 0xb7, 0xd7, 0x7d, 0xf0, 0xb4, 0xb1,
	0x04, 0x9b, 0xa0, 0xd1, 0xed, 0xf5, 0x76, 0xf6, 0x3d, 0xf4, 0xfc, 0xa9, 0x42, 0x76, 0xee, 0x37,
	0x0c, 0x87, 0xfe, 0xfd, 0xea, 0xdf, 0x5f, 0x96, 0x3f, 0x81, 0x1f, 0xeb, 0x83, 0x93, 0x13, 0x99,
	0xfd, 0x85, 0xb0, 0x48, 0x58, 0x92, 0xe3, 0x48, 0xc4, 0x8c, 0x4b, 0x24, 0x58, 0xd6, 0xf0, 0xc2,
	0x92, 0x81, 0xb0, 0xd2, 0x7b, 0x7f, 0xfc, 0xf0, 0xe7, 0xeb, 0x6a, 0xa9, 0x61, 0x80, 0xdb, 0x94,
	0xe9, 0xeb, 0xc6, 0x9c, 0x9d, 0x4c, 0xe6, 0x6f, 0xee, 0xd4, 0x7b, 0xea, 0xdf, 0x4d, 0xbd, 0x9f,
	0xfb, 0xc6, 0x61, 0x55, 0xe5, 0xf8, 0xde, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x61, 0x7a,
	0xf4, 0xe5, 0x09, 0x00, 0x00,
}

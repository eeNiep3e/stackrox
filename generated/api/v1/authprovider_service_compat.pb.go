// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package v1

func (m *GetAuthProviderRequest) Size() int                      { return m.SizeVT() }
func (m *GetAuthProviderRequest) Clone() *GetAuthProviderRequest { return m.CloneVT() }
func (m *GetAuthProviderRequest) Marshal() ([]byte, error)       { return m.MarshalVT() }
func (m *GetAuthProviderRequest) Unmarshal(dAtA []byte) error    { return m.UnmarshalVT(dAtA) }

func (m *GetAuthProvidersRequest) Size() int                       { return m.SizeVT() }
func (m *GetAuthProvidersRequest) Clone() *GetAuthProvidersRequest { return m.CloneVT() }
func (m *GetAuthProvidersRequest) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *GetAuthProvidersRequest) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *GetLoginAuthProvidersResponse) Size() int                             { return m.SizeVT() }
func (m *GetLoginAuthProvidersResponse) Clone() *GetLoginAuthProvidersResponse { return m.CloneVT() }
func (m *GetLoginAuthProvidersResponse) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *GetLoginAuthProvidersResponse) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

func (m *GetLoginAuthProvidersResponse_LoginAuthProvider) Size() int { return m.SizeVT() }
func (m *GetLoginAuthProvidersResponse_LoginAuthProvider) Clone() *GetLoginAuthProvidersResponse_LoginAuthProvider {
	return m.CloneVT()
}
func (m *GetLoginAuthProvidersResponse_LoginAuthProvider) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *GetLoginAuthProvidersResponse_LoginAuthProvider) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *GetAuthProvidersResponse) Size() int                        { return m.SizeVT() }
func (m *GetAuthProvidersResponse) Clone() *GetAuthProvidersResponse { return m.CloneVT() }
func (m *GetAuthProvidersResponse) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *GetAuthProvidersResponse) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *PostAuthProviderRequest) Size() int                       { return m.SizeVT() }
func (m *PostAuthProviderRequest) Clone() *PostAuthProviderRequest { return m.CloneVT() }
func (m *PostAuthProviderRequest) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *PostAuthProviderRequest) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *UpdateAuthProviderRequest) Size() int                         { return m.SizeVT() }
func (m *UpdateAuthProviderRequest) Clone() *UpdateAuthProviderRequest { return m.CloneVT() }
func (m *UpdateAuthProviderRequest) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *UpdateAuthProviderRequest) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *ExchangeTokenRequest) Size() int                    { return m.SizeVT() }
func (m *ExchangeTokenRequest) Clone() *ExchangeTokenRequest { return m.CloneVT() }
func (m *ExchangeTokenRequest) Marshal() ([]byte, error)     { return m.MarshalVT() }
func (m *ExchangeTokenRequest) Unmarshal(dAtA []byte) error  { return m.UnmarshalVT(dAtA) }

func (m *ExchangeTokenResponse) Size() int                     { return m.SizeVT() }
func (m *ExchangeTokenResponse) Clone() *ExchangeTokenResponse { return m.CloneVT() }
func (m *ExchangeTokenResponse) Marshal() ([]byte, error)      { return m.MarshalVT() }
func (m *ExchangeTokenResponse) Unmarshal(dAtA []byte) error   { return m.UnmarshalVT(dAtA) }

func (m *AvailableProviderTypesResponse) Size() int                              { return m.SizeVT() }
func (m *AvailableProviderTypesResponse) Clone() *AvailableProviderTypesResponse { return m.CloneVT() }
func (m *AvailableProviderTypesResponse) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *AvailableProviderTypesResponse) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *AvailableProviderTypesResponse_AuthProviderType) Size() int { return m.SizeVT() }
func (m *AvailableProviderTypesResponse_AuthProviderType) Clone() *AvailableProviderTypesResponse_AuthProviderType {
	return m.CloneVT()
}
func (m *AvailableProviderTypesResponse_AuthProviderType) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *AvailableProviderTypesResponse_AuthProviderType) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}
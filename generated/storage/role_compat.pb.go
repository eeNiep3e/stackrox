// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *Role) Size() int                   { return m.SizeVT() }
func (m *Role) Clone() *Role                { return m.CloneVT() }
func (m *Role) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Role) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *PermissionSet) Size() int                   { return m.SizeVT() }
func (m *PermissionSet) Clone() *PermissionSet       { return m.CloneVT() }
func (m *PermissionSet) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *PermissionSet) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *SimpleAccessScope) Size() int                   { return m.SizeVT() }
func (m *SimpleAccessScope) Clone() *SimpleAccessScope   { return m.CloneVT() }
func (m *SimpleAccessScope) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *SimpleAccessScope) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *SimpleAccessScope_Rules) Size() int                       { return m.SizeVT() }
func (m *SimpleAccessScope_Rules) Clone() *SimpleAccessScope_Rules { return m.CloneVT() }
func (m *SimpleAccessScope_Rules) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *SimpleAccessScope_Rules) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *SimpleAccessScope_Rules_Namespace) Size() int { return m.SizeVT() }
func (m *SimpleAccessScope_Rules_Namespace) Clone() *SimpleAccessScope_Rules_Namespace {
	return m.CloneVT()
}
func (m *SimpleAccessScope_Rules_Namespace) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *SimpleAccessScope_Rules_Namespace) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *EffectiveAccessScope) Size() int                    { return m.SizeVT() }
func (m *EffectiveAccessScope) Clone() *EffectiveAccessScope { return m.CloneVT() }
func (m *EffectiveAccessScope) Marshal() ([]byte, error)     { return m.MarshalVT() }
func (m *EffectiveAccessScope) Unmarshal(dAtA []byte) error  { return m.UnmarshalVT(dAtA) }

func (m *EffectiveAccessScope_Namespace) Size() int                              { return m.SizeVT() }
func (m *EffectiveAccessScope_Namespace) Clone() *EffectiveAccessScope_Namespace { return m.CloneVT() }
func (m *EffectiveAccessScope_Namespace) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *EffectiveAccessScope_Namespace) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *EffectiveAccessScope_Cluster) Size() int                            { return m.SizeVT() }
func (m *EffectiveAccessScope_Cluster) Clone() *EffectiveAccessScope_Cluster { return m.CloneVT() }
func (m *EffectiveAccessScope_Cluster) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *EffectiveAccessScope_Cluster) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }
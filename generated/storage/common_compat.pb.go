// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *KeyValuePair) Size() int                   { return m.SizeVT() }
func (m *KeyValuePair) Clone() *KeyValuePair        { return m.CloneVT() }
func (m *KeyValuePair) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *KeyValuePair) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }
// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *TelemetryConfiguration) Size() int                      { return m.SizeVT() }
func (m *TelemetryConfiguration) Clone() *TelemetryConfiguration { return m.CloneVT() }
func (m *TelemetryConfiguration) Marshal() ([]byte, error)       { return m.MarshalVT() }
func (m *TelemetryConfiguration) Unmarshal(dAtA []byte) error    { return m.UnmarshalVT(dAtA) }
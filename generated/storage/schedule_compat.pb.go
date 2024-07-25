// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *Schedule) Size() int                   { return m.SizeVT() }
func (m *Schedule) Clone() *Schedule            { return m.CloneVT() }
func (m *Schedule) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Schedule) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *Schedule_WeeklyInterval) Size() int                       { return m.SizeVT() }
func (m *Schedule_WeeklyInterval) Clone() *Schedule_WeeklyInterval { return m.CloneVT() }
func (m *Schedule_WeeklyInterval) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *Schedule_WeeklyInterval) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *Schedule_DaysOfWeek) Size() int                   { return m.SizeVT() }
func (m *Schedule_DaysOfWeek) Clone() *Schedule_DaysOfWeek { return m.CloneVT() }
func (m *Schedule_DaysOfWeek) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Schedule_DaysOfWeek) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *Schedule_DaysOfMonth) Size() int                    { return m.SizeVT() }
func (m *Schedule_DaysOfMonth) Clone() *Schedule_DaysOfMonth { return m.CloneVT() }
func (m *Schedule_DaysOfMonth) Marshal() ([]byte, error)     { return m.MarshalVT() }
func (m *Schedule_DaysOfMonth) Unmarshal(dAtA []byte) error  { return m.UnmarshalVT(dAtA) }
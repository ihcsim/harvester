package util

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type ProtoMarshaler struct{}

func NewProtoMarshaler() *ProtoMarshaler {
	return &ProtoMarshaler{}
}

func (m *ProtoMarshaler) MarshalTimestamp(ts *timestamp.Timestamp) ([]byte, error) {
	return proto.Marshal(ts)
}

func (m *ProtoMarshaler) UnmarshalTimestamp(data []byte) (*timestamp.Timestamp, error) {
	ts := &timestamp.Timestamp{}
	err := proto.Unmarshal(data, ts)
	return ts, err
}

func MarshalProtoMessage(msg proto.Message) ([]byte, error) {
	return proto.Marshal(msg)
}

func UnmarshalProtoMessage(data []byte, msg proto.Message) error {
	return proto.Unmarshal(data, msg)
}

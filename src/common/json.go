package common

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type JsonProtoWrapper struct {
	data proto.Message
}

func NewJsonProtoWrapper(data proto.Message) *JsonProtoWrapper {
	return &JsonProtoWrapper{data: data}
}

func (w *JsonProtoWrapper) MarshalJSON() ([]byte, error) {
	opt := protojson.MarshalOptions{
		EmitDefaultValues: true,
	}

	return opt.Marshal(w.data)
}

type JsonRawWrapper struct {
	data []byte
}

func NewJsonRawWrapper(data []byte) *JsonRawWrapper {
	return &JsonRawWrapper{data: data}
}

func (w *JsonRawWrapper) MarshalJSON() ([]byte, error) {
	return w.data, nil
}

package common

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
)

func InitPbMessage(msg proto.Message) error {
	if msg == nil {
		return errors.New("msg is nil")
	}

	return initPbMessage_(reflect.ValueOf(msg), reflect.Value{}, 0)
}

func initPbMessage_(rv reflect.Value, rvBelongStruct reflect.Value, index int) error {
	switch rv.Kind() {
	case reflect.Ptr:
		if rv.IsNil() {
			if !rv.CanSet() {
				return fmt.Errorf("can not set value of %v", rv.Type())
			}
			elem := reflect.New(rv.Type().Elem()).Elem()
			rv.Set(elem.Addr())
		}
		if err := initPbMessage_(rv.Elem(), rvBelongStruct, index); err != nil {
			return err
		}
	case reflect.Struct:
		// pb中只有结构体的指针
		for i := 0; i < rv.NumField(); i++ {
			ft := rv.Type().Field(i)
			if !ft.IsExported() { //不初始化未导出变量
				continue
			}
			if err := initPbMessage_(rv.Field(i), rv, i); err != nil {
				return err
			}
		}
	case reflect.Slice:
		if !rv.CanSet() {
			return fmt.Errorf("can not set value of %v", rv.Type())
		}
		if rv.IsNil() {
			slice := reflect.MakeSlice(rv.Type(), 0, 0)
			rv.Set(slice)
		}

		if rv.Len() == 0 {
			// 初始化一个默认元素
			elem := reflect.New(rv.Type().Elem()).Elem()
			if err := initPbMessage_(elem, rvBelongStruct, index); err != nil {
				return err
			}
			rv.Set(reflect.Append(rv, elem))
		}

	case reflect.Map:
		if rv.IsNil() {
			if !rv.CanSet() {
				return fmt.Errorf("can not set value of %v", rv.Type())
			}
			mp := reflect.MakeMap(rv.Type())
			rv.Set(mp)
		}
		if rv.Len() == 0 {
			// 初始化一个默认元素
			key := reflect.New(rv.Type().Key()).Elem()
			if err := initPbMessage_(key, rvBelongStruct, index); err != nil {
				return err
			}

			val := reflect.New(rv.Type().Elem()).Elem()
			if err := initPbMessage_(val, rvBelongStruct, index); err != nil {
				return err
			}
			rv.SetMapIndex(key, val)
		}
	case reflect.Interface:
		//接口类型，处理oneof
		if !rvBelongStruct.IsValid() {
			return errors.New("deal Interface{} failed")
		}

		if !rvBelongStruct.CanAddr() {
			return errors.New("deal Interface{}, not a pointer")
		}

		// 获取proto.Message
		protoMsg, ok := rvBelongStruct.Addr().Interface().(proto.Message)
		if !ok {
			return errors.New("deal Interface{}, get proto.Message failed")
		}

		//通过MessageInfo获取oneOf接口
		msgInfo, ok := protoMsg.ProtoReflect().Type().(*protoimpl.MessageInfo)
		if !ok {
			return errors.New("deal Interface{}, get protoimpl.MessageInfo failed")
		}

		//尝试赋值
		for _, v := range msgInfo.OneofWrappers {
			rtOneOf := reflect.TypeOf(v)
			if rtOneOf.Implements(rv.Type()) {
				rvOneOf := reflect.New(rtOneOf.Elem())
				if err := initPbMessage_(rvOneOf, rvBelongStruct, index); err != nil {
					return err
				}

				rv.Set(rvOneOf)
				break
			}
		}
		return nil

	default:
	}
	return nil
}

package encoding

import (
	"fmt"
	"reflect"
)

func deepCopy(dst reflect.Value, src reflect.Value) error {
	switch dst.Kind() {
	case reflect.Invalid:
		return fmt.Errorf("Invalid Kind %v", dst)
	case reflect.Ptr, reflect.Interface:
		if src.IsNil() || !src.Elem().IsValid() {
			return nil
		}
		dst.Set(reflect.New(src.Elem().Type()))
		deepCopy(dst.Elem(), src.Elem())
	case reflect.Array, reflect.Slice:
		if src.IsNil() {
			return nil
		}
		dst.Set(reflect.MakeSlice(dst.Type(), src.Len(), src.Cap()))
		// reflect.Copy(dst, src)
		for i := 0; i < src.Len(); i++ {
			deepCopy(dst.Index(i), src.Index(i))
		}
	case reflect.Struct:
		dstTyp := dst.Type()
		for i := 0; i < dst.NumField(); i++ {
			// 防止乱序
			fieldName := dstTyp.Field(i).Name
			deepCopy(dst.Field(i), src.FieldByName(fieldName))
		}
	case reflect.Map:
		if src.IsNil() {
			return nil
		}
		dst.Set(reflect.MakeMap(dst.Type()))
		for _, key := range src.MapKeys() {
			srcValue := src.MapIndex(key)
			dstValue := reflect.New(srcValue.Type()).Elem()
			deepCopy(dstValue, srcValue)
			dst.SetMapIndex(key, dstValue)
		}
	default:
		dst.Set(src)
	}
	return nil
}

// DeepCopy object `from` to `to`
func DeepCopy(to interface{}, from interface{}) (copied interface{}, err error) {
	copy := reflect.New(reflect.ValueOf(to).Type()).Elem()
	if err = deepCopy(copy, reflect.ValueOf(from)); err != nil {
		return nil, err
	}
	// fmt.Println(copy.Interface())
	return copy.Interface(), nil
}

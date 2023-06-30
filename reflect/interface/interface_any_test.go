package demo

import (
	"fmt"
	"reflect"
	"testing"
)

type A struct {
	Name string
	Age  int
}

func TestInterface2any(t *testing.T) {
	src1 := []interface{}{1, 2, 3}
	var dst1 []int
	if err := convertInterfaceValue(src1, &dst1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dst1)
	}

	src2 := []interface{}{0.1, 0.2, 0.3}
	var dst2 []float64
	if err := convertInterfaceValue(src2, &dst2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dst2)
	}

	src3 := []interface{}{A{Name: "Alex"}}
	dst3 := []A{}
	convertInterfaceValue(src3, &dst3)
	fmt.Println(dst3)

}

func convertInterfaceValue(srcSlice interface{}, dstSlicePtr interface{}) error {
	// 获取源切片和目标切片的类型信息
	srcSliceValue := reflect.ValueOf(srcSlice)
	dstSliceValue := reflect.ValueOf(dstSlicePtr).Elem()
	srcSliceType := srcSliceValue.Type().Elem()
	dstSliceType := dstSliceValue.Type().Elem()

	// 检查源切片是否是一个切片类型
	if srcSliceValue.Kind() != reflect.Slice {
		return fmt.Errorf("srcSlice 不是一个切片类型")
	}

	// 检查目标切片是否是一个切片类型
	if dstSliceValue.Kind() != reflect.Slice {
		return fmt.Errorf("dstSlice 不是一个切片类型的指针")
	}

	// 检查目标切片的元素类型是否和源切片的元素类型相同或者可以进行类型转换
	if !dstSliceType.AssignableTo(srcSliceType) && !dstSliceType.ConvertibleTo(srcSliceType) {
		return fmt.Errorf("目标切片的元素类型 %v 和源切片的元素类型 %v 不匹配", dstSliceType, srcSliceType)
	}

	// 创建一个新的目标切片，长度为源切片的长度
	targetSlice := reflect.MakeSlice(reflect.SliceOf(dstSliceType), srcSliceValue.Len(), srcSliceValue.Cap())

	// 循环遍历源切片并进行类型转换
	for i := 0; i < srcSliceValue.Len(); i++ {
		// 获取源切片的当前元素
		srcElem := srcSliceValue.Index(i)
		if srcElem.Type().Kind() == reflect.Interface {
			srcElem = srcElem.Elem()
		}
		// 进行类型转换并设置到目标切片中
		if srcElem.Type().ConvertibleTo(dstSliceType) {
			targetSlice.Index(i).Set(srcElem.Convert(dstSliceType))
		} else if srcElem.Type().AssignableTo(dstSliceType) {
			targetSlice.Index(i).Set(srcElem.Convert(dstSliceType))
		} else {
			return fmt.Errorf("第 %d 个元素无法转换为 %v 类型", i+1, dstSliceType)
		}
	}

	// 将目标切片赋值给 dstSlice 指向的变量
	dstSliceValue.Set(targetSlice)

	return nil
}

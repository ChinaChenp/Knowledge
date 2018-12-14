package main

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	// TagName is the deepcopier struct tag name.
	TagName = "json"
	// FieldOptionName is the from field option name for struct tag.
	FieldOptionName = "field"
	// ContextOptionName is the context option name for struct tag.
	//ContextOptionName = "context"
	// SkipOptionName is the skip option name for struct tag.
	SkipOptionName = "-"
	// ForceOptionName is the skip option name for struct tag.
	ForceOptionName = "force"
)

type (
	// TagOptions is a map that contains extracted struct tag options.
	TagOptions map[string]string

	// Options are copier options.
	Options struct {
		// Context given to WithContext() method.
		Context map[string]interface{}
		// Reversed reverses struct tag checkings.
		Reversed bool
	}
)

// DeepCopier deep copies a struct to/from a struct.
type DeepCopier struct {
	dst interface{}
	src interface{}
	ctx map[string]interface{}
}

// Copy sets source or destination.
func Copy1(src interface{}) *DeepCopier {
	return &DeepCopier{src: src}
}

// WithContext injects the given context into the builder instance.
func (dc *DeepCopier) WithContext(ctx map[string]interface{}) *DeepCopier {
	dc.ctx = ctx
	return dc
}

// To sets the destination.
func (dc *DeepCopier) To(dst interface{}) error {
	dc.dst = dst
	return process(dc.dst, dc.src, Options{Context: dc.ctx})
}

// From sets the given the source as destination and destination as source.
func (dc *DeepCopier) From(src interface{}) error {
	dc.dst = dc.src
	dc.src = src
	return process(dc.dst, dc.src, Options{Context: dc.ctx, Reversed: true})
}

// process handles copy.
func process(dst interface{}, src interface{}, args ...Options) error {
	var (
		options       = Options{}
		srcValue      = reflect.Indirect(reflect.ValueOf(src))
		dstValue      = reflect.Indirect(reflect.ValueOf(dst))
		srcFieldNames = getFieldNames(src)
	)

	if len(args) > 0 {
		options = args[0]
	}

	if !dstValue.CanAddr() {
		return fmt.Errorf("destination %+v is unaddressable", dstValue.Interface())
	}

	fmt.Println(srcFieldNames)
	for _, f := range srcFieldNames {
		var (
			srcFieldValue               = srcValue.FieldByName(f)
			srcFieldType, srcFieldFound = srcValue.Type().FieldByName(f)
			srcFieldName                = srcFieldType.Name
			dstFieldName                = srcFieldName
			tagOptions                  TagOptions
		)

		if !srcFieldFound {
			continue
		}

		if options.Reversed {
			tagOptions = getTagOptions(srcFieldType.Tag.Get(TagName))
			if v, ok := tagOptions[FieldOptionName]; ok && v != "" {
				dstFieldName = v
			}
		} else {
			if name, opts := getRelatedField(dst, srcFieldName); name != "" {
				dstFieldName, tagOptions = name, opts
			}
		}

		if _, ok := tagOptions[SkipOptionName]; ok {
			continue
		}

		var (
			dstFieldType, dstFieldFound = dstValue.Type().FieldByName(dstFieldName)
			dstFieldValue               = dstValue.FieldByName(dstFieldName)
		)

		if !dstFieldFound {
			continue
		}

		// Valuer -> value
		if isNullableType(srcFieldType.Type) {
			// We have same nullable type on both sides
			if srcFieldValue.Type().AssignableTo(dstFieldType.Type) {
				dstFieldValue.Set(srcFieldValue)
				continue
			}
		}

		if dstFieldValue.Kind() == reflect.Interface {
			continue
		}

		// Ptr -> Value
		if srcFieldType.Type.Kind() == reflect.Ptr && !srcFieldValue.IsNil() && dstFieldType.Type.Kind() != reflect.Ptr {
			indirect := reflect.Indirect(srcFieldValue)

			if indirect.Type().AssignableTo(dstFieldType.Type) {
				dstFieldValue.Set(indirect)
				continue
			}
		}

		// Other types
		types := srcFieldType.Type
		fmt.Println("other", srcFieldType.Type.Kind(), srcFieldType.Type, types)
		fmt.Println("other", dstFieldType.Type.Kind(), dstFieldType.Type)

		if srcFieldType.Type.AssignableTo(dstFieldType.Type) {
			dstFieldValue.Set(srcFieldValue)
		} else {
			// fn := srcFieldType.Type.ConvertibleTo(dstFieldType.Type)
			// fmt.Println("not set %v", fn)
		}
	}
	//
	//for _, m := range srcMethodNames {
	//	name, opts := getRelatedField(dst, m)
	//	if name == "" {
	//		continue
	//	}
	//
	//	if _, ok := opts[SkipOptionName]; ok {
	//		continue
	//	}
	//
	//	method := reflect.ValueOf(src).MethodByName(m)
	//	if !method.IsValid() {
	//		return fmt.Errorf("method %s is invalid", m)
	//	}
	//
	//	var (
	//		dstFieldType, _ = dstValue.Type().FieldByName(name)
	//		dstFieldValue   = dstValue.FieldByName(name)
	//		_, withContext  = opts[ContextOptionName]
	//		_, force        = opts[ForceOptionName]
	//	)
	//
	//	args := []reflect.Value{}
	//	if withContext {
	//		args = []reflect.Value{reflect.ValueOf(options.Context)}
	//	}
	//
	//	var (
	//		result          = method.Call(args)[0]
	//		resultInterface = result.Interface()
	//		resultValue     = reflect.ValueOf(resultInterface)
	//		resultType      = resultValue.Type()
	//	)
	//
	//	// Value -> Ptr
	//	if dstFieldValue.Kind() == reflect.Ptr && force {
	//		ptr := reflect.New(resultType)
	//		ptr.Elem().Set(resultValue)
	//
	//		if ptr.Type().AssignableTo(dstFieldType.Type) {
	//			dstFieldValue.Set(ptr)
	//		}
	//
	//		continue
	//	}
	//
	//	// Ptr -> value
	//	if resultValue.Kind() == reflect.Ptr && force {
	//		if resultValue.Elem().Type().AssignableTo(dstFieldType.Type) {
	//			dstFieldValue.Set(resultValue.Elem())
	//		}
	//
	//		continue
	//	}
	//
	//	if resultType.AssignableTo(dstFieldType.Type) && result.IsValid() {
	//		dstFieldValue.Set(result)
	//	}
	//}

	return nil
}

// getTagOptions parses deepcopier tag field and returns options.
func getTagOptions(value string) TagOptions {
	options := TagOptions{}

	for _, opt := range strings.Split(value, ";") {
		o := strings.Split(opt, ":")

		// deepcopier:"keyword; without; value;"
		if len(o) == 1 {
			options[o[0]] = ""
		}

		// deepcopier:"key:value; anotherkey:anothervalue"
		if len(o) == 2 {
			options[strings.TrimSpace(o[0])] = strings.TrimSpace(o[1])
		}
	}

	return options
}

// getRelatedField returns first matching field.
func getRelatedField(instance interface{}, name string) (string, TagOptions) {
	var (
		value      = reflect.Indirect(reflect.ValueOf(instance))
		fieldName  string
		tagOptions TagOptions
	)

	for i := 0; i < value.NumField(); i++ {
		var (
			vField     = value.Field(i)
			tField     = value.Type().Field(i)
			tagOptions = getTagOptions(tField.Tag.Get(TagName))
		)

		if tField.Type.Kind() == reflect.Struct && tField.Anonymous {
			if n, o := getRelatedField(vField.Interface(), name); n != "" {
				return n, o
			}
		}

		if v, ok := tagOptions[FieldOptionName]; ok && v == name {
			return tField.Name, tagOptions
		}

		if tField.Name == name {
			return tField.Name, tagOptions
		}
	}

	return fieldName, tagOptions
}

// getMethodNames returns instance's method names.
func getMethodNames(instance interface{}) []string {
	var methods []string

	t := reflect.TypeOf(instance)
	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}

	return methods
}

// getFieldNames returns instance's field names.
func getFieldNames(instance interface{}) []string {
	var (
		fields []string
		v      = reflect.Indirect(reflect.ValueOf(instance))
		t      = v.Type()
	)

	if t.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < v.NumField(); i++ {
		var (
			vField = v.Field(i)
			tField = v.Type().Field(i)
		)

		// Is exportable?
		if tField.PkgPath != "" {
			continue
		}

		if tField.Type.Kind() == reflect.Struct && tField.Anonymous {
			fields = append(fields, getFieldNames(vField.Interface())...)
			continue
		}

		fields = append(fields, tField.Name)
	}

	return fields
}

// isNullableType returns true if the given type is a nullable one.
func isNullableType(t reflect.Type) bool {
	return t.ConvertibleTo(reflect.TypeOf((*driver.Valuer)(nil)).Elem())
}

func shallowCopy(dst, src interface{}) error {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)

	fmt.Println(srcType.String(), dstType.String())

	if srcType.AssignableTo(dstType) {
		srcValue  := reflect.Indirect(reflect.ValueOf(src))
		dstValue  := reflect.Indirect(reflect.ValueOf(dst))

		if !dstValue.CanAddr() {
			return fmt.Errorf("destination %+v is unaddressable", dstValue.Interface())
		}
		dstValue.Set(srcValue)
	} else {
		return fmt.Errorf("can't assign able")
	}
	return nil
}

type (
	Rel struct {
		Int int
		Str string  `deepcopier:"skip"`
	}

	Rel1 struct {
		Int int
		Str string
	}

	Src struct {
		Int         int
		IntPtr      *int
		Slice       []string
		SlicePtr    *[]string
		Map         map[string]interface{}
		MapPtr      *map[string]interface{}
		SliceStruct []*Rel
		Struct      Rel
		StructPtr   *Rel
		Skipped     string `deepcopier:"skip"`
	}

	Dst struct {
		Int         int
		IntPtr      *int
		Slice       []string
		SlicePtr    *[]string
		Map         map[string]interface{}
		MapPtr      *map[string]interface{}
		SliceStruct []*Rel1
		Struct      Rel
		StructPtr   *Rel
		Skipped     string `deepcopier:"skip"`
	}

	Renamed struct {
		MyInt       int                     `deepcopier:"field:Int"`
		MyIntPtr    *int                    `deepcopier:"field:IntPtr"`
		MySlice     []string                `deepcopier:"field:Slice"`
		MySlicePtr  *[]string               `deepcopier:"field:SlicePtr"`
		MyMap       map[string]interface{}  `deepcopier:"field:Map"`
		MyMapPtr    *map[string]interface{} `deepcopier:"field:MapPtr"`
		MyStruct    Rel                     `deepcopier:"field:Struct"`
		MyStructPtr *Rel                    `deepcopier:"field:StructPtr"`
		Skipped     string                  `deepcopier:"skip"`
	}
)

func main() {
	var (
		integer = 1
		rel     = Rel{Int: 1, Str: "fuck"}
		slc     = []string{"one", "two", "three"}
		mp      = map[string]interface{}{"one": 1}
		ss      = []*Rel{{Int: 1, Str: "fuck1"}, {Int: 2, Str: "fuck2"}, {Int: 3, Str: "fuck3"}, {Int: 4, Str: "fuck4"}}
	)

	src := &Src{
		Int:         integer,
		IntPtr:      &integer,
		Slice:       slc,
		SlicePtr:    &slc,
		Map:         mp,
		MapPtr:      &mp,
		SliceStruct: ss,
		Struct:      rel,
		StructPtr:   &rel,
		Skipped:     "I should be skipped",
	}

	//srcRenamed := &Renamed{
	//	MyInt:       integer,
	//	MyIntPtr:    &integer,
	//	MySlice:     slc,
	//	MySlicePtr:  &slc,
	//	MyMap:       mp,
	//	MyMapPtr:    &mp,
	//	MyStruct:    rel,
	//	MyStructPtr: &rel,
	//	Skipped:     "I should be skipped",
	//}

	//
	// To()
	//

	dst := &Dst{}
	//Copy1(dst).From(src)
	//err := shallowCopy(&dst, src)
	err := Copy(dst, src)
	fmt.Println(err)
	fmt.Println(src)
	fmt.Println(dst)

	//src.Slice = append(src.Slice, "four")
	//src.Map["two"] = 2
	//
	//fmt.Println(src)
	//fmt.Println(dst)
	//fmt.Println(dst.SliceStruct)
}

func Copy(dst, src interface{}) error {
	if src == nil {
		return nil
	}

	// Make the interface a reflect.Value
	from := reflect.Indirect(reflect.ValueOf(src))
	to := reflect.Indirect(reflect.ValueOf(dst))

	// Make a copy of the same type as the original.
	//cpy := reflect.New(original.Type()).Elem()

	// Recursively copy the original.
	copyRecursive(from, to)

	// Return the copy as an interface.
	return nil
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(original, cpy reflect.Value) {
	// check for implement deepcopy.Interface
	//if original.CanInterface() {
	//	if copier, ok := original.Interface().(Interface); ok {
	//		cpy.Set(reflect.ValueOf(copier.DeepCopy()))
	//		return
	//	}
	//}

	//if original.Type().AssignableTo(cpy.Type()) {
	//	fmt.Println("can set")
	//	cpy.Set(original)
	//	return
	//}

	// handle according to original's Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to.
		originalValue := original.Elem()

		// if  it isn't valid, return.
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())

	case reflect.Interface:
		// If this is a nil, don't do anything
		if original.IsNil() {
			return
		}
		// Get the value for the interface, not the pointer.
		originalValue := original.Elem()

		// Get the value by calling Elem().
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		t, ok := original.Interface().(time.Time)
		if ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		// Go through each field of the struct and copy it.
		for i := 0; i < original.NumField(); i++ {
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(original.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}
		// Make a new slice and copy each element.
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}

	case reflect.Map:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)

			copyKey := reflect.New(key.Type().Elem())
			//copyKey := Copy(key.Interface())
			copyRecursive(key, copyKey)
			cpy.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		cpy.Set(original)
	}
}


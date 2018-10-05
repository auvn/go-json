package jsonutil

type Object map[string]interface{}

func (o Object) Value(key string) (Value, bool) {
	value, ok := o[key]
	return Value{Value: value}, ok
}

type Value struct {
	Value interface{}
}

func (v Value) Int() int {
	val, _ := v.Value.(float64)
	return int(val)
}

func (v Value) String() string {
	val, _ := v.Value.(string)
	return val
}

func (v Value) Object() Object {
	val, _ := v.Value.(map[string]interface{})
	return val
}

func (v Value) Objects() []Object {
	valsSlice, _ := v.Value.([]interface{})
	oo := make([]Object, len(valsSlice))
	for i := range valsSlice {

		val, _ := valsSlice[i].(map[string]interface{})
		oo[i] = val
	}
	return oo
}

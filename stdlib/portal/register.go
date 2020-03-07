package libportal

var ELEMENTS = map[string]interface{}{}

func register(k string, v interface{}) {
	ELEMENTS[k] = v
}

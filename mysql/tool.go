package mysql


func MapGet(m map[string]interface{},key string,defaultval interface{})   interface{} {
	var res interface{} = nil
	var ok bool = false
    if res,ok = m[key] ; !ok {
    	res = defaultval
	}
	return res;
}

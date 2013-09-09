package util

//判断相等
func Equals(a, b interface{}) bool {
	switch a.(type) {
	case int:
		_, btype := b.(int)
		if btype {
			return a == b
		} else {
			return false
		}
	case uint:
		_, btype := b.(uint)
		if btype {
			return a == b
		} else {
			return false
		}
	case int64:
		_, btype := b.(int64)
		if btype {
			return a == b
		} else {
			return false
		}
	case uint64:
		_, btype := b.(uint64)
		if btype {
			return a == b
		} else {
			return false
		}
	case byte:
		_, btype := b.(byte)
		if btype {
			return a == b
		} else {
			return false
		}
	case float32:
		_, btype := b.(float32)
		if btype {
			return a == b
		} else {
			return false
		}
	case string:
		_, btype := b.(string)
		if btype {
			return a == b
		} else {
			return false
		}
	default:
		return false
	}
}

//生成列表序号
func GenerateSeqNo(seq, step, pos int) int {
	return seq + 1 + (step * (pos - 1))
}

//生成查询参数连接url
func GenerateParamUrl(p map[string]string) string {
	purl := ""
	for k, v := range p {
		purl += "&" + k + "=" + v
	}
	return purl
}

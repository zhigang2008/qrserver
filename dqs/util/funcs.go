package util

import (
	"fmt"
	"strconv"
	"strings"
)

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
func GenerateParamUrl(p map[string]interface{}) string {
	purl := ""
	for k, o := range p {
		purl += "&" + k + "="
		switch v := o.(type) {
		case int:
			purl += strconv.Itoa(v)
		case bool:
			purl += strconv.FormatBool(v)
		case int64:
			purl += strconv.FormatInt(v, 10)
		case uint64:
			purl += strconv.FormatUint(v, 10)
		default:
			purl += fmt.Sprintf("%s", v)

		}
	}
	return purl
}

//判断是否包含在数组中
func Contain(c interface{}, b interface{}) bool {
	switch c.(type) {
	case []string:
		n, ok := c.([]string)
		if ok {
			for _, v := range n {
				if strings.TrimSpace(fmt.Sprintf("%v", v)) == strings.TrimSpace(fmt.Sprintf("%v", b)) {
					return true
				}
			}
		}
	case []int:
		n, ok := c.([]int)
		if ok {
			for _, v := range n {
				if strings.TrimSpace(fmt.Sprintf("%v", v)) == strings.TrimSpace(fmt.Sprintf("%v", b)) {
					return true
				}
			}
		}
	case []int32:
		n, ok := c.([]int32)
		if ok {
			for _, v := range n {
				if strings.TrimSpace(fmt.Sprintf("%v", v)) == strings.TrimSpace(fmt.Sprintf("%v", b)) {
					return true
				}
			}
		}
	case []int64:
		n, ok := c.([]int64)
		if ok {
			for _, v := range n {
				if strings.TrimSpace(fmt.Sprintf("%v", v)) == strings.TrimSpace(fmt.Sprintf("%v", b)) {
					return true
				}
			}
		}
	default:
		return false
	}

	return false
}

/**/
func HasRoles(userRoles []string, roles ...string) bool {
	if userRoles != nil {
		for _, role := range roles {
			if Contain(userRoles, role) {
				return true
			}
		}
	} else {
		return false
	}
	return false
}

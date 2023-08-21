package myutil

// Set，与 java 中的 Set 类似，特点是在添加数组的时候，不允许出现重复的 string 元素
type Set map[string]bool

// 添加单个 string 元素
func (s Set) Add(item string) {
	s[item] = true
}

// 添加 []string 元素
func (s Set) AddAll(items []string) {
	for _, item := range items {
		s.Add(item)
	}
}

// 判断该元素是否包含在 set 内
func (s Set) Contains(item string) bool {
	return s[item]
}

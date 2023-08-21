package myutil

// Set，与 java 中的 Set 类似，特点是在添加数组的时候，不允许出现重复的 string 元素
type Set map[string]bool

func (s Set) Add(item string) {
	s[item] = true
}

func (s Set) Contains(item string) bool {
	return s[item]
}

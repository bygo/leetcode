package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs[1:] {
		i := 0
		for i < len(str) && i < len(prefix) && prefix[i] == str[i] {
			i++
		}
		prefix = str[:i]
	}

	return prefix
}

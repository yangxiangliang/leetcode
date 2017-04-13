package leetcode

import "strings"

func wordPattern(pattern string, str string) bool {
    keyTargetMap := map[string]string{}
    TargetStrSet := map[string]bool{}

    strList := strings.Split(str, " ")
    if len(pattern) != len(strList) {
        return false
    }
    for index, char := range pattern {
        keyStr := string(char)
        valueStr := strList[index]
        if targetStr, exist := keyTargetMap[keyStr]; exist {
            if targetStr != valueStr {
                return false
            }
        } else if TargetStrSet[valueStr] { // cannot map 2 diff strings to same char
            return false
        } else {
            keyTargetMap[keyStr] = valueStr
            TargetStrSet[valueStr] = true
        }
    }
    return true
}

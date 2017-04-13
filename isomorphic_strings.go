package leetcode

func isIsomorphic1(s string, t string) bool {
    r1 := [256]int{}
    r2 := [256]int{}
    for i := 0; i < len(s); i ++ {
        r1[s[i]] += i + 1 // cannot do += i, because i may be 0, e.g. s="ab", t="aa"
        r2[t[i]] += i + 1
        if r1[s[i]] != r2[t[i]] {
            return false
        }
    }
    return true
}

func isIsomorphic2(s string, t string) bool {
    charMap := map[string]string{}
    targetStrMap := map[string]bool{}
    for index, char := range s {
        tempStr1 := string(char)
        tempStr2 := string(t[index])
        if targetStr, exist := charMap[tempStr1]; exist {
            if targetStr != tempStr2 {
                return false
            }
        } else if targetStrMap[tempStr2] { // cannot map 2 different char to same char
                return false
        }else {
               charMap[tempStr1] = tempStr2
               targetStrMap[tempStr2] = true
            }
    }
    return true
}

package validanagram

import "sort"

// Sorting-based solution
func isAnagramSort(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sRunes := []rune(s)
    tRunes := []rune(t)
    
    sort.Slice(sRunes, func(i, j int) bool { return sRunes[i] < sRunes[j] })
    sort.Slice(tRunes, func(i, j int) bool { return tRunes[i] < tRunes[j] })
    
    return string(sRunes) == string(tRunes)
}

// Array-based solution (for ASCII only)
func isAnagramASCII(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    counts := [256]int{}
    
    for i := 0; i < len(s); i++ {
        counts[s[i]]++
    }
    
    for i := 0; i < len(t); i++ {
        counts[t[i]]--
        if counts[t[i]] < 0 {
            return false
        }
    }
    
    return true
}
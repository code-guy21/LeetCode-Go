package validanagram

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    charMap := make(map[rune]int)

    for _, char := range s {
        charMap[char]++
    }

    for _, char := range t {
        if value, exists := charMap[char]; exists && value > 0 {
            charMap[char]--
        } else {
            return false
        }
    }

    return true
}
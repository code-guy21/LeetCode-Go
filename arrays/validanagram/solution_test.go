package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
    tests := []struct {
        s    string
        t    string
        want bool
    }{
        {"anagram", "nagaram", true},
        {"rat", "car", false},
        {"", "", true},
        {"ab", "a", false},
        {"café", "fécà", false },
        {"汉字", "字汉", true},
        {"😀😂", "😂😀", true},
    }

    for _, tt := range tests {
        got := isAnagram(tt.s, tt.t)
        if got != tt.want {
            t.Errorf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
        }
    }
}

func BenchmarkIsAnagram(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isAnagram("anagram", "nagaram")
    }
}
package validanagram

import "testing"

func TestAllAnagramVariations(t *testing.T) {
    tests := []struct {
        s    string
        t    string
        want bool
    }{
        {"anagram", "nagaram", true},
        {"rat", "car", false},
        {"", "", true},
        {"ab", "a", false},
        {"café", "fécà", false},
        {"汉字", "字汉", true},
        {"😀😂", "😂😀", true},
    }

    for _, tt := range tests {
        // Test sorting solution
        got := isAnagramSort(tt.s, tt.t)
        if got != tt.want {
            t.Errorf("isAnagramSort(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
        }

        // Test ASCII solution only for ASCII strings
        if isASCII(tt.s) && isASCII(tt.t) {
            got = isAnagramASCII(tt.s, tt.t)
            if got != tt.want {
                t.Errorf("isAnagramASCII(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
            }
        }
    }
}

// Helper function to check if string contains only ASCII characters
func isASCII(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] > 127 {
            return false
        }
    }
    return true
}

func BenchmarkAnagramVariations(b *testing.B) {
    // Test cases
    testCases := []struct {
        name string
        s    string
        t    string
    }{
        {"Short-ASCII", "listen", "silent"},
        {"Long-ASCII", "pneumonoultramicroscopicsilicovolcanoconiosis", "micropneumonoultrascopicsilicovolcanoconiosis"},
        {"Unicode", "汉字测试", "测试汉字"},
        {"Mixed", "Hello世界", "世界Hello"},
    }

    for _, tc := range testCases {
        b.Run("Sort-"+tc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                isAnagramSort(tc.s, tc.t)
            }
        })

        if isASCII(tc.s) && isASCII(tc.t) {
            b.Run("ASCII-"+tc.name, func(b *testing.B) {
                for i := 0; i < b.N; i++ {
                    isAnagramASCII(tc.s, tc.t)
                }
            })
        }
    }
}
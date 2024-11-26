# Valid Anagram - Problem Solution

## Problem Description

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

### Constraints:

- Both strings can contain any Unicode character
- Case sensitivity matters
- Both strings can be empty

### Examples:

```
Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false
```

## Solution Implementations

### 1. Hash Map Solution (Original)

```go
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
```

**Time Complexity:** O(n) where n is the length of the strings
**Space Complexity:** O(k) where k is the number of unique characters

### 2. Sorting Solution

```go
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
```

**Time Complexity:** O(n log n) due to sorting
**Space Complexity:** O(n) for creating rune slices

### 3. ASCII Optimization (for ASCII-only strings)

```go
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
```

**Time Complexity:** O(n)
**Space Complexity:** O(1) - fixed size array

## Running Tests

### Basic Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v
```

### Run Specific Tests

```bash
# Run main anagram tests
go test -run TestIsAnagram

# Run all variation tests
go test -run TestAllAnagramVariations
```

### Performance Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run benchmarks with memory allocation statistics
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkAnagram
```

## Test Files Structure

```
validanagram/
├── solution.go        # Original solution
├── solution_test.go   # Original tests
├── variations.go      # Additional implementations
└── variations_test.go # Tests for variations
```

## Go-Specific Considerations

1. **String Handling:**

   - Go strings are UTF-8 encoded
   - Use `[]rune` for Unicode support
   - Use byte arrays for ASCII-only optimization

2. **Performance Tips:**
   - Preallocate maps when size is known
   - Use arrays instead of maps for ASCII
   - Consider memory vs CPU trade-offs

## Edge Cases Covered

- Empty strings
- Different lengths
- Unicode characters
- ASCII-only strings
- Repeated characters
- Same characters different order
- Special characters and emojis

## Interview Tips

1. Always check string lengths first
2. Discuss Unicode vs ASCII trade-offs
3. Consider memory usage for different approaches
4. Explain why certain optimizations work

## Common Follow-up Questions

1. How would you handle very large strings?
2. What if we need case-insensitive comparison?
3. How would you optimize for ASCII-only strings?
4. What about concurrent processing for very large strings?

## Related Problems

- Group Anagrams
- Valid Palindrome
- Find All Anagrams in a String

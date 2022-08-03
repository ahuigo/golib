package main

// https://play.golang.org/p/Ega9qgD4Qz
import (
	"fmt"
	"path"
	"unicode/utf8"
)

func main() {
	// path.Match stops at a '/', whereas wildcard.Match considers the path as a flat namespace.
	pattern := "/*/*/file/**"
	text := "/ahui/go/file/abc/bcd"

	doesMatch := Match(pattern, text)
	fmt.Printf("=====Wildcard package match===== \nPattern:   \"%s\" \nText:      \"%s\" \nDoesMatch: %v", pattern, text, doesMatch)

	doesMatch, _ = path.Match(pattern, text)
	fmt.Printf("\n\n=====Path package match===== \nPattern:   \"%s\" \nText:      \"%s\" \nDoesMatch: %v", pattern, text, doesMatch)
}

func Match(pattern, name string) (matched bool) {
Pattern:
	for len(pattern) > 0 {
		var star bool
		var chunk string
		star, chunk, pattern = scanChunk(pattern)
		// fmt.Printf("star=%v,chunk=%v,pattern=%v\n", star, chunk, pattern)
		if star && chunk == "" {
			// Trailing * matches rest of string.
			return true
		}
		// Look for match at current position.
		t, ok := matchChunk(chunk, name)
		// if we're the last chunk, make sure we've exhausted the name
		// otherwise we'll give a false result even if we could still match
		// using the star
		if ok && (len(t) == 0 || len(pattern) > 0) {
			name = t
			continue
		}
		if star {
			// Look for match skipping i+1 bytes.
			for i := 0; i < len(name); i++ {
				t, ok := matchChunk(chunk, name[i+1:])
				if ok {
					// if we're the last chunk, make sure we exhausted the name
					if len(pattern) == 0 && len(t) > 0 {
						continue
					}
					name = t
					continue Pattern
				}
			}
		}
		return false
	}
	return len(name) == 0
}

// scanChunk gets the next segment of pattern, which is a non-star string
// possibly preceded by a star.
func scanChunk(pattern string) (star bool, chunk, rest string) {
	for len(pattern) > 0 && pattern[0] == '*' {
		pattern = pattern[1:]
		star = true
	}
	inrange := false
	var i int
Scan:
	for i = 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '*':
			if !inrange {
				break Scan
			}
		}
	}
	return star, pattern[0:i], pattern[i:]
}

// matchChunk checks whether chunk matches the beginning of s.
// If so, it returns the remainder of s (after the match).
// Chunk is all single-character operators: literals, char classes, and ?.
func matchChunk(chunk, s string) (rest string, ok bool) {
	for len(chunk) > 0 {
		if len(s) == 0 {
			return
		}
		switch chunk[0] {
		case '?':
			_, n := utf8.DecodeRuneInString(s)
			s = s[n:]
			chunk = chunk[1:]
		default:
			if chunk[0] != s[0] {
				return
			}
			s = s[1:]
			chunk = chunk[1:]
		}
	}
	return s, true
}

package main

/*
  Given a list of unique words, return all the pairs of the distinct indices (i, j) in the given list, so that the concatenation of the two words words[i] + words[j] is a palindrome.

  Example 1:
    Input: words = ["abcd","dcba","lls","s","sssll"]
    Output: [[0,1],[1,0],[3,2],[2,4]]
    Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]

  Example 2:
    Input: words = ["bat","tab","cat"]
    Output: [[0,1],[1,0]]
    Explanation: The palindromes are ["battab","tabbat"]

  Example 3:
    Input: words = ["a",""]
    Output: [[0,1],[1,0]]

  Constraints:
    1. 1 <= words.length <= 5000
    2. 0 <= words[i].length <= 300
    3. words[i] consists of lower-case English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/palindrome-pairs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie
func palindromePairs(words []string) [][]int {
	trie := NewTrie()
	for i, word := range words {
		trie.Insert(word, i)
	}

	out := [][]int{}
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			if isPalindrome(word[j:]) {
				if k := trie.Search(reverse(word[:j])); k != -1 && k != i {
					out = append(out, []int{i, k})
				}
			}
			if j > 0 && isPalindrome(word[:j]) {
				if k := trie.Search(reverse(word[j:])); k != -1 {
					out = append(out, []int{k, i})
				}
			}
		}
	}
	return out
}

func reverse(word string) string {
	chs := []byte(word)
	n := len(word)
	l, r := 0, n-1
	for l < r {
		chs[l], chs[r] = chs[r], chs[l]
		l++
		r--
	}
	return string(chs)
}

func isPalindrome(str string) bool {
	n := len(str)
	l, r := 0, n-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

type Trie struct {
	child [26]*Trie
	index int
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
		index: -1,
	}
}

func (t *Trie) Insert(word string, index int) {
	n := t
	for i := range word {
		idx := int(word[i] - 'a')
		if n.child[idx] == nil {
			n.child[idx] = NewTrie()
		}
		n = n.child[idx]
	}
	n.index = index
}

func (t *Trie) Search(word string) int {
	n := t
	for i := range word {
		idx := int(word[i] - 'a')
		if n.child[idx] == nil {
			return -1
		}
		n = n.child[idx]
	}
	return n.index
}

// Hash
func palindromePairs2(words []string) [][]int {
	revWord := func(word string) string {
		bytes := []byte(word)
		for i := 0; i < len(bytes)>>1; i++ {
			bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
		}
		return string(bytes)
	}

	isPalindrome := func(s string) bool {
		for i := 0; i < len(s)>>1; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}

	dict := make(map[string]int)
	for i, word := range words {
		dict[word] = i
	}

	var out [][]int
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			if isPalindrome(word[j:]) {
				if index, ok := dict[revWord(word[:j])]; ok && index != i {
					out = append(out, []int{i, index})
				}
			}
			if j != 0 && isPalindrome(word[:j]) {
				if index, ok := dict[revWord(word[j:])]; ok && index != i {
					out = append(out, []int{index, i})
				}
			}
		}
	}
	return out
}

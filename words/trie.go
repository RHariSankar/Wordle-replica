package words

import (
	"backend/resources"
	"log"
	"strings"
	"sync"
)

type trieNode struct {
	char      string
	letters   map[string]*trieNode
	isWordEnd bool
	count     int
}

type trie struct {
	head *trieNode
}

func (t *trie) insert(word string) {
	t.head.insert(word)
}

func (t *trie) Find(word string) bool {
	return t.head.find(word)
}

func (t *trie) PrintTrie() {
	t.head.printTrieNode(0)
}

func (t *trieNode) insert(word string) {
	wordLength := len(word)
	if wordLength == 0 {
		return
	}
	firstChar := string(word[0])
	node, ok := t.letters[firstChar]
	if ok {
		if wordLength == 1 {
			node.isWordEnd = true
			node.count++
			return
		}
		node.count++
		node.insert(word[1:wordLength])
	} else {
		newTrieNode := &trieNode{
			char:      firstChar,
			letters:   make(map[string]*trieNode),
			isWordEnd: false,
			count:     1,
		}
		t.letters[firstChar] = newTrieNode
		if wordLength == 1 {
			newTrieNode.isWordEnd = true
			return
		}
		newTrieNode.insert(word[1:wordLength])
	}
}

func (t *trieNode) find(word string) bool {
	wordlength := len(word)
	if wordlength == 0 {
		return t.isWordEnd
	}
	firstChar := string(word[0])
	node, ok := t.letters[firstChar]
	if !ok {
		return false
	}
	return node.find(word[1:wordlength])
}

func (t *trieNode) printTrieNode(level int) {
	if t == nil {
		return
	}
	for k, e := range t.letters {
		log.Println(strings.Repeat("_", level), k, e.isWordEnd, e.count)
		e.printTrieNode(level + 1)
	}
}

var trieInstance *trie

var lock = &sync.Mutex{}

// returns the singleton trie instance.
func GetInstance() *trie {

	if trieInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if trieInstance == nil {
			words := resources.AllWords
			trieInstance = &trie{
				head: &trieNode{
					letters:   make(map[string]*trieNode),
					isWordEnd: false,
					count:     0,
				},
			}
			for i := 0; i < len(words); i++ {
				trieInstance.insert(words[i])
			}
		}
	}

	return trieInstance
}

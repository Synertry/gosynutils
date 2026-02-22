/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package datastruct

import "strings"

// AlphabetSize is the number of possible characters in the trie
// we need this to keep the number of children per node low and constant
const AlphabetSize = 26

// TrieNode represents a node in the trie
type TrieNode struct {
	children [AlphabetSize]*TrieNode
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node.
// It is a data structure for storing dictionary with strings.
// Source: https://youtu.be/H-6-8_p88r0 (JamieGo) // heavily modified
type Trie struct {
	root *TrieNode
}

// NewTrie will create a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			// children: [AlphabetSize]*TrieNode(make([]*TrieNode, AlphabetSize)),
		},
	}
}

// Add will take in a word and add it to the trie.
// This method is idempotent, meaning if the word already exists, there will be no change.
func (t *Trie) Add(word string) {
	// word = strings.ToLower(strings.TrimSpace(word)) // Normalize the word by trimming spaces and converting to lowercase

	nodeCurrent := t.root
	for _, r := range word {
		charIndex := r - 'a' // Assuming lowercase letters a-z
		if nodeCurrent.children[charIndex] == nil {
			nodeCurrent.children[charIndex] = &TrieNode{
				// children: [AlphabetSize]*TrieNode(make([]*TrieNode, AlphabetSize)),
			} // Create a new node if it doesn't exist
		}
		nodeCurrent = nodeCurrent.children[charIndex]
	}

	nodeCurrent.isEnd = true // Mark the end of the word
}

// SafeAdd builds upon the Add method, but additionally trims and lowercases the input words, as well as checks for empty strings.
func (t *Trie) SafeAdd(word string) {
	// Normalize the word by trimming spaces and converting to lowercase
	word = strings.ToLower(strings.TrimSpace(word))
	if word == "" {
		return // Do not add empty strings
	}

	t.Add(word) // Call the Add method to add the word to the trie
}

// Find will take in a word and RETURN true if that word is included in the trie
// same walking logic as Add, but we don't need to create nodes
func (t *Trie) Find(word string) bool {
	// word = strings.ToLower(strings.TrimSpace(word))

	nodeCurrent := t.root
	for _, r := range word {
		nodeCurrent = nodeCurrent.children[r-'a'] // Assuming lowercase letters a-z
		if nodeCurrent == nil {
			return false // character is not found, word does not exist
		}
	}

	return nodeCurrent.isEnd // Return true if we reached the end of a word
}

// SafeFind builds upon the Find method, but additionally trims and lowercases the input words, as well as checks for empty strings.
func (t *Trie) SafeFind(word string) bool {
	// Normalize the word by trimming spaces and converting to lowercase
	word = strings.ToLower(strings.TrimSpace(word))
	if word == "" {
		return false // Do not search for empty strings
	}

	return t.Find(word) // Call the Find method to check if the word exists in the trie
}

// Delete will take in a word and remove it from the trie.
// This method will not panic if the word does not exist, it will simply do nothing.
// If any nodes become seperated from the trie after deletion, they will be garbage collected.
func (t *Trie) Delete(word string) {
	// word = strings.ToLower(strings.TrimSpace(word))

	nodeCurrent := t.root
	var nodesStack []*TrieNode // Stack to keep track of nodes for potential deletion

	for _, r := range word {
		charIndex := r - 'a' // Assuming lowercase letters a-z
		if nodeCurrent.children[charIndex] == nil {
			return // Word does not exist, nothing to delete
		}
		nodesStack = append(nodesStack, nodeCurrent) // Push current node onto the stack
		nodeCurrent = nodeCurrent.children[charIndex]
	}

	if !nodeCurrent.isEnd {
		return // Word does not exist, nothing to delete
	}

	nodeCurrent.isEnd = false // Mark the end of the word as false

	// Clean up nodes if they are no longer needed
	for i := len(nodesStack) - 1; i >= 0; i-- {
		if nodeCurrent.isEnd || !t.isLeaf(nodeCurrent) {
			break // If the current node has children, we cannot delete it
		} else {
			nodesStack[i].children[word[i]-'a'] = nil // Remove the child node
			nodeCurrent = nodesStack[i]               // Move up the stack
		}
	}
}

// isLeaf checks if the current node has no children
func (t *Trie) isLeaf(nodeCurrent *TrieNode) bool {
	for _, child := range nodeCurrent.children {
		if child != nil { // If any child exists, it's not the end and therefore not a leaf node
			return false
		}
	}
	return true
}

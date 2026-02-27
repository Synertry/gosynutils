module bench

go 1.26.0

require (
	github.com/Synertry/gosynutils v0.0.4
	github.com/dghubble/trie v0.1.0
	github.com/hashicorp/go-immutable-radix v1.3.1
	github.com/sauerbraten/radix v0.0.0-20150210222551-4445e9cd8982
)

require github.com/hashicorp/golang-lru v1.0.2 // indirect

replace gosynutils => ../../.

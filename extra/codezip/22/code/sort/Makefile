.SUFFIXES:

.SUFFIXES: .txt .go

%.txt: %.go
	go run $< > $@ 2>&1

all: $(addsuffix .txt, $(basename $(wildcard *.go)))

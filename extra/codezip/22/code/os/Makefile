.SUFFIXES:

.SUFFIXES: .txt .go

inspecting: inspecting.go
	-sudo rm inspecting
	go build inspecting.go
	sudo chown root inspecting
	sudo chmod u+s inspecting

inspecting.txt: inspecting
	./inspecting 2>&1 | tee inspecting.txt

%.txt: %.go
	go run $< > $@ 2>&1

all: $(addsuffix .txt, $(basename $(wildcard *.go)))

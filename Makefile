ORG_NAME = kelsonic-networks
TOOL_NAME = kelca

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -rf ${TOOL_NAME}

.PHONY : build
build:
	go build -o $(TOOL_NAME) ./
	
	src



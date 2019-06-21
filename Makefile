GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
NAME=action-cable-cli
OUTPUT ?= dist/$(NAME)

build:
		$(GOBUILD) -o $(OUTPUT) cmd/action-cable-cli/main.go
deps:
		$(GOGET) github.com/gorilla/websocket
		$(GOGET) github.com/rivo/tview
run:
		$(GOBUILD) -o $(OUTPUT) cmd/action-cable-cli/main.go
		ACTION_CABLE_URL=wss://remy-ws.glitch.me/ ./$(OUTPUT)

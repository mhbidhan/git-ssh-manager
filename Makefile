.PHONY: build install

build:
	go build -o bin/gsm ./src
	@echo "✅ Built binary moved to ./bin/gsm"

install: build
	@echo "📦 Installing gsm to ~/bin..."
	mkdir -p ~/bin
	sudo cp bin/gsm $(HOME)/.local/share/bin
	rm -rf ./bin
	@echo "✅ Installation complete"

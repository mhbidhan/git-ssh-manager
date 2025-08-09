build:
	go build -o bin/gsm ./src
	@echo "✅ Built binary moved to ./bin/gsm"

install: build
	@echo "📦 Installing gsm to ~/bin..."
	cp bin/gsm ~/.local/bin/
	chmod +x ~/.local/bin/gsm
	@echo "✅ Installation complete"

export_zip: build
	@echo "Exporting zip file..."
	cp bin/gsm ./
	cp ./src/build_utils/unix/install.sh ./
	zip -r gsm.zip ./gsm ./install.sh
	cp ./gsm.zip ../
	rm -rf ./bin
	rm -rf ./gsm
	rm -rf ./install.sh
	@echo "✅ Installation complete"

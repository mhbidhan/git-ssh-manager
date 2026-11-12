BINARY := gsm
BUILD_DIR := bin

build:
	go build -o $(BUILD_DIR)/$(BINARY) ./cmd
	@echo "✅ Built binary at ./$(BUILD_DIR)/$(BINARY)"

install: build
	@echo "📦 Installing $(BINARY) to ~/bin..."
	mkdir -p ~/bin
	cp $(BUILD_DIR)/$(BINARY) ~/bin/$(BINARY)
	@echo "✅ Installation complete"

clean:
	rm -rf $(BUILD_DIR) ./$(BINARY) ./gsm.zip
	@echo "🧹 Cleaned build artifacts"

export_zip: build
	@echo "Exporting zip file..."
	cp $(BUILD_DIR)/$(BINARY) ./$(BINARY)
	zip -r gsm.zip ./$(BINARY)
	cp ./gsm.zip ../
	rm -f ./$(BINARY) ./gsm.zip
	@echo "✅ Export complete"

# Run using project root (bin/..) as a temporary config location for testing
test-run: build
	GSM_HOME="$(CURDIR)" ./$(BUILD_DIR)/$(BINARY) $(ARGS)
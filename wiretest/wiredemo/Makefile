.PHONY: all
all: clean aicent kitexgen wiregen swag

.PHONY: clean
clean: 
	@echo "=clean="
	@echo "do shome things"

.PHONY: swag
swag:
	@echo "=swag="
	@swag init  -g cmd/server/main.go -o ./docs --parseDependency

.PHONY: aicent
aicent:
	@echo "=aicent="
	@aic-ent dao --yamlpath internal/repository/idl/ --model wiredemo/internal --outpath internal/repository/

.PHONY: kitexgen
kitexgen:
	@echo "=kitexgen="
	@sh kitex_gen.sh

.PHONY: wiregen
wiregen:
	@echo "=wiregen="
	@wire ./...
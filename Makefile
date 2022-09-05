MINT_CONFIG=mint.json
EXPERIMENT=$(EXPERIMENT)

worker:
	go run $(EXPERIMENT)/worker/main.go

starter:
	go run $(EXPERIMENT)/starter/main.go $(MINT_CONFIG)
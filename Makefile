all: $(APP) $(PUB)

$(APP):
	go build -o $@ main.go

$(PUB):
	go build -o $@ Publisher/publisher.go

run: $(APP) $(PUB)
	bash run.sh
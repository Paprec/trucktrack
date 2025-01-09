build: 
	docker build --no-cache --tag=trucktrack/led -f docker/dockerfile/Dockerfile .

run:
	docker compose -f docker/docker-compose.yml up -d

# run : 
# 	go run cmd/main.go
build: 
	docker build --no-cache --tag=trucktrack/led -f docker/dockerfile/Dockerfile .

run:
	docker compose -f docker/docker-compose.yml up -d

down:
	docker compose -f docker/docker-compose.yml down

logs:
	docker logs trucktrack-led
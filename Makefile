# build: 
# 	docker build --no-cache --tag=trucktrack/led -f docker/dockerfile/Dockerfile .

# # run:
# # 	docker run --name trucktrack-led

run : 
	go run cmd/main.go
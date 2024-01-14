.PHONY: build build-frontend build-backend

build:
		echo "Building..." && \
    make -j 2 build-frontend build-backend && \
		echo "Finished building"

start:
		echo "Starting..." && \
		make -j 2 start-frontend start-backend && \
		echo "Finished starting"

build-frontend:
		echo "Building frontend..." && \
		cd ./frontend && \
		npm ci && \
		npm run build && \
		echo "Finished building frontend"

build-backend:
		echo "Building backend..." && \
		cd ./backend && \
		go build -o ./bin/go-todo && \
		echo "Finished building backend"

start-frontend:
		echo "Starting frontend..." && \
		cd ./frontend && \
		npm run preview && \
		echo "Finished starting frontend" &

start-backend:
		echo "Starting backend..." && \
		cd ./backend && \
		./bin/go-todo && \
		echo "Finished starting backend" &
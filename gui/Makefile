init:
	npx create-react-app react-ui
	cd ui && PUBLIC_URL=http://127.0.0.1:9999 npm run build
build:
	go build -o bin/gui cmd/main.go

run: build
	bin/gui --listen 127.0.0.1:9999 --build react-ui/build
init:
	npx create-react-app react-ui --template cra-template-material-ui-only
	cd react-ui && PUBLIC_URL=http://127.0.0.1:9999 npm run build

build:
	cd react-ui && PUBLIC_URL=http://127.0.0.1:9999 npm run build
	go build -o bin/gui cmd/main.go

start-react:
	cd react-ui && PUBLIC_URL=http://127.0.0.1:9999 npm run start

start-server:
	go run cmd/main.go --listen 127.0.0.1:9999 --build react-ui/build

run: build
	bin/gui --listen 127.0.0.1:9999 --build react-ui/build

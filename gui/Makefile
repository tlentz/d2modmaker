init:
	npx create-react-app react-ui --template redux-typescript
	cd react-ui && PUBLIC_URL=http://127.0.0.1:9999 npm run build

build:
	cd react-ui && PUBLIC_URL=http://127.0.0.1:9999 npm run build
	go build -o bin/gui cmd/main.go

build_server:
	go build -o bin/gui cmd/main.go

run: build
	bin/gui --listen 127.0.0.1:9999 --build react-ui/build
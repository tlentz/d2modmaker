init:
	npx create-react-app react-ui --template cra-template-material-ui-only
	cd react-ui && PUBLIC_URL=http://127.0.0.1:8148 npm run build

# build server
bs: 
	go build -o bin/gui cmd/main.go

# start server
ss:
	bin/gui

# build react
br:
	cd react-ui && PUBLIC_URL=http://127.0.0.1:8148 npm run build

# start react for dev mode
sr:
	cd react-ui && PUBLIC_URL=http://127.0.0.1:8148 npm run start

# build react and build server
build: br bs

# run the server
run: ss

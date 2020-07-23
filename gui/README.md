# Go: Create React App

This is a demo project which shows one of possible implementations of integration between the regular server on Golang and 
React application created and built using [create-react-app](https://github.com/facebook/create-react-app).

## How can I build it?

It's supposed here that you have `Golang` and `Node.JS` installed on your computer. We are also using `make` to simplify
build flow. First of all you need to create new React application. It's not delivered as part of the source code to be sure 
latest version of `create-react-app`. 

So, clone the project and create test app:

```shell script
git clone https://github.com/tlentz/d2modmaker/gui.git .
make init
```

Ok, now we have our test project. Let's build our server now:

```shell script
make build
```

## How I can run it?

```shell script
make run
```

Visit the web page:

```shell script
open http://127.0.0.1:9999
```

## License

Please, take a look at the [LICENSE](https://github.com/tlentz/d2modmaker/gui/blob/master/LICENSE) file for the 
details about this aspect of the project.
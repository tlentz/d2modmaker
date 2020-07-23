package client

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

//OpenBrowser - Opens your default browser, depending on the OS you are on.
func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

//LaunchServer starts up the server
func LaunchServer() string {
	port := ":8080"
	fmt.Printf(`
---Lauching---
`)
	fmt.Println("To exit this programme just CTRL-C at console, or exit via the website GUI")
	fmt.Println("Opening browser...")
	OpenBrowser("http://127.0.0.1" + port + "/")
	err := http.ListenAndServe(port, nil) // setting listening port

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened")
	return port
}

var mainpage = `<span>hi</span>`

// func main() {
// 	// jsonFile, _ := ioutil.ReadFile("cfg.json")
// 	// json.Unmarshal([]byte(jsonFile), &N)
// 	http.HandleFunc("/", mainpage)
// 	// http.HandleFunc("/process", process)
// 	// http.HandleFunc("/run", run)
// 	serverPort := front.LaunchServer()
// 	err2 := http.ListenAndServe(serverPort, nil) // setting listening port
// 	if err2 != nil {
// 		fmt.Println(err2)
// 		log.Fatal("ListenAndServe: ", err2)
// 	}
// }

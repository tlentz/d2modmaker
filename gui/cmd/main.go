package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/tlentz/d2modmaker/gui/api"
	"github.com/tlentz/d2modmaker/gui/server"
	webview "github.com/webview/webview"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version = "1.0.0"
var uiURL = "http://localhost:3000/"

func main() {
	cfg := struct {
		BuildPath string
		Listen    string
	}{}

	kp := kingpin.New(filepath.Base(os.Args[0]), "Demo of create-react-app intergration into golang http server")
	kp.Version(version)
	kp.Flag("listen", "Which address should be listened").Required().StringVar(&cfg.Listen)
	kp.Flag("build", "Path to the build directory of the project created using create-react-app").Required().StringVar(&cfg.BuildPath)
	kp.HelpFlag.Short('h')

	if _, err := kp.Parse(os.Args[1:]); err != nil {
		kp.Usage(os.Args[1:])
		os.Exit(1)
	}

	buildPath := path.Clean(cfg.BuildPath)
	staticPath := path.Join(buildPath, "/static/")

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))
	mux.Handle("/api", api.Handler())
	mux.Handle("/", server.Handler(buildPath))

	srv := &http.Server{
		Addr:    cfg.Listen,
		Handler: mux,
	}

	errs := make(chan error, 1)
	go func() {
		fmt.Println("Starting", cfg.Listen)
		errs <- srv.ListenAndServe()
	}()

	launchWebView()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		fmt.Println("Sutting down...")
		os.Exit(0)
	case err := <-errs:
		fmt.Println("Failed to start server:", err.Error())
		os.Exit(1)
	}

	shutdown, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdown); err != nil {
		fmt.Println("Failed to shutdown server:", err.Error())
		os.Exit(1)
	}
}

func launchWebView() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("D2 Mod Maker")
	w.SetSize(960, 540, webview.HintNone)
	w.Navigate(uiURL)
	w.Run()
}

//OpenBrowser - Opens your default browser, depending on the OS you are on.
func openBrowser(url string) {
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

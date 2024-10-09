package main

import (
	"fmt"
	"os"
	"strings"

	webview "github.com/lexesv/go-webview-gui"
)

type WindowSize struct {
	width  int
	height int
}

type App struct {
	window     webview.WebView
	appname    string
	size       WindowSize
	maximized  bool
	url        string
	folderPath string
	debug      bool
}

// Url is prioritized over folderPath
func NewApp(appname string, width int, height int, maximized bool, url string, folderPath string, debug bool) *App {
	size := WindowSize{width, height}
	return &App{webview.New(debug, true), appname, size, maximized, url, folderPath, debug}
}

func (app *App) runWithUrl() {
	defer app.window.Destroy()
	app.window.SetTitle(app.appname)
	app.window.SetSize(app.size.width, app.size.height, webview.HintNone)
	app.window.Navigate(app.url)
	if app.maximized {
		app.window.Maximize()
	}
	app.window.Run()
}

func (app *App) runWithLocalWebsite() {
	html, err := os.ReadFile(app.folderPath + "index.html")
	if err != nil {
		fmt.Println("Error reading >> index.html <<! This file is required.")
		return
	}

	css, _ := os.ReadFile(app.folderPath + "style.css")
	js, _ := os.ReadFile(app.folderPath + "script.js")

	fullcode := string(html) + "<style>" + string(css) + "</style>" + "<script>" + string(js) + "</script>"

	defer app.window.Destroy()
	app.window.SetTitle(app.appname)
	app.window.SetSize(app.size.width, app.size.height, webview.HintNone)
	app.window.SetHtml(fullcode)
	if app.maximized {
		app.window.Maximize()
	}
	app.window.Run()
}

func (app *App) Run() {
	if app.isValidUrl() {
		app.runWithUrl()
	} else {
		app.runWithLocalWebsite()
	}
}

func (app *App) isValidUrl() bool {
	return strings.HasPrefix(app.url, "http://") || strings.HasPrefix(app.url, "https://")
}

func main() {
	app := NewApp("My App", 800, 600, true, "", "./web/", true)
	app.Run()
}

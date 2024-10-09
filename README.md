# WebToDesk

`WebToDesk` ist ein Go-basiertes lightweight Tool, das eine Desktop-Anwendung erstellt, die eine Website oder eine lokale Web-App in einem nativen Fenster anzeigt. Es nutzt das `webview`-Package, um ein einfaches GUI-Fenster zu erstellen, in dem entweder eine externe URL oder lokale HTML-, CSS- und JavaScript-Dateien geladen werden.

## Funktionen

- **Externen Link öffnen**: Lädt eine URL und zeigt sie in einem WebView-Fenster an.
- **Lokale Web-App öffnen**: Lädt eine lokale HTML-Datei mit optionalen CSS- und JavaScript-Dateien.
- **Fenstergrößenanpassung und Maximierung**: Initialisiert das Fenster mit einer benutzerdefinierten Größe und optional als maximiertes Fenster.
- **Debugging-Modus**: Aktiviert den Debugging-Modus für Entwicklungszwecke.

## Voraussetzungen

- **Go** (mindestens Version 1.16)
- Das `webview`-Go-Package. Installiere es mit:
  ```bash
  go get github.com/lexesv/go-webview-gui
  ```

## Installation

1. Klone das Repository:
   ```bash
   git clone https://github.com/dein-benutzername/WebToDesk.git
   cd WebToDesk
   ```
2. Installiere die Abhängigkeiten:
   ```bash
   go mod tidy
   ```

## Verwendung

### Beispielaufruf

Zum Starten der Anwendung öffne das Terminal und führe die Anwendung mit `go run` aus oder erstelle eine ausführbare Datei mit `go build`:

```bash
go run main.go
```

Alternativ kannst du mit `go build` eine ausführbare Datei erstellen und dann die App starten:

```bash
go build -o MyWebApp
./MyWebApp
```

### Hauptoptionen

Die Hauptfunktion `NewApp` nimmt folgende Parameter entgegen:

- **appname**: Der Name der Anwendung, der im Fenster angezeigt wird.
- **width, height**: Breite und Höhe des Fensters in Pixeln.
- **maximized**: Wenn `true`, wird das Fenster maximiert.
- **url**: Die URL, die im Fenster angezeigt werden soll (falls gesetzt, hat diese Vorrang vor `folderPath`).
- **folderPath**: Pfad zum lokalen Ordner, der die Dateien `index.html`, `style.css`, und `script.js` enthalten sollte.
- **debug**: Aktiviert den Debug-Modus, falls `true`.

### Parameterbeispiel

```go
app := NewApp("My App", 800, 600, true, "https://example.com", "./web/", true)
app.Run()
```

- **`"My App"`**: Name der Anwendung.
- **`800, 600`**: Fenstergröße.
- **`true`**: Fenster maximiert.
- **`"https://example.com"`**: Externe URL (priorisiert über `folderPath`).
- **`"./web/"`**: Lokaler Pfad, der `index.html` enthalten sollte.
- **`true`**: Debugging aktiviert.

### Beispiel: URL-basiertes Fenster öffnen

```go
app := NewApp("My Online App", 1024, 768, false, "https://example.com", "", false)
app.Run()
```

### Beispiel: Lokale HTML-Anwendung laden

```go
app := NewApp("My Local App", 1024, 768, true, "", "./localApp/", true)
app.Run()
```

Hier wird eine lokale Anwendung im Ordner `./localApp/` geöffnet. Der Ordner sollte mindestens die Datei `index.html` enthalten. CSS- und JavaScript-Dateien werden optional eingebunden, falls vorhanden.

## Fehlersuche

Falls beim Laden der HTML-Dateien Probleme auftreten, prüfe, ob der Pfad korrekt ist und die benötigten Dateien (`index.html`, `style.css`, `script.js`) vorhanden sind.

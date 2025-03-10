# REST API

## Ziel

Baue eine JSON HTTP API in Go, mit welcher man eine Liste von Bestellungen aufrufen, um Einträge erweitern und aus der man Einträge löschen kann.

### Zusatz

- **Eingabevalidierung**  
  Überprüfe beim Erstellen von Bestellungen, dass die Werte für die Felder korrekt sind. Z.B. darf der Name oder das Produkt nicht leer sein.

- **Erweiterung des Order-Models**  
  Erweiterung des Models um Extras, welche man zu einer Bestellung dazu haben möchte, wie z.B. Getränke. Getränke können wiederum auch in verschiedenen Größen bestellt werden. Eine weitere Erweiterung wäre, ob das Gericht in vegetarischer Variante bestellt werden soll oder nicht. 

- **Produktauswahl**  
  Produkte dürfen nur aus einer Festen Liste ausgewählt weren. Diese Liste sollte über einen Extra Endpunkt abrufbar sein. Bei der Erstellung muss geprüft werden, ob das übergebene Produkt sich in der Liste von Produkten befindet.

- **Export-Endpunk**  
  Es soll einen Endpunkt geben, welcher alle Bestellungen zusammenfast. Also wenn z.B. 3 mal das selbe Produkt in der selben Variante bestellt wurde, dann soll dies in der Ausgabe zusammengefasst werden als "3x Produkt A".

## Ressourcen

### REST API Fundament

#### Videos

- https://youtu.be/xpeQz7Hsfz0
- https://youtu.be/v7h238iam8U

#### Artikel

- https://www.ibm.com/de-de/topics/rest-apis

### Dokumentation

- Allgemeine Referenzen zu HTTP: https://http.dev
- JSON in Go by Example: https://gobyexample.com/json
- Go JSON Package Dokumentation: https://pkg.go.dev/encoding/json
- HTTP Server in Go by Example: https://gobyexample.com/http-servers
- Go HTTP Package Dokumentation: https://pkg.go.dev/net/http

## Tools

- Thunder Client: https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

// Hier speichern wir alle Order-Einträge in einer
// Globalen variable. Somit können alle Funktionen
// auf diese Liste zugreifen.
var orders = []*Order{}

func main() {
	// Wir erstellen einen neuen ServeMux, welcher einen simplen
	// HTTP-Router zur Verfügung stellt.
	mux := http.NewServeMux()

	// Handler-Registierung für das auflisten, erstellen und löschen
	// von Order-Einträgen.
	mux.HandleFunc("GET /orders", getOrders)
	mux.HandleFunc("POST /orders", addOrder)
	mux.HandleFunc("DELETE /orders/{idx}", deleteOrder)

	fmt.Println("Starte server auf http://127.0.0.1:8080 ...")
	// Start des Servers auf Port der Adresse 127.0.0.1:8080.
	http.ListenAndServe("127.0.0.1:8080", mux)
}

func getOrders(resp http.ResponseWriter, req *http.Request) {
	// Wir erstellen einen neuen JSON-Encoder, welcher bein encodieren
	// in die Response schreibt.
	enc := json.NewEncoder(resp)
	// Mit dem Encoder encodieren wir die Liste in als JSON in die Response.
	enc.Encode(orders)
}

func addOrder(resp http.ResponseWriter, req *http.Request) {
	// Wir erstellen einen neuen JSON-Decoder, welcher aus dem Body der
	// Request liest.
	dec := json.NewDecoder(req.Body)

	// Wir erstellen ein neues Order-Objekt, in welches der JSON-Decoder
	// vom Request body hinein-dekodieren kann.
	var newOrder Order
	err := dec.Decode(&newOrder)
	if err != nil {
		// Gibt das Dekodieren einen Fehler zurück, so antworten wir direkt
		// auf die Request mit einem Status Code "Bad Request" (400) und senden
		// die Nachricht des Fehlers als Body in der Response zurück.
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(err.Error()))
		// Danach steigen wir direkt aus der Funktion aus.
		return
	}

	// Wenn alles geklappt hat, dann fügen wir das neue Order-Objekt
	// der Liste von Orders hinzu.
	orders = append(orders, &newOrder)

	// Final können wir nun mit einem Status-Code "OK" (200) antwotten.
	resp.WriteHeader(http.StatusOK)
}

func deleteOrder(resp http.ResponseWriter, req *http.Request) {
	// Wir entnehmen den Index des zu löschenden Elements aus dem Pfad-Parameter
	// der Request.
	// (Oben definiert in DELETE /orders/{idx})
	idxStr := req.PathValue("idx")

	// Den Index in Form eines Strings wandeln wir nun in einen Integer um.
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		// Gibt Umwandeln einen Fehler zurück, so antworten wir direkt
		// auf die Request mit einem Status Code "Bad Request" (400) und senden
		// die Nachricht des Fehlers als Body in der Response zurück.
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(err.Error()))
		// Danach steigen wir direkt aus der Funktion aus.
		return
	}

	if idx < 0 || idx >= len(orders) {
		// Ist der angegebene Index kleiner als 0 oder größer als die Menge der Elemente
		// in der Liste der Orders, so antworten wir direkt auf die Request mit einem
		// Status Code "Bad Request" (400) und senden entsprechend eine eigene Fehler-Nachricht
		// zurück.
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("index must be larger or equal to 0 and smaller than the size of the list"))
		// Danach steigen wir direkt aus der Funktion aus.
		return
	}

	// Wenn alles geplappt hat, dann können wir mit der Hilfsfunktion slices.Delete
	// das gewünschte Element aus der Liste entfernen.
	orders = slices.Delete(orders, idx, idx+1)

	// Final können wir nun mit einem Status-Code "OK" (200) antwotten.
	resp.WriteHeader(http.StatusOK)
}

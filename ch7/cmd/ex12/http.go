// Change the handler for /list to print its output as an HTML table, not text. You may find the
// html/template package (§4.6) useful.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const templ = `<html lang="en">
<head></head>
<body>
	<table>
		<thead>
			<tr>
				<td>Item</td>
				<td>Price</td>
			</tr>
		</thead>
		<tbody>
			{{range $key, $value := .}}
				<tr>
					<td>{{$key}}</td>
					<td>{{$value}}</td>
				</tr>
			{{end}}
		</tbody>
	</table>
</body>
</html>`

var table = template.Must(template.New("table").Parse(templ))

var mu sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 50}
	http.HandleFunc("/", db.index)
	http.HandleFunc("/price", db.show)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type database map[string]dollars

func (db database) index(w http.ResponseWriter, req *http.Request) {
	table.Execute(w, db)
}

func (db database) show(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item, price := req.URL.Query().Get("item"), req.URL.Query().Get("price")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is required\n")
		return
	} else if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item %q already exists\n", item)
		return
	}

	dollarPrice, err := toDollars(price)
	if err != nil {
		priceError(dollarPrice, &w)
		return
	}

	mu.Lock()
	db[item] = dollarPrice
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "created %s with price %s\n", item, dollarPrice)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item, price := req.URL.Query().Get("item"), req.URL.Query().Get("price")
	if _, ok := db[item]; !ok {
		itemNotExistError(item, &w)
		return
	}

	dollarPrice, err := toDollars(price)
	if err != nil {
		priceError(dollarPrice, &w)
		return
	}

	mu.Lock()
	db[item] = dollarPrice
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "updated %s with price %s\n", item, dollarPrice)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		itemNotExistError(item, &w)
		return
	}
	mu.Lock()
	delete(db, item)
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item %q deleted", item)
}

func priceError(price dollars, w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(*w, "invalid price %q\n", price)
}

func itemNotExistError(item string, w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusNotFound)
	fmt.Fprintf(*w, "no such item %q\n", item)
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func toDollars(price string) (dollars, error) {
	floatPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return 0, err
	}
	return dollars(float32(floatPrice)), nil
}

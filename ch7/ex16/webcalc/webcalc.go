// Write a web-based calculator program.
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"the_go_programming_language/ch7/ex16/eval"
)

func main() {
	http.HandleFunc("/", evaluate)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func evaluate(w http.ResponseWriter, r *http.Request) {
	URLExpr := r.FormValue("expr")
	if URLExpr == "" {
		fmt.Fprintf(w, "usage: /?expr=sin(x)+y&x=30&y=12\nDo not include spaces.")
		return
	}
	// + is used to represent spaces in query strings and must be restored
	URLExpr = strings.ReplaceAll(URLExpr, " ", "+")
	expr, vars, err := parseAndCheck(URLExpr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	env, err := parseEnv(vars, r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%.6g", expr.Eval(env))
}

// parseAndCheck parses the URL expression and returns a map of its variables to be populated
func parseAndCheck(URLExpr string) (expr eval.Expr, vars map[eval.Var]bool, err error) {
	vars = make(map[eval.Var]bool)
	expr, err = eval.Parse(URLExpr)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing %q: %v", URLExpr, err)
	}
	expr.Check(vars)
	if err != nil {
		return nil, nil, fmt.Errorf("checking %q: %v", URLExpr, err)
	}
	return
}

// parseEnv populates a map of expression variables with their corresponding values from the query
func parseEnv(vars map[eval.Var]bool, form url.Values) (env eval.Env, err error) {
	env = make(eval.Env)
	for k := range vars {
		def, ok := form[k.String()]
		if !ok {
			return nil, fmt.Errorf("variable %q is not defined", k)
		}
		val, err := strconv.ParseFloat(def[0], 64)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse variable %q: %s", k, def)
		}
		env[k] = val
	}
	return
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"eval"
)

func main() {
	http.HandleFunc("/evaluate", evaluate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func evaluate(w http.ResponseWriter, req *http.Request) {
	query := parseQuery(req.URL.RawQuery)

	e, ok := query["expr"]
	if !ok || e == "" {
		http.Error(w, "input error: no \"expr\" provided", http.StatusBadRequest)
		return
	}

	ex, err := eval.Parse(e)
	if err != nil {
		http.Error(w, fmt.Sprintf("parse error: %s", err), http.StatusBadRequest)
		return
	}

	vars := map[eval.Var]bool{}
	err = ex.Check(vars)
	if err != nil {
		http.Error(w, fmt.Sprintf("check error: %s", err), http.StatusBadRequest)
		return
	}

	env := eval.Env{}
	for v := range vars {
		val, ok := query[string(v)]

		if !ok || val == "" {
			http.Error(w, fmt.Sprintf("input error: no value provided for \"%s\"", v), http.StatusBadRequest)
			return
		}

		pval, err := strconv.ParseFloat(val, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("value parse error: %s\n\n", err), http.StatusBadRequest)
			return
		}

		env[v] = pval
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, ex.Eval(env))
}

func parseQuery(q string) map[string]string {
	parsed := make(map[string]string)

	params := strings.Split(q, "&")
	for _, p := range params {
		elems := strings.Split(p, "=")
		parsed[elems[0]] = elems[1]
	}

	return parsed
}

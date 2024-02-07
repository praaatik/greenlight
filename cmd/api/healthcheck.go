package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=> healthcheckHandler")

	// data := map[string]string{
	// 	"status":      "available",
	// 	"environment": app.config.env,
	// 	"version":     version,
	// }

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		// app.logger.Println(err)
		// http.Error(w, "The server encountered a problem", http.StatusInternalServerError)
		app.serverErrorReponse(w, r, err)
	}

	// js := `{"status": "available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)
	//
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(js))
	//
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)
}

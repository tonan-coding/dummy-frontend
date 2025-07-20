package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v3"
	"net/http"
)

func main() {
	cmd := &cli.Command{
		Name:  "dummy-frontend",
		Usage: "Dummy frontend implementation",
		Flags: []cli.Flag{},
		Action: func(ctx context.Context, command *cli.Command) error {
			r := mux.NewRouter()
			r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("home_page"))
			})
			http.ListenAndServe(":8080", r)
			return nil
		},
	}
	cmd.Run(context.Background(), nil)
}

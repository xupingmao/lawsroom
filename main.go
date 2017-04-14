package main

import (
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	app := cli.NewApp()
	app.Name = "Lawsroom"
	app.Usage = "Law's room."
	app.Author = "Cloud"
	app.Email = "cloud@txthinking.com"
	app.Version = "0.9.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "listen",
			Value: "",
			Usage: "Listen address.",
		},
		cli.StringSliceFlag{
			Name:  "origin",
			Usage: "Allow origins for CORS, can repeat more times.",
		},
	}
	app.Action = func(c *cli.Context) error {
		// SAE的端口号
		port := os.Getenv("PORT")

		if port != "" {
			return run(":" + port, c.GlobalStringSlice("origin"))
		}

		if c.String("listen") == "" {
			return cli.NewExitError("Listen address is empty.", 86)
		}
		return run(c.String("listen"), c.GlobalStringSlice("origin"))
	}
	app.Run(os.Args)
}

func run(listen string, origins []string) error {
	r := mux.NewRouter()
	r.Methods("GET").Path("/signal/_/{id}").Handler(getSignalHandle(origins))
	r.Methods("GET").Path("/random").HandlerFunc(redirect)
	r.Methods("GET").Path("/room/{roomID}").HandlerFunc(redirect)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	}))
	n.Use(negroni.NewStatic(assetFS()))
	n.UseHandler(r)

	s := &http.Server{
		Addr:         listen,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		// IdleTimeout:  120 * time.Second,
		Handler:      n,
	}
	// return s.ListenAndServeTLS("server.crt", "server.key");
	return s.ListenAndServe()
}

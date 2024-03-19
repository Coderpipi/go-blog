package main

import (
	"blog-admin/config"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.InitConfig()

	app := cli.NewApp()
	app.Name = "blog-admin"
	app.Usage = "blog"
	app.Commands = []cli.Command{
		{
			Name:  "api",
			Usage: "start web server",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "migrate",
			Usage: "do db migrate",
			Action: func(c *cli.Context) error {

				return nil
			},
		},
	}
	sigComplete := make(chan struct{})
	go func() {
		defer close(sigComplete)
		err := app.Run(os.Args)
		if err != nil {
			log.Fatal("app run failed, err: ", err.Error())
		}
	}()

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-sigTerm:
		log.Println("receive stop signal")
	case <-sigComplete:
	}
}

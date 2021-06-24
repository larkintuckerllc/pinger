package main

import (
	"log"
	"os"

	"github.com/larkintuckerllc/pinger/internal/pinger"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pinger",
		Usage: "exports ping metrics",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "ip",
				Usage:    "ip address",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "location",
				Usage:    "GKE cluster location",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pod",
				Usage:    "K8s pod name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "project",
				Usage:    "Google project",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			ips := c.StringSlice("ip")
			location := c.String("location")
			pod := c.String("pod")
			project := c.String("project")
			err := pinger.Execute(project, location, pod, ips)
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

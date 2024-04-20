package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ugursogukpinar/go-cybersecurity-tools/portscanner"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "PortScanner",
		Usage: "[hostname] [start] [end]",
		Action: func(ctx *cli.Context) error {
			if ctx.Args().Len() < 3 {
				return errors.New("invalid arguments")
			}

			host := ctx.Args().Get(0)

			startPort, err := strconv.ParseUint(ctx.Args().Get(1), 10, 16)
			if err != nil {
				return err
			}

			endPort, err := strconv.ParseUint(ctx.Args().Get(2), 10, 16)
			if err != nil {
				return err
			}

			fmt.Println(startPort, endPort, host)

			openPorts, err := portscanner.GetOpenPorts(host, uint16(startPort), uint16(endPort))
			if err != nil {
				return err
			}

			fmt.Println(openPorts)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

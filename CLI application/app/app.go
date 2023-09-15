package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Return app cli
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"

	app.Usage = "Busca IPs e Nome de Servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs",
			Flags:  flags,
			Action: buscarIPS,
		},
		{
			Name:   "servidores",
			Usage:  "Busca Servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIPS(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

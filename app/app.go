package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet!"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na Internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "nameservers",
			Usage:  "Busca Nomes de Servidor na Internet",
			Flags:  flags,
			Action: buscarNameServers,
		},
	}

	return app

}

func buscarIps(c *cli.Context) {

	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Printf("IPs do dominio %s:\n\n", host)

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarNameServers(c *cli.Context) {

	host := c.String("host")

	nms, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Printf("NameServers do dominio %s:\n\n", host)

	for _, ns := range nms {
		fmt.Println(ns.Host)
	}

}
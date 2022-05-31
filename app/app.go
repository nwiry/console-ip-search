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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na Internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: buscarIps,
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

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

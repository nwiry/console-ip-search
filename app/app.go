package app

import "github.com/urfave/cli"

func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet!"
	return app
	
}

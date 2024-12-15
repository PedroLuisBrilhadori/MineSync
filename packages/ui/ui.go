package ui

import (
	"minesync/packages/compress"
	"minesync/packages/finder"
)


func ShowServerList() {
	finder.FindServersFiles();
	// mostrar servidores para o usuário
}

func CompressServerFolder() {
	compress.CompressFiles();
}
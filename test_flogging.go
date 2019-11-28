package main

import (
	"os"
	"github.com/hyperledger/fabric/common/flogging"
)

var logger *flogging.FabricLogger

func main() {
	fd, _ := os.Create("log01.log")
	logging, err := flogging.New(flogging.Config{LogSpec:"info", Writer: fd,
		Format:"%{color}%{time:2006-01-02 15:04:05.000 MST}%{color} [%{module}][%{level:.4s}] [%{shortfunc}] \"%{message}\""})
	if err != nil {
		panic(err)
	}

	logger = logging.Logger("flogging")
	logger.Info("test info logger")
	logger.Error("test error logger")
	logger.Debug("test debug logger")
	logger.Error("test error logger 2 ")
	logger.Warn("test error logger 2 ")

	Get()
}

func Get() {
	logger.Info("test Get")
}

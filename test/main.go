package main

import (
	"github.com/SilenMark/log"
	l "log"
)

func main() {
	log.Info("test Info", "test Info")
	log.Infof("Info is: %v", log.BLUE)
	log.Infoln("test Infoln", "test Info")
	log.Warn("test Warn", "test Warn")
	log.Warnf("Warn is: %v", log.YELLOW)
	log.Warnln("test Warnln", "test Warn")
	log.Erro("test Erro", "test Erro")
	log.Errof("Erro is: %v", log.RED)
	log.Erroln("test Erroln", "test Erro")
	log.Succ("test Succ", "test Succ")
	log.Succf("Succ is: %v", log.GREEN)
	log.Succln("test Succln", "test Succ")
	l.Print("aaa", "bbb")
	l.Println("aaa", "bbb")

}

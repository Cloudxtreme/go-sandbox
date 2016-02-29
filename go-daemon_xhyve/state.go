// https://github.com/sevlyar/go-daemon/blob/master/sample/sample.go
package xhyve

import (
	"log"
	"os"
	"syscall"
	"time"

	"github.com/sevlyar/go-daemon"
)

func Daemon(sig os.Signal) error {
	// signal.SIGCHLD
	// signal.SIGINT
	// signal.SIGKILL
	// signal.SIGTRAP
	// signal.SIGQUIT

	ctx := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon goxhyve]"},
	}

	if sig == nil {
		d, err := ctx.Search()
		if err != nil {
			log.Fatalln("Unable send signal to the daemon:", err)
			return err
		}
		daemon.SendCommands(d)
	} else {
		daemon.SetSigHandler(killHandler, sig)
	}

	_, err := ctx.Reborn()
	if err != nil {
		return err
	}
	defer ctx.Release()

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	go worker()

	err = daemon.ServeSignals()
	if err != nil {
		return err
	}
	log.Println("daemon terminated")
	return nil
}

var (
	stop = make(chan struct{})
	done = make(chan struct{})
)

func worker() {
	for {
		time.Sleep(time.Second)
		if _, ok := <-stop; ok {
			break
		}
	}
	done <- struct{}{}
}

func killHandler(sig os.Signal) error {
	log.Println("terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}

// func reloadHandler(sig os.Signal) error {
// 	log.Println("configuration reloaded")
// 	return nil
// }

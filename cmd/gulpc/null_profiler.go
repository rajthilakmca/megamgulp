/*
** Copyright [2013-2015] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
*/
// +build !linux !profile

package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	log "code.google.com/p/log4go"
)

func startProfiler(stoppable Stoppable) error {
	go waitForSignals(stoppable)
	return nil
}

func waitForSignals(stoppable Stoppable) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-ch
		log.Info("Received signal: %s", sig.String())
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			stoppable.Stop()
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}
}

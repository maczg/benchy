package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var loadCmd = cobra.Command{
	Use:   "load",
	Short: "Start load test server",
	Long:  `Start load test server"`,
	Run: func(cmd *cobra.Command, args []string) {

		wg := sync.WaitGroup{}
		nreq := uint64(0)

		log.Printf("Starting load test")

		for {
			wg.Add(1)

			go func() {
				defer wg.Done()
				atomic.AddUint64(&nreq, 1)
				res, err := http.Get("http://benchy-benchy.apps.okd.cloudnative.lab/cpuintensive?n=30")
				if err != nil {
					log.Printf("[Req#%d] Error: %s", nreq, err)
					return
				}
				defer res.Body.Close()
				log.Printf("[Req#%d] being processed", nreq)
			}()

			time.Sleep(time.Duration(1) * time.Second)
		}

		wg.Wait()

	},
}

func init() {

}

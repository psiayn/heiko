package cmd

import (
	"fmt"
	"sync"
	"github.com/spf13/cobra"
	"github.com/psiayn/heiko/internal/scheduler"
	"github.com/psiayn/heiko/internal/config"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new heiko job",
	Run: func(cmd *cobra.Command, args []string) {
		task_arr := configuration.Jobs
		nodes := configuration.Nodes

		fmt.Println("len of nodes = ", len(task_arr))
		tasks := make(chan config.Task)

		var wg sync.WaitGroup
		wg.Add(len(task_arr))
		go scheduler.RandomScheduler(tasks, nodes, &wg)
		for _, task := range task_arr {
			tasks <- task
		}
		wg.Wait()

	},
}

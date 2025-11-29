package cmd

import (
	"fmt"
	"os"

	"github.com/dandimuzaki/project-app-task-list-cli-nama/cmd/dto"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/cmd/handler"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/service"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/utils"
	"github.com/spf13/cobra"
)

var (
	activity string
	status string
	priority string
	id uint

	rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "Apikasi To-Do List Sederhana dengan Golang CLI",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	service := service.NewTaskService()
	handler := handler.NewTaskHandler(service)

	var listCmd = &cobra.Command{
		Use: "list",
		Short: "View list of tasks",
		Run: func(cmd *cobra.Command, args []string) {
			res, msg := handler.GetAllTask()
			if msg != "" {
				fmt.Println(msg)
			} else {
				utils.Table(res)
			}
		},
	}

	rootCmd.AddCommand(listCmd)

	var viewCmd = &cobra.Command{
		Use: "view",
		Short: "View a task",
		Run: func(cmd *cobra.Command, args []string) {
			res, msg := handler.GetTaskByID(id)
			if msg != "" {
				fmt.Println(msg)
			} else {
				utils.Card(res)
			}
		},
	}

	viewCmd.Flags().UintVarP(&id, "id", "i", 0, "task id")
	rootCmd.AddCommand(viewCmd)

	var createCmd = &cobra.Command{
		Use: "create",
		Short: "Create a new task",
		Run: func(cmd *cobra.Command, args []string) {
			req := dto.CreateTaskRequest{activity, priority}
			res := handler.CreateTask(req)
			fmt.Println(res.Message)
		},
	}

	createCmd.Flags().StringVarP(&activity, "activity", "a", "", "task activity")
	createCmd.Flags().StringVarP(&priority, "priority", "p", "", "task priority")
	rootCmd.AddCommand(createCmd)

	var updateCmd = &cobra.Command{
		Use: "update",
		Short: "Update activity, status, or priority of a task",
		Run: func(cmd *cobra.Command, args []string) {
			req := dto.UpdateTaskRequest{activity, status, priority}
			res := handler.UpdateTask(id, req)
			fmt.Println(res.Message)
		},
	}

	updateCmd.Flags().UintVarP(&id, "id", "i", 0, "task id")
	updateCmd.Flags().StringVarP(&activity, "activity", "a", "", "task activity")
	updateCmd.Flags().StringVarP(&status, "status", "s", "", "task status")
	updateCmd.Flags().StringVarP(&priority, "priority", "p", "", "task priority")
	rootCmd.AddCommand(updateCmd)

	var finishCmd = &cobra.Command{
		Use: "finish",
		Short: "Finish a task",
		Run: func(cmd *cobra.Command, args []string) {
			req := dto.UpdateTaskRequest{"", "Finished", ""}
			res := handler.UpdateTask(id, req)
			fmt.Println(res.Message)
		},
	}

	finishCmd.Flags().UintVarP(&id, "id", "i", 0, "task id")
	rootCmd.AddCommand(finishCmd)

	var deleteCmd = &cobra.Command{
		Use: "delete",
		Short: "Delete a task",
		Run: func(cmd *cobra.Command, args []string) {
			res := handler.DeleteTask(id)
			fmt.Println(res.Message)
		},
	}

	deleteCmd.Flags().UintVarP(&id, "id", "i", 0, "task id")
	rootCmd.AddCommand(deleteCmd)
}
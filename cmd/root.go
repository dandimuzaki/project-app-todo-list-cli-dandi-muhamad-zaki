package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/dandimuzaki/project-app-task-list-cli-nama/cmd/dto"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/cmd/handler"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/service"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/utils"
	"github.com/spf13/cobra"
)

var (
	keyword string
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
			query := dto.NewQueryRequest("", "", "")
			res, msg := handler.GetAllTask(*query)
			if msg != "" {
				fmt.Println(msg)
			} else {
				utils.Table(res)
			}
		},
	}

	rootCmd.AddCommand(listCmd)

	var searchCmd = &cobra.Command{
		Use: "search",
		Short: "Search tasks with keyword based on task activity",
		Run: func(cmd *cobra.Command, args []string) {
			if keyword == "" {
				fmt.Println(utils.ErrorMessage(errors.New("keyword is required")))
				return
			}

			query := dto.NewQueryRequest(keyword, "", "")
			res, msg := handler.GetAllTask(*query)
			if msg != "" {
				fmt.Println(msg)
			} else if len(res) == 0 {
				fmt.Println("\033[33mNo record is found\033[0m")
			} else {
				fmt.Printf("\033[32mSearch result for %v\033[0m\n", keyword)
				utils.Table(res)
			}
		},
	}

	searchCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "keyword for task activity")
	rootCmd.AddCommand(searchCmd)

	var filterCmd = &cobra.Command{
		Use: "filter",
		Short: "Filter tasks by status or priority",
		Run: func(cmd *cobra.Command, args []string) {
			if status != "" && priority != "" {
				fmt.Println(utils.ErrorMessage(errors.New("only filter by one field")))
				return
			}

			if status != "" && !utils.IsValidStatus(status) {
				fmt.Println(utils.ErrorMessage(utils.ErrInvalidStatus))
				return
			}

			if priority != "" && !utils.IsValidPriority(priority) {
				fmt.Println(utils.ErrorMessage(utils.ErrInvalidPriority))
				return
			}

			query := dto.NewQueryRequest("", status, priority)
			res, msg := handler.GetAllTask(*query)
			if msg != "" {
				fmt.Println(msg)
			} else if status != "" {
				fmt.Printf("\033[32mFilter tasks by status: %v\033[0m\n", utils.Uppercase(status))
				utils.Table(res)
			} else if priority != "" {
				fmt.Printf("\033[32mFilter tasks by priority: %v\033[0m\n", utils.Uppercase(priority))
				utils.Table(res)
			}
		},
	}

	filterCmd.Flags().StringVarP(&status, "status", "s", "", "task status")
	filterCmd.Flags().StringVarP(&priority, "priority", "p", "", "task priority")
	rootCmd.AddCommand(filterCmd)

	var viewCmd = &cobra.Command{
		Use: "view",
		Short: "View a task by ID",
		Run: func(cmd *cobra.Command, args []string) {
			if id<=0 {
				fmt.Println(utils.ErrorMessage(errors.New("id must be greater than 0 and is required")))
				return
			}

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
			if priority == "" && activity == "" {
				fmt.Println(utils.ErrorMessage(errors.New("task activity and priority is required")))
				return
			}

			if activity == "" {
				fmt.Println(utils.ErrorMessage(errors.New("task activity is required")))
				return
			}

			if priority == "" {
				fmt.Println(utils.ErrorMessage(errors.New("task priority is required")))
				return
			}

			if priority != "" && !utils.IsValidPriority(priority) {
				fmt.Println(utils.ErrorMessage(utils.ErrInvalidPriority))
				return
			}

			req := dto.NewCreateTaskRequest(activity, priority)
			res := handler.CreateTask(*req)
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
			if id<=0 {
				fmt.Println(utils.ErrorMessage(errors.New("id must be greater than 0 and is required")))
				return
			}

			if activity == "" && status == "" && priority == "" {
				fmt.Println(utils.ErrorMessage(errors.New("no field is updated")))
				return
			}

			if status != "" && !utils.IsValidStatus(status) {
				fmt.Println(utils.ErrorMessage(utils.ErrInvalidStatus))
				return
			}

			if priority != "" && !utils.IsValidPriority(priority) {
				fmt.Println(utils.ErrorMessage(utils.ErrInvalidPriority))
				return
			}

			req := dto.NewUpdateTaskRequest(activity, status, priority)
			res := handler.UpdateTask(id, *req)
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
			if id<=0 {
				fmt.Println(utils.ErrorMessage(errors.New("id must be greater than 0 and is required")))
				return
			}

			req := dto.NewUpdateTaskRequest("", "Finished", "")
			res := handler.UpdateTask(id, *req)
			fmt.Println(res.Message)
		},
	}

	finishCmd.Flags().UintVarP(&id, "id", "i", 0, "task id")
	rootCmd.AddCommand(finishCmd)

	var deleteCmd = &cobra.Command{
		Use: "delete",
		Short: "Delete a task",
		Run: func(cmd *cobra.Command, args []string) {
			if id<=0 {
				fmt.Println(utils.ErrorMessage(errors.New("id must be greater than 0 and is required")))
				return
			}

			res := handler.DeleteTask(id)
			fmt.Println(res.Message)
		},
	}

	deleteCmd.Flags().UintVarP(&id, "id", "i", 0, "task id")
	rootCmd.AddCommand(deleteCmd)
}
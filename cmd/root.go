package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	showAll bool
	priority string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}

func taskFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		color.Red("Unable to locate home directory")
		os.Exit(1)
	}
	return filepath.Join(home, ".cli-tasks.json")
}

func loadTasks() ([]Task, error) {
	file := taskFile()

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	if len(data) == 0 {
		return []Task{}, nil
	}

	err = json.Unmarshal(data, &tasks)

	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(taskFile(), data, 0644)
}

func findTask(tasks []Task, id string) int {
	for i, task := range tasks {
		if strings.HasPrefix(task.ID, id) {
			return i
		}
	}
	return -1
}

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "CLI Task Manager",
	Long: color.CyanString(`
╔══════════════════════════════════════╗
║          CLI TASK MANAGER           ║
║        Powered by Cobra + Go        ║
╚══════════════════════════════════════╝

A simple terminal task manager.
`),
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(statsCmd)

	listCmd.Flags().BoolVar(&showAll, "all", false, "Show all tasks")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "medium", "Task priority")
}

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		p := strings.ToLower(priority)

		if p != "low" && p != "medium" && p != "high" {
			color.Red("Priority must be low, medium or high")
			return
		}

		tasks, err := loadTasks()
		if err != nil {
			color.Red(err.Error())
			return
		}

		task := Task{
			ID:        uuid.New().String()[:8],
			Title:     args[0],
			Priority:  p,
			Done:      false,
			CreatedAt: time.Now(),
		}

		tasks = append(tasks, task)

		if err := saveTasks(tasks); err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("✓ Task added [%s]", task.ID)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := loadTasks()
		if err != nil {
			color.Red(err.Error())
			return
		}

		table := tablewriter.NewWriter(os.Stdout)

		table.SetHeader([]string{
			"ID",
			"Title",
			"Priority",
			"Status",
			"Created",
		})

		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgCyanColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgCyanColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgCyanColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgCyanColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgCyanColor, tablewriter.Bold},
		)

		count := 0

		for _, t := range tasks {

			if !showAll && t.Done {
				continue
			}

			status := "Pending"

			if t.Done {
				status = "Done"
			}

			table.Append([]string{
				t.ID,
				t.Title,
				strings.ToUpper(t.Priority),
				status,
				t.CreatedAt.Format("2006-01-02"),
			})

			count++
		}

		if count == 0 {
			color.Yellow("No tasks found")
			return
		}

		table.Render()
	},
}

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark task complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := loadTasks()
		if err != nil {
			color.Red(err.Error())
			return
		}

		idx := findTask(tasks, args[0])

		if idx < 0 {
			color.Red("Task not found")
			return
		}

		tasks[idx].Done = true

		if err := saveTasks(tasks); err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("✓ Task completed")
	},
}

var deleteCmd = &cobra.Command{
	Use:     "delete [id]",
	Aliases: []string{"rm", "del"},
	Short:   "Delete task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := loadTasks()
		if err != nil {
			color.Red(err.Error())
			return
		}

		idx := findTask(tasks, args[0])

		if idx < 0 {
			color.Red("Task not found")
			return
		}

		tasks = append(tasks[:idx], tasks[idx+1:]...)

		if err := saveTasks(tasks); err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("✓ Task deleted")
	},
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove completed tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := loadTasks()
		if err != nil {
			color.Red(err.Error())
			return
		}

		var remaining []Task

		for _, t := range tasks {
			if !t.Done {
				remaining = append(remaining, t)
			}
		}

		if err := saveTasks(remaining); err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("✓ Completed tasks removed")
	},
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show task statistics",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := loadTasks()
		if err != nil {
			color.Red(err.Error())
			return
		}

		total := len(tasks)
		completed := 0
		high := 0

		for _, t := range tasks {
			if t.Done {
				completed++
			}
			if t.Priority == "high" {
				high++
			}
		}

		pending := total - completed

		progress := 0.0
		if total > 0 {
			progress = float64(completed) / float64(total) * 100
		}

		fmt.Println()
		fmt.Println("Total Tasks      :", total)
		fmt.Println("Completed Tasks  :", completed)
		fmt.Println("Pending Tasks    :", pending)
		fmt.Println("High Priority    :", high)
		fmt.Println("Progress         :", strconv.FormatFloat(progress, 'f', 2, 64)+"%")
		fmt.Println()
	},
}

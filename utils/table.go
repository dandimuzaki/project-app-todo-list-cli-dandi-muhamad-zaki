package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/dandimuzaki/project-app-task-list-cli-nama/cmd/dto"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

func Table(data []dto.TaskResponse) {
	var format []dto.TaskResponse
	for _, d := range data {
		switch {
		case strings.ToLower(d.Status) == "finished":
			d.Status = "\033[32m" + d.Status
			format = append(format, d)
		case strings.ToLower(d.Status) == "on progress":
			d.Status = "\033[34m" + d.Status
			format = append(format, d)
		case strings.ToLower(d.Status) == "on hold":
			d.Status = "\033[33m" + d.Status
			format = append(format, d)
		default:
			d.Status = "\033[37m" + d.Status
			format = append(format, d)
		}
	}

	var formatted []dto.TaskResponse
	for _, d := range format {
		switch {
		case strings.ToLower(d.Priority) == "low":
			d.Priority = "\033[34m" + d.Priority
			formatted = append(formatted, d)
		case strings.ToLower(d.Priority) == "normal":
			d.Priority = "\033[32m" + d.Priority
			formatted = append(formatted, d)
		case strings.ToLower(d.Priority) == "urgent":
			d.Priority = "\033[33m" + d.Priority
			formatted = append(formatted, d)
		case strings.ToLower(d.Priority) == "critical":
			d.Priority = "\033[31m" + d.Priority
			formatted = append(formatted, d)
		default:
			d.Priority = "\033[37m" + d.Priority
			formatted = append(formatted, d)
		}
	}

	colorCfg := renderer.ColorizedConfig{
		Header: renderer.Tint{
			FG: renderer.Colors{color.FgGreen, color.Bold}, // Green bold headers
			BG: renderer.Colors{color.BgHiWhite},
		},
		Column: renderer.Tint{
			FG: renderer.Colors{color.FgWhite}, // Default cyan for rows
			Columns: []renderer.Tint{
				{FG: renderer.Colors{color.FgMagenta}}, // Magenta for column 0
				{},                                     // Inherit default (cyan)
			},
		},
		Footer: renderer.Tint{
			FG: renderer.Colors{color.FgYellow, color.Bold}, // Yellow bold footer
			Columns: []renderer.Tint{
				{},                                      // Inherit default
				{FG: renderer.Colors{color.FgHiYellow}}, // High-intensity yellow for column 1
				{},                                      // Inherit default
			},
		},
		Border:    renderer.Tint{FG: renderer.Colors{color.FgWhite}}, // White borders
		Separator: renderer.Tint{FG: renderer.Colors{color.FgWhite}}, // White separators
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewColorized(colorCfg)),
		tablewriter.WithConfig(tablewriter.Config{
			Row: tw.CellConfig{
				Formatting:   tw.CellFormatting{AutoWrap: tw.WrapNormal}, // Wrap long content
				Alignment:    tw.CellAlignment{Global: tw.AlignLeft},     // Left-align rows
				ColMaxWidths: tw.CellWidth{Global: 25},
			},
			Footer: tw.CellConfig{
				Alignment: tw.CellAlignment{PerColumn: []tw.Align{tw.AlignLeft, tw.AlignRight, tw.AlignLeft, tw.AlignLeft}},
			},
		}),
	)
	
	table.Header([]string{"ID", "Activity", "Status", "Priority"})
	table.Bulk(formatted)
	table.Footer([]string{"", "", "Total", fmt.Sprintf("%v", len(format))})
	table.Render()
}

func Card(data dto.TaskResponse) {
	formatted := [][]string{
		{"ID", fmt.Sprintf("%v", data.ID)},
		{"Activity", data.Activity},
		{"Status", data.Status},
		{"Priority", data.Priority},
	}

	// Configure colors: green headers, cyan/magenta rows, yellow footer
	colorCfg := renderer.ColorizedConfig{
		Column: renderer.Tint{
			FG: renderer.Colors{color.FgCyan}, // Default cyan for rows
			Columns: []renderer.Tint{
				{FG: renderer.Colors{color.FgMagenta}}, // Magenta for column 0
				{},                                     // Inherit default (cyan)
				{FG: renderer.Colors{color.FgHiRed}},   // High-intensity red for column 2
			},
		},
		Footer: renderer.Tint{
			FG: renderer.Colors{color.FgYellow, color.Bold}, // Yellow bold footer
			Columns: []renderer.Tint{
				{},                                      // Inherit default
				{FG: renderer.Colors{color.FgHiYellow}}, // High-intensity yellow for column 1
				{},                                      // Inherit default
			},
		},
		Border:    renderer.Tint{FG: renderer.Colors{color.FgWhite}}, // White borders
		Separator: renderer.Tint{FG: renderer.Colors{color.FgWhite}}, // White separators
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewColorized(colorCfg)),
		tablewriter.WithConfig(tablewriter.Config{
			Header: tw.CellConfig{
				Merging:   tw.CellMerging{Mode: tw.MergeHorizontal},
			},
			Row: tw.CellConfig{
				Formatting:   tw.CellFormatting{AutoWrap: tw.WrapNormal}, // Wrap long content
				Alignment:    tw.CellAlignment{Global: tw.AlignLeft},     // Left-align rows
				ColMaxWidths: tw.CellWidth{Global: 25},
			},
			Footer: tw.CellConfig{
				Alignment: tw.CellAlignment{Global: tw.AlignRight},
			},
		}),
	)
	table.Bulk(formatted)
	table.Render()
}
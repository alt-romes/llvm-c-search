package main

import (
	"fmt"
	"os"
    "errors"
    "encoding/gob"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	Titl, Desc string
}

func (i item) Title() string       { return i.Titl }
func (i item) Description() string { return i.Desc }
func (i item) FilterValue() string { return i.Titl /* TODO: add Desc */ }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, nil
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {

    ssrpath := "saved_search_results.binary"

    var hits []item

    // Use saved search results when available, else crawl
    saved_search_results, err := os.Open(ssrpath)
    if errors.Is(err, os.ErrNotExist) {
        saved_search_results, err = os.OpenFile(ssrpath, os.O_WRONLY|os.O_CREATE, 0600)
        if err != nil { panic(err) }

        enc := gob.NewEncoder(saved_search_results)

        hits = search();

        err = enc.Encode(hits); if err != nil { panic(err) }

    } else {

        fmt.Println("Using previously saved search results.\nTo search again remove the 'saved_search_results.binary' file.")

        enc := gob.NewDecoder(saved_search_results)

        err = enc.Decode(&hits)
        if err != nil {
            fmt.Println("Error reading from saved_search_results.binary file.\nTry deleting it to re-search.")
            os.Exit(1)
        }
    }

    // TODO: How to cast list of type to list of interface
    items := make([]list.Item, len(hits))
    for i := range hits {
        items[i] = hits[i]
    }

	m := model{list: list.NewModel(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "LLVM C"

	p := tea.NewProgram(m)
	p.EnterAltScreen()

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

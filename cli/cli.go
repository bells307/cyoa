package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/bells307/cyoa/story"
)

func Run(arcs story.ArcMap, entrypoint string) error {
	cli, err := newCli(arcs)
	if err != nil {
		return err
	}

	return cli.run(entrypoint)
}

type cli struct {
	arcs     story.ArcMap
	arc_tmpl *template.Template
}

func newCli(arcs story.ArcMap) (*cli, error) {
	tmpl, err := template.ParseFiles("views/story_cli.txt")
	if err != nil {
		return nil, err
	}

	return &cli{
		arcs:     arcs,
		arc_tmpl: tmpl,
	}, nil
}

func (c *cli) run(entrypoint string) error {
	currArc, err := c.findArc(entrypoint)
	if err != nil {
		return err
	}

	for currArc != nil {
		currArc, err = c.processArc(currArc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *cli) processArc(arc *story.Arc) (*story.Arc, error) {
	// Печатаем арку
	err := c.arc_tmpl.Execute(os.Stdout, arc)
	if err != nil {
		return nil, err
	}

	// Получаем пользовательский ввод
	opt, err := getUserOption(arc)
	if err != nil {
		return nil, err
	}

	// Возвращаем следующую арку
	return c.findArc(opt.Arc)
}

func (c *cli) findArc(arcName string) (*story.Arc, error) {
	if arc, ok := c.arcs[arcName]; ok {
		return &arc, nil
	} else {
		return nil, errors.New("arc not found")
	}
}

func getUserOption(arc *story.Arc) (*story.Option, error) {
	for {
		if len(arc.Options) == 0 {
			return nil, nil
		}

		fmt.Print("Your option: ")
		var in string
		fmt.Scanln(&in)

		// Находим опцию в арке
		num, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Incorrect input: the value must be an integer")
			continue
		}

		num--

		found := false
		for i := range arc.Options {
			if i == num {
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Incorrect input: the option %s not found\n", in)
			continue
		}

		return &arc.Options[num], nil
	}
}

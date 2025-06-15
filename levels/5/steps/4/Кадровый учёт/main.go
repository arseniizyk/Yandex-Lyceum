package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Company struct {
	Workers []Worker
}

type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

func NewWorker(name, position string, salary, experience uint) *Worker {
	return &Worker{name, position, salary, experience}
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	worker := NewWorker(name, position, salary, experience)
	c.Workers = append(c.Workers, *worker)

	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	var positions = map[string]int{
		"директор":       5,
		"зам. директора": 4,
		"начальник цеха": 3,
		"мастер":         2,
		"рабочий":        1,
	}

	slices.SortStableFunc(c.Workers, func(a, b Worker) int {
		aSalary := a.Salary * a.Experience
		bSalary := b.Salary * b.Experience

		if aSalary != bSalary {
			return cmp.Compare(bSalary, aSalary)
		}

		return cmp.Compare(positions[b.Position], positions[a.Position])
	})

	var sorted []string
	for _, worker := range c.Workers {
		salary := worker.Salary * worker.Experience
		workerInfo := fmt.Sprintf("%s — %d — %s", worker.Name, salary, worker.Position)
		sorted = append(sorted, workerInfo)
	}

	return sorted, nil
}

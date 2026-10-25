// Композиция и развязка: sysadmin, programmer и company
package main

import "fmt"

type administrator interface {
	administrate(system string)
}

type developer interface {
	develop(system string)
}

type adminlist struct {
	list []administrator
}

func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}

func (l *adminlist) Dequeue() administrator {
	a := l.list[0]
	l.list = l.list[1:]
	return a
}

type devlist struct {
	list []developer
}

func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

type sysadmin struct {
	name string
}

func (s *sysadmin) administrate(system string) {
	fmt.Println(s.name, "администрирует", system)
}

type programmer struct {
	name string
}

func (p *programmer) develop(system string) {
	fmt.Println(p.name, "разрабатывает", system)
}

type company struct {
	administrator
	developer
}

func main() {

	var admins adminlist
	var devs devlist

	admins.Enqueue(&sysadmin{name: "Алексей"})

	devs.Enqueue(&programmer{name: "Борис"})
	devs.Enqueue(&programmer{name: "Виктор"})

	cmp := company{
		administrator: admins.Dequeue(),
		developer:     devs.Dequeue(),
	}

	admins.Enqueue(cmp)
	devs.Enqueue(cmp)

	tasks := []struct {
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	for _, task := range tasks {

		if task.needsAdmin {
			adm := admins.Dequeue()
			adm.administrate(task.system)
			continue
		}

		dev := devs.Dequeue()
		dev.develop(task.system)
	}
}

/*
 * @Author: yujiajie
 * @Date: 2024-03-18 11:05:56
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 11:06:07
 * @FilePath: /gateway/core/rungroup/group.go
 * @Description:
 */
package rungroup

type Group struct {
	actors []actor
}

type actor struct {
	execute   func() error
	interrupt func(error)
}

func (g *Group) Add(execute func() error, interrupt func(error)) {
	g.actors = append(g.actors, actor{execute: execute, interrupt: interrupt})
}

func (g *Group) Run() error {
	if len(g.actors) == 0 {
		return nil
	}

	errors := make(chan error, len(g.actors))
	for _, a := range g.actors {
		go func(a actor) {
			errors <- a.execute()
		}(a)
	}

	err := <-errors

	for _, a := range g.actors {
		a.interrupt(err)
	}

	for i := 1; i < cap(errors); i++ {
		<-errors
	}

	return err
}

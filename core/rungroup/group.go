/*
 * @Author: yujiajie
 * @Date: 2024-03-18 11:05:56
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-28 16:03:09
 * @FilePath: /Gateway/core/rungroup/group.go
 * @Description: 用于管理协程同步，优雅关闭
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

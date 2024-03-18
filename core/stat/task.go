/*
 * @Author: yujiajie
 * @Date: 2024-03-18 16:23:03
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 16:23:27
 * @FilePath: /gateway/core/stat/task.go
 * @Description:
 */
package stat

import "time"

type Task struct {
	Drop        bool
	Duration    time.Duration
	Description string
}

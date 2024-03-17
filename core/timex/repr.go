/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-12 23:52:31
 * @LastEditTime: 2023-12-12 23:53:04
 * @LastEditors: yujiajie
 */
package timex

import (
	"fmt"
	"time"
)

func ReprOfDuration(duration time.Duration) string {
	return fmt.Sprintf("%.1fms", float32(duration)/float32(time.Millisecond))
}

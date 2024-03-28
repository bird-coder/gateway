/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-12 23:52:31
 * @LastEditTime: 2024-03-28 16:03:16
 * @LastEditors: yujiajie
 */
package timex

import (
	"fmt"
	"time"
)

// 将时长转为毫秒
func ReprOfDuration(duration time.Duration) string {
	return fmt.Sprintf("%.1fms", float32(duration)/float32(time.Millisecond))
}

// email test
package util

import (
	"testing"
)

func TestSendToMail(t *testing.T) {
	Send("1398154729@qq.com","My Go Email Test","Email from test,hello world.")
}
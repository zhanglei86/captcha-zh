package captcha

import (
	"captcha-zh/config"

	"strconv"
	"time"
	"testing"
)

func BenchmarkTopic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rd := strconv.Itoa(time.Now().Nanosecond())
		tp := RandTopic()
		Draw(tp.Subject, config.PATH_CONFIG_IMAGE_TEMP + rd + config.SEPARATOR_VERTICAL_LINE + tp.Result + ".png")
	}
}

func BenchmarkDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Draw("十 二 减 去 13 等 于", config.PATH_CONFIG_IMAGE_TEMP + "result.gif")
	}
}

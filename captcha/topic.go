package captcha

import (
	"captcha-zh/config"

	"strings"
	"strconv"
	"math/rand"
	"time"
)

const (
	captchaLen int = 7
)

var (
	num2chinese []string = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	operator2chinese map[string][]string = map[string][]string{
		"+": []string{"加", "加 上"},
		"-": []string{"减", "减 掉"},
	}
	eql2chinese []string = []string{"是", "等 于", "是 多 少", "是 多 少 呢"}
)

type NumSt struct {
	Size int
	Cn   string
}

type Topic struct {
	Subject string
	Result  string
}

func random(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max-min) + min
}

func Random(min, max int) int {
	i := random(min, max)
	return i
}

func RandomGen() chan int{
	out := make(chan int)
	go func() {
		for {
			rand.Seed(time.Now().Unix())
			out <- rand.Intn(100)
		}
	}()
	return out
}

func randFont() string {
	rand.Seed(int64(time.Now().Nanosecond()))
	return config.Fonts[rand.Intn(len(config.Fonts))]
}

/*
num最大109
 */
func Num2Cn(num int) NumSt {
	// case:个位数
	if num < 11 {
		return NumSt{1, num2chinese[num]}
	}

	var numSt NumSt

	a, b := num/10, num%10
	// case:十几
	if a == 1 {
		numSt.Size = 2
		numSt.Cn = strings.Join([]string{num2chinese[10], num2chinese[b]}, " ")
		return numSt
	}
	// case:几十
	if b == 0 {
		numSt.Size = 2
		numSt.Cn = strings.Join([]string{num2chinese[a], num2chinese[10]}, " ")
		return numSt
	}

	// case:其他
	numSt.Size = 3
	numSt.Cn = strings.Join([]string{num2chinese[a], num2chinese[10], num2chinese[b]}, " ")

	return numSt
}

func TopicParse(le NumSt, rt NumSt, operator string, optLen int) string {
	var currentLen int = le.Size + rt.Size
	optArr := operator2chinese[operator]
	if currentLen == 6 {
		return strings.Join([]string{le.Cn, optArr[0], rt.Cn}, " ")
	}
	if currentLen == 5 && optLen == 2 {
		return strings.Join([]string{le.Cn, optArr[1], rt.Cn}, " ")
	}
	eqlLen := captchaLen - optLen - currentLen
	return strings.Join([]string{le.Cn, optArr[optLen-1], rt.Cn, eql2chinese[eqlLen-1]}, " ")
}

func RandNumParse(num int) NumSt {
	if random(0, 10) < 5 {
		return Num2Cn(num)
	}
	return NumSt{1, strconv.Itoa(num)}
}

func RandTopic() Topic {
	var le, rt, result int
	var operator string

	operateInt := random(0, 10)
	if operateInt < 5 {
		operator = "+"
		le, rt = random(0, 10), random(0, 10)
		result = le + rt
	} else {
		operator = "-"
		le = random(10, 20)
		rt = random(0, le)
		result = le - rt
	}

	// 随机
	optLen := 2
	if random(0, 10) < 5 {
		optLen = 1
	}

	var topic Topic
	topic.Subject = TopicParse(RandNumParse(le), RandNumParse(rt), operator, optLen)
	topic.Result = strconv.Itoa(result)

	return topic
}
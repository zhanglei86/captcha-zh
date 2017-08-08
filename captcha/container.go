package captcha

import (
	"captcha-zh/config"

	"sync"
	"errors"
	"strings"
	"os"
)

/**
图片容器逻辑，包括预生成
 */

type Container struct {
	captchaList		[]string	//图片列表
	locker			*sync.Mutex	//锁
	pointerIndex	int			//计数器，消费1次累计1，配合next使用
	consumption		int			//消费次数，会触发更新逻辑
}

var CaptchaContainer *Container

func init() {
	CaptchaContainer = &Container{
		[]string{},
		new(sync.Mutex),
		0,
		0,
	}
}

func (c *Container) Lock() {
	c.locker.Lock()
}

func (c *Container) Unlock() {
	c.locker.Unlock()
}

// Append items to the end of the list
func (c *Container) Append(items ...string) {
	c.captchaList = append(c.captchaList, items...)
}

func (c *Container) UpdateNeed() bool {

	// 阀值条件，跟消费计数器比较
	ts := config.TConfig.CaptchaSys.Threshold
	if c.consumption < ts {
		return false
	}
	//需要更新后，计数器立马清零，接受新的请求计数
	c.consumption = 0

	return true
}

// Append items to the end of the list and remove old items from the front.
// At the same time move pointer
func (c *Container) UpdateBak(items ...string) []string {

	itemsSize, listSize := len(items), len(c.captchaList)
	c.Append(items...)

	captchaList := make([]string, len(c.captchaList))
	copy(captchaList, c.captchaList)
	c.captchaList = captchaList[itemsSize:]

	c.pointerIndex = (c.pointerIndex - itemsSize) % listSize
	if c.pointerIndex < 0 {
		c.pointerIndex += listSize
	}

	return captchaList[:itemsSize]
}

func (c *Container) Update(items ...string) {
	// 旧的
	oldCaptchaList := c.captchaList

	// 初始化
	c.pointerIndex = 0
	c.captchaList = nil
	// 加入新的
	c.Append(items...)

	// 删除旧的
	var fileName string
	for _, captcha := range oldCaptchaList {
		fileName = strings.Split(captcha, config.SEPARATOR_VERTICAL_LINE)[0]
		os.Remove(config.TConfig.Paths.Path + config.PATH_CONFIG_IMAGE_TEMP + fileName)
	}
}

// Get next item by index
func (c *Container) Next() (string, error) {
	c.Lock()
	defer c.Unlock()

	// 自增
	c.consumption += 1

	if len(c.captchaList) == 0 {
		return "", errors.New("No item found")
	}

	// %的算法，数组内取值
	index := c.pointerIndex % len(c.captchaList)
	c.pointerIndex = index + 1

	return c.captchaList[index], nil
}

package captcha

import (
	"captcha-zh/config"

	"sync"
	"errors"
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
	ts := config.TConfig.CaptchaSys.Threshold

	if c.consumption < ts {
		return false
	}

	//更新后，计数器清零
	c.consumption = 0

	return true
}

// Append items to the end of the list and remove old items from the front.
// At the same time move pointer
func (c *Container) Update(items ...string) []string {

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

// Get next item by index
func (c *Container) Next() (string, error) {
	c.Lock()
	defer c.Unlock()

	c.consumption += 1
	if len(c.captchaList) == 0 {
		return "", errors.New("No item found")
	}
	index := c.pointerIndex % len(c.captchaList)
	c.pointerIndex = index + 1
	return c.captchaList[index], nil
}

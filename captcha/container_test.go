package captcha

import (
	"testing"
	"sync"
)

func TestCaptchaInsert(t *testing.T) {
	CaptchaContainer.Append("1", "2")
	if len(CaptchaContainer.captchaList) != 2 || CaptchaContainer.pointerIndex != 0 {
		t.Fatalf("Insert with unexcepted")
	}
}

func TestCaptchaNext(t *testing.T) {
	// 初始化逻辑
	CaptchaContainer.Append("1", "2")

	item, err := CaptchaContainer.Next()
	if err != nil {
		t.Fatalf("Next with err: %s", err.Error())
	}
	if item != "1" {
		t.Fatalf("Next with unexcepted")
	}
	item, _ = CaptchaContainer.Next()
	if item != "2" {
		t.Fatalf("Next with unexcepted")
	}
}

func TestCaptchaUpdate(t *testing.T) {

	// 初始化逻辑
	CaptchaContainer = &Container{
		[]string{"1", "2"},
		new(sync.Mutex),
		0,
		0,
	}

	a := CaptchaContainer.Update("3")
	if a[0] != "1" || len(a) != 1 {
		t.Fatalf("Update return with unexcepted")
	}
	item, _ := CaptchaContainer.Next()
	if item != "3" {
		t.Fatalf("Next with unexcepted")
	}
	item, _ = CaptchaContainer.Next()
	if item != "2" {
		t.Fatalf("Next with unexcepted")
	}
}

func TestCaptchaLock(t *testing.T) {

	// 初始化逻辑
	CaptchaContainer = &Container{
		[]string{"1", "2"},
		new(sync.Mutex),
		0,
		0,
	}

	CaptchaContainer.Lock()
	go func() {
		// =3
		CaptchaContainer.Update("4")
		CaptchaContainer.Unlock()
	}()
	item, _ := CaptchaContainer.Next()
	if item != "3" {
		t.Fatalf("Next with unexcepted")
	}
}

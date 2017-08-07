package config

const (
	//这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	DATE_FORMAT_DEFAULT string = "2006-01-02 15:04:05"

	// 分隔符
	SEPARATOR_UNDER_LINE = "_"
	SEPARATOR_VERTICAL_LINE = "|"
	SEPARATOR_SPACE = " "

	// 绝对路径，根
	PATH_ROOT = "/Volumes/IntelSSD/Dev/GoRepository/src/captcha-zh/"
	// 配置文件路径
	PATH_CONFIG = PATH_ROOT + "resource/conf/config.json"
	// 字库文件路径
	PATH_CONFIG_FONT = PATH_ROOT + "resource/assets/fonts/"
	// 背景图片路径
	PATH_CONFIG_IMAGE_BG = PATH_ROOT + "resource/assets/images/bg.gif"
	// 测试用临时图片路径
	PATH_CONFIG_IMAGE_TEMP = PATH_ROOT + "bin/tmp/"
)

var (
	Fonts []string
)
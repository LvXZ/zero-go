module robots

go 1.15

require (
	gitee.com/go-package/carbon v1.3.4
	github.com/confluentinc/confluent-kafka-go v1.6.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/wire v0.4.0
	github.com/phachon/go-logger v0.0.0-20191215032019-86e4227f71ea
	github.com/robfig/cron v1.2.0
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/tidwall/gjson v1.6.3
	go-basiclib v0.0.21
	go-common v1.20.1
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.5
)

replace (
	go-basiclib => git.bilibili.co/datacenter/go-basiclib v0.0.21
	go-common => git.bilibili.co/platform/go-common v1.20.1
)

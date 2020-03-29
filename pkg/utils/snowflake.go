package utils

import (
	"github.com/CHYinYiHang/gocron/pkg/logging"
	"github.com/bwmarrin/snowflake"
)

/*
snowflake ID 算法是 twitter 使用的唯一 ID 生成算法，为了满足 Twitter 每秒上万条消息的请求，使每条消息有唯一、有一定顺序的 ID ，且支持分布式生成。

snowflake ID 的结构是一个 64 bit 的 int 型数据。
1 bit：不使用，可以是 1 或 0

41 bit：记录时间戳 (当前时间戳减去用户设置的初始时间，毫秒表示)，可记录最多 69 年的时间戳数据

10 bit：用来记录分布式节点 ID，一般每台机器一个唯一 ID，也可以多进程每个进程一个唯一 ID，最大可部署 1024 个节点

12 bit：序列号，用来记录不同 ID 同一毫秒时的序列号，最多可生成 4096 个序列号


假设在一个节点 (机器) 上，节点 ID 唯一，并发时有多个线程去生成 ID。
满足以上条件时，如果多个进程在同一毫秒内生成 ID，那么序列号步进 (加一)，这里要保证序列号的操作并发安全，使同一毫秒内生成的 ID 拥有不同序列号。如果序列号达到上限，则等待这一毫秒结束，在新的毫秒继续步进。

这样保证了：
所有生成的 ID 按时间趋势递增 （有序）
整个分布式系统内不会产生重复 ID （唯一）

yiyihang 2020/03/20
单元测试：大概毫秒可以生成4096个唯一id，1毫秒=100W/ns
BenchmarkSnowFlake-6     4891089               245 ns/op
BenchmarkSnowFlake-6     4890499               245 ns/op
BenchmarkSnowFlake-6     4910473               245 ns/op
BenchmarkSnowFlake-6     4930611               244 ns/op
BenchmarkSnowFlake-6     4910408               245 ns/op
*/

//id生成器对象
var SfNode *snowflake.Node
var InitErr error

//初始化Id生成器
//参数：@1当前节点ID
func LoadSnowflakeNode(workerId int64) {
	//fmt.Println(CurrWork.Generate().String())
	SfNode, InitErr = snowflake.NewNode(workerId)
	if InitErr != nil {
		logging.Error(InitErr, "初始化ID生成器失败")
	}
}

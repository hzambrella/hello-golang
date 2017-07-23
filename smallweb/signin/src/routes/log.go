package routes

import "engine/logz"

//日志系统有助于排查错误。本程序的日志只是简单的打印到终端。柳丁的可以保存成文件，以前有个项目日志甚至专门有日志服务器，可以筛选错误类型。日志系统见柳丁的说明。
//TODO:Warn级别和Panic级别
//TODO:序列化为文件或日志服务器
var logl = logz.NewLogDebug(true)

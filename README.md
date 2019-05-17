# parse_dart_lock_file

解析pubspec.lock文件反向获取依赖

使用方法:

下载release中对应系统的包(mac是darwin)

解压

linux/mac需要在命令行中执行`sudo chmod 777 文件名` 或者+x 这个根据你自己的需要

调用:
如果和flutter/dart目录同级,可以直接执行: `read-dart-lock-file`  
如果不同级则执行: `read-dart-lock-file ./read-dart-lock-file ~/Documents/GitHub/flutter_ijkplayer/example/pubspec.lock`

这样可以反向解析以获取当前的真实插件版本号

---


工程中带有一个build_release.go 这个编译后的二进制工具可以帮助快速构建go的应用程序
`go build build_release.go` // 编译获取二进制  
`./build_release read-dart-lock-file.go` 这样能快速获取win/mac/darwin的tar.gz格式的压缩包, 这一步会删除read-dart-lock-file文件,无论是否是本次编译生成的

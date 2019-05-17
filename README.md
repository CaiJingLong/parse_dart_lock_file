# parse_dart_lock_file

解析pubspec.lock文件反向获取依赖

## 使用方法

下载release中对应系统的包(mac是darwin)

解压

linux/mac需要在命令行中执行`sudo chmod 777 文件名` 或者+x 这个根据你自己的需要

调用:
如果和flutter/dart目录同级,可以直接执行: `read-dart-lock-file`  
如果不同级则执行: `read-dart-lock-file ./read-dart-lock-file ~/Documents/GitHub/flutter_ijkplayer/example/pubspec.lock`

这样可以反向解析以获取当前的真实插件版本号

```bash
$ ./read-dart-lock-file ~/Documents/GitHub/flutter_ijkplayer/example/pubspec.lock
dependencies:
  photo: 0.3.3
  glob: 1.1.7
  intl_translation: 0.17.3
  string_scanner: 1.0.4
  args: 1.5.1
  async: 2.1.0
  dart_style: 1.2.4
  flutter_ijkplayer: 
    path: ..
  test_api: 0.2.4
  crypto: 2.0.6
  cupertino_icons: 0.1.2
  matcher: 0.12.5
  package_config: 1.0.5
  vector_math: 2.0.8
  collection: 1.14.11
  kernel: 0.3.14
  pedantic: 1.5.0
  petitparser: 2.2.1
  photo_manager: 0.3.3
  convert: 2.1.1
  charcode: 1.1.2
  source_span: 1.5.5
  stack_trace: 1.9.3
  term_glyph: 1.1.0
  typed_data: 1.1.6
  yaml: 2.1.15
  path: 1.6.2
  pub_semver: 1.4.2
  front_end: 0.1.14
  intl: 0.15.8
  meta: 1.1.6
  oktoast: 2.1.6
  quiver: 2.0.2
  stream_channel: 2.0.0
  analyzer: 0.35.4
  boolean_selector: 1.0.4
  watcher: 0.9.7+10
```

## 其他

工程中带有一个build_release.go 这个编译后的二进制工具可以帮助快速构建go的应用程序
`go build build_release.go` // 编译获取二进制  
`./build_release read-dart-lock-file.go` 这样能快速获取win/mac/darwin的tar.gz格式的压缩包, 这一步会删除read-dart-lock-file文件,无论是否是本次编译生成的

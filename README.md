# Gsgm - GNU Single Game Manager

## 介绍

目前 Gsgm 要联合 Lutris 一起使用，与 Lutris 一起能够很方便的管理自定义游戏库。

Gsgm 不止能够完成全自动批量安装游戏的功能，还有诸如 Lutris 记录备份、Lutris 封面管理等，还有游戏库组管理(需要 `lutris-git v0.5.13+`)

Gsgm 的特点和功能：

1. 根据规则批量导入游戏，只需要遵循规则，然后输入 `gsgm install -l <游戏库位置>` 即可（详细请看下文
2. 同步 Lutris 的游戏记录和设置，即便 Lutris 的记录丢了，我们也可以通过游戏库中，每个游戏下的 `.gsgm` 文件恢复游戏库的 Lutris 数据
3. 在 Lutris 上，封面的管理相当随意，只有联合 Steam 等平台的游戏封面才能很好的被管理。Gsgm 脱离 Lutris，可以主动生成 Lutris 的 3 种显示封面，并代为管理

*注意：使用 Gsgm 的时候，请关闭 Lutris!!!，Gsgm 使用过程中可能会操作 Lutris 数据库，可能会导致运行中的 Lutris 发生错误*

## 游戏库规则

*请仔细阅读，这个结构至关重要！！！*

```
游戏库位置
├── 游戏A
│   ├── ...                 -------- 其他文件夹和文件
│   ├── .gsgm
│   │     ├── info.json     -------- 游戏的信息
│   │     ├── setting.json  -------- 游戏的运行设置
│   │     ├── history.json  -------- 游玩历史
│   │     ├── cover.xxx     -------- 游戏封面，可以是 [png | jpg | jpeg]
│   └── bin
│        └── 游戏A.exe       -------- 游戏可执行文件，可以识别游戏根目录和二级目录的可执行文件
├── 游戏B
│   ├── ...                 -------- 其他文件夹和文件
│   ├── .gsgm
│   │     └── ...
│   └── 游戏B.exe            -------- 游戏可执行文件
├── @文件夹A                 -------- 文件夹的命名是单个 @ 开头，例如 `@3D大作`、`@ 2D游戏`，文件夹可以嵌套
│   ├── 游戏A
│   │   └── ...
│   └── @文件夹B             -------- 文件夹可以嵌套
│       └── ...
├── @_文件夹C                -------- `@_` 开头，如 `@_3D大作`，例如这里面的游戏会被分到 `@_文件夹C` 组中
│   ├── 游戏A
│   ├── 游戏C
│   │   └── ...
│   ├── 游戏D
│   │   └── ...
│   └── @文件夹B             -------- 文件夹可以嵌套
│       └── ...
├── @@忽略文件夹A             -------- 忽略文件夹以两个 @@ 开头，例如 `@@游戏脚本`、`@@ 风灵月影工具包`
├── @@忽略文件夹B
└── ...
```

info.json

```json5
{
    "id": 1778238210896297984, // `gsgm init 随机生成 gsgm id，每个游戏独一无二，不可重复！
    "initTime": 1712800002
}
```

setting.json

```json5
{
    "execute": "run.exe",// 游戏启动文件，路径相对于游戏根目录
    "prefixAlone": false,// wineprefix 是否隔离
    "localeCharSet": "zh_CN.UTF-8",// 字符集编码，默认 zh_CN.UTF-8
    "platform": "Windows",// 游戏平台，默认 Windows，可选值：[Windows | Linux]
    "runner": "wine"
}
```

## 使用方式

### ArchLinux

```sh
yay -U go-gsgm-{具体版本}.pkg.tar.zst 
```

### 原生安装

> 需要有一定 Linux 基础

1. 下载 go-gsgm 二进制文件
2. 生成对应的代码补全脚本 `go-gsgm completion (bash|zsh|fish) > completion.sh`
3. 生成 lutris 的 post exit 脚本 `gsgm gen lupes > gsgm-lupes`

## 文档

具体命令的介绍，可以通过 `gsgm xxx -h` 获取

```sh
Gsgm Linux 游戏管理工具

Usage:
  gsgm [command]

Available Commands:
  check       检查当前游戏或者游戏库目录是否合法
  clean       清理 Lutris 中 Gsgm 游戏
  completion  Generate the autocompletion script for the specified shell
  gen         生成辅助性的脚本命令
  help        Help about any command
  init        初始化单个游戏或者一整个 Gsgm 库
  install     安装游戏
  scan        扫描游戏当前库
  sync        同步游戏时长

Flags:
  -h, --help      help for gsgm
  -v, --verbose   verbose output
      --version   version for gsgm

Use "gsgm [command] --help" for more information about a command.
```

```sh
检查当前游戏或者游戏库目录是否合法

Usage:
  gsgm check [flags]

Flags:
  -h, --help         help for check
  -l, --is-library   是否是游戏库

Global Flags:
  -v, --verbose   verbose output
```

```sh
清理 Lutris 中 Gsgm 游戏

Usage:
  gsgm clean [flags]

Flags:
  -h, --help   help for clean

Global Flags:
  -v, --verbose   verbose output
```

```sh
Generate the autocompletion script for gsgm for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  gsgm completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh

Flags:
  -h, --help   help for completion

Global Flags:
  -v, --verbose   verbose output

Use "gsgm completion [command] --help" for more information about a command.
```

```sh
Help provides help for any command in the application.
Simply type gsgm help [path to command] for full details.

Usage:
  gsgm help [command] [flags]

Flags:
  -h, --help   help for help

Global Flags:
  -v, --verbose   verbose output
```

```sh
初始化单个游戏或者一整个 Gsgm 库

Usage:
  gsgm init [flags]

Flags:
  -h, --help         help for init
  -l, --is-library   是否是 Gsgm 游戏库

Global Flags:
  -v, --verbose   verbose output
```

```sh
安装游戏

Usage:
  gsgm install [flags]

Flags:
  -f, --force   覆盖安装
  -h, --help    help for install
  -l, --lib     是否是 Gsgm 游戏库
  -s, --safe    安装，但不添加图片

Global Flags:
  -v, --verbose   verbose output
```

```sh
扫描游戏当前库

Usage:
  gsgm scan [flags]

Flags:
  -h, --help   help for scan

Global Flags:
  -v, --verbose   verbose output
```

## 常见操作举例

清空Gsgm的游戏库

> 这里将清除所有通过 Gsgm 安装的游戏数据

```shell
gsgm clean
```

### 初始化游戏库

```shell
gsgm init -l /path/to/library

gsgm init /path/to/single
```

### 检查游戏结构

```shell
gsgm check -l /path/to/library

gsgm check /path/to/single
```

### 安装游戏库

```shell
# 覆盖安装
gsgm install -f /path/to/library

# 单独安装某个游戏
gsgm install --lib=false /path/to/single

# 安全模式安装（不安装图片
gsgm install -f -s /path/to/library
```

### 罗列游戏

```shell
gsgm scan /path/to/library
```


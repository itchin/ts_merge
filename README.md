### ts_merge
适用于windows，批量合并手机uc浏览器等多种场景下m3u8视频缓存生成的ts流文件

### 二进制下载
https://github.com/itchin/ts_merge/releases

### 使用说明
程序应放在.m3u8文件同一目录，且.ts文件目录也在同一目录下，目录结构如下：

. <br/>
--- video_0.m3u8 <br/>
--- vide0_1.m3u8 <br/>
+-- 1635251256948 <br/>
|   --- 0.ts <br/>
|   --- 1.ts <br/>
|   --- 2.ts <br/>
+-- 1635251853491 <br/>
|   --- 0.ts <br/>
|   --- 1.ts <br/>
|   --- 2.ts <br/>
--- ts_merge.exe <br/>
--- video_0.bat <br/>
--- video_1.bat <br/>

执行ts_merge.exe，将读取.m3u8，生成对应的.bat文件。之后批量执行.bat，将.ts合并。

### 其它
1、原理 <br/>
使用windows cmd命令 copy/b 0.ts + 1.ts + 2.ts + ... + n.ts video.ts 将多个.ts合并为一个

2、为什么不只生成一个.bat? <br/>
可以用连接符&将多条copy/b命令连接，但本人亲测如果.m3u8文件稍多一点时，命令太长而执行失败。

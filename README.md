# excel转json<br>
## 支持服务器端和客户端<br>
## 服务器端：生成的是标准JSON<br>
## 客户端：生成的是阉割版的,表头+数据的方式<br><br/>
![演示图](process.png)<br>
<br/><br/>
<pre>
config.json
{
  "data":[
    {
      "name":"内网165",
      "inPath":"C:/Users/Administrator/Desktop/new_svn/shuzhi/数据表/",
      "serverOutPath":"C:/Users/Administrator/Desktop/new_svn/a/out1/s/",
      "clientOutPath":"C:/Users/Administrator/Desktop/new_svn/a/out1/c/",
      "structPath":"C:/Users/Administrator/Downloads/1/ttttt/u/",
      "clientZip":1
    }
  ,
    {
      "name":"公网181",
      "inPath":"C:/Users/Administrator/Desktop/new_svn/shuzhi/数据表/",
      "serverOutPath":"C:/Users/Administrator/Desktop/new_svn/a/out2/s/",
      "clientOutPath":"C:/Users/Administrator/Desktop/new_svn/a/out2/c/",
      "structPath":"C:/Users/Administrator/Downloads/1/ttttt/u/",
      "serverZip":0
    }
  ]
}
</pre>
name：平台标志<br/>
inPath：excel目录<br/>
serverOutPath：服务器端生成json路径<br/>
[clientOutPath]：客户端生成json路径<br/>
[structPath]：结构体输出路径<br/>
[clientZip]：客户端是否打个ZIP包,1打包，0不打包. 此配置可选，删除此配置默认为不打包<br/>
[serverZip]:服务器是否打包，同上

其中，clientOutPath，structPath，clientZip，serverZip 都是选填。不填则不生成
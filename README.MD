## 使用说明


### Windows:

```
C2> 
Agent: d2cfbad734a347a2932bdf061ba44ae1 Arrived!

C2> interact d2cfbad734a347a2932bdf061ba44ae1
Type help or ? to list commands.

d2cfbad734a347a2932bdf061ba44ae1> shell dir
[+] Shell Command Results:


 驱动器 C 中的卷没有标签。
 卷的序列号是 4C91-9DB8

 C:\Users\web\Desktop 的目录

2021/10/31  00:50    <DIR>          .
2021/10/31  00:50    <DIR>          ..
2021/08/25  21:27            73,802 1.exe
2021/08/26  21:20             9,988 1111.xml
2021/08/26  21:05                88 222.txt
2021/10/31  13:30             9,728 C2.exe


d2cfbad734a347a2932bdf061ba44ae1> sleep 0
[+] Shell Command Results:
sleep:0

```

### Linux:

```
kHKofWxMkmTFLiku> shell ping -c 5 8.8.8.8
kHKofWxMkmTFLiku>
[+] Shell Command Results:

PING 8.8.8.8 (8.8.8.8): 56 data bytes
64 bytes from 8.8.8.8: icmp_seq=0 ttl=112 time=55.691 ms
64 bytes from 8.8.8.8: icmp_seq=1 ttl=112 time=50.363 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=112 time=55.485 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=112 time=55.257 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=112 time=54.776 ms

--- 8.8.8.8 ping statistics ---
5 packets transmitted, 5 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 50.363/54.314/55.691/1.999 ms

```

更新记录

```
2021/11/6
[!]修改C#通讯逻辑
[+]新增GoClient(Linux)
```

## 通讯过程


![](/img/1.png)


## 免责声明

本项目只用于个人交流学习，利用此工具造成任何后果，均由使用者本人负责，作者不承担任何责任。

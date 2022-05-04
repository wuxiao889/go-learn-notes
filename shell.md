# mv

 也可以使用带选项的mv命令，把多个文件移动到一个目录中，如

```shell
$ mv a b c -t d
$ mv -t d a b c
# 其中，-t后面紧接着的就是要移动到的目录，并且不能有多个目录出现
```
打印移动信息

```shell
mv -v *.txt /home/office
```

提示是否覆盖文件

```shell
mv -i file_1.txt /home/office
```

反选

``mv !(child1|child2) child1``
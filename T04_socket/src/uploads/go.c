1.3.3 编译程序
假设之前介绍的Hello, world代码被保存为了hello.go，并位于~/goyard目录下，那么可以用以
下命令行编译并直接运行该程序：
$ cd ~/goyard
$ go run hello.go # 直接运行
Hello, world. 你好，世界！

使用这个命令，会将编译、链接和运行3个步骤合并为一步，运行完后在当前目录下也看不
到任何中间文件和最终的可执行文件。如果要只生成编译结果而不自动运行，我们也可以使用 Go
命令行工具的build命令：
图灵社区会员 soooldier(soooldier@live.com) 专享 尊重版权
$ cd ~/goyard
$ go build hello.go
$ ./hello
Hello, world. 你好，世界！

可以看出， Go命令行工具是一个非常强大的源代码管理工具。我们将在第4章中详细讲解Go
命令行工具所包含的更多更强大的功能。
从根本上说， Go命令行工具只是一个源代码管理工具，或者说是一个前端。真正的Go编译
器和链接器被Go命令行工具隐藏在后面，我们可以直接使用它们：
$ 6g helloworld.go
$ 6l helloworld.6
$ ./6.out
Hello, world. 你好，世界！
6g和6l是64位版本的Go编译器和链接器，对应的32位版本工具为8g和8l。 Go还有另外一个
GCC版本的编译器，名为 gccgo，但不在本书的讨论范围内。



1.6.2 GDB调试
不用设置什么编译选项，Go语言编译的二进制程序直接支持GDB调试，比如之前用go build
calc编译出来的可执行文件calc，就可以直接用以下命令以调试模式运行：
$ gdb calc
因为GDB的标准用法与Go没有特别关联，这里就不详细展开了，有兴趣的读者可以自行查
看对应的文档。需要注意的是， Go编译器生成的调试信息格式为DWARFv3，只要版本高于7.1
的GDB应该都支持它。


<calcproj>
├─<src>
├─<calc>
  ├─calc.go
├─<simplemath>
  ├─add.go
  ├─add_test.go
  ├─sqrt.go
  ├─sqrt_test.go
├─<bin>
├─<pkg>＃包将被安装到此处
了能够构建这个工程，需要先把这个工程的根目录加入到环境变量GOPATH中。
假设calcproj目录位于~/goyard下，则应编辑~/.bashrc文件，并添加下面这行代码：
export GOPATH=~/goyard/calcproj
然后执行以下命令应用该设置：
$ source ~/.bashrc

GOPATH和PATH环境变量一样，也可以接受多个路径，并且路径和路径之间用冒号分割。
设置完GOPATH后，现在我们开始构建工程。假设我们希望把生成的可执行文件放到
calcproj/bin目录中，需要执行的一系列指令如下：
$ cd ~/goyard/calcproj
$ mkdir bin
$ cd bin
$ go build calc

顺利的话，将在该目录下发现生成的一个叫做calc的可执行文件，执行该文件以查看帮助信
息并进行算术运算：
$ ./calc
USAGE: calc command [arguments] ...
The commands are:
addAddition of two values.
sqrtSquare root of a non-negative value.

$ ./calc add 2 3
Result: 5
	
$ ./calc sqrt 9
Result: 3
	
go test simplemath

	
从上面的构建过程中可以看到，真正的构建命令就一句：
go build calc



<sorter>
├─<src>
├─<sorter>
├─sorter.go
├─<algorithms>
├─<qsort>
├─qsort.go
├─qsort_test.go
├─<bubblesort>
├─bubblesort.go
├─bubblesort_test.go
├─<pkg>
├─<bin>


$ go build sorter.go
$ ./sorter -i unsorted.dat -o sorted.dat -a bubblesort
infile = unsorted.dat outfile = sorted.dat algorithm = bubblesort
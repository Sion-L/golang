```go
1、math.Floor函数 -> math包
   拿到一个浮点数，将其向下取舍为最接近的整数，然后返回该整合数
   fmt.Println(math.Floor(2.75)) => 输出2

2、strings.Title函数 -> strings包
   拿到一个字符串，将它所包含的每个单词的第一个大写("将其转换为首大写字母")，并返回大写字母开头的字符串
   fmt.Println(strings.Title("hello world")) => 输出Hello World

3、go的符文：rune，用于表示单个字符
   字符串的字面量用双引号包围
   rune的字面量用单引号包围
   go使用Unicode标准来存储rune，rune被保存为数字代码，而不是字符本身，用fmt.Println输出的数字代码，而不是原始字符
   fmt.Println('A') => 输出为65
   rune字面量也可以用转移序列，用来表示程序代码中难以包含的字符：
   fmt.Println('\t') => 输出为9
   fmt.Println('\n') => 输出为10
   fmt.Println('\\') => 输出为92

4、go是静态类型的，这意味着它甚至在程序运行之前就知道值的类型是什么

5、reflect.TypeOf函数 -> reflect包
   TypeOf函数可以查看输出值的类型
   fmt.Println(reflect.TypeOf(42)) => 输出为int
   float64类型：类型中的64是因为要用64位的数据来保存数字，这意味着在四舍五入之前，float值可以相当精确，但也不是无限精确

6、变量：包含值的一块存储

7、短声明变量：声明变量时就知道它的初始值是什么

8、命名规则
   名称必须以字母开头，并且可以有任意数量的额外的字母和数字。
   如果变量、函数或类型的名称以大写字母开头，则认为它是导出的，可以从当前包之外的包访问它。（这就是为什么fmt.Println中的P是大写的：这样它就可以在main包或任何其他包中使用。）如果变量/函数/类型的名称是以小写字母开头的，则认为该名称是未导出的，只能在当前包中使用。

9、go社区遵循的额外的命名规则
   如果一个名称由多个单词组成，那么第一个单词之后的每个单词都应该首字母大写，并且它们应该连接在一起，中间没有空格，比如topPrice、RetryConnection，等等。（名称的第一个字母只有在你想从包中导出时才应大写。）这种样式通常称为驼峰大小写，因为大写字母看起来像驼峰。
   当名称的含义在上下文中很明显时，Go社区的惯例是缩写它：用i代替index，用max代替maximum，等等
    注：只有名称是以大写字母开头的变量、函数或类型才被认为是可导出的：可以从当前包之外的包访问。
```



```go
1、time.Time函数 -> time包的Time类型函数可以表示日期和时间.
   var nian time.Time = time.Now()	// time.Now()返回一个表示当前日期和时间的time.Time值
   var year int = nian.Year()         // time.Time值有一个返回年份的year方法
   fmt.Println(year)    =>   输出2021
含义：time.Now()函数返回当前日期和时间的新Time值，我们将其存储在now变量中，然后对now引用的值调用year方法，year方法返回一个代表年的整数
注：方法是与特定类型的值关联的函数

2、strings.Relacer函数 -> strings包有一个Replacer类型，可以在字符串中搜索子字符串，并且在每次子字符串出现的地方用另一个字符串替换它
   test := "G# r#cks!"      
   word := strings.NewReplacer("#", "o")   // 这里返回一个strings.Replace，其设置为将每个"#"替换为"o"
   fixed := word.Replace(test)		// 对strings.Replace调用Replace方法，并传递一个字符串来进行替换
   fmt.Println(fixed) 	 	=>  输出Go rocks！
含义：strings.NewReplacer函数接受要替换的字符串（"#"）和要替换为的字符串（"o"）的参数，并返回strings.Replacer。当我们将一个字符串传递给Replacer值的Replace方法时，它将返回一个完成了替换的字符串。
注：.表示右边的东西属于左边

3、_：空白标识符
   golang中定义的变量必须被使用，但我们不打算使用某个变量的值时，我们可以使用Go的空白标识符。为空白标识符分配一个值实际上会丢弃它

4、os.stdin函数、bufio.NewReader函数以及ReadString方法、_忽略错误
   fmt.Println("please enter: ")
   reader := bufio.NewReader(os.Stdin)		// 设置从键盘获取文本的"缓冲读取器"
   input, _ := reader.ReadString('\n')		// 返回用户输入的所有内容
   fmt.Println(input)		// 打印用户输入的内容
      // 含义：从程序的标准输入中读取（接收和存储）输入的方法，所有的键盘输入都使用标准输入。行reader：=bufio.NewReader（os.Stdin）将bufio.Reader保存在reader变量中
      为了实际获得用户的输入，我们调用Reader的ReadString方法。ReadString方法需要一个带有rune（字符）的参数来标记输入的结束。我们想要读取用户输入的所有内容，直到他们按下<Enter>，所以我们给ReadString一个换行符
      第二个变量的值我们不需要，则用_去处理错误

5、log.Fatal函数 -> 将一条消息记录到终端并停止程序运行。（在这个上下文中，“Fatal”意味着报告一个错误，并“杀死”你的程序。）
   fmt.Println("please enter: ")
   reader := bufio.NewReader(os.Stdin)		// 设置从键盘获取文本的"缓冲读取器"
   input, _ := reader.ReadString('\n')		// 返回用户输入的所有内容，err存储错误返回值
   log.Fatal(err)    // 记录错误返回值
   fmt.Println(input)	   // 错误返回值为nil

6、os.Stat函数 -> 该函数返回一个os.FileInfo值，可能还返回错误值。然后它对FileInfo值调用Size方法来获取文件大小

7、stings.TrimSpace函数 ->删除字符串开头和结尾的所有空白字符（换行符、制表符和常规空格）
   s := "\t hello world \n"
   fmt.Println(strings.TrimSpace(s))      // 输出为hello world

8、块：Go代码可以分为块，即代码段。块通常由大括号（{}）包围
   文件快 -> 包块 -> 函数块 -> if块


   
```


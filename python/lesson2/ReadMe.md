## 开发环境的搭建
	开发环境的搭建就是 **安装python的解释器**
	Python的解释器分类：
		Cpython（官方）
			用C语言编写的Python解释器
	
	步骤：
		>1.下载安装包 -3.x, -2.x 互不兼容
		>2.安装（自定义安装：1.添加到环境变量、2.自定义安装路径）
		>3.打开命令行，输入python
## 基本语法
	1.在Python中严格区分大小写
	2.Python中每一行就是一条语句，每条语句以换行符结束
	3.Python中的每一条语句不要过长（规范中建议每行不要超过80个字符）
		"rulers":[80],
	4.一条语句可以分多行编写，语句后边以\结尾
	5.Python是缩进严格的语言，所以在Python中不要随便缩进
	6.在Python中使用#来表示注释，#后的内容都属于注释，注释的内容将会被解释器所忽略
		我们可以通过注释来对程序进行解释说明

## 字面量和变量
	字面量就是一个一个的值:1,2,'hello'；字面量所表示的意思就是他的字面意思
	变量（variable）变量可以用来保存字面量，并且变量中保存的字面量是不定的
		变量本身没有任何意义，他会根据不同的字面量表示不同的意思
	一般我们在开发时，很少使用字面量；都是将字面量保存到变量中，通过变量来引用字面量

## 变量和标识符
	
# python中使用变量，不要声明，直接为变量赋值即可
a = 10
print(a)

# 不能使用没有进行过赋值的变量
# print(b)报错 NameError: name 'b' is not defined

# python是一个动态类型的语言，可以为变量赋任何类型的值，也可以修改任意变量的值
a = 'hello'
print(a)

# ctrl + / 注释
# ctrl + 回车 换行

# 标识符
# 在Python中所有可以自主命名的内容都属于标识符
# 比如：变量名、函数名、类名
# 标识符命名规范：
# 	1.字母、数字、_ 组成，但是数字不能开头
#	2.不能是保留字或关键字，也不建议使用Python中的函数名：
	# 例如
	# print = 123
	# print(print)
	# 此时print会被覆盖，为int型123
#1. 下划线命名法 max_length
#2. 帕斯卡命名法（大驼峰）MaxLength

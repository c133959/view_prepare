a = {'数学': 95, '语文': 89, '英语': 90}
for k in a.keys():
	print(k, end='---')
print('\n##########')
for v in a.values():
	print(v, end='---')
print('\n##########')
for k,v in a.items():
	print('key:',k, 'value:', v)
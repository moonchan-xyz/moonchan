def snake_to_camel(snake_case):
    words = snake_case.split('_')  # 将下划线分隔的单词拆分为列表
    camel_case = ''.join(word.title() for word in words)
    return camel_case

keys = [
'shortcode',
'url',
'static_url',
'visible_in_picker',
'category',
]

for i in keys:
  i = i.strip()
  k = snake_to_camel(i)
  print(f'{k} string `json:"{i}"`')
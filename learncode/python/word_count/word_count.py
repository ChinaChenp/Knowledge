def word_count(filename):
    try:
        with open(filename) as files:
            contents = files.read()
    except FileNotFoundError:
        #print("file not find")
        pass
        return 0
    else:
        words = contents.split()
        # for word in words:
        #     print(word)
        return len(words)

num = word_count('slice.txt')
print(num)

#多个文件，包含不存在文件
filename_list = ['shit.txt', 'slice.txt', 'fuck.txt']
for filename in filename_list:
    print(word_count(filename))
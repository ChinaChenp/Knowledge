#vim环境搭建方法
- 第一步: [lua安装地址](http://www.lua.org/download.html) 
    - `curl -R -O http://www.lua.org/ftp/lua-5.3.4.tar.gz ` 
    - `tar zxf lua-5.3.3.tar.gz  `
    - `cd lua-5.3.3  `
    - `make linux test`
    - `make install ` 
- 第二步: [vim安装地址](http://www.vim.org/download.php#unix)  
    - `git clone https://github.com/vim/vim.git`
    - `cd vim/src`
    - 指定lua地址和上面安装地址保持一致 
    ```
    ./configure --prefix=/usr --with-features=huge --enable-rubyinterp --enable-pythoninterp --enable-luainterp --with-lua-prefix=/usr/local| grep lua
    ```
    - `make && make install`  
- 第三步：`curl https://j.mp/spf13-vim3 -L > spf13-vim.sh && sh spf13-vim.sh`  
    - `tar -axf vim74.tar.bz2 -C ~` 
    - `mv .vimrc .spf13-vim-3/.vimrc` 替换现有的vimrc  
    - `ln -s /home/chenping/.spf13-vim-3/.vimrc .vimrc` 重新生成软连接
- 第四步：让环境变量生效 `source ~/.bashrc `  
- 第五步：高亮函数
    - mkdir -p ~/.vim/after/syntax  
    - cp cpp.vim c.vim ~/.vim/after/syntax 
    - 下载地址：[cpp.vim](https://github.com/ChinaChenp/Knowledge/tree/master/tools/vim74/cpp.vim)、[c.vim](https://github.com/ChinaChenp/Knowledge/tree/master/tools/vim74/c.vim)

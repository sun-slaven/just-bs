DE:
针对于前端开发环境的安装:
Tips:   http://www.cnblogs.com/Alex--Yang/p/4217577.html
1.安装node.js
2.利用npm安装grunt和bower
3.执行bower install安装第三方包
4.执行grunt命令 执行前端自动化压缩/打包

调试:
  执行grunt watch 监听js和css文件改动,自动合并压缩.第三方库的引入需要手动执行 grunt bower:install 然后 grunt default

Tips:	由于/bower_component和/node_modules文件太多,所以不放入vcs,所以需要一个grunt-bower-task 自动将第三方库的js和css抽取到app/lib/grunt_bower目录下,如果抽取不到一些文件,修改bower_component下每个组件的.bower.json 的"main" 配置

PE:

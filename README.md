DE:

 1. 知识点需求：
	Angularjs 1.x
	Less框架
	npm、bower、grunt等工具的使用
	chrome浏览器常用调试手段
	git版本管理
 2. 开发平台：
	编辑器采用sublime2并安装angularjs插件；
	基于ubuntu14.04系统开发；
	浏览器采用chrome浏览器进行调试。
 3. 针对于前端开发环境的安装:
	Tips:   http://www.cnblogs.com/Alex--Yang/p/4217577.html
	（1）安装node.js
	（2）利用npm安装grunt和bower
	（3）根据项目源码中的bower.json和package.json文件执行“bower install”和“npm install” 命令安装第三方包
	（4）执行“grunt”命令 执行前端自动化压缩/打包
 4. 调试:
	 由于本部分为前台代码，若先进行前台调试，可安装基于nodejs的http-server作为静态服务器，后期可搭建后台基于go语言的内置服务器，然后实现系统整体运行。
	  开发时要首先执行“grunt watch”命令监听js和css文件改动,自动合并压缩.第三方库的引入需要手动执行 “grunt bower:install” 然后 “grunt default”
Tips:	
	  由于/bower_component和/node_modules文件太多,所以不放入vcs,所以需要一个grunt-bower-task 自动将第三方库的js和css抽取到app/lib/grunt_bower目录下,如果抽取不到一些文件,修改bower_component下每个组件的.bower.json 的"main" 配置。
	
PE:
	系统引入的第三方库的js文件自动打包压缩为dest目录下的thirdparty.js文件。项目本身的js文件打包压缩为dest下的application.js文件，css文档合并为application.css

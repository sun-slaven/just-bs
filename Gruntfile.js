//暴露接口
module.exports = function(grunt) {
    //包裹
    grunt.initConfig({
        //载入pkg用于读取某些配置
        pkg: grunt.file.readJSON('package.json'),
        //以下为具体插件配置:

        //js语法错误检测
        jshint: {
            options: {
                asi: true //省略分号不检查
                    //trailing: false
            },
            files: ['Gruntfile.js', 'app/js/**/*.js']
        },
        //将less编译成css
        less: {
            development: {
                options: {
                    paths: ["app/stylesheets/"], //定义@import加载文件的路径
                    compress: true, //压缩编译之后的css文件
                    expand: true
                },
                files: {
                    "dest/application.css": "app/stylesheets/application.less"
                }
            }
        },

        //监听js和css变动(除去第三方库的加入),自动执行合并压缩任务,command: grunt watch
        watch: {
            js: {
                files: ['app/js/**/*.js', 'Gruntfile.js'],
                tasks: ['concat:two', 'uglify:two'],
                options: {
                    //livereload: true,
                }
            },
            less: {
                files: ['app/stylesheets/*.less'],
                tasks: ['less'] //less先合并然后编译成css然后压缩
            }
        },
        //命令: grunt bower:install ,将bower_component中的第三方库抽取js和css到app下
        bower: {
            install: {
                options: {
                    targetDir: "app/lib/grunt_bower/",
                    layout: "byComponent",
                    install: true, //自执行task
                    verbose: false,
                    cleanTargetDir: false,
                    bowerOptions: {}
                }
            }
        },
        concat: {
            options: {
                separator: '\n',
                stripBanners: true
            },
            //合并第三方库js文件
            one: {
                src: [
                    'app/lib/grunt_bower/angular/angular.js',
                    'app/lib/grunt_bower/angular-cookies/angular-cookies.js',
                    'app/lib/grunt_bower/angular-resource/angular-resource.js',
                    'app/lib/grunt_bower/angular-route/angular-route.js',
                    'app/lib/grunt_bower/angular-animate/angular-animate.js',
                    'app/lib/grunt_bower/jquery/jquery.js',
                    'app/lib/grunt_bower/bootstrap/bootstrap.js'
                ],
                dest: 'dest/thirdparty.js'
            },
            //合并app下的js
            two: {
                //由于js之间的依赖关系,src里文件的顺序不能乱改
                src: [
                    'app/js/just/global.js',
                    'app/js/just/angular/*/*.js', //通配符
                    'app/js/just/application.js'
                ],
                //合并后的js目录
                dest: 'dest/all.js',
            }
        },
        //压缩js
        uglify: {
            options: {
                //文件头部输出信息
                banner: '/*! <%= pkg.name %> - v<%= pkg.version %> - ' + '<%= grunt.template.today("yyyy-mm-dd") %>\n' + '* Copyright (c) <%= grunt.template.today("yyyy") %> <%= pkg.author %>; */\n'
            },
            one: {
                src: 'dest/thirdparty.js',
                dest: 'dest/thirdparty.min.js'
            },
            two: {
                src: 'dest/all.js',
                dest: 'dest/all.min.js'
            }
        },
    });
    //载入指定插件任务
    grunt.loadNpmTasks('grunt-contrib-less');
    grunt.loadNpmTasks('grunt-contrib-jshint');
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.loadNpmTasks('grunt-bower-task');
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-concat');
    //注册自定义命名task
    grunt.registerTask('default', ['concat', 'uglify', 'less', 'watch']);
};

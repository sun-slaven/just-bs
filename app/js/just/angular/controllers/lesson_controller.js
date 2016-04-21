GlobalModules.add_controller('lesson')
angular.module('just.controllers.lesson', [])
    .controller('LessonController', ['$rootScope', '$scope',
        function($rootScope, $scope) {
            $scope.lesson = {
                title: "java程序设计",
                teacher: "藤苇",
                length: '1小时12分钟',
                description: "本课程从最基本的概念开始讲起，步步深入，带领大家学习HTML、CSS样式基础知识，了解各种常用标签的意义以及基本用法，后半部分讲解CSS样式代码添加，为后面的案例课程打下基础"
            }
            $scope.items_array = [{
                name: "教师",
                value: "藤苇"
            }, {
                name: "长度",
                value: "1小时12分钟"
            }]
            $scope.chapters = [{
                name: "第1章 Html介绍",
                desc: "本章节主要讲解html和css样式的关系，以及html标签、html文件结构、head标签,最后讲解了在html中的注释代码的作用"
            }, {
                name: "第2章 认识标签(第一部分)",
                desc: "学完这一章节将对标签的使用有了一些初步的认识，可以使用标签制作出一篇简单的文章网页。下一章节我们将进一步学习另外一部分html标签。"
            }, {
                name: "第3章 认识标签(第二部分)",
                desc: "本章节主要讲解列表、div及table标签使用，学完本章，我们可以在网页上展示一些信息列表及表格数据，使网页上的信息更加丰富"
            }]
            $scope.video_url = 'http://7xt49i.com2.z0.glb.clouddn.com/%E5%8D%8A%E7%B3%96%E4%B8%BB%E4%B9%89.mp4'
            $scope.pdf_url = 'http://7xt49i.com2.z0.glb.clouddn.com/AngularJS%E6%9D%83%E5%A8%81%E6%95%99%E7%A8%8B%28www.Linuxidc.com%E6%95%B4%E7%90%86%29.pdf'
            $scope.zip_url = 'http://7xt49i.com2.z0.glb.clouddn.com/pack.zip'
            $scope.need_learn = true;
            if ($scope.need_learn) {
                $scope.btn_content = "开始学习";
            } else {
                $scope.btn_content = "继续学习";
                $scope.progress_info_percent = 10;
                $scope.progress_info_hour = 1;
                $scope.progress_info_minute = 10;
            }

            $scope.start_or_continue = function() {
                if ($scope.need_learn) {
                    $scope.need_learn = false;
                    $scope.btn_content = "继续学习";
                    $scope.progress_info_percent = 10;
                    $scope.progress_info_hour = 1;
                    $scope.progress_info_minute = 10;
                } else {
                    $scope.show_resource = true;
                }
            }

        }
    ])

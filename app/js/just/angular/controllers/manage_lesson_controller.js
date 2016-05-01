GlobalModules.add_controller('manage_lesson')
angular.module('just.controllers.manage_lesson', [])
    .controller('ManageLessonController', ['$rootScope', '$scope', '$log', '$filter', 'LessonsService', 'CommonUtil', 'FileService', function($rootScope, $scope, $log, $filter, LessonsService, CommonUtil, FileService) {
        $scope.active_type = 'creat_lesson'
        $scope.change_active = function(type) {
            $scope.active_type = type;
        }
        $scope.itemsByPage = 2;


        //TODO
        //$scope.useful_lessons = CommonUtil.getMyCreatedLessons()


        $scope.show_modal = function() {
            // var myOtherModal = $rootScope.custom_modal($scope,'')
            // myOtherModal.$promise.then(myOtherModal.show);
            $rootScope.strap_modal({
                title: 'title',
                content: '..',
                show: true
            });
        };

        $scope.edit = function(lesson) {
            $rootScope.strap_modal({
                title: 'title',
                content: '..',
                show: true
            });
        }

        $scope.delete = function(lesson) {
            $scope.modal_ok = function() {
                console.log('ok')
            }
            $rootScope.strap_modal({
                scope: $scope,
                title: '删除课程',
                content: '确定要删除该课程吗',
            });
        }


        //学院专业联动 由于使用ng-include 产生子scope,所以使用#$watch无法达到效果
        $scope.colleges = $rootScope.all_colleges;
        $scope.majors = $rootScope.all_majors;
        $scope.chosen_college = null;
        $scope.chosen_major = null;
        $scope.change_college = function(college) {
            if (college) {
                $scope.majors = college.major_list;
            } else {
                $scope.majors = $rootScope.all_majors;
            }
            $scope.new_lesson.college_id = college.id;
        }
        $scope.change_major = function(major) {
            if ($scope.chosen_college) {
                $scope.majors = $scope.chosen_college.major_list;
                return;
            };
            if (major) {
                angular.forEach($rootScope.all_colleges, function(college) {
                    if (college.id == major.college_id) {
                        $scope.chosen_college = college;
                    };
                })
            } else {
                $scope.majors = $rootScope.all_majors;
            }
            $scope.new_lesson.major_id = major.id;
        }


        //新建lesson
        $scope.new_lesson = {
            icon: null,
            name: "",
            college_id: null,
            major_id: null,
            introduction: '',
            description: '',
            outline_list: [],
            temp_outline_list: [],
            wish: '',
            temp_outline: {
                chapter: '',
                name: '',
                introduction: '',
            },
            uploaded_file: {},
            open_outline_plus_modal: function() {
                $scope.modal_title = "创建提纲";
                $scope.form = $scope.modalForm
                $scope.modal_content_url = '/manage_lesson/_new_lesson_modal_content';
                $scope.modal_type = 'open_outline_plus_modal';
                $scope.new_lesson.temp_outline = {
                        chapter: '',
                        name: '',
                        introduction: '',
                    },
                    $scope.modal_ok = function() {
                        if ($scope.modal_type == 'open_outline_plus_modal') {
                            //如果直接将$scope.new_lesson.temp_outline则放进去的值依然会根据watch改变,和"="一样,故用copy API,=本质是引用复制,后者是新创建一个对象然后进行深度值复制
                            $scope.new_lesson.temp_outline_list.push($scope.new_lesson.temp_outline);
                            $scope.new_lesson.temp_outline = {
                                chapter: '',
                                name: '',
                                introduction: '',
                            }
                        } else {
                            var keepGoing = true;
                            angular.forEach($scope.new_lesson.temp_outline_list, function(item, index) {
                                if (keepGoing) {
                                    $scope.new_lesson.temp_outline_list.splice(index, 1, $scope.new_lesson.temp_outline);
                                    keepGoing = false;
                                }

                            })
                        }
                    }
                $rootScope.strap_modal({
                    scope: $scope, //将scope传入,便可以在modal中调用本scope的方法
                })
            },
            open_outline_edit_modal: function(temp_outline) {
                $scope.modal_title = "修改提纲";
                $scope.modal_content_url = '/manage_lesson/_new_lesson_modal_content';
                $scope.modal_type = 'open_outline_edit_modal';
                $scope.new_lesson.temp_outline = angular.copy(temp_outline);
                $rootScope.strap_modal({
                    scope: $scope
                })
            },
            remove_temp_outline: function(temp_outline) {
                var keepGoing = true;
                angular.forEach($scope.new_lesson.temp_outline_list, function(item, index) {
                    if (keepGoing && angular.equals(item, temp_outline)) {
                        $scope.new_lesson.temp_outline_list.splice(index, 1);
                        keepGoing = false;
                    }
                })
            }

        }

        // collapse default value
        $scope.new_lesson.temp_outline_list.active_outline_index = -1;



        //upload to qiniu
        $scope.selectFiles = [];
        $scope.onFileSelect = function($files) {
            for (var i = 0; i < $files.length; i++) {
                var suffix = $files[i].name.substr($files[i].name.indexOf('.')).toLowerCase(); //文件后缀
                var type = CommonUtil.adjustFileType(suffix);//image/video/file
                var fileObj = {
                    suffix: suffix,
                    type: type
                }
                get_token_and_start( i ,$files[i], fileObj)
            };
        }
        var get_token_and_start = function(i, file, fileObj) {
            FileService.get_file_token(fileObj).$promise.then(function(resp) {
                 var offsetx = $scope.selectFiles.length;
                $scope.selectFiles[i + offsetx] = {
                    file: file,
                    key: resp.key,
                    token: resp.token
                };
                $scope.start($scope.selectFiles[i + offsetx])
            })
        }

        $scope.create_lesson = function(resp) {
            //1.上传文件
            angular.forEach($scope.selectFiles, function(ready_file) {
                $scope.start(ready_file);
            })
            LessonsService.create_lesson($scope.new_lesson, function(resp) {
                console.log(resp)
            })
        }




    }])

GlobalModules.add_controller('manage_lesson')
angular.module('just.controllers.manage_lesson', [])
    .controller('ManageLessonController', ['$rootScope', '$scope', '$qupload', '$log', '$filter', 'LessonsService', 'CommonUtil', function($rootScope, $scope, $qupload, $log, $filter, LessonsService, CommonUtil) {
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

        $scope.create_lesson = function(resp) {
                LessonsService.create_lesson($scope.new_lesson, function() {

                })
            }
            //学院专业联动 由于使用ng-include 产生子scope,所以使用#$watch无法达到效果
        $scope.colleges = $rootScope.all_colleges;
        $scope.majors = $rootScope.all_majors;
        $scope.chosen_college = null;
        $scope.chosen_major = null;
        console.log("test")
        $scope.change_college = function(college) {
            if (college) {
                $scope.majors = college.major_list;
            } else {
                $scope.majors = $rootScope.all_majors;
            }
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
        }



        //新建lesson
        $scope.new_lesson = {
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
                };
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

        var start = function(index) {
            $scope.selectFiles[index].progress = {
                p: 0
            };
            $scope.selectFiles[index].upload = $qupload.upload({
                key: '<your qiniu file key>',
                file: $scope.selectFiles[index].file,
                token: '<your qiniu UpToken>'
            });
            $scope.selectFiles[index].upload.then(function(response) {
                $log.info(response);
            }, function(response) {
                $log.info(response);
            }, function(evt) {
                $scope.selectFiles[index].progress.p = Math.floor(100 * evt.loaded / evt.totalSize);
            });
        };

        $scope.abort = function(index) {
            $scope.selectFiles[index].upload.abort();
            $scope.selectFiles.splice(index, 1);
        };

        $scope.onFileSelect = function($files) {
            var offsetx = $scope.selectFiles.length;
            for (var i = 0; i < $files.length; i++) {
                $scope.selectFiles[i + offsetx] = {
                    file: $files[i]
                };
                start(i + offsetx);
            }
        };

    }])

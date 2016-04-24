GlobalModules.add_controller('manage_lesson')
angular.module('just.controllers.manage_lesson', [])
    .controller('ManageLessonController', ['$rootScope', '$scope', '$qupload', '$log', '$filter', function($rootScope, $scope, $qupload, $log, $filter) {
        $scope.active_type = 'creat_lesson'
        $scope.change_active = function(type) {
            $scope.active_type = type;
        }

        $scope.useful_lessons = [{
            name: 'AngularJS',
            created_time: new Date('2016-04-19 20:56:06'),
            updated_time: new Date('2016-04-20 20:56:06'),
            subscribe_amount: 1
        }, {
            name: 'Go',
            created_time: new Date('2016-04-19 20:55:11'),
            updated_time: new Date('2016-04-20 20:55:11'),
            subscribe_amount: 2
        }, {
            name: 'Bootstrap',
            created_time: new Date('2016-04-19 19:47:33'),
            updated_time: new Date('2016-04-20 19:47:33'),
            subscribe_amount: 3
        }]
        $scope.itemsByPage = 2;

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


        //新建lesson
        $scope.new_lesson = {
            name: "",
            introduction: '',
            outline_list: [],
            temp_outline_list: [],
            wishes: '',
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

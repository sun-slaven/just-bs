GlobalModules.add_controller('manage_lesson')
angular.module('just.controllers.manage_lesson', [])
    .controller('ManageLessonController', ['$rootScope', '$scope', '$log', '$filter', '$q', '$interval', 'LessonService', 'LessonsService', 'CommonUtil', 'FileService', 'QiniuUpload', 'UuidService', function($rootScope, $scope, $log, $filter, $q, $interval, LessonService, LessonsService, CommonUtil, FileService, QiniuUpload, UuidService) {
        $scope.active_type = 'creat_lesson'
        $scope.change_active = function(type) {
            $scope.active_type = type;
        }
        $scope.itemsByPage = 5;


        //TODO
        //$scope.useful_lessons = CommonUtil.getMyCreatedLessons()
        // show all created lessons
        $scope.useful_lessons = [{
            major: {
                id: "fc71592a-0ba7-11e6-b512-3e1d05defe78",
                name: "空军"
            },
            college: {
                id: "5b18f62d-b360-4f1a-9899-5d69a71325a1",
                name: "国防学院"
            },
            comment_sum: 1,
            create_time: "2016-05-02 23:55:23",
            description: "",
            experiment: "string",
            icon: {
                url: "http://7xnz7k.com1.z0.glb.clouddn.com/"
            },
            id: "337c0a43-bdd8-480b-875b-a27668be23fd",
            introduction: "学习基本数据库操作知识",
            mark_sum: 2,
            name: "计算机基础",
            syllabus: "",
            update_time: "2016-05-06 23:13:34",
            video_url: "",
            wish: "希望你们好好学"
        }]


        //update lessons
        $scope.edit = function(lesson) {
            $scope.modal_title = "更新课程";
            $scope.modal_content_url = '/manage_lesson/_update_lesson_modal';
            $scope.edit_lesson = lesson;
            $scope.update_lesson_modal_style = {
                width: '70%'
            };
            $scope.modal_ok = function() {
                if ($scope.upload.get_token_promise_array.length) {
                    $scope.upload.do_upload(function() {
                        LessonService.update_lesson($scope.edit_lesson, function(resp) {
                            $scope.upload.get_token_promise_array = [];
                            $rootScope.reload()
                        })
                    })
                } else {
                    LessonService.create_lesson($scope.edit_lesson, function(resp) {
                        $rootScope.reload()
                    })
                }
            }
            $scope.modal_cancel = function() {
                $scope.update_lesson_modal_style = null;
            }
            $rootScope.strap_modal({
                scope: $scope
            });
        }


        //update lesson outline ,isolate  controller
        $scope.outlineEdit = {
            open_outline_partial: function(lesson) {
                $scope.open_outline_partial = true;
                $scope.outline_edit_lesson = lesson;
                $scope.active_type = 'lesson_outline';
            }
        }

        //delete lessons
        $scope.delete = function(lesson) {
            $scope.modal_ok = function() {
                LessonService.delete_lesson(lesson.id, function(resp) {
                    $rootScope.alert_modal("", "课程:" + lesson.name + " 删除成功")
                })
            }
            $rootScope.strap_modal({
                scope: $scope,
                title: '删除课程',
                content: '确定要删除该课程吗',
            });
        }

        //create lessons
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
        $scope.new_lesson = {
            teacher_id: $rootScope.current_user.id,
            icon_url: null,
            name: "",
            college_id: null,
            major_id: null,
            introduction: '',
            description: '',
            video_url: '',
            //outline_list: [],
            temp_outline_list: [],
            wish: '',
            temp_outline: {
                order: null,
                name: '',
                content: '',
            },
            attachment_list: [],
            open_outline_plus_modal: function() {
                $scope.modal_title = "创建提纲";
                $scope.form = $scope.modalForm
                $scope.modal_content_url = '/manage_lesson/_new_lesson_chapter_modal';
                $scope.modal_type = 'open_outline_plus_modal';
                $scope.new_lesson.temp_outline = {
                        order: null,
                        name: '',
                        introduction: '',
                    },
                    $scope.modal_ok = function() {
                        if ($scope.modal_type == 'open_outline_plus_modal') {
                            //如果直接将$scope.new_lesson.temp_outline则放进去的值依然会根据watch改变,和"="一样,故用copy API,=本质是引用复制,后者是新创建一个对象然后进行深度值复制
                            $scope.new_lesson.temp_outline_list.push($scope.new_lesson.temp_outline);
                            $scope.new_lesson.temp_outline = {
                                order: null,
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
                $scope.modal_content_url = '/manage_lesson/_new_lesson_chapter_modal';
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
        $scope.upload = {
            start_flag: false,
            get_token_promise_array: [],
            //选择多个文件得到token并推入promise数组
            onFileSelect: function($files) {
                for (index in $files) {
                    var suffix_info_obj = QiniuUpload.get_suffix_info_obj($files[index]);
                    var get_token_promise_obj = {
                        suffix_info_obj: suffix_info_obj,
                        file: $files[index],
                        file_name: $files[index].name,
                        get_token_promise: QiniuUpload.get_token(suffix_info_obj).then(function(resp) {
                            return {
                                key: UuidService.newuuid(suffix_info_obj.suffix),
                                token: resp.token
                            };
                        })
                    }
                    this.get_token_promise_array.push(get_token_promise_obj)
                }
            },
            do_upload: function(callback1) {
                this.start_flag = true;
                var upload_fun = function(file, token_obj, callback2) {
                    QiniuUpload.upload(file, token_obj).then(function(resp) {
                        if (callback2) { callback2(resp) };
                        $rootScope.alert_modal("", file.name + "上传成功!");
                    }, function(error) {
                        console.log(error)
                    }, function(evt) {
                        if (file) {
                            file.progress.p = Math.floor(100 * evt.loaded / evt.totalSize);
                        };
                    })
                }
                var use_loop_upload_fun = function(resolve, reject) {
                    $scope.index_flag = 0; //等待标志,解决所有请求还没结束但是index == length执行callback的问题
                    for (index in $scope.upload.get_token_promise_array) {
                        use_loop_upload_fun_by_index(index, resolve, reject);
                    }
                }
                var use_loop_upload_fun_by_index = function(index, resolve, reject) {
                    $scope.upload.get_token_promise_array[index].file.progress = {
                        p: 0
                    };
                    //promise 使用上一个promise的返回结果
                    $scope.upload.get_token_promise_array[index].get_token_promise.then(function(token_obj) {
                        console.log(token_obj)
                        upload_fun($scope.upload.get_token_promise_array[index].file, token_obj, function(resp) {
                            switch ($scope.upload.get_token_promise_array[index].suffix_info_obj.type) {
                                case 'icon':
                                    $scope.new_lesson.icon_url = resp.key
                                    break;
                                case 'video':
                                    $scope.new_lesson.video_url = resp.key //icon url
                                    break;
                                case 'attachment':
                                    var attachment = {
                                        name: $scope.upload.get_token_promise_array[index].file_name,
                                        url: resp.key
                                    }
                                    $scope.new_lesson.attachment_list.push(attachment)
                                    break;
                            }
                            $scope.index_flag += 1;
                            if (index == $scope.upload.get_token_promise_array.length - 1) {
                                var tm = $interval(function() {
                                    if (index == $scope.index_flag - 1) {
                                        //do the last things
                                        $interval.cancel(tm);
                                        $scope.upload.get_token_promise_array = []
                                        resolve();
                                    }
                                }, 500)
                            }

                        })
                    })
                }
                $q(use_loop_upload_fun).then(function() {
                    if (callback1) { callback1() };
                })

            },
            abort: function(file, get_token_promise_array, indexInArray) {
                QiniuUpload.abort(file, get_token_promise_array, indexInArray)
            }
        }

        $scope.create_lesson = function() {
            if ($scope.upload.get_token_promise_array.length) {
                $scope.upload.do_upload(function() {
                    LessonsService.create_lesson($scope.new_lesson, function(resp) {
                        $scope.upload.get_token_promise_array = [];
                        console.log(resp)
                    })
                })
            } else {
                LessonsService.create_lesson($scope.new_lesson, function(resp) {
                    console.log(resp)
                })
            }
        }


    }])

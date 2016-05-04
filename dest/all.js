var GlobalModules = (function() {
    var services = []
    var controllers = []
    var directives = []
    var actions = []

    function add_service(service) { services.push(service) }

    function add_controller(controller) { controllers.push(controller) }

    function add_directive(directive) { directives.push(directive) }

    function add_action(action) { actions.push(action) }

    function get(others) {
        var all = []
        services.forEach(function(service) { all.push("just.services." + service) })
        controllers.forEach(function(controller) { all.push("just.controllers." + controller) })
        directives.forEach(function(directive) { all.push("just.directives." + directive) })
        actions.forEach(function(action) { all.push("just.actions." + action) })
        return all.concat(others)
    }

    return {
        add_service: add_service,
        add_controller: add_controller,
        add_directive: add_directive,
        add_action: add_action,
        get: get
    }
})();


// local storage

var justConst = (function() {
    function get_meta(name) {
        return $('meta[name=' + name + ']').attr('content');
    }
})();

angular.module('just.constants', []).constant('JustConst', justConst);

GlobalModules.add_controller('admin')
angular.module('just.controllers.admin', [])
    .controller('AdminController', ['$rootScope', '$scope', 'LessonService',
        function($rootScope, $scope, LessonService) {
            $scope.active_type = 'manager_user';
            $scope.change_active = function(attr) {
                $scope.active_type = attr;
            }
            $scope.itemsByPage = 10;
            //manager user
            $scope.all_users = [{
                "id": "aa5eba0a-703c-4801-955b-1f44997738fe",
                "name": "小泡子仔",
                "email": "992444037@qq.com",
                "role_name": "STUDENT",
                created_time: new Date()
            }, {
                "id": "aa5eba0a-703c-4801-955b-1f44997738fe",
                "name": "slaven",
                "email": "893196569@qq.com",
                "role_name": "STUDENT",
                created_time: new Date()
            }]


            $scope.delete_user = function(user) {

            }



            //manager lessons
            $scope.all_lessons = [{
                "name": "数据库",
                created_time: new Date(),
                updated_time: new Date(),
                "major": {
                    "id": "0260bb7c-2e93-4a7d-895d-59fac58fdbc6",
                    "name": "物联网工程"
                },
                "college": {
                    "id": "b6a0808f-b87a-44ca-b850-9545a3f3f089",
                    "name": "计算机学院",
                    "major_list": {
                        "id": "0260bb7c-2e93-4a7d-895d-59fac58fdbc6",
                        "name": "物联网工程"
                    }
                },
                "teacher": {
                    "id": "aa5eba0a-703c-4801-955b-1f44997738fe",
                    "name": "小泡子仔",
                    "email": "992444037@qq.com",
                    "role_name": "STUDENT",
                }
            }]

            $scope.delete_lesson = function(lesson) {
                LessonService.delete_lesson(lesson.id,function(resp){
                    $rootScope.alert_modal("", "课程:" + lesson.name + " 删除成功")
                })
            }


        }
    ])

GlobalModules.add_controller('header')
angular.module('just.controllers.header', [])
    .controller('HeaderController', ['$rootScope', '$scope', 'UserService',
        function($rootScope, $scope, UserService) {

            //log out
            $scope.sign_out = function() {
                    $rootScope.confirm_modal("确认退出吗?", $scope, function() {
                        UserService.sign_out($rootScope.current_user,function() {
                            $rootScope.go('/login');
                            $rootScope.current_user = null;
                        })
                    })
                }
                //nav-head controller
            $scope.header_search = {
                input_show: false,
                search_info: '',
                open: function() {
                    if (this.input_show == false && this.search_info == '') {
                        this.input_show = true;
                    } else {
                        if (this.can_submit()) {
                            this.submit()
                        } else {
                            this.close();
                        }
                    }
                },
                close: function() {
                    this.input_show = false
                    this.search_info = ''
                },
                can_submit: function() {
                    if (this.search_info) {
                        return true
                    }
                },
                submit: function() {
                    console.log("submit")
                    this.close()
                }

            }
            $scope.go_me = function() {
                $rootScope.go('/users/' + $rootScope.current_user.id + '/me')
            }

            $scope.go_manager_lessons = function() {
                $rootScope.go('/users/' + $rootScope.current_user.id + '/manage_lesson')
            }
        }
    ])

GlobalModules.add_controller('lesson')
angular.module('just.controllers.lesson', [])
    .controller('LessonController', ['$rootScope', '$scope', '$routeParams', 'LessonService', 'CommentsService', 'MarkService',
        function($rootScope, $scope, $routeParams, LessonService, CommentsService, MarkService) {

            $scope.active_type = 'show_outline'
            $scope.change_active = function(type) {
                $scope.active_type = type;
            }
            if ($routeParams.lesson_id) {
                LessonService.get_lesson($routeParams.lesson_id, function(resp) {
                    console.log(resp)
                    $rootScope.current_lesson = resp
                })
                CommentsService.get_comments($routeParams.lesson_id, function(resp) {
                    console.log(resp)
                    $scope.comments = resp;
                })
            };


            $scope.submit_comment = function() {
                CommentsService.add_comments({
                    course_id: $rootScope.current_lesson.id,
                    content: $scope.my_comment
                }, function(resp) {
                    console.log(resp)
                    $rootScope.alert_modal("提示", "评论成功")
                    CommentsService.get_comments($rootScope.current_lesson.id, function(resp) {
                        $scope.comments = resp;
                    })
                })
            }
            $scope.video_url = 'http://7xnz7k.com1.z0.glb.clouddn.com/cxrs.MP4'
            $scope.pdf_url = 'http://7xt49i.com2.z0.glb.clouddn.com/AngularJS%E6%9D%83%E5%A8%81%E6%95%99%E7%A8%8B%28www.Linuxidc.com%E6%95%B4%E7%90%86%29.pdf'
            $scope.zip_url = 'http://7xt49i.com2.z0.glb.clouddn.com/pack.zip'
            if (!$routeParams.user_id) {
                $scope.need_learn = true;
            } else {
                $scope.need_learn = false;
            }
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
                    mark_and_learn();
                } else {
                    continue_learn()
                }
            }
            var mark_and_learn = function() {
                $scope.need_learn = false;
                $scope.btn_content = "继续学习";
                $scope.progress_info_percent = 10;
                $scope.progress_info_hour = 1;
                $scope.progress_info_minute = 10;
                MarkService.add_mark($rootScope.current_lesson.id, function(resp) {
                    console.log(resp)
                })
            }
            var continue_learn = function() {
                $scope.show_resource = true;

            }
        }
    ])

GlobalModules.add_controller('lessons')
angular.module('just.controllers.lessons', [])
    .controller('LessonsController', ['$rootScope', '$scope', '$timeout', 'LessonsService', 'CollegeMajorService', 'CommonUtil', function($rootScope, $scope, $timeout, LessonsService, CollegeMajorService, CommonUtil) {
        if ($rootScope.all_lessons == undefined) {
            LessonsService.lessons_list(function(resp) {
                $rootScope.all_lessons = resp;
                $scope.lessons = $rootScope.all_lessons
            })
        };
        $scope.lessons = $rootScope.all_lessons

        $scope.colleges = $rootScope.all_colleges;
        $scope.majors = $rootScope.all_majors;
        $scope.chosen_college = null;
        $scope.chosen_major = null;


        $scope.change_college = function(college) {
            if (college) {
                $scope.majors = college.major_list;
                $scope.lessons = CommonUtil.getLessonsByCollege(college.id);
            } else {
                $scope.chosen_major = null;
                $scope.majors = $rootScope.all_majors;
                $scope.lessons = $rootScope.all_lessons;
            }
        }
        $scope.change_major = function(major) {
            if (major) {
                angular.forEach($rootScope.all_colleges, function(college) {
                    if (college.id == major.college_id) {
                        $scope.chosen_college = college;
                        $scope.lessons = CommonUtil.getLessonsByMajor(major.id);
                    };
                })
            } else {
                $scope.majors = $rootScope.all_majors;
                $scope.lessons = $rootScope.all_lessons;
            }
        }

    }])

GlobalModules.add_controller('manage_lesson')
angular.module('just.controllers.manage_lesson', [])
    .controller('ManageLessonController', ['$rootScope', '$scope', '$log', '$filter', '$q', '$interval', 'LessonsService', 'CommonUtil', 'FileService', 'QiniuUpload', 'UuidService', function($rootScope, $scope, $log, $filter, $q, $interval, LessonsService, CommonUtil, FileService, QiniuUpload, UuidService) {
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
                chapter: '',
                name: '',
                introduction: '',
            },
            attachment_list: [],
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
                                case 'image':
                                    $scope.new_lesson.icon_url = resp.key
                                    break;
                                case 'video':
                                    $scope.new_lesson.video_url = resp.key //icon url
                                    break;
                                case 'file':
                                    var attachment = {
                                        name: $scope.upload.get_token_promise_array[index].file_name,
                                        file_url: resp.key
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
                        console.log("callback")
                        console.log(resp)
                    })
                })
            } else {
                LessonsService.create_lesson($scope.new_lesson, function(resp) {
                    console.log("callback")
                    console.log(resp)
                })
            }
        }

        // 18052769341 
        // 231691


    }])

GlobalModules.add_controller('me')
angular.module('just.controllers.me', [])
    .controller('MeController', ['$rootScope', '$scope', 'UserService', 'QiniuUpload', 'CommonUtil', 'FileService',
        function($rootScope, $scope, UserService, QiniuUpload, CommonUtil, FileService) {
            $scope.active_type = 'chosen_lessons'
            $scope.change_active = function(type) {
                $scope.active_type = type;
            }

            UserService.myLessons($rootScope.current_user, function(resp) {
                console.log(resp)
                $scope.chosen_lessons = resp
            })

            //upload avatar
            $scope.upload = {
                selectFileArray: null,
                selectFile: null,
                get_token_promise: {},
                onFileSelect: function($files) {
                    this.selectFileArray = $files;
                    this.selectFile = $files[0];
                    var suffix_info_obj = QiniuUpload.get_suffix_info_obj(this.selectFile);
                    this.get_token_promise = QiniuUpload.get_token(suffix_info_obj).then(function(resp) {
                        return {
                            key: resp.key,
                            token: resp.token
                        };
                    })
                },
                do_upload: function() {
                    $scope.upload.selectFile.progress = {
                        p: 0
                    };
                    var upload_fun = function(token_obj) {
                            console.log(token_obj)
                            QiniuUpload.upload($scope.upload.selectFile, token_obj).then(function(resp) {
                                console.log(resp)
                                    //QiniuUpload.save_file_to_db()
                                    //TODO
                                $rootScope.alert_modal("成功", "上传文件成功");
                            }, function(error) {
                                console.log(error)
                            }, function(evt) {
                                if ($scope.upload.selectFile) {
                                    $scope.upload.selectFile.progress.p = Math.floor(100 * evt.loaded / evt.totalSize);
                                };
                            })
                        }
                        //promise 使用上一个promise的返回结果
                    this.get_token_promise.then(function(token_obj) {
                        upload_fun(token_obj)
                    })
                },
                abort: function() {
                    QiniuUpload.abort(this.selectFile, this.selectFileArray)
                    $scope.upload.selectFile = null;
                }
            }
        }
    ])

GlobalModules.add_controller('user')
angular.module('just.controllers.user', ['ngCookies'])
    .controller('UserController', ['$rootScope', '$scope', 'UserService',
        function($rootScope, $scope, UserService) {
            if ($rootScope.current_user) {
                $rootScope.go('/')
            }
            $scope.form_type = 'login';

            $scope.change_active = function(attr) {
                $scope.form_type = attr;
            }
            $scope.error_infos = []
                //login
            $scope.user = {
                email: $rootScope.get_storage('email') || '',
                password: $rootScope.get_storage('password') || '',
                remember_me: true
            }
            $scope.can_submit = function() {
                if ($scope.user.name == '') {
                    return false
                };
                if ($scope.user.password == '') {
                    return false
                };
                return true
            }
            $scope.submit = function() {
                    if ($scope.can_submit()) {
                        if ($scope.form_type == 'login') {
                            if ($scope.user.remember_me) {
                                $rootScope.set_storage('email', $scope.user.email)
                                $rootScope.set_storage('password', $scope.user.password)
                            } else {
                                $rootScope.set_storage('email', null)
                                $rootScope.set_storage('password', null)
                            }
                            UserService.sign_in($scope.user, function(resp) {
                                $rootScope.set_cache('token', resp.token)
                                //管理员账户单独一个界面
                                $rootScope.go("/admin/show");
                                // if (resp.user.role_name == 'ADMIN') {
                                //     $rootScope.go("/admin/show");
                                // }else{
                                //     $rootScope.go("/users/" + $rootScope.current_user.id + "/me");
                                // }
                            })

                        }
                    }
                }
                //register
            $scope.register = {
                name: '',
                email: '',
                password: '',
                password_again: ''
            }
            $scope.can_register = function() {
                if ($scope.register.name && $scope.register.email) {
                    if ($scope.register.password && ($scope.register.password === $scope.register.password_again)) {
                        return true;
                    };
                };
                return false
            }
            $scope.register_ok = function() {
                if ($scope.can_register()) {
                    UserService.register($scope.register, function(resp) {
                        $rootScope.alert_modal("提示", "邮件已经发送到邮箱,请登陆邮箱确认后登陆")
                    })
                }
            };

        }
    ])

GlobalModules.add_directive('just_unique_validation')
angular.module('just.directives.just_unique_validation', [])
    .directive('justUniqueValidation', ['$rootScope', function($rootScope) {
        //使用必须为: just_chapter_exist_validation
        return {
            require: 'ngModel', //要求节点上必须使用到ng-modal
            restrict: 'A',
            scope: {
                target_array_or_object: '=targetArrayOrObject'
            },
            link: function(scope, iElm, iAttrs, controller) {
                iElm.on('input', function(event) {
                    scope.$apply(function() {
                        var keepGoing = true;
                        angular.forEach(scope.target_array_or_object, function(item) {
                            if (keepGoing) {
                                if (item.chapter) {
                                    if (item.chapter == iElm.val()) {
                                        controller.$setValidity('chapterExistValidation', false)
                                        keepGoing = false;
                                    } else {
                                        controller.$setValidity('chapterExistValidation', true)
                                    }
                                }
                                //others
                            };

                        })
                    })
                })
            }
        };
    }]);

GlobalModules.add_directive('just_video')
angular.module('just.directives.just_video', [])
    .directive('justVideo', ['$rootScope', function($rootScope) {
        // Runs during compile
        return {
            // name: '',
            // priority: 1,
            // terminal: true,
            scope: {
                video_url: '=videoUrl' //directive中的属性必须在此处''内是驼峰式写法
            },
            // controller: function($scope, $element, $attrs, $transclude) {},
            // require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
            restrict: 'E', // E = Element, A = Attribute, C = Class, M = Comment
            // template: '',
            templateUrl: '/app/partials/directives_template/just_video.html',
            replace: true,
            // transclude: true,
            // compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
            link: function($scope, iElm, iAttrs, controller) {
                // $('video#test-vid').bind("progress", function(e) {
                //     console.log(e.total + ' ' + e.loaded + ' ' + e.lengthComputable);
                // });
            }
        };
    }]);

GlobalModules.add_directive('just_card')
angular.module('just.directives.just_card', [])
    //directive的命名必须小写开头,使用为<just-card>
    .directive('justCard', ['$rootScope','$location', function($rootScope,$location) {
        // Runs during compile
        return {
            // name: '',
            // priority: 1,
            // terminal: true,
            scope: {
                lesson: '='
            }, // {} = isolate, true = child, false/undefined = no change
            // controller: function($scope, $element, $attrs, $transclude) {},
            // require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
            restrict: 'E', // E = Element, A = Attribute, C = Class, M = Comment
            // template: '',
            templateUrl: '/app/partials/directives_template/just_card.html',
            replace: true,
            transclude: true,
            // compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
            link: function($scope, iElm, iAttrs, controller) {
                $scope.choose_lesson = function(lesson) {
                    if ($location.path() == '/lessons/index') {{
                        $rootScope.go('/lessons/'+lesson.id+'/show')
                    }}else{
                        $rootScope.go('/users/'+$rootScope.current_user.id+ '/lessons/'+lesson.id+'/show')
                    }
                }
            }
        };
    }]);

angular.module('just.filters', [])
    .filter('cut', function() {
        return function(value, wordwise, max, tail) {
            if (!value) return '';

            max = parseInt(max, 10);
            if (!max) return value;
            if (value.length <= max) return value;

            value = value.substr(0, max);
            if (wordwise) {
                var lastspace = value.lastIndexOf(' ');
                if (lastspace != -1) {
                    value = value.substr(0, lastspace);
                }
            }

            return value + (tail || ' …');
        };
    })
    .filter('password', [function() {
        return function(str) {
            if (!str) return '';
            var result = ''
            for (i = 0; i < str.length; i++) {
                result += '*'
            }
            return result
        }
    }]);

angular.module('just.route_config', []).
provider('RouteConfig', function() {
    this.$get = function() {
        var all_configs = [];
        var partial_url = function(url) {
            return '/app/partials/' + url + '.html';
        }
        var base_config = [{
            path: '/',
            templateUrl: partial_url('user/login'),
            controller: 'UserController'
        },{
            path: '/login',
            templateUrl: partial_url('user/login'),
            controller: 'UserController'
        }];

        var me_config = [{
            path: '/users/:user_id/me',
            templateUrl: partial_url('me/show'),
            controller: 'MeController'
        }, {
            path: '/users/:user_id/lessons/:lesson_id/show',
            templateUrl: partial_url('lessons/show'),
            controller: 'LessonController'
        }]

        var lessons_config = [{
            path: '/lessons/index',
            templateUrl: partial_url('lessons/index'),
            controller: 'LessonsController'
        }, {
            path: '/lessons/:lesson_id/show',
            templateUrl: partial_url('lessons/show'),
            controller: 'LessonController'
        }]

        var manager_lesson_config = [{
            path: '/users/:user_id/manage_lesson',
            templateUrl: partial_url('manage_lesson/show'),
            controller: 'ManageLessonController'
        }]

        var admin_config = [{
            path: '/admin/show',
            templateUrl: partial_url('admin/show'),
            controller: 'AdminController'
        }]

        add_config(base_config);
        add_config(me_config);
        add_config(lessons_config);
        add_config(manager_lesson_config)
        add_config(admin_config)

        function add_config(config) {
            all_configs = all_configs.concat(config);
        }

        function get() {
            return all_configs;
        }

        return {
            get: get
        }
    }
});

GlobalModules.add_service('anchorSmoothScroll')
angular.module('just.services.anchorSmoothScroll', [])
    .service('AnchorSmoothScrollService', function() {
        this.scrollTo = function(eID) {
            var startY = currentYPosition();
            var stopY = elmYPosition(eID);
            var distance = stopY > startY ? stopY - startY : startY - stopY;

            // if (distance < 100) {
            //     scrollTo(0, stopY);
            //     return;
            // }
            var speed = Math.round(distance / 100);
            if (speed >= 20) speed = 20;
            var step = Math.round(distance / 50);
            var leapY = stopY > startY ? startY + step : startY - step;
            var timer = 0;
            if (stopY > startY) {
                for (var i = startY; i < stopY; i += step) {
                    setTimeout("window.scrollTo(0, " + leapY + ")", timer * speed);
                    leapY += step;
                    if (leapY > stopY) leapY = stopY;
                    timer++;
                }
                return;
            }
            for (var i = startY; i > stopY; i -= step) {
                setTimeout("window.scrollTo(0, " + leapY + ")", timer * speed);
                leapY -= step;
                if (leapY < stopY) leapY = stopY;
                timer++;
            }

            function currentYPosition() {
                // Firefox, Chrome, Opera, Safari
                if (self.pageYOffset) return self.pageYOffset;
                // Internet Explorer 6 - standards mode
                if (document.documentElement && document.documentElement.scrollTop)
                    return document.documentElement.scrollTop;
                // Internet Explorer 6, 7 and 8
                if (document.body.scrollTop) return document.body.scrollTop;
                return 0;
            }

            function elmYPosition(eID) {
                var elm = document.getElementById(eID);
                var y = elm.offsetTop;
                var node = elm;
                while (node.offsetParent && node.offsetParent != document.body) {
                    node = node.offsetParent;
                    y += node.offsetTop;
                }
                return y;
            }
        };
    });

GlobalModules.add_service('college_major')
angular.module('just.services.college_major', []).
factory('CollegeMajorService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var college_majorAPI = $resource('/api/v1/colleges', {}, {
            get_college_major: {method: 'get' , isArray: true}
        })


        function get_college_major(success) {
            college_majorAPI.get_college_major({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            get_college_major: get_college_major
        }
    }
])

GlobalModules.add_service('comments')
angular.module('just.services.comments', []).
factory('CommentsService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var commentsAPI = $resource('/api/v1/courses/:course_id/comments', {course_id: '@course_id'}, {
            delete_comments: {method: 'delete' , isArray: false},
            get_comments: {method: 'get' , isArray: true},
            add_comments: {method: 'post' , isArray: false},
        })


        function delete_comments(lesson_id,success) {

            commentsAPI.delete_comments({course_id:  lesson_id}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function get_comments(lesson_id,success) {
            commentsAPI.get_comments({}, {course_id:  lesson_id}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function add_comments(obj,success) {
            commentsAPI.add_comments({}, {
                course_id: obj.course_id,
                content: obj.content
            }, function(resp) {
                if (success) { success(resp) }
            })
        }        


        return {
            delete_comments: delete_comments,
            get_comments: get_comments,
            add_comments: add_comments
        }
    }
])

GlobalModules.add_service('common_util')
angular.module('just.services.common_util', []).
factory('CommonUtil', ['$rootScope', 'LessonsService',
    function($rootScope, LessonsService) {

        var init_if_need = function() {
            if ($rootScope.all_lessons == undefined) {
                LessonsService.lessons_list(function(resp) {
                    $rootScope.all_lessons = resp;
                })
            };
        }

        var getLessonsByCollege = function(college_id) {
            var college_lessons = [];
            init_if_need()
            angular.forEach($rootScope.all_lessons, function(lesson) {
                if (lesson.college.id == college_id) {
                    college_lessons.push(lesson);
                };
            })
            return college_lessons;
        }

        var getLessonsByMajor = function(Major_id) {
            var major_lessons = [];
            init_if_need()
            angular.forEach($rootScope.all_lessons, function(lesson) {
                if (lesson.major.id == Major_id) {
                    major_lessons.push(lesson);
                };
            })
            return major_lessons;
        }

        var getMyCreatedLessons = function() {
            var created_lessons = [];
            init_if_need()
            angular.forEach($rootScope.all_lessons, function(lesson) {
                if (lesson.teacher.id == $rootScope.current_user.id) {
                    created_lessons.push(lesson);
                };
            })
            return created_lessons;
        }

        var adjustFileType = function(fileSuffix){
            image_array = ['.bmp','.png','.gif','.jpg','.jpeg']
            video_array = ['.vob','.avi','.rmvb','.asf','.wmv','.mp4']
            if (image_array.indexOf(fileSuffix) > -1) {
                return 'image'
            }
            if (video_array.indexOf(fileSuffix) > -1) {
                return 'video'
            };
            return 'file'
        }



        return {
            getLessonsByCollege: getLessonsByCollege,
            getLessonsByMajor: getLessonsByMajor,
            getMyCreatedLessons: getMyCreatedLessons,
            adjustFileType: adjustFileType
        }
    }
])

GlobalModules.add_service('file')
angular.module('just.services.file', []).
factory('FileService', ['$rootScope', '$resource',
    function($rootScope, $resource) {
        var fileTokenAPI = $resource('/api/v1/files/tokens', {}, {
            file_token: { method: 'post' },
        })

        var fileAPI = $resource('/api/v1/files', {}, {
            save_file: { method: 'post' },
        })

        function get_file_token(fileObj) {
            return fileTokenAPI.file_token({},fileObj)
        }

        function save_file(success) {
            fileAPI.save_file({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            get_file_token: get_file_token,
            save_file: save_file
        }
    }
])

GlobalModules.add_service('lesson')
angular.module('just.services.lesson', []).
factory('LessonService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {
        var lessonAPI = $resource('/api/v1/courses/:course_id', {course_id : '@course_id'}, {
            delete_lesson: {method: 'delete' , isArray: false},
            get_lesson: {method: 'get' , isArray: false},
            delete_lesson: {method: 'delete' , isArray: false},
        })


        function get_lesson(course_id,success) {
            lessonAPI.get_lesson({}, {course_id: course_id}, function(resp) {
                if (success) { success(resp) }
            })
        }



        function delete_lesson(course_id,success) {
            lessonAPI.delete_lesson({}, {
                course_id: course_id
            }, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            get_lesson: get_lesson,
            delete_lesson: delete_lesson
        }
    }
])

GlobalModules.add_service('lessons')
angular.module('just.services.lessons', []).
factory('LessonsService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var lessonsAPI = $resource('/api/v1/courses', {}, {
            lessons_list: { method: 'get', isArray: true, cache: false },
            create_lesson: { method: 'post', isArray: false },
        })

        function lessons_list(success) {
            lessonsAPI.lessons_list({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function create_lesson(new_lesson, success) {

            var lesson_obj = {
                name: new_lesson.name,
                teacher_id: new_lesson.teacher_id,
                icon_url: new_lesson.icon_url,
                video_url: new_lesson.video_url,
                description: new_lesson.description,
                introduction: new_lesson.introduction,
                wish: new_lesson.wish,
                college_id: new_lesson.college_id,
                major_id: new_lesson.major_id,
                outline: new_lesson.temp_outline_list,
                attachment: new_lesson.attachment_list
            }
            console.log(lesson_obj)
            // lessonsAPI.create_lesson({}, lesson_obj, function(resp) {
            //     if (success) { success(resp) }
            // })
        }



        return {
            lessons_list: lessons_list,
            create_lesson: create_lesson
        }
    }
])

GlobalModules.add_service('mark')
angular.module('just.services.mark', []).
factory('MarkService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var marksAPI = $resource('/api/v1/courses/:course_id/marks', { course_id: '@course_id' }, {
            add_mark: { method: 'post', isArray: false },
        })

        var markAPI = $resource('/api/v1/courses/:course_id/marks/', {}, {
            cancel_mark: { method: 'post', isArray: false },
        })

        function add_mark(course_id, success) {
            console.log(course_id)
            marksAPI.add_mark({}, {
                course_id: course_id
            }, function(resp) {
                if (success) { success(resp) }
            })
        }

        function cancel_mark(course_id, success) {
            markAPI.cancel_mark({}, {
                course_id: course_id
            }, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            add_mark: add_mark,
            cancel_mark: cancel_mark
        }
    }
])

GlobalModules.add_service('qiniu_upload')
angular.module('just.services.qiniu_upload', []).
factory('QiniuUpload', ['$rootScope', '$resource', '$http', '$qupload', 'FileService', 'CommonUtil',

    function($rootScope, $resource, $http, $qupload, FileService, CommonUtil) {

        var get_suffix_info_obj = function(selectFile) {
            var suffix = selectFile.name.substr(selectFile.name.indexOf('.')).toLowerCase(); //文件后缀
            var type = CommonUtil.adjustFileType(suffix); //image/video/file
            var fileObj = {
                suffix: suffix,
                type: type
            }
            return fileObj;
        }

        var get_token = function(fileObj) {
            return FileService.get_file_token(fileObj).$promise
        }

        var upload = function(file, token_obj) {
            file.upload = $qupload.upload({
                file: file,
                key: token_obj.key,
                token: token_obj.token
            });
            return file.upload //返回一个promise
        };

        var abort = function(file, files, indexAttr) {
            if (indexAttr > -1) {
                files.splice(index, 1)
            } else {
                for (index in files) {
                    if (files[index].name == file.name) {
                        files.splice(index, 1);
                        break;
                    };
                }
            }
            file.upload.abort();
        };

        var save_file_to_db = function() {}

        return {
            get_suffix_info_obj: get_suffix_info_obj,
            get_token: get_token,
            upload: upload,
            abort: abort,
            save_file_to_db: save_file_to_db
        }
    }
])

GlobalModules.add_service('user')
angular.module('just.services.user', []).
factory('UserService', ['$rootScope', '$resource', '$cookies',
    function($rootScope, $resource, $cookies) {
        var userAPI = $resource('/api/v1/tokens', {}, {
            sign_in: { method: 'post' },
            signout: { method: 'delete' }
        })
        var registerAPI = $resource('/api/v1/users', {}, {
            register: { method: 'post' }
        })

        var myLessonsAPI = $resource('/api/v1/users/:user_id/courses', { user_id: '@user_id' }, {
                myLessons: { method: 'get', isArray: true }
            })

        var UserInfoApi = $resource('/api/v1/users/:user_id/', {}, {
            updateUser: {method: 'patch'},
            getUser: {method: 'get'}    //暂时无用
        })
            //992444037@qq.com  123456   STUDENT
            //158274194@qq.com   123456  TEACHER
            //893196569@qq.com  123456   ADMIN
        function sign_in(user, success) {
            userAPI.sign_in({}, {
                email: user.email || $rootScope.get_storage('email'),
                password: user.password || $rootScope.get_storage('password')
            }, function(resp) {
                set_token(resp.token);
                set_user(resp.user);
                if (success) { success(resp) }
            }, function(error) {
                console.log(error)
            })
        }

        function sign_out(user, success) {
            userAPI.sign_out({}, {
                user: user
            }, function(resp) {
                set_user(null)
                set_token(null);
                $rootScope.clear_cache()
                $rootScope.reload(); //route reload
                if (success) { success(resp) }
            })
        }

        function register(user, success) {
            registerAPI.register({}, {
                user_name: user.name,
                email: user.email,
                password: user.password,
                password2: user.password_again
            }, function(resp) {
                if (success) { success(resp) }
            })
        }

        function set_user(new_user) {
            $rootScope.current_user = new_user
            $cookies.putObject('current_user', new_user)
        }

        function set_token(token) {
            $cookies.loginTokenCookie = token;
        }

        function myLessons(user, callback) {
            myLessonsAPI.myLessons({}, { user_id: user.id }, function(resp) {
                if (callback) {
                    callback(resp)
                };
            })
        }

        function updateUser(user,callback){
            UserInfoApi.updateUser({},{
                user_id: user.id
            },function(resp){
                console.log(resp)
            })
        }


        return {
            sign_in: sign_in,
            sign_out: sign_out,
            register: register,
            myLessons: myLessons
        }
    }
])

GlobalModules.add_service('uuid')
angular.module('just.services.uuid', []).
factory('UuidService', [
    function() {
        function s4() {
            return Math.floor((1 + Math.random()) * 0x10000)
                .toString(16)
                .substring(1);
        }

        return {
            newuuid: function(prefix) {
                // http://www.ietf.org/rfc/rfc4122.txt
                var s = [];
                var hexDigits = "0123456789abcdef";
                for (var i = 0; i < 36; i++) {
                    s[i] = hexDigits.substr(Math.floor(Math.random() * 0x10), 1);
                }
                s[14] = "4"; // bits 12-15 of the time_hi_and_version field to 0010
                s[19] = hexDigits.substr((s[19] & 0x3) | 0x8, 1); // bits 6-7 of the clock_seq_hi_and_reserved to 01
                s[8] = s[13] = s[18] = s[23] = "-";
                return s.join("") + prefix;
            },
            newguid: function(prefix) {
                return s4() + s4() + '-' + s4() + '-' + s4() + '-' +
                    s4() + '-' + s4() + s4() + s4() + prefix;
            }
        }
    }
])

var version_timestamp = "?v" + Date.parse(new Date());
/**
 *  Module
 *
 * application.js
 */
angular.module('just', GlobalModules.get([
    'ngRoute', 'ngResource', 'ngCookies', 'ngAnimate', 'ui.bootstrap', 'smart-table', 'angularQFileUpload', 'mgcrea.ngStrap', 'angularLocalStorage',
    'just.route_config',
    'just.constants',
    'just.filters'
])).config(['$httpProvider', '$routeProvider', '$locationProvider', '$sceDelegateProvider', 'RouteConfigProvider', '$modalProvider',
    function($httpProvider, $routeProvider, $locationProvider, $sceDelegateProvider, RouteConfigProvider, $modalProvider) {
        //同源策略:在本站访问外站资源时,需要添加到信任名单中,不然就会加载错误.video
        $sceDelegateProvider.resourceUrlWhitelist([
            'self', 'http://7xt49i.com2.z0.glb.clouddn.com/**',
            'http://7xnz7k.com1.z0.glb.clouddn.com/**'
        ]);
        //使用过滤器将所有请求都加上token
        $httpProvider.interceptors.push(function($cookies) {
            return {
                'request': function(config) {
                    config.headers['token'] = $cookies.loginTokenCookie;
                    return config;
                }
            };
        });

        var all_configs = RouteConfigProvider.$get().get()
        angular.forEach(all_configs, function(conf) {
            $routeProvider.when(conf.path, {
                templateUrl: conf.templateUrl + version_timestamp,
                controller: conf.controller
            })
        })
        $routeProvider.otherwise({
            redirectTo: '/login'
        });
        //disable get method cache globally
        //initialize get if not there
        if (!$httpProvider.defaults.headers.get) {
            $httpProvider.defaults.headers.get = {};
        }
        //disable IE ajax request caching
        $httpProvider.defaults.headers.get['If-Modified-Since'] = 'Mon, 26 Jul 1997 05:00:00 GMT';
        $httpProvider.defaults.headers.get['Cache-Control'] = 'no-cache';
        $httpProvider.defaults.headers.get['Pragma'] = 'no-cache';


        // $locationProvider.html5Mode(true); // remove # in the url
        // $locationProvider.hashPrefix = '!';
        //修改modal的全局配置
        angular.extend($modalProvider.defaults, {
            animation: 'am-fade-and-scale',
            html: true,
            templateUrl: '/app/partials/common_modal.html',
            show: true
        });
    }
]).run(['$rootScope', '$location', '$routeParams', '$modal', '$cacheFactory', 'AnchorSmoothScrollService', 'storage', 'CollegeMajorService', 'LessonsService', '$alert', 'UserService', '$cookies', function($rootScope, $location, $routeParams, $modal, $cacheFactory, AnchorSmoothScrollService, storage, CollegeMajorService, LessonsService, $alert, UserService, $cookies) {
    //路由以及$location
    $rootScope.partial = function(partial_name) {
        return "app/partials/" + partial_name + ".html" + version_timestamp;
    }
    $rootScope.go = function(url) {
        $location.url(url)
    }
    $rootScope.reload = function(bool) {
        if (bool) { location.reload() } else { $route.reload() }
    }
    $rootScope.location_path = function() {
        return $location.path();
    }

    //cache
    var cache = $cacheFactory('just_cache')
    $rootScope.get_cache = function(key) {
        return cache.get(key);
    }
    $rootScope.set_cache = function(key, value) {
        cache.put(key, value);
    }
    $rootScope.clear_cache = function() {
            if (cache.get('$http')) {
                cache.get('$http').removeAll();
            };
            cache.removeAll();
        }
        //localStorage
    $rootScope.get_storage = function(key) {
        return storage.get(key);
    }
    $rootScope.set_storage = function(key, value) {
        storage.set(key, value);
    }
    $rootScope.clear_storage = function() {
        storage.clearAll();
    }

    //role
    $rootScope.is_student = function() {
        return $rootScope.current_user.role_name == 'STUDENT';
    }
    $rootScope.is_teacher = function() {
        return $rootScope.current_user.role_name == 'TEACHER';
    }
    $rootScope.is_admin = function() {
        return $rootScope.current_user.role_name == 'ADMIN';
    }



    //滚动到顶部
    $rootScope.scrollTo = function(eID) {
        AnchorSmoothScrollService.scrollTo(eID);
    }

    //bootstrap  customer modals
    $rootScope.strap_modal = function(modal_obj) {
        return $modal(modal_obj)
    }
    $rootScope.confirm_modal = function(content, scope, success) {
            scope.modal_ok = success;
            $rootScope.strap_modal({
                content: content,
                title: "提示".concat(' <i class="fa fa-info-circle" aria-hidden="true"></i>'),
                scope: scope
            });
        }
        //alert
    $rootScope.alert_modal = function(title, content) {
        return $alert({
            title: title.concat(' <i class="fa fa-info-circle" aria-hidden="true"></i>'),
            content: content,
            placement: 'top-right',
            type: 'info',
            show: true
        })
    }

    // 防止页面刷新,从cookie里取出当前对象.cookie在页面刷新时并不会清空
    if ($cookies.getObject('current_user')) {
        $rootScope.current_user = $cookies.getObject('current_user');
    }
    //init college major info
    if ($rootScope.college_major == undefined) {
        $rootScope.all_colleges = []
        $rootScope.all_majors = []
        CollegeMajorService.get_college_major(function(response) {
            for (var i = 0; i < response.length; i++) {
                $rootScope.all_colleges.push(response[i])
                for (index in response[i].major_list) {
                    response[i].major_list[index].college_id = response[i].id;
                    $rootScope.all_majors.push(response[i].major_list[index])
                }
            }
        });
    }

    $rootScope.$on('$routeChangeSuccess', function(evt, next, current) {
        //refuse change the url to /# then header show
        if ($location.path() == '/' || $location.path() == '/login') {
            $rootScope.current_user = null;
        }
    })


}])

GlobalModules.add_controller('admin')
angular.module('just.controllers.admin', [])
    .controller('AdminController', ['$rootScope', '$scope', '$timeout', 'LessonService', 'UserService', 'LessonsService',
        function($rootScope, $scope, $timeout, LessonService, UserService, LessonsService) {
            $scope.active_type = 'manager_user';
            $scope.change_active = function(attr) {
                $scope.active_type = attr;
            }

            $scope.itemsByPage = 10;
            $scope.filterGetters = {
                college_name: function(value) {
                    return value.college.name
                },
                major_name: function(value) {
                    return value.major.name
                },
                college_name: function(value) {
                    return value.college.name
                },
                teacher_name: function(value) {
                    if (value.teacher == null) {
                        return null;
                    };
                    return value.teacher.name
                }
            }

            //manager user
            var getAllUsers = function() {
                UserService.getAllUsers(function(resp) {
                    $scope.all_users = resp
                })
            }
            getAllUsers();

            $scope.initPassword = function(user) {
                UserService.initPassword(user, function(resp) {
                    $rootScope.alert_modal('', '密码重置成功,已将重置密码发送到用户邮箱')
                })
            }
            $scope.delete_user = function(user) {
                UserService.deleteUser(user, function(resp) {
                    getAllUsers();
                    $rootScope.alert_modal('', '该用户已被删除')
                })
            }


            //manager courses
            $scope.all_displayed_lessons = [] //st-table needs to show existed values
            var getAllLessons = function() {
                LessonsService.lessons_list(function(resp) {
                    $scope.all_asy_lessons = [].concat(resp) //st-safe-src needs to show ajax values
                })
            }
            getAllLessons();

            $scope.delete_lesson = function(lesson) {
                LessonService.delete_lesson(lesson.id, function(resp) {
                    getAllLessons();
                    $rootScope.alert_modal("", "课程:" + lesson.name + " 删除成功")
                })
            }

        }
    ])

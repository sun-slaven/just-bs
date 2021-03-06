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
                                //管理员账户单独一个界面
                                if (resp.user.role_name == 'ADMIN') {
                                    $rootScope.go("/admin/show");
                                }else{
                                    $rootScope.go("/users/" + $rootScope.current_user.id + "/me");
                                }
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

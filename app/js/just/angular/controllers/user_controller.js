GlobalModules.add_controller('user')
angular.module('just.controllers.user', ['ngCookies'])
    .controller('UserController', ['$rootScope', '$scope', 'UserService',
        function($rootScope, $scope, UserService) {
            if ($rootScope.user) {
                $rootScope.go('/')
            }
            $scope.form_type = 'login';

            $scope.change_active = function(attr) {
                $scope.form_type = attr;
            }
            $scope.error_infos = []
                //login
            $scope.user = {
                email: $rootScope.get_cookie('email') || '',
                password: $rootScope.get_cookie('password') || '',
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
                                $rootScope.set_cookie('email', $scope.user.email)
                                $rootScope.set_cookie('password', $scope.user.password)
                            };
                            UserService.sign_in($scope.user, function(resp) {
                                $rootScope.icon = resp.user.icon;
                                $rootScope.show_header = true;
                                $rootScope.go("/users/" + $rootScope.user.id + "/me");
                            })

                        }
                    }
                }
                //register
            $scope.register = {
                email: '',
                password: '',
                password_again: ''
            }
            $scope.can_register = function() {
                if ($scope.register.email) {
                    if ($scope.register.password && ($scope.register.password === $scope.register.password_again)) {
                        return true;
                    };
                };
                return false
            }
            $scope.register = function() {
                if ($scope.can_register()) {
                    $rootScope.current_user = $scope.register
                    $rootScope.go("/users/1/show")
                };
            }

        }
    ])

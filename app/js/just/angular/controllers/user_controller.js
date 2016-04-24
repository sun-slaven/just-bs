GlobalModules.add_controller('user')
angular.module('just.controllers.user', ['ngCookies'])
    .controller('UserController', ['$rootScope', '$scope', '$cookies', 'UserService',
        function($rootScope, $scope, $cookies, UserService) {
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
                name: '',
                password: '',
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
                            UserService.sign_in($scope.user, function(resp) {
                                $rootScope.show_header = true;
                                $rootScope.go("/users/1/show")
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

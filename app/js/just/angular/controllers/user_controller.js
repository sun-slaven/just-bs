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

            $scope.user = {
                name: '',
                password: '',
                remember_me: true
            }
            $scope.can_submit = function() {
                if ($scope.user.name == '') {
                    return false };
                if ($scope.user.password == '') {
                    return false };
                return true
            }
            $scope.submit = function() {
                if ($scope.can_submit()) {
                    if ($scope.form_type == 'login') {
                        // UserService.login($scope.user, function(resp) {
                        //     //$rootScope.go("")
                        // })
                    }
                }
            }
        }
    ])

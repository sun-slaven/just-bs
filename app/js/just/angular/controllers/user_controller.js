GlobalModules.add_controller('user')
angular.module('just.controllers.user', ['ngCookies'])
    .controller('UserController', ['$rootScope', '$scope', '$cookies', 'UserService',
        function($rootScope, $scope, $cookies, UserService) {
            return
        }
    ])

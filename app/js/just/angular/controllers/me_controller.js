GlobalModules.add_controller('me')
angular.module('just.controllers.me', ['ngCookies'])
    .controller('MeController', ['$rootScope', '$scope', '$cookies',
        function($rootScope, $scope, $cookies) {
        }
    ])

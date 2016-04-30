GlobalModules.add_controller('me')
angular.module('just.controllers.me', [])
    .controller('MeController', ['$rootScope', '$scope', '$cookies', 'UserService',
        function($rootScope, $scope, $cookies, UserService) {
            $scope.active_type = 'chosen_lessons'
            $scope.change_active = function(type) {
                $scope.active_type = type;
            }

            UserService.myLessons($rootScope.current_user, function(resp) {
                $scope.chosen_lessons = resp
            })



        }
    ])

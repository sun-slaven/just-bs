GlobalModules.add_controller('me')
angular.module('just.controllers.me', ['ngCookies'])
    .controller('MeController', ['$rootScope', '$scope', '$cookies',
        function($rootScope, $scope, $cookies) {
            $scope.active_type = 'chosen_lessons'
            $scope.change_active = function(type) {
                $scope.active_type = type;
            }
            $scope.chosen_lessons = [{
                name: "course1",
                desc: "desc1",
                img_url: "app/images/login_background.jpg",
                duration: "duration1"
            }, {
                name: "course2",
                desc: "desc2",
                img_url: "app/images/login_background.jpg",
                duration: "duration2"
            }, {
                name: "course3",
                desc: "desc3",
                img_url: "app/images/login_background.jpg",
                duration: "duration3"
            }, {
                name: "course4",
                desc: "desc4",
                img_url: "app/images/login_background.jpg",
                duration: "duration4"
            }, {
                name: "course5",
                desc: "desc5",
                img_url: "app/images/login_background.jpg",
                duration: "duration5"
            }, {
                name: "course6",
                desc: "desc6",
                img_url: "app/images/login_background.jpg",
                duration: "duration6"
            }]

        }
    ])

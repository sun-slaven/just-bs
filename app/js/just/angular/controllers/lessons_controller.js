GlobalModules.add_controller('lessons')
angular.module('just.controllers.lessons', [])
    .controller('LessonsController', ['$rootScope', '$scope', 'LessonsService', function($rootScope, $scope, LessonsService) {
        
        LessonsService.lessons_list(function(resp){
            $scope.lessons = resp;
        })

        $scope.institutes = [{
            id: 1,
            name: "计算机",
        }, {
            id: 2,
            name: "土木工程"
        }]
        $scope.teachers = [{
            id: 1,
            name: "教师1"
        }, {
            id: 2,
            name: "教师2"
        }]

        // $scope.lessons = [{
        //     name: "course1",
        //     desc: "desc1",
        //     img_url: "app/images/login_background.jpg",
        //     duration: "duration1"
        // }, {
        //     name: "course2",
        //     desc: "desc2",
        //     img_url: "app/images/login_background.jpg",
        //     duration: "duration2"
        // }, {
        //     name: "course3",
        //     desc: "desc3",
        //     img_url: "app/images/login_background.jpg",
        //     duration: "duration3"
        // }, {
        //     name: "course4",
        //     desc: "desc4",
        //     img_url: "app/images/login_background.jpg",
        //     duration: "duration4"
        // }, {
        //     name: "course5",
        //     desc: "desc5",
        //     img_url: "app/images/login_background.jpg",
        //     duration: "duration5"
        // }, {
        //     name: "course6",
        //     desc: "desc6",
        //     img_url: "app/images/login_background.jpg",
        //     duration: "duration6"
        // }]

        $scope.$watch('chosen_college', function(newValue, oldValue, scope) {
            console.log(newValue);
        })
        $scope.$watch('chosen_major', function(newValue, oldValue, scope) {
            console.log(newValue);
        })
    }])

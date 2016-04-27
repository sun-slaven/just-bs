GlobalModules.add_controller('lessons')
angular.module('just.controllers.lessons', [])
    .controller('LessonsController', ['$rootScope', '$scope', 'LessonsService', 'CollegeMajorService', function($rootScope, $scope, LessonsService, CollegeMajorService) {

        $scope.colleges = $rootScope.all_colleges;
        $scope.majors = $rootScope.all_majors;
        $scope.chosen_college = null;
        $scope.chosen_major = null;
        $scope.$watch('chosen_college', function(newValue, oldValue) {
            if (newValue) {
                $scope.majors = newValue.major_list;
            } else {
                $scope.majors = $rootScope.all_majors;
            }
        })
        $scope.$watch('chosen_major', function(newValue, oldValue, scope) {
            if (newValue) {
                angular.forEach($rootScope.all_colleges, function(college) {
                    if (college.id == newValue.college_id) {
                        $scope.chosen_college = college;
                    };
                })
            } else {
                $scope.majors = $rootScope.all_majors;
            }
        })
    }])

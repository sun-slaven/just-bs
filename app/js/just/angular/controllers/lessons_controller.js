GlobalModules.add_controller('lessons')
angular.module('just.controllers.lessons', [])
    .controller('LessonsController', ['$rootScope', '$scope', 'LessonsService', 'CollegeMajorService', 'CommonUtil', function($rootScope, $scope, LessonsService, CollegeMajorService, CommonUtil) {
        $scope.lessons = $rootScope.all_lessons;
        $scope.colleges = $rootScope.all_colleges;
        $scope.majors = $rootScope.all_majors;
        $scope.chosen_college = null;
        $scope.chosen_major = null;


        $scope.change_college = function(college) {
            if (college) {
                $scope.majors = college.major_list;
                $scope.lessons = CommonUtil.getLessonsByCollege(college.id);
            } else {
                $scope.chosen_major = null;
                $scope.majors = $rootScope.all_majors;
                $scope.lessons = $rootScope.all_lessons;
            }
        }
        $scope.change_major = function(major) {
            if ($scope.chosen_college) {
                $scope.majors = $scope.chosen_college.major_list;
                return;
            };
            if (major) {
                angular.forEach($rootScope.all_colleges, function(college) {
                    if (college.id == major.college_id) {
                        $scope.chosen_college = college;
                        $scope.lessons = CommonUtil.getLessonsByMajor(major.id);
                    };
                })
            } else {
                $scope.majors = $rootScope.all_majors;
                $scope.lessons = $rootScope.all_lessons;
            }
        }


/*        $scope.$watch('chosen_college', function(newValue, oldValue) {
            if (newValue && $scope.chosen_major.college_id == newValue.id) {
                return;
            };
            if (newValue) {
                $scope.majors = newValue.major_list;
                $scope.lessons = CommonUtil.getLessonsByCollege(newValue.id);
            } else {
                $scope.chosen_major = null;
                $scope.majors = $rootScope.all_majors;
                $scope.lessons = $rootScope.all_lessons;
            }
        })
        $scope.$watch('chosen_major', function(newValue, oldValue, scope) {
            if (newValue) {
                angular.forEach($rootScope.all_colleges, function(college) {
                    if (college.id == newValue.college_id) {
                        $scope.chosen_college = college;
                        console.log(CommonUtil.getLessonsByMajor(newValue.id))
                        $scope.lessons = CommonUtil.getLessonsByMajor(newValue.id);
                    };
                })
            } else {
                $scope.majors = $rootScope.all_majors;
                $scope.lessons = $rootScope.all_lessons;
            }
        })
*/


    }])

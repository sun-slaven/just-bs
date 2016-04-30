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

    }])

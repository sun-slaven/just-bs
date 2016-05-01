GlobalModules.add_service('common_util')
angular.module('just.services.common_util', []).
factory('CommonUtil', ['$rootScope', 'LessonsService',
    function($rootScope, LessonsService) {

        var init_if_need = function() {
            if ($rootScope.all_lessons == undefined) {
                LessonsService.lessons_list(function(resp) {
                    $rootScope.all_lessons = resp;
                })
            };
        }

        var getLessonsByCollege = function(college_id) {
            var college_lessons = [];
            init_if_need()
            angular.forEach($rootScope.all_lessons, function(lesson) {
                if (lesson.college.id == college_id) {
                    college_lessons.push(lesson);
                };
            })
            return college_lessons;
        }

        var getLessonsByMajor = function(Major_id) {
            var major_lessons = [];
            init_if_need()
            angular.forEach($rootScope.all_lessons, function(lesson) {
                if (lesson.major.id == Major_id) {
                    major_lessons.push(lesson);
                };
            })
            return major_lessons;
        }

        var getMyCreatedLessons = function() {
            var created_lessons = [];
            init_if_need()
            angular.forEach($rootScope.all_lessons, function(lesson) {
                if (lesson.teacher.id == $rootScope.current_user.id) {
                    created_lessons.push(lesson);
                };
            })
            return created_lessons;
        }

        var adjustFileType = function(fileSuffix){
            image_array = ['.bmp','.png','.gif','.jpg','.jpeg']
            video_array = ['.vob','.avi','.rmvb','.asf','.wmv','.mp4']
            if (image_array.indexOf(fileSuffix) > -1) {
                return 'image'
            }
            if (video_array.indexOf(fileSuffix) > -1) {
                return 'video'
            };
            return 'file'
        }



        return {
            getLessonsByCollege: getLessonsByCollege,
            getLessonsByMajor: getLessonsByMajor,
            getMyCreatedLessons: getMyCreatedLessons,
            adjustFileType: adjustFileType
        }
    }
])

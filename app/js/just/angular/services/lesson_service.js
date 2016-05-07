GlobalModules.add_service('lesson')
angular.module('just.services.lesson', []).
factory('LessonService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {
        var lessonAPI = $resource('/api/v1/courses/:course_id', {course_id : '@course_id'}, {
            delete_lesson: {method: 'delete' , isArray: false},
            get_lesson: {method: 'get' , isArray: false},
        })


        function get_lesson(course_id,success) {
            lessonAPI.get_lesson({}, {course_id: course_id}, function(resp) {
                if (success) { success(resp) }
            })
        }



        function delete_lesson(course_id,success) {
            lessonAPI.delete_lesson({}, {
                course_id: course_id
            }, function(resp) {
                console.log(resp)
                if (success) { success(resp) }
            })
        }



        return {
            get_lesson: get_lesson,
            delete_lesson: delete_lesson
        }
    }
])

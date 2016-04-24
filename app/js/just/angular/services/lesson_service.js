GlobalModules.add_service('lesson')
angular.module('just.services.lesson', []).
factory('LessonService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var lessonAPI = $resource('/api/v1/courses/:course_id', {}, {
            delete_lesson: {method: 'delete' , isArray: false},
            get_lesson: {method: 'get' , isArray: false},
            delete_lesson: {method: 'delete' , isArray: false},
        })


        function get_lesson(success) {
            lessonAPI.get_lesson({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function delete_lesson(success) {
            lessonAPI.delete_lesson({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            get_lesson: get_lesson,
            delete_lesson: delete_lesson
        }
    }
])

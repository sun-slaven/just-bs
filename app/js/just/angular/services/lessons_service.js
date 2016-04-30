GlobalModules.add_service('lessons')
angular.module('just.services.lessons', []).
factory('LessonsService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var lessonsAPI = $resource('/api/v1/courses', {}, {
            lessons_list: { method: 'get', isArray: true, cache: false },
            create_lesson: { method: 'post', isArray: false },
        })

        function lessons_list(success) {
            lessonsAPI.lessons_list({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function create_lesson(success) {
            lessonsAPI.create_lesson({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            lessons_list: lessons_list,
            create_lesson: create_lesson
        }
    }
])

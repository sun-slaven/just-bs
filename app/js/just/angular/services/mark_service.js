GlobalModules.add_service('mark')
angular.module('just.services.mark', []).
factory('MarkService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var marksAPI = $resource('/api/v1/courses/:course_id/marks', {}, {
            add_mark: {method: 'post' , isArray: false},
        })

        var markAPI = $resource('/api/v1/courses/:course_id/marks/:mark_id', {}, {
            cancel_mark: {method: 'post' , isArray: false},
        })

        function add_mark(success) {
            lessonAPI.get_lesson({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function cancel_mark(success) {
            markAPI.cancel_mark({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            add_mark: add_mark,
            cancel_mark: cancel_mark
        }
    }
])

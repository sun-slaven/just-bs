GlobalModules.add_service('mark')
angular.module('just.services.mark', []).
factory('MarkService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var marksAPI = $resource('/api/v1/courses/:course_id/marks', { course_id: '@course_id' }, {
            add_mark: { method: 'post', isArray: false },
        })

        var markAPI = $resource('/api/v1/courses/:course_id/marks/', {}, {
            cancel_mark: { method: 'post', isArray: false },
        })

        function add_mark(course_id, success) {
            console.log(course_id)
            marksAPI.add_mark({}, {
                course_id: course_id
            }, function(resp) {
                if (success) { success(resp) }
            })
        }

        function cancel_mark(course_id, success) {
            markAPI.cancel_mark({}, {
                course_id: course_id
            }, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            add_mark: add_mark,
            cancel_mark: cancel_mark
        }
    }
])

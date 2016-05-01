GlobalModules.add_service('comments')
angular.module('just.services.comments', []).
factory('CommentsService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var commentsAPI = $resource('/api/v1/courses/:course_id/comments', {course_id: '@course_id'}, {
            delete_comments: {method: 'delete' , isArray: false},
            get_comments: {method: 'get' , isArray: true},
            add_comments: {method: 'post' , isArray: false},
        })


        function delete_comments(lesson_id,success) {
            commentsAPI.delete_comments({course_id:  lesson_id}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function get_comments(lesson_id,success) {
            commentsAPI.get_comments({}, {course_id:  lesson_id}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function add_comments(obj,success) {
            commentsAPI.add_comments({}, {
                course_id: obj.course_id,
                comment: obj.comment
            }, function(resp) {
                if (success) { success(resp) }
            })
        }        


        return {
            delete_comments: delete_comments,
            get_comments: get_comments,
            add_comments: add_comments
        }
    }
])

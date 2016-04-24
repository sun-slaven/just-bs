GlobalModules.add_service('comments')
angular.module('just.services.comments', []).
factory('CommentsService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var commentsAPI = $resource('/api/v1/courses/:course_id/comments', {}, {
            delete_comments: {method: 'delete' , isArray: false},
            get_comments: {method: 'get' , isArray: false},
            add_comments: {method: 'post' , isArray: false},
        })


        function delete_comments(success) {
            commentsAPI.delete_comments({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function get_comments(success) {
            commentsAPI.get_comments({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function add_comments(success) {
            commentsAPI.add_comments({}, {}, function(resp) {
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

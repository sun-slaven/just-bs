GlobalModules.add_service('user')
angular.module('just.services.user', []).
factory('UserService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var userAPI = $resource('/users/:action', {}, {
            sign_in: { method: 'post', params: { action: 'sign_in' } },
            signout: { method: 'delete', params: { action: 'sign_out' } },
        })


        function get(success) {

        }


        function sign_in(user, success) {
            userAPI.sign_in({}, {
                user: user
            }, function(resp) {
                set_user(resp)
                reset_csrf_token(resp.csrfToken)
                if (success) { success(resp) }
            })
        }

        function sign_out(success) {
            // userAPI.sign_out({}, {
            //     user: user
            // }, function(resp) {
            //     set_user(null)
            //     $rootScope.reload();//route reload
            //     reset_csrf_token(resp.csrfToken)
                if (success) { success(resp) }
            // })

        }

        function set_user(new_user) {
            $rootScope.clear_cache()
            $rootScope.account = new_user
            $rootScope.set_cache('user', new_user)
        }

        function reset_csrf_token(new_token) {
            $http.defaults.headers.common["X-CSRF-Token"] = new_token
        }

        return {
            get: get,
            sign_in: sign_in,
            sign_out: sign_out,
        }
    }
])

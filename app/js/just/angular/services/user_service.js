GlobalModules.add_service('user')
angular.module('just.services.user', []).
factory('UserService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {
        var userAPI = $resource('/api/v1/tokens', {}, {
            sign_in: { method: 'post' },
            signout: { method: 'delete' }
        })
        var registerAPI = $resource('/api/v1/users', {}, {
                register: { method: 'post' }
            })
            //992444037@qq.com  123456
        function sign_in(user, success) {
            userAPI.sign_in({}, {
                email: user.email,
                password: user.password
            }, function(resp) {
                set_user(resp.user);
                if (success) { success(resp) }
            })
        }

        function sign_out(success) {
            userAPI.sign_out({}, {
                user: user
            }, function(resp) {
                set_user(null)
                $rootScope.clear_cache()
                $rootScope.reload(); //route reload
                if (success) { success(resp) }
            })
        }

        function register(user, success) {
            registerAPI.register({}, {
                email: user.email,
                password: user.password
            }, function(resp) {
                set_user(resp.user);
                if (success) { success(resp) }
            })
        }

        function set_user(new_user) {
            $rootScope.current_user = new_user
        }

        return {
            sign_in: sign_in,
            sign_out: sign_out,
            register: register
        }
    }
])

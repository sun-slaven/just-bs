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

        var myLessonsAPI = $resource('/api/v1/users/courses', {}, {
                myLessons: { method: 'get' ,isArray: true}
        })
            //992444037@qq.com  123456   STUDENT
            //158274194@qq.com   123456  TEACHER
            //893196569@qq.com  123456   ADMIN
        function sign_in(user, success) {
            userAPI.sign_in({}, {
                email: user.email || $rootScope.get_storage('email'),
                password: user.password || $rootScope.get_storage('password')
            }, function(resp) {
                set_user(resp.user);
                if (success) { success(resp) }
            },function(error){
                console.log(error)
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
                user_name: user.name,
                email: user.email,
                password: user.password,
                password2: user.password_again
            }, function(resp) {
                console.log(resp)
                if (success) { success(resp) }
            })
        }

        function set_user(new_user) {
            $rootScope.current_user = new_user
        }

        function myLessons(user,callback){
            myLessonsAPI.myLessons({},function(resp){
                
                if (callback) {callback(resp)};
            })
        }

        return {
            sign_in: sign_in,
            sign_out: sign_out,
            register: register,
            myLessons: myLessons
        }
    }
])

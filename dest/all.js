var GlobalModules = (function() {
    var services = []
    var controllers = []
    var directives = []
    var actions = []

    function add_service(service) { services.push(service) }

    function add_controller(controller) { controllers.push(controller) }

    function add_directive(directive) { directives.push(directive) }

    function add_action(action) { actions.push(action) }

    function get(others) {
        var all = []
        services.forEach(function(service) { all.push("just.services." + service) })
        controllers.forEach(function(controller) { all.push("just.controllers." + controller) })
        directives.forEach(function(directive) { all.push("just.directives." + directive) })
        actions.forEach(function(action) { all.push("just.actions." + action) })
        return all.concat(others)
    }

    return {
        add_service: add_service,
        add_controller: add_controller,
        add_directive: add_directive,
        add_action: add_action,
        get: get
    }
})();


// local storage

var justConst = (function() {
    function get_meta(name) {
        return $('meta[name=' + name + ']').attr('content');
    }
})();

angular.module('just.constants', []).constant('JustConst', justConst);

GlobalModules.add_controller('me')
angular.module('just.controllers.me', ['ngCookies'])
    .controller('MeController', ['$rootScope', '$scope', '$cookies',
        function($rootScope, $scope, $cookies) {
        }
    ])

GlobalModules.add_controller('user')
angular.module('just.controllers.user', ['ngCookies'])
    .controller('UserController', ['$rootScope', '$scope', '$cookies', 'UserService',
        function($rootScope, $scope, $cookies, UserService) {
            if ($rootScope.user) {
                $rootScope.go('/')
            }
            $scope.form_type = 'login';

            $scope.change_active = function(attr) {
                $scope.form_type = attr;
            }

            $scope.user = {
                name: '',
                password: '',
                remember_me: true
            }
            $scope.can_submit = function() {
                if ($scope.user.name == '') {
                    return false };
                if ($scope.user.password == '') {
                    return false };
                return true
            }
            $scope.submit = function() {
                if ($scope.can_submit()) {
                    if ($scope.form_type == 'login') {
                            $rootScope.go("/users/1/list")
                    }
                }
            }
        }
    ])

angular.module('just.filters', [])
    .filter('cut', function() {
        return function(value, wordwise, max, tail) {
            if (!value) return '';

            max = parseInt(max, 10);
            if (!max) return value;
            if (value.length <= max) return value;

            value = value.substr(0, max);
            if (wordwise) {
                var lastspace = value.lastIndexOf(' ');
                if (lastspace != -1) {
                    value = value.substr(0, lastspace);
                }
            }

            return value + (tail || ' …');
        };
    })
    .filter('password', [function() {
        return function(str) {
            if (!str) return '';
            var result = ''
            for (i = 0; i < str.length; i++) {
                result += '*'
            }
            return result
        }
    }]);

angular.module('just.route_config', []).
provider('RouteConfig', function() {
    this.$get = function() {
        var all_configs = [];
        var partial_url = function(url) {
            return '/app/partials/' + url + '.html';
        }
        var base_config = [{
            path: '/login',
            templateUrl: partial_url('user/login'),
            controller: 'UserController'
        }, ];

        var me_config = [{
            path: '/users/:user_id/list',
            templateUrl: partial_url('me/index'),
            controller: 'MeController'
        }]

        add_config(base_config);
        add_config(me_config);

        function add_config(config) {
            all_configs = all_configs.concat(config);
        }

        function get() {
            return all_configs;
        }

        return {
            get: get
        }
    }
});

GlobalModules.add_service('user')
angular.module('just.services.user', []).
factory('UserService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        function get(success) {

        }


        function sign_in(account, success) {

        }

        function sign_out(success) {

        }

        function set_account(new_account) {
            account = new_account
            $rootScope.account = new_account
            $rootScope.set_shop(new_account ? new_account.shop : {})
        }


        return {
            get: get,
            sign_in: sign_in,
            sign_out: sign_out,
        }
    }
])

var version_timestamp = "?v" + Date.parse(new Date());
/**
 *  Module
 *
 * application.js
 */
angular.module('just', GlobalModules.get([
    'ngRoute', 'ngResource', 'ngCookies','ngAnimate',
    'just.route_config',
    'just.constants',
    'just.filters'
])).config(['$routeProvider', 'RouteConfigProvider',
    function($routeProvider, RouteConfigProvider) {
        var all_configs = RouteConfigProvider.$get().get()
        angular.forEach(all_configs, function(conf) {
            $routeProvider.when(conf.path, {
                templateUrl: conf.templateUrl + version_timestamp,
                controller: conf.controller
            })
        })
        $routeProvider.otherwise({
            redirectTo: '/login'
        });
    }
]).run(['$rootScope', '$location', function($rootScope, $location) {
    $rootScope.partial = function(partial_name) {
        return "app/partials/" + partial_name + ".html" + version_timestamp;
    }
    $rootScope.go = function(url) {
        $location.url(url)
    }
    $rootScope.reload = function(bool) {
        if (bool) { location.reload() } else { $route.reload() }
    }
    $rootScope.get_cache = function(key) {
        return $cacheFactory.get(key);
    }
    $rootScope.set_cache = function(key, value) {
        $cacheFactoryput(key, value);
    }
    $rootScope.clear_cache = function() {
        $cacheFactory.get('$http').removeAll();
        $cacheFactory.removeAll();
    }

    $rootScope.header_search = {
        input_show: false,
        search_info: '',
        open: function() {
            if (this.input_show == false && this.search_info == '') {
                this.input_show = true;
            } else {
                if (this.can_submit()) {
                    this.submit()
                } else {
                    this.close();
                }
            }
        },
        close: function() {
            this.input_show = false
            this.search_info = ''
        },
        can_submit: function() {
            if (this.search_info) {
                return true
            }
        },
        submit: function() {
            console.log("submit")
            this.close()
        }

    }
}])

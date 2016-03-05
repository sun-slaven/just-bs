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
;var justConst = (function() {
    function get_meta(name) {
        return $('meta[name=' + name + ']').attr('content');
    }
})();

angular.module('just.constants', []).constant('JustConst', justConst);
;GlobalModules.add_controller('user')
angular.module('just.controllers.user', ['ngCookies'])
    .controller('UserController', ['$rootScope', '$scope', '$cookies', 'UserService',
        function($rootScope, $scope, $cookies, UserService) {
            return
        }
    ])
;angular.module('just.filters', [])
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
;angular.module('just.route_config', []).
provider('RouteConfig', function() {
    this.$get = function() {
        var all_configs = [];

        var base_config = [{
            path: '/login',
            templateUrl: '/app/partials/user/login.html',
            controller: 'UserController'
        }, ];

        var branch_config = [{
            path: '/branches/:branch_id/eat_in_hall',
            templateUrl: '/webpos/partials/branches/eat_in_hall.html',
            controller: 'branchEatInHallController'
        }]

        add_config(base_config);
        //     add_config(branch_config);


        function get() {
            return all_configs;
        }

        function add_config(config) {
            all_configs = all_configs.concat(config);
        }
        return {
            get: get
        }
    }
});
;GlobalModules.add_service('user')
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
;var version_timestamp = "?v" + Date.parse(new Date());
/**
 *  Module
 *
 * application.js
 */
angular.module('just', GlobalModules.get([
    'ngRoute', 'ngResource', 'ngCookies',
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
]).run(['$rootScope', function($rootScope) {
    $rootScope.partial = function(partial_name) {
        return "/just/partials/" + partial_name + ".html" + version_timestamp;
    }
}])

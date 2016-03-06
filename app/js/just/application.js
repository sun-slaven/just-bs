var version_timestamp = "?v" + Date.parse(new Date());
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

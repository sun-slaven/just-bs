var version_timestamp = "?v" + Date.parse(new Date());
/**
 *  Module
 *
 * application.js
 */
angular.module('just', GlobalModules.get([
    'ngRoute', 'ngResource', 'ngCookies', 'ngAnimate', 'ui.bootstrap',
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
]).run(['$rootScope', '$location', 'AnchorSmoothScrollService', function($rootScope, $location, AnchorSmoothScrollService) {
    //路由以及$location
    $rootScope.partial = function(partial_name) {
        return "app/partials/" + partial_name + ".html" + version_timestamp;
    }
    $rootScope.go = function(url) {
        $location.url(url)
    }
    $rootScope.reload = function(bool) {
        if (bool) { location.reload() } else { $route.reload() }
    }
    $rootScope.location_path = function() {
        return $location.path();
    }

    //cache
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
        //nav-head 搜索
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
        //滚动到顶部
    $rootScope.scrollTo = function(eID) {
        AnchorSmoothScrollService.scrollTo(eID);
    }

}])

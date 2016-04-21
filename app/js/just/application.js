var version_timestamp = "?v" + Date.parse(new Date());
/**
 *  Module
 *
 * application.js
 */
angular.module('just', GlobalModules.get([
    'ngRoute', 'ngResource', 'ngCookies', 'ngAnimate', 'ui.bootstrap','smart-table','angularQFileUpload',
    'just.route_config',
    'just.constants',
    'just.filters'
])).config(['$routeProvider', '$sceDelegateProvider', 'RouteConfigProvider',
    function($routeProvider, $sceDelegateProvider, RouteConfigProvider) {
        //同源策略:在本站访问外站资源时,需要添加到信任名单中,不然就会加载错误.video
        $sceDelegateProvider.resourceUrlWhitelist([
            'self', 'http://7xt49i.com2.z0.glb.clouddn.com/**'
        ]);
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
        //nav-head controller
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
    $rootScope.go_me = function() {
        $rootScope.go('/users/1/show')
    }

    //滚动到顶部
    $rootScope.scrollTo = function(eID) {
        AnchorSmoothScrollService.scrollTo(eID);
    }

}])

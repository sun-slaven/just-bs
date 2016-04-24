var version_timestamp = "?v" + Date.parse(new Date());
/**
 *  Module
 *
 * application.js
 */
angular.module('just', GlobalModules.get([
    'ngRoute', 'ngResource', 'ngCookies', 'ngAnimate', 'ui.bootstrap', 'smart-table', 'angularQFileUpload', 'mgcrea.ngStrap',
    'just.route_config',
    'just.constants',
    'just.filters'
])).config(['$routeProvider', '$sceDelegateProvider', 'RouteConfigProvider', '$modalProvider',
    function($routeProvider, $sceDelegateProvider, RouteConfigProvider, $modalProvider) {
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
        //修改modal的全局配置
        angular.extend($modalProvider.defaults, {
            animation: 'am-fade-and-scale',
            html: true,
            templateUrl: '/app/partials/common_modal.html',
            show: true
        });
    }
]).run(['$rootScope', '$location', '$modal', '$cacheFactory', '$cookies', '$cookieStore', 'AnchorSmoothScrollService', function($rootScope, $location, $modal, $cacheFactory, $cookies, $cookieStore, AnchorSmoothScrollService) {
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
    var cache = $cacheFactory('just_cache')
    $rootScope.get_cache = function(key) {
        return cache.get(key);
    }
    $rootScope.set_cache = function(key, value) {
        cache.put(key, value);
    }
    $rootScope.clear_cache = function() {
            if (cache.get('$http')) {
                cache.get('$http').removeAll();
            };
            cache.removeAll();
        }
        //cookie
    $rootScope.get_cookie = function(key) {
        return $cookieStore.get(key);
    }
    $rootScope.set_cookie = function(key, value) {
        $cookieStore.put(key, value);
    }
    $rootScope.remove_cookie = function(key) {
        $cookieStore.remove(key);
    }
    $cookieStore.put("test","test")
    console.log($rootScope.get_cookie("test"))

    //滚动到顶部
    $rootScope.scrollTo = function(eID) {
        AnchorSmoothScrollService.scrollTo(eID);
    }

    //bootstrap  customer modals
    $rootScope.strap_modal = function(modal_obj) {
        return $modal(modal_obj)
    }
    $rootScope.confirm_modal = function(content, scope, success) {
        scope.modal_ok = success;
        $rootScope.strap_modal({
            content: content,
            title: "提示".concat(' <i class="fa fa-info-circle" aria-hidden="true"></i>'),
            scope: scope
        });
    }


    //if show header
    $rootScope.show_header = false;
    if ($rootScope.get_cache('current_user')) {
        $rootScope.show_header = true;
    };


}])

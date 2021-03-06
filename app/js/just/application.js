var version_timestamp = "?v" + Date.parse(new Date());
/**
 *  Module
 *
 * application.js
 */
angular.module('just', GlobalModules.get([
    'ngRoute', 'ngResource', 'ngCookies', 'ngAnimate', 'ui.bootstrap', 'smart-table', 'angularQFileUpload', 'mgcrea.ngStrap', 'angularLocalStorage', 'angular-loading-bar',
    'just.route_config',
    'just.constants',
    'just.filters'
])).config(['$httpProvider', '$routeProvider', '$locationProvider', '$sceDelegateProvider', 'RouteConfigProvider', '$modalProvider', 'cfpLoadingBarProvider',
    function($httpProvider, $routeProvider, $locationProvider, $sceDelegateProvider, RouteConfigProvider, $modalProvider, cfpLoadingBarProvider) {
        //同源策略:在本站访问外站资源时,需要添加到信任名单中,不然就会加载错误.video
        $sceDelegateProvider.resourceUrlWhitelist([
            'self',
            'http://7xnz7k.com1.z0.glb.clouddn.com/**'
        ]);
        //使用过滤器将所有请求都加上token和时间戳
        //function中可以注入$rootscope
        $httpProvider.interceptors.push(function($q, $cookies, $rootScope) {
            return {
                request: function(config) {
                    if (config.url.indexOf('/api/v1/') > -1) {
                        config.headers['Authorization'] = JSON.stringify($cookies.getObject('token'));
                    }
                    config.requestTimestamp = new Date().getTime();
                    return config;
                },
                requestError: function(err) {
                    return $q.reject(err);
                },
                response: function(res) {
                    return res;
                },
                responseError: function(err) {
                    if (400 === err.status) {
                        $rootScope.alert_modal("error", err.data.message)
                    } else if (401 === err.status) {
                        window.location = '/';
                    } else if (501 === err.status) {
                        $rootScope.alert_modal("error", err.data.message)
                    }
                    return $q.reject(err);
                }
            };
        });

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

        //angular-loading-bar
        cfpLoadingBarProvider.includeSpinner = true


        // $locationProvider.html5Mode(true); // remove # in the url
        // $locationProvider.hashPrefix = '!';
        //修改modal的全局配置
        angular.extend($modalProvider.defaults, {
            animation: 'am-fade-and-scale',
            html: true,
            templateUrl: '/app/partials/common_modal.html',
            show: true
        });
    }
]).run(['$rootScope', '$location', '$route', '$routeParams', '$modal', '$cacheFactory', 'AnchorSmoothScrollService', 'storage', 'CollegeMajorService', 'LessonsService', '$alert', 'UserService', '$cookies', function($rootScope, $location, $route, $routeParams, $modal, $cacheFactory, AnchorSmoothScrollService, storage, CollegeMajorService, LessonsService, $alert, UserService, $cookies) {
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
        //localStorage
    $rootScope.get_storage = function(key) {
        return storage.get(key);
    }
    $rootScope.set_storage = function(key, value) {
        storage.set(key, value);
    }
    $rootScope.clear_storage = function() {
        storage.clearAll();
    }

    //role
    $rootScope.is_student = function() {
        return $rootScope.current_user.role_name == 'STUDENT';
    }
    $rootScope.is_teacher = function() {
        return $rootScope.current_user.role_name == 'TEACHER';
    }
    $rootScope.is_admin = function() {
        return $rootScope.current_user.role_name == 'ADMIN';
    }



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
        //alert
    $rootScope.alert_modal = function(title, content) {
        return $alert({
            title: title.concat(' <i class="fa fa-info-circle" aria-hidden="true"></i>'),
            content: content,
            placement: 'top-right',
            type: 'info',
            show: true
        })
    }

    // 防止页面刷新,从cookie里取出当前对象.cookie在页面刷新时并不会清空
    if ($cookies.getObject('current_user')) {
        $rootScope.current_user = $cookies.getObject('current_user');
    }

    $rootScope.$on('$routeChangeSuccess', function(evt, next, current) {
        //refuse change the url to /# then header show
        if ($location.path() == '/' || $location.path() == '/login') {
            $rootScope.current_user = null;
        } else {
            //init college major info
            if ($rootScope.college_major == undefined) {
                $rootScope.all_colleges = []
                $rootScope.all_majors = []
                CollegeMajorService.get_college_major(function(response) {
                    for (var i = 0; i < response.length; i++) {
                        $rootScope.all_colleges.push(response[i])
                        for (index in response[i].major_list) {
                            response[i].major_list[index].college_id = response[i].id;
                            $rootScope.all_majors.push(response[i].major_list[index])
                        }
                    }
                });
            }

        }


    })


}])

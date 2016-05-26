GlobalModules.add_directive('just_video')
angular.module('just.directives.just_video', [])
    .directive('justVideo', ['$rootScope', '$window', '$cookies', function($rootScope, $window, $cookies) {
        // Runs during compile
        return {
            // name: '',
            // priority: 1,
            // terminal: true,
            scope: {
                video_url: '=videoUrl', //directive中的属性必须在此处''内是驼峰式写法
                video_process: '=videoProcess'
            },
            // controller: function($scope, $element, $attrs, $transclude) {},
            // require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
            restrict: 'E', // E = Element, A = Attribute, C = Class, M = Comment
            // template: '',
            templateUrl: '/app/partials/directives_template/just_video.html',
            replace: true,
            // transclude: true,
            // compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
            link: function($scope, element, iAttrs, controller) {
                $scope.$watch('video_url', function(newValue) {
                    if ($scope.video_url.indexOf('.swf') > -1) {
                        $scope.isSwf = true;
                    } else {
                        $scope.isSwf = false;
                    }
                })
                if ($scope.video_url != $rootScope.current_lesson.video_url) return;
                $('video').on('loadedmetadata', function() {
                    if ($scope.video_process) {
                        var process_seconds = this.duration * $scope.video_process
                        this.currentTime = process_seconds; //视频当前播放位置
                    };
                    $scope.duration = this.duration; //视频总长度
                })
                $('video').on('timeupdate', function() {
                    $scope.video_process = this.currentTime / $scope.duration;
                })

                $rootScope.$on('$routeChangeSuccess', function(evt, next, current) {
                        if (current.loadedTemplateUrl.indexOf('/lessons/show.html') > -1 && $scope.video_process) {
                            var client = new XMLHttpRequest();
                            client.open("POST", "/api/v1/courses/" + $rootScope.current_lesson.id + "/records", true); //同步ajax请求
                            client.setRequestHeader("Content-type", "application/json");
                            client.setRequestHeader("Authorization", JSON.stringify($cookies.getObject('token')))
                            client.send(JSON.stringify({ process: $scope.video_process }));
                        };
                    })
                    //页面关闭事件
                    //TODO  has some problems
                    //window.onunload = remember_progress;  

                var remember_progress = function() {
                    if ("sendBeacon" in navigator) {
                        //Beacon API
                        navigator.sendBeacon("/api/v1/courses/" + $rootScope.current_lesson.id + "/records", { process: $scope.video_process });
                    } else {
                        var client = new XMLHttpRequest();
                        client.open("POST", "/api/v1/courses/" + $rootScope.current_lesson.id + "/records", false); //同步ajax请求
                        client.setRequestHeader("Content-type", "application/json");
                        client.setRequestHeader("Authorization", JSON.stringify($cookies.getObject('token')))
                        client.send(JSON.stringify({ process: $scope.video_process }));
                    }
                }
            }
        };
    }])
    .directive('justSwf', function() {
        return {
            restrict: 'E',
            link: function(scope, element, attrs) {
                var url = scope.$eval(attrs.src);
                element.replaceWith('<object type="application/x-shockwave-flash" data="' + url + '"></object>');
            }
        };
    });

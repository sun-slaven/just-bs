GlobalModules.add_controller('lesson')
angular.module('just.controllers.lesson', [])
    .controller('LessonController', ['$rootScope', '$scope', 'CommentsService',
        function($rootScope, $scope, CommentsService) {

            $scope.active_type = 'show_outline'
            $scope.change_active = function(type) {
                $scope.active_type = type;
            }
            CommentsService.get_comments($rootScope.current_lesson.id, function(resp) {
                $scope.comments = resp;
            })
            $scope.submit_comment = function() {
                CommentsService.add_comments({
                    course_id: $rootScope.current_lesson.id,
                    comment: $scope.my_comment
                }, function(resp) {
                    $rootScope.alert_modal({
                        title: "提示",
                        content: "评论成功",
                        placement: 'top-right',
                        type: 'info',
                        show: true
                    })
                    CommentsService.get_comments($rootScope.current_lesson.id, function(resp) {
                        $scope.comments = resp;
                    })
                })
            }
            $scope.video_url = 'http://7xt49i.com2.z0.glb.clouddn.com/%E5%8D%8A%E7%B3%96%E4%B8%BB%E4%B9%89.mp4'
            $scope.pdf_url = 'http://7xt49i.com2.z0.glb.clouddn.com/AngularJS%E6%9D%83%E5%A8%81%E6%95%99%E7%A8%8B%28www.Linuxidc.com%E6%95%B4%E7%90%86%29.pdf'
            $scope.zip_url = 'http://7xt49i.com2.z0.glb.clouddn.com/pack.zip'
            $scope.need_learn = true;
            if ($scope.need_learn) {
                $scope.btn_content = "开始学习";
            } else {
                $scope.btn_content = "继续学习";
                $scope.progress_info_percent = 10;
                $scope.progress_info_hour = 1;
                $scope.progress_info_minute = 10;
            }

            $scope.start_or_continue = function() {
                if ($scope.need_learn) {
                    $scope.need_learn = false;
                    $scope.btn_content = "继续学习";
                    $scope.progress_info_percent = 10;
                    $scope.progress_info_hour = 1;
                    $scope.progress_info_minute = 10;
                } else {
                    $scope.show_resource = true;
                }
            }

        }
    ])

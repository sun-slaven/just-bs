GlobalModules.add_controller('lesson')
angular.module('just.controllers.lesson', [])
    .controller('LessonController', ['$rootScope', '$scope', '$routeParams', 'LessonService', 'CommentsService', 'MarkService', 'ChaptersService',
        function($rootScope, $scope, $routeParams, LessonService, CommentsService, MarkService, ChaptersService) {
            $scope.active_type = 'show_comment';
            $scope.change_active = function(type) {
                $scope.active_type = type;
            }
            if ($routeParams.lesson_id) {
                LessonService.get_lesson($routeParams.lesson_id, function(resp) {
                    learn_status_callback();//show btn status
                    $scope.video_url = resp.video_url;
                })
                CommentsService.get_comments($routeParams.lesson_id, function(resp) {
                    $scope.comments = resp;
                })
                ChaptersService.get_chapters($routeParams.lesson_id, function(resp) {
                    $scope.lesson_outline_list = resp
                })
            };


            $scope.submit_comment = function() {
                CommentsService.add_comments({
                    course_id: $rootScope.current_lesson.id,
                    content: $scope.my_comment
                }, function(resp) {
                    $rootScope.alert_modal("提示", "评论成功")
                    CommentsService.get_comments($rootScope.current_lesson.id, function(resp) {
                        $scope.comments = resp;
                        $rootScope.current_lesson.comment_sum = resp.length
                    })
                })
            }

            $scope.delete_comment = function(comment) {
                CommentsService.delete_comments($rootScope.current_lesson.id, comment.id, function(resp) {
                    CommentsService.get_comments($routeParams.lesson_id, function(resp) {
                        $scope.comments = resp;
                    })
                    $rootScope.alert_modal("提示", "评论删除成功")
                })
            }

            var learn_status_callback = function() {
                if ($rootScope.current_lesson.mark_status == 'N') {
                    $scope.need_learn = true;
                } else {
                    $scope.need_learn = false;
                }
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
                        mark_and_learn();
                    } else {
                        continue_learn()
                    }
                }
                var mark_and_learn = function() {
                    $scope.need_learn = false;
                    $scope.btn_content = "继续学习";
                    MarkService.add_mark($rootScope.current_lesson.id, function(resp) {})
                }
                var continue_learn = function() {
                    $scope.show_resource = true;
                }

                $scope.show_chapter_video = function(chapter){
                    $scope.video_url = chapter.video_url;
                    $scope.show_resource = true;
                }


            }
        }
    ])

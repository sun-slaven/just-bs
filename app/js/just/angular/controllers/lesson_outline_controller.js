GlobalModules.add_controller('lesson_outline')
angular.module('just.controllers.lesson_outline', [])
    .controller('LessonOutlineController', ['$rootScope', '$scope', 'ChaptersService',
        function($rootScope, $scope, ChaptersService) {
            if (!$scope.outline_edit_lesson) {
                return
            }
            //just required by table
            $scope.lesson_outline = []
            //每次请求都返回一个对象.所以需要重新去拉取所有章节.此处可优化
            var get_chapters = function() {
                ChaptersService.get_chapters($scope.outline_edit_lesson.id, function(resp) {
                    $scope.lesson_outline_list = resp
                })
            }
            get_chapters();
            $scope.new_chapter = {
                order: '',
                name: '',
                content: '',
            };
            $scope.chapter_plus = function() {
                $scope.edit_chapter = angular.copy($scope.new_chapter);
                $scope.modal_title = "添加章节";
                $scope.modal_content_url = '/manage_lesson/_update_lesson_chapter_modal';
                $scope.modal_ok = function() {
                    ChaptersService.add_chapter($scope.outline_edit_lesson.id, $scope.edit_chapter, function(resp) {
                        get_chapters();
                        $scope.edit_chapter = {}
                        $scope.new_chapter = {}
                    })
                }
                $scope.modal_cancel = function() {
                    $scope.edit_chapter = {}
                    $scope.new_chapter = {}
                }
                $rootScope.strap_modal({
                    scope: $scope
                })
            };
            $scope.chapter_edit = function(chapter) {
                //deep copy ,must click ok to effect
                $scope.edit_chapter = angular.copy(chapter);
                $scope.modal_title = "修改章节";
                $scope.modal_content_url = '/manage_lesson/_update_lesson_chapter_modal';
                $scope.modal_ok = function() {
                    ChaptersService.update_chapter($scope.outline_edit_lesson.id, $scope.edit_chapter, function(resp) {
                        get_chapters();
                    })
                }
                $scope.modal_cancel = function() {
                    $scope.edit_chapter = {}
                }
                $rootScope.strap_modal({
                    scope: $scope
                })
            };
            $scope.chapter_delete = function(chapter) {
                ChaptersService.delete_chapter($scope.outline_edit_lesson.id, chapter.id, function(resp) {
                    $scope.edit_chapter = {}
                    get_chapters();
                    $rootScope.alert_modal('', '已成功删除章节')
                })
            };
        }
    ])

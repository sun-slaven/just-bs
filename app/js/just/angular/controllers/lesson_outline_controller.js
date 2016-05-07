GlobalModules.add_controller('lesson_outline')
angular.module('just.controllers.lesson_outline', [])
    .controller('LessonOutlineController', ['$rootScope', '$scope', 'ChaptersService',
        function($rootScope, $scope, ChaptersService) {
            //$scope.outline_edit_lesson;
            // $scope.lesson_outline_list = ChaptersService.get_chapters($scope.outline_edit_lesson.id, function(resp) {
            //     console.log(resp)
            // })
            $scope.lesson_outline_list = [{
                id: 1,
                name: 'name',
                content: 'content',
                order: 1,
                "create_time": "2016-01-12 05:20:11"
            }, {
                "id": "string",
                "name": "概述",
                "content": "你应该这样这样...",
                "order": 0,
                "create_time": "2016-01-12 05:20:11"
            }];
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
                        console.log(resp)
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
                        console.log(resp)
                        $scope.edit_chapter = {}
                    })
                }
                $scope.modal_cancel = function() {
                    $scope.edit_chapter = {}
                }
                $rootScope.strap_modal({
                    scope: $scope
                })
            };
            $scope.chapter_delete = function() {
                console.log('delete')
            };
        }
    ])

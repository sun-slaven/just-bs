GlobalModules.add_controller('lesson_outline')
angular.module('just.controllers.lesson_outline', [])
    .controller('LessonOutlineController', ['$rootScope', '$scope', 'ChaptersService', 'QiniuUpload', 'FileService', 'UuidService',
        function($rootScope, $scope, ChaptersService, QiniuUpload, FileService, UuidService) {
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
                video_name: '',
                video_url: ''
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
                        $scope.upload.clearAll();
                    })
                }
                $scope.modal_cancel = function() {
                    $scope.edit_chapter = {}
                    $scope.new_chapter = {}
                    $scope.upload.clearAll();
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
                        $scope.upload.clearAll();
                        get_chapters();
                    })
                }
                $scope.modal_cancel = function() {
                    $scope.edit_chapter = {}
                    $scope.upload.clearAll();
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

            //upload chapter video
            $scope.upload = {
                selectFileArray: null,
                selectFile: null,
                get_token_promise: {},
                onFileSelect: function($files) {
                    this.selectFileArray = $files;
                    this.selectFile = $files[0];
                    var suffix_info_obj = QiniuUpload.get_suffix_info_obj(this.selectFile);
                    this.get_token_promise = QiniuUpload.get_token(suffix_info_obj).then(function(resp) {
                        return {
                            key: UuidService.newuuid(suffix_info_obj.suffix),
                            token: resp.token
                        };
                    })
                },
                do_upload: function() {
                    $scope.upload.selectFile.progress = {
                        p: 0
                    };
                    var upload_fun = function(token_obj) {
                            QiniuUpload.upload($scope.upload.selectFile, token_obj).then(function(resp) {
                                $scope.edit_chapter.video_url = resp.key;
                                $scope.edit_chapter.video_name = $scope.upload.selectFile.name;
                                $rootScope.alert_modal("success", "视频上传成功");
                            }, function(error) {
                                console.log(error)
                            }, function(evt) {
                                if ($scope.upload.selectFile) {
                                    $scope.upload.selectFile.progress.p = Math.floor(100 * evt.loaded / evt.totalSize);
                                };
                            })
                        }
                        //promise 使用上一个promise的返回结果
                    this.get_token_promise.then(function(token_obj) {
                        upload_fun(token_obj)
                    })
                },
                abort: function() {
                    QiniuUpload.abort(this.selectFile, this.selectFileArray)
                    $scope.upload.selectFile = null;
                },
                clearAll: function(){
                    $scope.upload.selectFileArray = $scope.upload.selectFile = null;
                    $scope.upload.get_token_promise = {};
                }
            }



        }
    ])

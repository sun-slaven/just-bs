GlobalModules.add_controller('me')
angular.module('just.controllers.me', [])
    .controller('MeController', ['$rootScope', '$scope', 'UserService', 'QiniuUpload', 'CommonUtil', 'FileService',
        function($rootScope, $scope, UserService, QiniuUpload, CommonUtil, FileService) {
            $scope.active_type = 'chosen_lessons'
            $scope.change_active = function(type) {
                $scope.active_type = type;
            }

            UserService.myLessons($rootScope.current_user, function(resp) {
                $scope.chosen_lessons = resp
            })

            //upload avatar
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
                            key: resp.key,
                            token: resp.token
                        };
                    })
                },
                do_upload: function() {
                    $scope.upload.selectFile.progress = {
                        p: 0
                    };
                    var upload_fun = function(token_obj) {
                            console.log(token_obj)
                            QiniuUpload.upload($scope.upload.selectFile, token_obj).then(function(resp) {
                                console.log(resp)
                                    //QiniuUpload.save_file_to_db()
                                    //TODO
                                $rootScope.alert_modal("成功", "上传文件成功");
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
                }
            }
        }
    ])

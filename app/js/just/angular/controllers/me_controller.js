GlobalModules.add_controller('me')
angular.module('just.controllers.me', [])
    .controller('MeController', ['$rootScope', '$scope', 'UserService', 'QiniuUpload', 'FileService', 'UuidService',
        function($rootScope, $scope, UserService, QiniuUpload, FileService, UuidService) {
            $scope.active_type = 'chosen_lessons'
            $scope.change_active = function(type) {
                $scope.active_type = type;
            }

            UserService.myMarkedLessons($rootScope.current_user, function(resp) {
                $scope.chosen_lessons = resp
            })

            var updateUser = function(updateUserObj) {
                UserService.updateUser($rootScope.current_user, updateUserObj, function(resp) {})
            }


            $scope.updatePassword = {
                new_password: null,
                new_password_again: null,
                update: function() {
                    if (this.new_password == this.new_password_again) {
                        updateUser({
                            password: this.new_password
                        })
                        $rootScope.alert_modal("", "密码修改成功,请重新登陆!");
                        $rootScope.go('/login')
                    }
                }
            }

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
                                updateUser({
                                    icon_url: resp.key
                                })
                                $rootScope.alert_modal("", "头像修改成功");
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

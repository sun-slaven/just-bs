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
            $scope.onFileSelect = function($files) {
                selectFile = $files.pop();
                var suffix_info_obj = QiniuUpload.get_suffix_info_obj(selectFile);
                var token_obj = QiniuUpload.get_token(suffix_info_obj)
                selectFile.progress = {
                    p: 0
                };
                QiniuUpload.upload(selectFile, token_obj).then(function(resp) {
                    console.log(resp)
                }, function(error) {
                    console.log(error)
                }, function(evt) {
                    selectFile.progress.p = Math.floor(100 * evt.loaded / evt.totalSize);
                })
                //QiniuUpload.save_file_to_db()
                //TODO
            }



        }
    ])

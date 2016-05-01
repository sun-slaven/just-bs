GlobalModules.add_service('qiniu_upload')
angular.module('just.services.qiniu_upload', []).
factory('QiniuUpload', ['$rootScope', '$resource', '$http', '$qupload', 'FileService', 'CommonUtil'

    function($rootScope, $resource, $http, $qupload, FileService, CommonUtil) {

        var get_suffix_info_obj = function(selectFile) {
            var suffix = selectFile.name.substr(selectFile.name.indexOf('.')).toLowerCase(); //文件后缀
            var type = CommonUtil.adjustFileType(suffix); //image/video/file
            var fileObj = {
                suffix: suffix,
                type: type
            }
            return fileObj;
        }

        var get_token = function(fileObj) {
            FileService.get_file_token(fileObj).$promise.then(function(resp) {
                var token_obj = {
                    key: resp.key,
                    token: resp.token
                };
                return token_obj;
            })
        }

        var upload = function(file, token_obj) {
            file.upload = $qupload.upload({
                file: file,
                key: token_obj.key,
                token: token_obj.token
            });
            return file.upload //返回一个promise
        };

        var abort = function(file, index_in_files) {
            fileupload.abort();
            $scope.selectFiles.splice(index_in_files, 1);
        };

        var save_file_to_db = function(){
        }

        return {
            get_suffix_info_obj: get_suffix_info_obj,
            get_token: get_token,
            upload: upload,
            abort: abort,
            save_file_to_db: save_file_to_db
        }
    }
])

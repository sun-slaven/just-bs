GlobalModules.add_service('qiniu_upload')
angular.module('just.services.qiniu_upload', []).
factory('QiniuUpload', ['$rootScope', '$resource', '$http', '$qupload', 'FileService', 'CommonUtil',

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
            return FileService.get_file_token(fileObj).$promise
        }

        var upload = function(file, token_obj) {
            console.log(token_obj)
            file.upload = $qupload.upload({
                file: file,
                key: token_obj.key,
                token: token_obj.token
            });
            return file.upload //返回一个promise
        };

        var abort = function(file, files, indexAttr) {
            if (indexAttr > -1) {
                files.splice(index, 1)
            } else {
                for (index in files) {
                    if (files[index].name == file.name) {
                        files.splice(index, 1);
                        break;
                    };
                }
            }
            if (file.upload) {
                file.upload.abort();
            };
            
        };


        return {
            get_suffix_info_obj: get_suffix_info_obj,
            get_token: get_token,
            upload: upload,
            abort: abort,
        }
    }
])

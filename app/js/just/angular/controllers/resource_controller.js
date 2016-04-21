GlobalModules.add_controller('resource')
angular.module('just.controllers.resource', ['angularQFileUpload', 'LocalStorageModule'])
    .controller('resourceCtrl', ['$rootScope', '$scope', '$qupload', function($rootScope, $scope, $qupload) {
        $scope.upload_files = [{
            name: 'AngularJS权威教程(www.Linuxidc.com整理).pdf',
            upload_time: new Date('2016-04-19 20:56:06')
        }, {
            name: 'pack.zip',
            upload_time: new Date('2016-04-19 20:55:11')
        }, {
            name: ' 半糖主义.mp4',
            upload_time: new Date('2016-04-18 19:47:33')
        }]
        $scope.itemsByPage = 2;

        // $http.get(你七牛的获取uptoken地址).success(function (data) {
        //     $scope.uptoken = data.uptoken;   //获取你的七牛uptoken
        //     $scope.prefix = data.prefix;    //获取你的七牛文件存储地址
        // });
        $scope.selectFiles = [];

        var start = function(index) {
            $scope.selectFiles[index].progress = {
                p: 0
            };
            $scope.selectFiles[index].upload = $qupload.upload({
                key: '<your qiniu file key>',
                file: $scope.selectFiles[index].file,
                token: '<your qiniu UpToken>'
            });
            $scope.selectFiles[index].upload.then(function(response) {
                $log.info(response);
            }, function(response) {
                $log.info(response);
            }, function(evt) {
                $scope.selectFiles[index].progress.p = Math.floor(100 * evt.loaded / evt.totalSize);
            });
        };

        $scope.abort = function(index) {
            $scope.selectFiles[index].upload.abort();
            $scope.selectFiles.splice(index, 1);
        };

        $scope.onFileSelect = function($files) {
            var offsetx = $scope.selectFiles.length;
            for (var i = 0; i < $files.length; i++) {
                $scope.selectFiles[i + offsetx] = {
                    file: $files[i]
                };
                start(i + offsetx);
            }
        };

    }])

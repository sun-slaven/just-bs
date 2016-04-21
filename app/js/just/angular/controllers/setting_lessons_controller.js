GlobalModules.add_controller('setting_lessons')
angular.module('just.controllers.setting_lessons', [])
    .controller('settingLessonsCtrl', ['$rootScope', '$scope', '$qupload', '$log',function($rootScope, $scope, $qupload,$log) {
        $scope.useful_lessons = [{
            name: 'AngularJS',
            created_time: new Date('2016-04-19 20:56:06'),
            updated_time: new Date('2016-04-20 20:56:06')
        }, {
            name: 'Go',
            created_time: new Date('2016-04-19 20:55:11'),
            updated_time: new Date('2016-04-20 20:55:11')
        }, {
            name: 'Bootstrap',
            created_time: new Date('2016-04-19 19:47:33'),
            updated_time: new Date('2016-04-20 19:47:33')
        }]
        $scope.itemsByPage = 2;

        $scope.show_edit_lesson_modal = false;

        $scope.edit = function(lesson){
            $scope.show_edit_lesson_modal = true;
        }

        $scope.delete = function(lesson){
            console.log("delete")
        }

        //upload to qiniu
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

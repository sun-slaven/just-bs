GlobalModules.add_controller('admin')
angular.module('just.controllers.admin', [])
    .controller('AdminController', ['$rootScope', '$scope', 'LessonService',
        function($rootScope, $scope, LessonService) {
            $scope.active_type = 'manager_user';
            $scope.change_active = function(attr) {
                $scope.active_type = attr;
            }
            $scope.itemsByPage = 10;
            //manager user
            $scope.all_users = [{
                "id": "aa5eba0a-703c-4801-955b-1f44997738fe",
                "name": "小泡子仔",
                "email": "992444037@qq.com",
                "role_name": "STUDENT",
                created_time: new Date()
            }, {
                "id": "aa5eba0a-703c-4801-955b-1f44997738fe",
                "name": "slaven",
                "email": "893196569@qq.com",
                "role_name": "STUDENT",
                created_time: new Date()
            }]


            $scope.delete_user = function(user) {

            }



            //manager lessons
            $scope.all_lessons = [{
                "name": "数据库",
                created_time: new Date(),
                updated_time: new Date(),
                "major": {
                    "id": "0260bb7c-2e93-4a7d-895d-59fac58fdbc6",
                    "name": "物联网工程"
                },
                "college": {
                    "id": "b6a0808f-b87a-44ca-b850-9545a3f3f089",
                    "name": "计算机学院",
                    "major_list": {
                        "id": "0260bb7c-2e93-4a7d-895d-59fac58fdbc6",
                        "name": "物联网工程"
                    }
                },
                "teacher": {
                    "id": "aa5eba0a-703c-4801-955b-1f44997738fe",
                    "name": "小泡子仔",
                    "email": "992444037@qq.com",
                    "role_name": "STUDENT",
                }
            }]

            $scope.delete_lesson = function(lesson) {
                LessonService.delete_lesson(lesson.id,function(resp){
                    $rootScope.alert_modal("", "课程:" + lesson.name + " 删除成功")
                })
            }


        }
    ])

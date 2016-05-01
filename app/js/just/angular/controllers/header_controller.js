GlobalModules.add_controller('header')
angular.module('just.controllers.header', [])
    .controller('HeaderController', ['$rootScope', '$scope', 'UserService',
        function($rootScope, $scope, UserService) {

            //log out
            $scope.sign_out = function() {
                    $rootScope.confirm_modal("确认退出吗?", $scope, function() {
                        $rootScope.current_user = null;//can delete , in service
                        $rootScope.go('/login');
                        UserService.sign_out(function() {
                            $rootScope.go('/login');
                        })
                     })
                }
                //nav-head controller
            $scope.header_search = {
                input_show: false,
                search_info: '',
                open: function() {
                    if (this.input_show == false && this.search_info == '') {
                        this.input_show = true;
                    } else {
                        if (this.can_submit()) {
                            this.submit()
                        } else {
                            this.close();
                        }
                    }
                },
                close: function() {
                    this.input_show = false
                    this.search_info = ''
                },
                can_submit: function() {
                    if (this.search_info) {
                        return true
                    }
                },
                submit: function() {
                    console.log("submit")
                    this.close()
                }

            }
            $scope.go_me = function() {
                $rootScope.go('/users/'+ $rootScope.current_user.id + '/me')
            }

            $scope.go_manager_lessons = function() {
                $rootScope.go('/users/1/manage_lesson')
            }
        }
    ])

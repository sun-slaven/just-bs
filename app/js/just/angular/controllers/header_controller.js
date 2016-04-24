GlobalModules.add_controller('header')
angular.module('just.controllers.header', [])
    .controller('HeaderController', ['$rootScope', '$scope', 'UserService',
        function($rootScope, $scope, UserService) {

            $rootScope.show_header = true;
            //log out
            $scope.log_out = function() {
                    $rootScope.confirm_modal("确认退出吗?", $scope, function() {
                        UserService.sign_out(function() {
                            $rootScope.show_header = false;
                            $rootScope.go('/');
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
                $rootScope.go('/users/1/show')
            }

            $scope.go_manager_lessons = function() {
                $rootScope.go('/users/1/manage_lesson')
            }
        }
    ])

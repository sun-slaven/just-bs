GlobalModules.add_directive('just_unique_validation')
angular.module('just.directives.just_unique_validation', [])
    .directive('justUniqueValidation', ['$rootScope', function($rootScope) {
        //使用必须为: just_chapter_exist_validation
        return {
            require: 'ngModel', //要求节点上必须使用到ng-modal
            restrict: 'A',
            scope: {
                target_array_or_object: '=targetArrayOrObject'
            },
            link: function(scope, iElm, iAttrs, controller) {
                iElm.on('input', function(event) {
                    scope.$apply(function() {
                        var keepGoing = true;
                        angular.forEach(scope.target_array_or_object, function(item) {
                            if (keepGoing) {
                                if (item.chapter) {
                                    if (item.chapter == iElm.val()) {
                                        controller.$setValidity('chapterExistValidation', false)
                                        keepGoing = false;
                                    } else {
                                        controller.$setValidity('chapterExistValidation', true)
                                    }
                                }
                                //others
                            };

                        })
                    })
                })
            }
        };
    }]);

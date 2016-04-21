GlobalModules.add_directive('just_modal')
angular.module('just.directives.just_modal', [])
    .directive('justModal', ['$rootScope', function($rootScope) {
        // Runs during compile
        return {
            // name: '',
            // priority: 1,
            // terminal: true,
            scope: {
                can_show: '=canShow'
            },
            // controller: function($scope, $element, $attrs, $transclude) {},
            // require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
            restrict: 'E', // E = Element, A = Attribute, C = Class, M = Comment
            // template: '',
            templateUrl: '/app/partials/directives_template/just_modal.html',
            replace: true,
            // transclude: true,
            // compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
            link: function($scope, iElm, iAttrs, controller) {
                console.log($scope.can_show)
                $scope.$watch($scope.can_show, function(newValue,oldValue) {
                    if (newValue) {
                        $('#myModal').modal();
                    };
                })
            }
        };
    }]);

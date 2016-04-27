GlobalModules.add_service('college_major')
angular.module('just.services.college_major', []).
factory('CollegeMajorService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var college_majorAPI = $resource('/api/v1/colleges', {}, {
            get_college_major: {method: 'get' , isArray: true}
        })


        function get_college_major(success) {
            college_majorAPI.get_college_major({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            get_college_major: get_college_major
        }
    }
])

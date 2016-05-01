GlobalModules.add_service('file')
angular.module('just.services.file', []).
factory('FileService', ['$rootScope', '$resource',
    function($rootScope, $resource) {
        var fileTokenAPI = $resource('/api/v1/files/tokens', {}, {
            file_token: { method: 'post' },
        })

        var fileAPI = $resource('/api/v1/files', {}, {
            save_file: { method: 'post' },
        })

        function get_file_token(fileObj) {
            return fileTokenAPI.file_token({},fileObj)
        }

        function save_file(success) {
            fileAPI.save_file({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            get_file_token: get_file_token,
            save_file: save_file
        }
    }
])

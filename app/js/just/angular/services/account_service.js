GlobalModules.add_service('account')
angular.module('just.services.account', []).
factory('AccountService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        function get(success) {

        }


        function sign_in(account, success) {

        }

        function sign_out(success) {

        }

        function set_account(new_account) {
            account = new_account
            $rootScope.account = new_account
            $rootScope.set_shop(new_account ? new_account.shop : {})
        }


        return {
            get: get,
            sign_in: sign_in,
            sign_out: sign_out,
        }
    }
])

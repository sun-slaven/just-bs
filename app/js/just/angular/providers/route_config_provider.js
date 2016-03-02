angular.module('just.route_config', []).
provider('RouteConfig', function() {
    this.$get = function() {
        var all_configs = [];

        var base_config = [{
            path: '/login',
            templateUrl: '/app/partials/login/login.html',
            controller: 'AccountController'
        }, ];

        var branch_config = [{
            path: '/branches/:branch_id/eat_in_hall',
            templateUrl: '/webpos/partials/branches/eat_in_hall.html',
            controller: 'branchEatInHallController'
        }]

        add_config(base_config);
        //     add_config(branch_config);


        function get() {
            return all_configs;
        }

        function add_config(config) {
            all_configs = all_configs.concat(config);
        }
        return {
            get: get
        }
    }
});

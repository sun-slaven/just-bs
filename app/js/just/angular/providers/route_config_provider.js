angular.module('just.route_config', []).
provider('RouteConfig', function() {
    this.$get = function() {
        var all_configs = [];
        var partial_url = function(url) {
            return '/app/partials/' + url + '.html';
        }
        var base_config = [{
            path: '/login',
            templateUrl: partial_url('user/login'),
            controller: 'UserController'
        }, ];

        var me_config = [{
            path: '/users/:user_id/list',
            templateUrl: partial_url('me/index'),
            controller: 'MeController'
        }]

        add_config(base_config);
        add_config(me_config);

        function add_config(config) {
            all_configs = all_configs.concat(config);
        }

        function get() {
            return all_configs;
        }

        return {
            get: get
        }
    }
});

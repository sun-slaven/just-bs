angular.module('just.route_config', []).
provider('RouteConfig', function() {
    this.$get = function() {
        var all_configs = [];
        var partial_url = function(url) {
            return '/app/partials/' + url + '.html';
        }
        var base_config = [{
            path: '/',
            templateUrl: partial_url('user/login'),
            controller: 'UserController'
        },{
            path: '/login',
            templateUrl: partial_url('user/login'),
            controller: 'UserController'
        }];

        var me_config = [{
            path: '/users/:user_id/me',
            templateUrl: partial_url('me/show'),
            controller: 'MeController'
        }, {
            path: '/users/:user_id/lessons/:lesson_id/show',
            templateUrl: partial_url('lessons/show'),
            controller: 'LessonController'
        }]

        var lessons_config = [{
            path: '/lessons/index',
            templateUrl: partial_url('lessons/index'),
            controller: 'LessonsController'
        }, {
            path: '/lessons/:lesson_id/show',
            templateUrl: partial_url('lessons/show'),
            controller: 'LessonController'
        }]

        var manager_lesson_config = [{
            path: '/users/:user_id/manage_lesson',
            templateUrl: partial_url('manage_lesson/show'),
            controller: 'ManageLessonController'
        }]

        var admin_config = [{
            path: '/admin/show',
            templateUrl: partial_url('admin/show'),
            controller: 'AdminController'
        }]

        add_config(base_config);
        add_config(me_config);
        add_config(lessons_config);
        add_config(manager_lesson_config)
        add_config(admin_config)

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

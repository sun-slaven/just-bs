GlobalModules.add_service('lessons')
angular.module('just.services.lessons', []).
factory('LessonsService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {

        var lessonsAPI = $resource('/api/v1/courses', {}, {
            lessons_list: { method: 'get', isArray: true, cache: false },
            create_lesson: { method: 'post', isArray: false },
        })

        function lessons_list(success) {
            lessonsAPI.lessons_list({}, {}, function(resp) {
                if (success) { success(resp) }
            })
        }

        function create_lesson(new_lesson, success) {
            lessonsAPI.create_lesson({}, {
                name: new_lesson.name,
                teacher_id: new_lesson.teacher_id,
                icon_url: new_lesson.icon_url,
                description: new_lesson.description,
                introduction: new_lesson.introduction,
                wish: new_lesson.wish,
                college_id: new_lesson.college_id,
                major_id: new_lesson.major_id,
                //outline: new_lesson.outline_list,
                //attachment: new_lesson.attachment_list
            }, function(resp) {
                if (success) { success(resp) }
            })
        }



        return {
            lessons_list: lessons_list,
            create_lesson: create_lesson
        }
    }
])

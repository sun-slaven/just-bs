GlobalModules.add_service('lesson')
angular.module('just.services.lesson', []).
factory('LessonService', ['$rootScope', '$resource', '$http',
    function($rootScope, $resource, $http) {
        var lessonAPI = $resource('/api/v1/courses/:course_id', {course_id : '@course_id'}, {
            delete_lesson: {method: 'delete' , isArray: false},
            get_lesson: {method: 'get' , isArray: false},
            update_lesson: {method: 'patch',isArray: false}
        })


        function get_lesson(course_id,success) {
            lessonAPI.get_lesson({}, {course_id: course_id}, function(resp) {
                if (success) { success(resp) }
            })
        }
        function delete_lesson(course_id,success) {
            lessonAPI.delete_lesson({}, {
                course_id: course_id
            }, function(resp) {
                console.log(resp)
                if (success) { success(resp) }
            })
        }
        function update_lesson(course,sunccess){
            lessonAPI.update_lesson({},{
                course_id: course.id,
                name: course.name,
                college_id: course.college.id,
                major_id: course.major.id,
                icon_url : course.icon.url,
                description: course.description,
                introduction: course.introduction,
                wish: course.wish
            },function(resp){
                if (callback) {callback(resp)};
            })
        }



        return {
            get_lesson: get_lesson,
            delete_lesson: delete_lesson,
            update_lesson: update_lesson
        }
    }
])

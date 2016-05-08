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
                if (success) { success(resp) }
            })
        }
        function update_lesson(course,callback){
            var icon_url = course.icon.url.replace('http://7xnz7k.com1.z0.glb.clouddn.com/','')
            lessonAPI.update_lesson({},{
                course_id: course.id,
                teacher_id: $rootScope.current_user.id,
                name: course.name,
                college_id: course.college.id,
                major_id: course.major.id,
                icon_url : icon_url,
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

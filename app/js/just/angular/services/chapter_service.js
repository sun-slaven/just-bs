GlobalModules.add_service('chapters')
angular.module('just.services.chapters', []).
factory('ChaptersService', ['$rootScope', '$resource',
    function($rootScope, $resource) {

        var chaptersAPI = $resource('/api/v1/courses/:course_id/chapters', { course_id: '@course_id' }, {
            get_chapters: { method: 'get', isArray: true },
            add_chapter: { method: 'post', isArray: false },
        })

        var chapterAPI = $resource('/api/v1/courses/:course_id/chapters/:chapter_id', { course_id: '@course_id' ,chapter_id: '@chapter_id'}, {
            update_chapter: { method: 'patch', isArray: true },
        })

        function get_chapters(lesson_id, success) {

            chaptersAPI.get_chapters({ course_id: lesson_id }, {}, function(resp) {
                if (success) { success(resp) }
            })
        }



        function add_chapter(lesson_id, chapter,success) {
            chaptersAPI.add_chapter({}, angular.extend({
                course_id: lesson_id,
            },chapter), function(resp) {
                if (success) { success(resp) }
            })
        }

        function update_chapter(lesson_id, chapter,success) {
            chapterAPI.update_chapter({}, angular.extend({
                course_id: lesson_id,
                chapter_id: chapter.id
            },chapter), function(resp) {
                if (success) { success(resp) }
            })
        }

        return {
            get_chapters: get_chapters,
            add_chapter: add_chapter,
            update_chapter: update_chapter
        }
    }
])

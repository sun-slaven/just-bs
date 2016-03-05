var justConst = (function() {
    function get_meta(name) {
        return $('meta[name=' + name + ']').attr('content');
    }
})();

angular.module('just.constants', []).constant('JustConst', justConst);

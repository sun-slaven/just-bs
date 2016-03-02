var GlobalModules = (function() {
    var services = []
    var controllers = []
    var directives = []
    var actions = []

    function add_service(service) { services.push(service) }

    function add_controller(controller) { controllers.push(controller) }

    function add_directive(directive) { directives.push(directive) }

    function add_action(action) { actions.push(action) }

    function get(others) {
        var all = []
        services.forEach(function(service) { all.push("just.services." + service) })
        controllers.forEach(function(controller) { all.push("just.controllers." + controller) })
        directives.forEach(function(directive) { all.push("just.directives." + directive) })
        actions.forEach(function(action) { all.push("just.actions." + action) })
        return all.concat(others)
    }

    return {
        add_service: add_service,
        add_controller: add_controller,
        add_directive: add_directive,
        add_action: add_action,
        get: get
    }
})();


// local storage

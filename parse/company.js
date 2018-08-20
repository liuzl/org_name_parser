var nf = nf || {};

nf.list = function() {
    var arr = [];
    for (var i = 0; i < arguments.length; i++) {
        arr.push(arguments[i]);
    }
    return arr;
}

nf.entity = function(type, value) {
    return {"type": type, "value": value};
}

nf.company = function(loc, name, type, scope) {
    return {"loc":loc, "name":name, "type":type, "scope":scope};
}


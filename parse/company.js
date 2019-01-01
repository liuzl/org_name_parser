var nf = nf || {};
var fmr = fmr || {};

fmr.list = function() {
    return arguments;
}

nf.entity = function(type, value) {
    return {"type": type, "value": value};
}

nf.company = function(loc, name, type, scope) {
    return {"loc":loc, "name":name, "type":type, "scope":scope};
}


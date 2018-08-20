var nf = nf || {};

nf.list = function() {
    var arr = [];
    for (var i = 0; i < arguments.length; i++) {
        arr.push(arguments[i]);
    }
    return arr;
}

nf.object = function(type, value) {
    return {"type": type, "value": value};
}

nf.company = function(loc, name, type, scope) {
    return {"loc":loc, "name":name, "type":type, "scope":scope};
}

nf.loc = function(l) {
    return {"loc_code": l};
}

nf.loc_province_city = function(p, c) {
    return {"province_code": p, "city_code": c};
}

nf.loc_province_county = function(p, county) {
    return {"province_code": p, "county_code": c};
}

nf.loc_city_county = function(city, county) {
    return {"city_code": city, "county_code": county};
}

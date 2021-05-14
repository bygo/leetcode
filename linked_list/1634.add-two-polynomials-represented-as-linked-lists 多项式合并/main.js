var addPoly = function (l, r) {
    if (l == undefined) {
        return r
    }

    if (r == undefined) {
        return l
    }
    if (l.power == r.power) {
        l.coefficient += r.coefficient
        if (l.coefficient == 0) {
            return addPoly(l.next, r.next)
        }
        l.next = addPoly(l.next, r.next)
        return l
    } else if (r.power < l.power) {
        l.next = addPoly(l.next, r)
        return l
    } else if (l.power < r.power) {
        r.next = addPoly(l, r.next)
        return r
    }
};

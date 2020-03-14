//  Author: Hang Yu
//  Create Time: 13/3/2020

/*  Check whether the characters in s can be replaced to get t */
function can_map(s, t) {
    // check empty string and mismatch length
    if ( s === "" || t === "" || s.length != t.length ) {
        return false;
    }
    mapping = {}
    for( var i = 0; i < s.length; i++) {
        char_s = s.charAt(i);
        char_t = t.charAt(i);
        // return False if this mapping conflicts with former one
        if (char_s in mapping) {
            if (char_t != mapping[char_s])
                return false;
        } else {
            mapping[char_s] = char_t;
        }
    }
    return true;
}

if (process.argv.length < 4) {
    console.log("false");
} else {
    var s = process.argv[2];
    var t = process.argv[3];
    if (can_map(s, t)) {
        console.log("true");
    } else {
        console.log("false");
    }
}

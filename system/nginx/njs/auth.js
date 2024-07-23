export default {rewrite_auth_set_cookie_header}

function rewrite_auth_set_cookie_header(r) {
    let rawCookies = r.variables['auth_cookie'];

    const regex = /(.*?); ((([eE][xX][pP][iI][rR][eE][sS])=([aA-zZ]{3}, [0-9]{2} [aA-zZ]{3} [0-9]{4} [0-9]{2}:[0-9]{2}:[0-9]{2} [aA-zZ]{3})(; |,|$))|(([dD][oO][mM][aA][iI][nN])=(.*?)(; |,|$))|(([pP][aA][tT][hH])=(.*?)(; |,|$))|(([hH][tT][tT][pP][oO][nN][lL][yY])(; |,|$))|(([sS][eE][cC][uU][rR][eE])(; |,|$))|(([sS][aA][mM][eE][sS][iI][tT][eE])=(.*?)(; |,|$))){1,}/g;

    if (rawCookies != undefined) {
        rawCookies = rawCookies.trim();

        if (rawCookies != "" && rawCookies.length > 0) {
            let headers = []
            let cookies = rawCookies.match(regex)

            for (let i = 0; i < cookies.length; i++) {
                let cookie = cookies[i].trim();

                if (cookie.slice(-1) == ','){
                    cookie = cookie.substring(0, cookie.length - 1);
                }

                if (cookie.length > 0) {
                    headers.push(cookie)
                }
            }

            r.headersOut['Set-Cookie'] = headers
        }
    }
}

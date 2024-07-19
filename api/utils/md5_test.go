package utils

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	sign := Md5ByString(fmt.Sprintf("%s%s%s%s", "jrc6ynvJoREDyf9", "mysqldump", "u3d%2Fcourse%2Fsql%2Fhetao-20221117134614.sql", "jrc6ynvJoREDyf9"))
	t.Log(sign)
}

// {"@timestamp":"2022-11-17T13:43:20.706+08:00","caller":"handler/loghandler.go:196","content":"[HTTP] 200 - GET /walnut3d-file-api/v1/object/exist?filename=u3d%2Fcourse%2Fsql%2Fhetao-20221117134319.sql\u0026scene=mysqldump
// \u0026sign=3b28243dc3b5830086ee31e944921356 - 106.75.84.144 - go-resty/2.7.0 (https://github.com/go-resty/resty)","duration":"0.2ms","level":"info","span":"e3cb982f23ec859a","trace":"a6fdf4edd6a638cca64f8ce4499d749e"}
// {"@timestamp":"2022-11-17T13:46:17.459+08:00","caller":"result/httpresult.go:45","content":"【API-ERR】 : 100008\nsign=449b6a9a60fa16c157155a1db6f72faa dsign=a565d85404e7dbdc3b3124c71a1b9a00 ","level":"error","span":"ada0
// 0c823395d075","trace":"72db8f01f172647b40fd0352b89b9df1"}
// {"@timestamp":"2022-11-17T13:46:17.460+08:00","caller":"handler/loghandler.go:196","content":"[HTTP] 200 - GET /walnut3d-file-api/v1/object/exist?filename=u3d%2Fcourse%2Fsql%2Fhetao-20221117134614.sql\u0026scene=mysqldump
// \u0026sign=449b6a9a60fa16c157155a1db6f72faa - 106.75.84.144 - go-resty/2.7.0 (https://github.com/go-resty/resty)","duration":"0.2ms","level":"info","span":"ada00c823395d075","trace":"72db8f01f172647b40fd0352b89b9df1"}

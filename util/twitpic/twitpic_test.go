package twitpic

import "testing"

const html = `
<!DOCTYPE html>
<html lang="en">
<head xmlns="http://www.w3.org/1999/xhtml"
      xmlns:og="http://ogp.me/ns#">

  <meta http-equiv="content-type" content="text/html; charset=utf-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <link rel="shortcut icon" type="image/x-icon" href="https://dfo9svwruwoho.cloudfront.net/images/favicon.ico">
  <title>@KBENSTEEN hierover  (:</title>

  <!-- Twitter Card -->
  <meta name="twitter:title" value="@KBENSTEEN hierover  (:" />
  <meta name="twitter:description" value="@KBENSTEEN hierover  (:" />
  <meta name="twitter:url" value="https://twitpic.com/9vkh19" />
  <meta name="twitter:card" value="summary_large_image" />
  <meta name="twitter:image" value="https://dn3pm25xmtlyu.cloudfront.net/photos/large/597218877.png?Expires=1604843739&Signature=BpTvwqaWyqouHMjsU-2Iiyyu4BeN5WrZuog4OrU8OASHwN6uHYlNonH9LizeRMaqOsPkDPfWniV738wOO1ppLZoGP7BEY31mesdZGUBvJo6H1OWRVSqwtRrwYqZ7UjwAcDLuFgiF66o-52GdLL4MPBEE5W8~sZBGX78EMP~4iVPgSG5XJ9SbsUo7ga6SjWefqFaZbRdqA14HvhmA8SPljeIORSBFvY~hqR~XOl7ujFrGzIEhNokOr4kLViqQ4H1VEKy~fSQh3G9g0aDXmi-pRLn9bHgq~WXi3i9~eEqyuZoer9P6t8~2BZclMUr3njfZ08BMfoa9UAzqKP6sihUwgA__&Key-Pair-Id=APKAJROXZ7FN26MABHYA" />
  <meta name="twitter:image:width" value="531" />
  <meta name="twitter:image:height" value="462" />
  <meta name="twitter:site" value="@twitpic" />
  <meta name="twitter:site:id" value="12925072" />
  <meta name="twitter:creator" value="@goosix" />
  <meta name="twitter:creator:id" value="30142929" />
</head>

<body>
    <img src="https://dn3pm25xmtlyu.cloudfront.net/photos/large/597218877.png?Expires=1604843739&Signature=BpTvwqaWyqouHMjsU-2Iiyyu4BeN5WrZuog4OrU8OASHwN6uHYlNonH9LizeRMaqOsPkDPfWniV738wOO1ppLZoGP7BEY31mesdZGUBvJo6H1OWRVSqwtRrwYqZ7UjwAcDLuFgiF66o-52GdLL4MPBEE5W8~sZBGX78EMP~4iVPgSG5XJ9SbsUo7ga6SjWefqFaZbRdqA14HvhmA8SPljeIORSBFvY~hqR~XOl7ujFrGzIEhNokOr4kLViqQ4H1VEKy~fSQh3G9g0aDXmi-pRLn9bHgq~WXi3i9~eEqyuZoer9P6t8~2BZclMUr3njfZ08BMfoa9UAzqKP6sihUwgA__&Key-Pair-Id=APKAJROXZ7FN26MABHYA" alt="@KBENSTEEN hierover  (:" style="width:531px;" />
</body>
</html>
`
const expectedURL = "https://dn3pm25xmtlyu.cloudfront.net/photos/large/597218877.png?Expires=1604843739&Signature=BpTvwqaWyqouHMjsU-2Iiyyu4BeN5WrZuog4OrU8OASHwN6uHYlNonH9LizeRMaqOsPkDPfWniV738wOO1ppLZoGP7BEY31mesdZGUBvJo6H1OWRVSqwtRrwYqZ7UjwAcDLuFgiF66o-52GdLL4MPBEE5W8~sZBGX78EMP~4iVPgSG5XJ9SbsUo7ga6SjWefqFaZbRdqA14HvhmA8SPljeIORSBFvY~hqR~XOl7ujFrGzIEhNokOr4kLViqQ4H1VEKy~fSQh3G9g0aDXmi-pRLn9bHgq~WXi3i9~eEqyuZoer9P6t8~2BZclMUr3njfZ08BMfoa9UAzqKP6sihUwgA__&Key-Pair-Id=APKAJROXZ7FN26MABHYA"
const expectedExt = "png"

func TestParse(t *testing.T) {
	img, err := Parse(html)
	if err != nil {
		t.Fatal("Parse returned unexpected error")
	}
	if img.URL != expectedURL {
		t.Errorf("wrong url: expected %s, got %s", expectedURL, img.URL)
	}
	if img.Ext != expectedExt {
		t.Errorf("wrong ext: expected %s, got %s", expectedExt, img.Ext)
	}
}

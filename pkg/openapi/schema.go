// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
	externalRef1 "github.com/unikorn-cloud/region/pkg/openapi"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+w9aW8bOZZ/hagdYKaxOkqyLNv6MnAn02mjk44RJ+ndib0GVXyS2K4ia0iWbcXQf1/w",
	"qFPUZcnp2d6gG4ilIh8f38V3sfQURDxJOQOmZDB6ClIscAIKhPkUxZlUIC5eX+Zf628JyEjQVFHOglHw",
	"cQbIjUMMJ9BB7zKp0BgQRvc4pgS9/vUKRZwpTBllU8RZPEcxfwCBIiwBRTMscKSXbF0zliVjEBJxgWbz",
	"dAZMtpBUWCiEGUHACHqgaoZwOUsPtbNaZoxeWKGES3XNhkcV6IgyFAObqlknaAVU455iNQtagUY7GJW7",
	"DVqBgH9lVAAJRkpk0ApkNIME693/RcAkGAX/0S0J17VPZfcuG4NgoED+ihMoibZYtAIuppjRr1iTbSNB",
	"q4MtVf0o14G+CN6p4L9DpDai7Matw7YA9SKICphuQ1o7DFECTNEJBbEC2RzcC+C6sCBBqh85oWBVTQBW",
	"8IonaabglZXED3aQecyZAmb+xGka08iwvPu71Pt6CuARJ2kM+s8EFCZYGfTqgh0sWoFMIdJP3O5IMArG",
	"4fHZ+AiG7TMMx+1Bf3zSPhuMB+3JoD8Zn+DhGAMEreCBi7uYY3LJeSyD0ZenYEIFPOA4Nh8IFRA5MlM2",
	"FSBl0ApSLlQwOg21GMGEPuq9fgl6Z/1Ob3ja6XXCbn8Q3BghUzzicTAKVJQGi9Z6gL1wOLR/v8OPwah3",
	"dnbWWCHsmP+6p0Er6J3o5czHXt+32k0rSHA0o8wQcBLjey4MaaKT4+Ep9El7cobH7cHxEWmf4SPcPu4d",
	"nRxPTk4H/eFYS0+Cp2aqhBgixa3kUakED0ZBNs6YyoJWcA9C2v30B51w4ITWMFMGo6NFIXoEJjiLld5u",
	"No5pdHF5Hsfcstxwm+FxnIvjYnGjAW0nk1FNvn4TVIGVx7qeODEs7LqTVlSeDp0ltbC7kSln0kl0Q5bt",
	"o+cLs9EQytlHaojUD/tH7fCkfdT72AtHg+PR4Pifmhc78K2hHg0brQGRwTAMyRDacDY8bg/Gg0Ebn4an",
	"7dPBZNyf4KPhSdgPShtp1sbQ752Rk3Yv1Oo0DHvt06gftQFOIBwOx2dHEdgp91SLA2XTK4VVJq2BtF8C",
	"+a6sf1ZlbQXS8fvJwye3N/thxqVyy7jv2z0jO/QeK7i4DEY58XqV9fW3ORW1dEqveN14dlBusPd8q/IB",
	"MPEZlXPUMCsdLeP1ufI5ZuLLdzvx3U58txN/nJ24eZ6hkH4rEVOpEJ80rYU05iJj9I4L1o5inpHbiAu4",
	"TTBlt+nd9JanwHBKbyOeJJzd4iiCVAGpmhRfOGBdmxmWaAzAUD7NBJMPNI51RDnJ4gmNY/2tnLNoJjjj",
	"mYznnWv23zxDCZ6jlMcxUgai5JmIwABIOKOKC0SVRJa+aMIF0oSIQaOx667GmLiw4HkeFQihRTmgzATn",
	"t27/Qcs+ua1TKKfOmJM5clOCrc+FHbZl0fLIw4cqBhNMNQ8sfJtdMBttIS4c7e1owkEixlWed7hmuOCO",
	"dVrRhEJMdhaqiLNJTKM9iZ9DWUF1XMqQyXZovCVOwMTWCMcCMJkjeKRSyW/NDYdXvgPp8jGMqxmIFspk",
	"huN4jtSMSpQAZlJjP0czfA/1fexK+QkXY0oIsP1IX4BZQftMgkCRAJMewLFEhBtBKjZQCJC2rjSGKcg/",
	"RiMesEQEGAWCxnOEMzXjgkqnD5b+eK6NV4QzaQdp/GsDr5nid8DyHVI2re9RRjwFY7EwQ+eXF4WiGTJp",
	"LWN/LWlzzRhEICUW8wp1EGdmijlVCAiUxlhNuEh2lQDKFAiG4ysQ9yD+oemznyxIA8hR2i8OzuIojiyh",
	"ohjT5Nvy+5yhjMFjCpE+lcwwxKMoEwJIndG4NlIJzCQFptwczMg10yNlFkUARPNFWxol5h10MbGQqGGo",
	"ZleEJbRQGgOWWiC0X4aoQljqZaiU2c4azLj6iWeM7Mc0xtXtRINZwbHKMQCkNKTFiWDM5rfl4Cfj2mkh",
	"mlBGUGned6Vgxpz2foU9qai9HClvrf1YdQxlaqatoIXmDt9vLPs+FHIbZPfgFFP7b/CYaqvloaqNjJbX",
	"dxbARgcvEHz2zk7Cdthrh72PYTgy/xfB5xk+jYZHJ2F7EOqYkQxw+4zgsH0yPDklk0EYkTNSBp/TzqAz",
	"o9NZAkkH98Kw05t2euF0XI3/ojT7CSc0ngej4IIpiNF/AWfoMsaKsixBp71h+BH97epuHuM7+CFo6Rky",
	"GA1aOrC5C0b9sBVM00zDivmURjh+xTNNhH4rSCDhYh6MhoNWkHACsVlEKsoihd5d9I9DHVrM5rIyrafj",
	"I0aMxJ2/e61xzcEc9XeIG57DzPXBhRu0u6iYsPDF0hT9dr//sdcfhYNR76iQFDwcTM76w7P20RDC9uCo",
	"12+PT0mvfdwnZ0fkeHg2PqmkKbJx1u+Hg/Z9r9M/7gzb0zRrH/ePO6fHnfC4fRIBGfSOB1W5cSwngt6D",
	"ZlUxOnCsNuH3eS/ULP7Z/dMPQx12F/z99fPF64tzDZa79C8BhxHj4xiClid6nuTiSmBMsfbJ7kAwI1sx",
	"ZdmjDrCxoFiT1x3Xvpi7FUj6Fd7QH4NRL2wFkk/UAxbw2Y4z6JTVmGAUONLoifdUqAzHLpmjn+VfaFv3",
	"QiJqhWi9hNoxSM2wQliACRmxovogcUEBlVVXqrDLnuT7O5sUuCryCevT/S6H4ALWjs1CpSCUq1WV6YYm",
	"JLcQykfouWqeaiGQSlA2DRa1tMSq+W4IurhEmBABUvohFZmMlYDMiE1wykTLgQ60/IC/XM7e5dU/Wzn5",
	"UhKzQOOmwJCPf4dILadJ3e7kKoauTKLkrJUV3lIFidwtc1OXp0WBLxYCz5fRNRnhjVKno9plWauazAMx",
	"x2Vhr3RsQ/KIVuP4Ll+rYhu3J8qVnrG1MDXm+mWj2LzDZ7NkXDm015M6P3saBb065cts8poyuuKoyBba",
	"sMEtQZlX0xqJzJ2qlbWpTVIV2DbX2IJmW5rFVeZQytmltVe/wHwZzNXVz4U9u4O5s+hxzB9k6cJWafei",
	"hKtI2gayVKf5TEyTPDmGKOU89ihymYxfh7obZhYtkufNxX+rLoVWHDQNAXE2NkfjZqf9bysiNRqsFJgy",
	"d/8MoytLq7sHcao5/CaAX003lT49ckQrAChTMC06bpbIW0Ddjb47nWE1Ij/3CKtJt45LKLuwUHqbjjOf",
	"Pj0b/f0OYo+MbjyNbdfHdmcE6Fg7SwlW8E2O5txvOsRJvPdZugtXn8tAWyxfybKLvAi5fACb+ECfHJl0",
	"GVpkowuNCWfwfmJCzioJKAluFq36d0V182Zx02QwJeuWvnjt96Ur5dJ1dDBArvLB3vMor0V/yGLw8SJ/",
	"jkQWAzLBt83V4k2HUaUO7dtg8VhzmoHS0JASeDKhkYGfprHNppuVFdcrAMsSQ+Wirg32jxtf5GJK3t4+",
	"Si6UWUN7CCbfbvtf+QRh+1BgNgWfSa5Uz32QgZEmlBaiTEswvQf9yMSUNrDkQrUQnSCtK3RCgaxYsKjP",
	"r1GVi9cfUD6wsjWqUEKnM9srzObo4vJ+oPd7cXk/RJTZeYwrbOukFf1aImdTfcpmgBWtquZpgUuVfSpK",
	"g1aQkdTDt4ZJKaWosqLjbYU0NxtEey3xajIuNwj5VgaoplUe2tUV04OaNTPOCuTq7tMxm+05YN6Ey9cW",
	"6KKSF/L2T6cgsKJsiuRcKkiQG+21WEU6aTtIdrSzvEZHGFeIwIQyIFqoTdOAtsraQMRYgVSb3VNHLZ+s",
	"rGokXuOT79uIjw7Zh4/WtOEn+PGt+RCMhkfGC8s/9jycqgYHW0QkbrhXNO82CeY9j7MEqrq6i2IZP6xs",
	"Omri+pN54o7QDZsuGpG28Cqsv7BY1WO0DoJnxo5RQu6FaIMdmS4ASb+Ci0NczSbBjzTJkrykjDPF2zLC",
	"MZRGbPmQySSI187RbFTW3JNcrCUaYwnDQRtYxAkQ06dAp5mw6Q7KJlwk9m99sBpIMneispQzFOOMRUY0",
	"7dBgFIznCjbqb0GoCttz3vmU2s+gpkRX0pfFOCRBaYPkCSyLHrImpH/YCqQPXIXkY85jwGxpbzlY3z52",
	"rvIt4fYGGAgauSp1AlJi7Zg0890cZ2rW9+zZD/UcaWMDDqplGoLHFDNibZ8Rx58/frx0Q7S8dJDpKpAm",
	"4a4lieQD35/r1XNvKHJdFuPM5uYtXHBGUeMnKCgs5nnPlwZu9eP88kIi0yyD1Axr4FxC6WWh8dytVXcp",
	"m81a1UrwbRRTYPrbZlU3YzJLtTsCeq4tWd0aFrYKmKbFQwdF9WYIBUnKBRY0nt9mDN9jGmNbxcknFqvm",
	"X0wFZqqxqvkuX7Jauq+0QCWgZpzc6qfGG1tCPQFCcQ6kbODx+dWeOnZTMj6DGGuaO0lD9uk4730xEDaf",
	"16ubRvbUkvK0f4vHEH/GceYNfuyR/ks2BjMYxXq0/jaDFlLz1Jlg0/qjBa/ohNC+gasrRZihMVwzygg8",
	"aufFCrsOk7X0G2XDSoHQS/7Pl7B9dt7+J25/vfnb30flp/Zt5+YpbA17i8qIH/7+F5+vtQMhqk3Wa3L1",
	"o6cAx7GLdg+chmgUBZ4apqfZBr7xJl956azeKTqGmLOpdIHkesFrLLosbTe7kXlzNeQlKLwlc5dpXmmh",
	"X3cL8RCULpfam8hLea0l5IvOyjxP5RTcOEpxrM+NchcCMJG2QVlQBR5PYK0F/FilSOWR6ybj5oOxHjib",
	"JprFhoim+9SccAk3BWmm4NEfUuWZ6QNJi9cm6tAVTw9ZtVV4+tY0h/ly28/kuKcCvEpwi3HGawCbqylb",
	"xUp3IGN3jD+wxu2Q6kdzhBJoPLZn1s1+Zll8KxNh44gPyxnhpyVZt630tpPGR1xFE6jbAds8G4Oyia3C",
	"1ydYQVsP9zcu+Hj57VsGPHh4LFRrRwtjjMquHYkSkntYkSVKMFM0yhMmDWfi/vqa/Of1dafyz74OwwqB",
	"eUkHYY1U2voJ+XHuF0nTQP0w467OQmri6bWp9Yax7cXcLbC9mK9K+2eM/iurAF9RAEg4MaHMxp3butIW",
	"O88hbtg5ru/bgd923w0doyZIqZJ8CxX7aG89OPWisuZ4O5/790y6JmSbJyOc/VXlHffXDLN5/SzQY2aA",
	"YzVzwaQNO7XbP6EKTQRPbNKFEWzCwWtWYGD33blmwX4xicJTj4ozhMWYKqHjXIWn7nYKIzYIWfZJ/OXq",
	"81yuchD+HK0/DNK8N49s9cIA2boVwMK82Zs0xmdYk8LXXsrW2fkdueLJ3z+jWdZbXDNPvknVd+VJv1Xt",
	"9xn7fXaB+Llr7UDfMRagUfFkuH+bgUscgZue2xiMCBAaGXNayXo383qtaou4ByUBppScoFeXn5Dtza3G",
	"Bwg60w4yHcFlql9EM6ogUplYcWylq1xeVmSQX11+kv7Ub56pX56NE54xo1+QziABgWOkRyPK0Jsf/dBc",
	"o/PBpGmaZnnXX97Xvh5VO8qgSH/coqvGEK8A7shxIMFcX3bM++OfZba2Mzr72q5pmr2zFxGW9/Hm8lNN",
	"UDubfcstV9tk7ZsrvxANi80fgIp+E6U3Ust3L9uq+iURn+i7ERVdf3P5SaIin4ywRBKA5anH91d+zV2l",
	"Xobam5SquLCyRk78/eX16yzeuN0Nae7wbxEWRP5Q7tSPWH594rCS8dlCbVoTt1hOjopdqW+0VWfs3vam",
	"xMhLQs0Di1o1yeHuk7TMxaGbfbWX+hupzpltY/izuTm2h+ObeDlmqTfuBpGXt/ZeEaITRJlUOI7B0+ef",
	"Xz7aAMQlE1r5bVUnOIWhWhU/Qvwy5vetu0z6x+ixI9phePj+yqsfSw0wlRHLbCxvfa1zL/UoGy4aj/IB",
	"CzXvjnXA62fgC7cSTQqP+IDgnZu9KK+0HRT8LxboukaoKsXdIEtvAvJO8bRb9otu3xPl7tTl8faSdJgF",
	"ru3FvOtgcyzuiFMwoWB2icNhxHtFHLbyADh0rFDYyeKG5GFBv7+q3YNcuhBHmWm8MSGSHlW+5MG1v6ii",
	"s3edJ+W7W3mwjSwB91zQPDTdPtfhN+UzJ+gSIoaLhxHNz0tbbMb9WCG9TqXbqsoukwOo+k7VO6ytSkah",
	"FWA2P5AztTZ83PHy6ksESDRvhdsrOiqOCc9mjaGi46z2kt6cBYJHd9qauavOe9O8OFB8Zt5mavhkyRzL",
	"DkImKV02hNhGVVvRTXF0pyXI5Ymr6AOZYWWKiOZi9gHw/6U4Bpv42zPASHgVB3sDfO+VfRYLE0Jtkfuy",
	"ZvQP5enbqtjS7W7TnolSAYUjXLQT5//m/nnHl7J3Pam+biDzxHM3kX71jH9dHAJbZ6IMoGVzZy58RJmg",
	"an6l6eQaY0yvXv3lHctYvLfCyll+FMm8zW4MWIBwb/aov4bElHNi/mDZ6HrgzJNXnMDSl59EHIyCmVKp",
	"HHW7titEzTs1Xna4mHYtyt37frc2X4dJEU/NtrQt0Rg9A6aZV7vkYh7ZNwBQNuH+wCdvaL4CcU8jME0g",
	"7jVK0rTh06Ks5C6ULpWOYjoBFM2jGK5ZghmeQgLMf4MJuTdLSZRkUrmQau76JN+9Re5FF0ZPr9kMMHHO",
	"I1UxVG6QVbCtvawh7PQ6oXF/rLYEo+CoE3aObGF4ZjjaxSnt3ve61f4g2X2qv+d80Y1W3tB6lV/WddTR",
	"CE7B4x3ryM3UvXNY5mRyaajqeuY+VS6q5g2Sb0Cdp/Rz730Vyfc1FIsrZI2XIvfDcJWdKcZ1V70TddEK",
	"BtvM3+91gmaV3kFX8b63yKxzdNB1lt8RZxYZHHSRpddYLVrB8YHZsu5VZ1Wba3oa/Nb2y41pUqv+oMOK",
	"/odySHfVjxQYUFvqpmubk92n4n3/uyrsgZBubZzq+XEDcw2S+3L8r0wjhUQYMXgor/OWVsPehQJSuYCi",
	"Da5YtiCXXG40IZeOjpc5jjWbkv98wHy13FV+YaC77ucFFks2qr+rjfpuovY0UWcHXWTpLaX/1ibqsU14",
	"lGm3xHzfngqepeZVxtTUrA9hd7pPxa+7LIqeQZ9nbL5HuFBh01/jdBzvqOUW2LP1/FXt92h2VdB93pT8",
	"XYu/Oxq7avG/+am+eZbnx66ML5B5XIFPprewYiae5wVk6kDG4Y/2Br4bm+/G5v+gy2Azh7L7lP/Y16K7",
	"sjcqz264xrut8gvSJhjyREuZBHer2NRN8T7tGa68EexZaYcPdkMf3HZ+cpt5ThJirxfo/ul0+k8f4G8+",
	"H5d/X2+XrIBH1VaVkXJNy18wso+iVQtRL6dnF3YnL65mjZcPf9ey/ydatlj8bwAAAP//eQhSKah1AAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(path.Dir(pathToFile), "https://raw.githubusercontent.com/unikorn-cloud/region/main/pkg/openapi/server.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

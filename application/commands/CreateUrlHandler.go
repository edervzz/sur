package commands

import (
	"hash/fnv"
	"math"
	"strconv"
	"sur/application/messages"
	"sur/application/validators"
	"sur/domain/entities"
	"sur/domain/ports"
	"time"
)

type CreateUrlHandler struct {
	repository ports.IUrlRepository
}

func (h *CreateUrlHandler) Handler(request messages.CreateUrlRequest) (*messages.CreateUrlResponse, error) {
	// 1. Request validated
	if err := validators.CreateUrlValidator(request); err != nil {
		return nil, err
	}
	// 2. Create url code
	request.Hash, request.Code = createUrlCode(request)
	// 3. Map entity and save repository
	url := entities.Url{
		TenantId: request.TenantId,
		Code:     request.Code,
		IsActive: true,
		Url:      request.Url,
		Hash:     request.Hash,
		CreateAt: time.Now().String(),
		CreateBy: "sys",
	}
	if err := h.repository.Create(url); err != nil {
		return nil, err
	}

	return &messages.CreateUrlResponse{
		Url: "https://surly.dev/" + url.Code,
	}, nil
}

func NewCreateUrlHandler(r ports.IUrlRepository) *CreateUrlHandler {
	return &CreateUrlHandler{
		repository: r,
	}
}

func createUrlCode(request messages.CreateUrlRequest) (uint32, string) {
	hash := fnv.New32a()
	hash.Write([]byte(request.Url))
	key := hash.Sum32()
	keys, codes := CreateShortUrl(key)
	codes = PerfomMirrorCodes(keys, codes)
	code := ""
	for i := 0; i < len(codes); i++ {
		code += codes[i]
	}

	return key, code
}

func CreateShortUrl(hash uint32) (map[int]string, map[int]string) {

	keyMap := make(map[int]string)
	codeMap := make(map[int]string)
	i := 0
	value := ""
	// get key (type string) and calculate pairs
	key := strconv.FormatUint(uint64(hash), 10)
	pairs := int(math.Round(float64(len(key)) / float64(2)))
	ckey := key
	// convert pairs into codes
	for ; pairs > 0; pairs-- {
		clen := len(ckey)
		if clen-2 >= 0 {
			value = ckey[clen-2 : clen]
			ckey = ckey[:clen-2]
		} else if clen-1 >= 0 {
			value = ckey[clen-1 : clen]
			ckey = ckey[:clen-1]
		} else {
			value = "0"
			ckey = ""
		}
		idx, _ := strconv.Atoi(value)
		code := GetCode(idx)
		keyMap[i] = value
		codeMap[i] = code
		i++
	}
	// reverse keyMap
	aux := keyMap
	keyMap = make(map[int]string)
	for i := 0; i < len(aux); i++ {
		keyMap[i] = aux[len(aux)-i-1]
	}
	// reverse codeMap
	aux2 := codeMap
	codeMap = make(map[int]string)
	for i := 0; i < len(aux2); i++ {
		codeMap[i] = aux2[len(aux2)-i-1]
	}

	return keyMap, codeMap
}

func PerfomMirrorCodes(keyMap map[int]string, codeMap map[int]string) map[int]string {
	ckey := make(map[int]int)
	for i, v := range keyMap {
		ckey[i], _ = strconv.Atoi(v)
	}
	// check mirrors
	even := 8
	times := int(len(codeMap) / 2)
	for i := 0; i < times; i++ {
		j := len(codeMap) - 1 - i
		if codeMap[i] == codeMap[j] {
			if ckey[i] < ckey[j] {
				codeMap[j] = strconv.Itoa(even)
			} else if ckey[i] > ckey[j] {
				codeMap[j] = strconv.Itoa(even - 2)
			}
		}
		even = even - 4
	}
	return codeMap
}

func GetCode(idx int) string {
	codes := make(map[int]string)
	codes[0] = "a"
	codes[1] = "e"
	codes[2] = "i"
	codes[3] = "o"
	codes[4] = "u"
	codes[5] = "A"
	codes[6] = "E"
	codes[7] = "I"
	codes[8] = "O"
	codes[9] = "U"

	codes[10] = "B"
	codes[11] = "C"
	codes[12] = "D"
	codes[13] = "F"
	codes[14] = "G"
	codes[15] = "H"
	codes[16] = "J"
	codes[17] = "K"
	codes[18] = "L"
	codes[19] = "M"
	codes[20] = "N"
	codes[21] = "P"
	codes[22] = "Q"
	codes[23] = "R"
	codes[24] = "S"
	codes[25] = "T"
	codes[26] = "V"
	codes[27] = "W"
	codes[28] = "X"
	codes[29] = "Y"
	codes[30] = "Z"

	codes[31] = "b"
	codes[32] = "c"
	codes[33] = "d"
	codes[34] = "f"
	codes[35] = "g"
	codes[36] = "h"
	codes[37] = "j"
	codes[38] = "k"
	codes[39] = "l"
	codes[40] = "m"
	codes[41] = "n"
	codes[42] = "p"
	codes[43] = "q"
	codes[44] = "r"
	codes[45] = "s"
	codes[46] = "t"
	codes[47] = "v"
	codes[48] = "w"
	codes[49] = "x"
	codes[50] = "y"
	codes[51] = "z"

	codes[52] = "B"
	codes[53] = "C"
	codes[54] = "D"
	codes[55] = "F"
	codes[56] = "G"
	codes[57] = "H"
	codes[58] = "J"
	codes[59] = "K"
	codes[60] = "L"
	codes[61] = "M"
	codes[62] = "N"
	codes[63] = "P"
	codes[64] = "Q"
	codes[65] = "R"
	codes[66] = "S"
	codes[67] = "T"
	codes[68] = "V"
	codes[69] = "W"
	codes[70] = "X"
	codes[71] = "Y"
	codes[72] = "Z"

	codes[73] = "b"
	codes[74] = "c"
	codes[75] = "d"
	codes[76] = "f"
	codes[77] = "g"
	codes[78] = "h"
	codes[79] = "j"
	codes[80] = "k"
	codes[81] = "l"
	codes[82] = "m"
	codes[83] = "n"
	codes[84] = "p"
	codes[85] = "q"
	codes[86] = "r"
	codes[87] = "s"
	codes[88] = "t"
	codes[89] = "v"
	codes[90] = "w"
	codes[91] = "x"
	codes[92] = "y"
	codes[93] = "z"

	codes[94] = "1"
	codes[95] = "3"
	codes[96] = "5"
	codes[97] = "7"
	codes[98] = "9"
	codes[99] = "0"

	return codes[idx]
}

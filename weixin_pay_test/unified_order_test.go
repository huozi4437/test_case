package wx_helper

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dchest/uniuri"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"testing"
	"time"
)

func TestUnifiedOrder(t *testing.T) {
	InitParams(&FixedParamsStc{AppId: "", AppSecret: "", MchId: "", MchKey: "",	NotifyUrl: "http://"})

	nonceStr := uniuri.NewLenChars(32, []byte("0123456789abcdefghijklmnopqrstuvwxyz"))
	resp, err := UnifiedOrder(&UnifiedOrderRequest{
		NonceStr:       nonceStr,
		Body:           "测试的",
		Attach:         "{}",
		OutTradeNo:     nonceStr,
		TotalFee:       1,
		SpbillCreateIp: "183.14.132.224",
		ProductId:      "test20191120",
	})
	if err != nil {
		fmt.Println("UnifiedOrder err", err)
		return
	}
	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
	fmt.Println("Code_url:", resp.CodeUrl)
}

func TestNewUUID(t *testing.T) {
	nonceStr := uuid.New().String()[:32]
	t.Logf("new uuid:%s len:%d\n", nonceStr, len(nonceStr))

	//nonceStr := uuid.NewV4().String()
	//t.Logf("new uuid:%s len:%d\n", nonceStr, len(nonceStr))
}

func TestBson(t *testing.T) {
	id := bson.NewObjectId().String()
	t.Logf("id:%s len:%d", id, len(id))
}

func TestRandNum(t *testing.T) {
	for i := 0; i < 10; i++ {
		n1 := rand.Int31n(999999)
		n2 := rand.Int()
		t.Logf("n1:%d n2:%d", n1, n2)
	}
}

func TestNumber(t *testing.T) {
	//tm := time.Now().Format("20060102150405Mon")
	//t.Logf("tm:%s", tm)
	orderNum := time.Now().Format("20060102") + fmt.Sprintf("%07d", 111)
	t.Logf("orderNum:%s", orderNum)
}

func TestHashStr(t *testing.T) {
	h := sha256.New()
	b := make([]byte, 512)
	n, err := h.Write(b)
	digest := h.Sum(nil)
	if err != nil {
		t.Fatalf("write err: %v", err)
	}

	hex.EncodeToString(digest)
	t.Logf("str:%s b:%v size:%d n:%d", hex.EncodeToString(digest), b, h.Size(), n)
}

func TestMd5Str(t *testing.T) {
	//h := md5.New()
	//h.Write([]byte(string(time.Now().UnixNano())))

	//t.Logf("md5: %v",hex.EncodeToString(h.Sum(nil)))
}
package wx_helper

import (
	"fmt"
	"testing"
)

func TestQueryOrder(t *testing.T) {
	//Init("wx00eae1aec8b32727", "e3e3e8ed110d5b43fd6157f71caffef7", "1490342292", "Al9wI3AERmvovZEUXbRNORnhZy42eQTL")
	resp, err := QueryOrder("4jbiasljq1eovowxde4u539n42nh148d")
	fmt.Println(err, resp)
}

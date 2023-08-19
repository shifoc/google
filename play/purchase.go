package play

import (
   "io"
   "net/http"
   "os"
   "strconv"
   "strings"
   "time"
)

type Native_Platform map[int64]string

var Platforms = Native_Platform{
   // com.google.android.youtube
   0: "x86",
   // com.miui.weather2
   1: "armeabi-v7a",
   // com.kakaogames.twodin
   2: "arm64-v8a",
}

func (n Native_Platform) String() string {
   var b []byte
   b = append(b, "native platform"...)
   for key, value := range n {
      b = append(b, '\n')
      b = strconv.AppendInt(b, key, 10)
      b = append(b, ": "...)
      b = append(b, value...)
   }
   return string(b)
}

const Sleep = 4 * time.Second

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(doc string) error {
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/purchase",
      strings.NewReader("doc=" + doc),
   )
   if err != nil {
      return err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

func (r Response) Write_File(name string) error {
   data, err := io.ReadAll(r.Body)
   if err != nil {
      return err
   }
   return os.WriteFile(name, data, 0666)
}

type Response struct {
   *http.Response
}

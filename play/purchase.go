package play

import (
   "io"
   "net/http"
   "os"
   "strings"
)

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

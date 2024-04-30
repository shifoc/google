package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (g GoogleCheckin) Sync(device GoogleDevice) error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(10, func(m *protobuf.Message) {
         for _, v := range device.Feature {
            m.Add(1, func(m *protobuf.Message) {
               m.AddBytes(1, []byte(v))
            })
         }
         for _, v := range device.Library {
            m.AddBytes(2, []byte(v))
         }
         for _, v := range device.Texture {
            m.AddBytes(4, []byte(v))
         }
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(15, func(m *protobuf.Message) {
         m.AddBytes(4, []byte(device.ABI))
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(18, func(m *protobuf.Message) {
         m.AddBytes(1, []byte("am-unknown")) // X-DFE-Client-Id
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(19, func(m *protobuf.Message) {
         m.AddVarint(2, google_play_store)
      })
   })
   m.Add(1, func(m *protobuf.Message) {
      m.Add(21, func(m *protobuf.Message) {
         m.AddVarint(6, gl_es_version)
      })
   })
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/sync",
      bytes.NewReader(m.Encode()),
   )
   if err != nil {
      return err
   }
   if err := g.x_dfe_device_id(req); err != nil {
      return err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}

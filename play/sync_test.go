package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Sync(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google-play"
   for _, platform := range ABIs {
      var checkin GoogleCheckin
      fmt.Println(platform)
      Phone.ABI = platform
      if err := checkin.Checkin(Phone); err != nil {
         t.Fatal(err)
      }
      err := os.WriteFile(home+"/"+platform+".bin", checkin.Data, 0666)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if err := checkin.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      if err := checkin.Sync(Phone); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

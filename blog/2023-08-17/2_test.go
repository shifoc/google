package play

import (
   "154.pages.dev/http/option"
   "fmt"
   "testing"
)

func Test_Account_Lookup(t *testing.T) {
   option.No_Location()
   option.Trace()
   setup, err := new_embedded_setup()
   if err != nil {
      t.Fatal(err)
   }
   account, err := setup.account_lookup("srpen6@gmail.com")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(account.host_gaps)
   fmt.Println()
   fmt.Println(account.TL())
}

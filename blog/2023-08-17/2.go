package play

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "strings"
)

func (e embedded_setup) account_lookup(email string) (*account_lookup, error) {
   body, err := func() (url.Values, error) {
      f_req, err := json.Marshal([]any{
         "",
         e.user_hash(),
         []any{},
         nil,
         "US",
         nil,
         nil,
         2,
         1,
         1,
         []any{},
         email,
      })
      if err != nil {
         return nil, err
      }
      v := make(url.Values)
      v.Set("f.req", string(f_req))
      // THE SECOND STRING CANNOT BE EMPTY, WE ARE FUCKED.
      v.Set("bgRequest", `["identifier",""]`)
      return v, nil
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://accounts.google.com/_/lookup/accountlookup",
      strings.NewReader(body.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(e.host_gaps)
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Header.Set("Google-Accounts-Xsrf", "1")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   {
      b, err := httputil.DumpResponse(res, true)
      if err != nil {
         return nil, err
      }
      println(string(b))
   }
   var lookup account_lookup
   lookup.host_gaps, err = host_gaps(res)
   if err != nil {
      return nil, err
   }
   lookup.body, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &lookup, nil
}

type account_lookup struct {
   // this is needed for /_/signin/challenge:
   host_gaps *http.Cookie
   // this is needed for /_/signin/challenge:
   body []byte
}

func (a account_lookup) TL() (string, error) {
   a.body = bytes.TrimPrefix(a.body, []byte(")]}'"))
   var body []any
   if err := json.Unmarshal(a.body, &body); err != nil {
      return "", err
   }
   return body[0].([]any)[2].(string), nil
}

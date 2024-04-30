# Google Play

Download APK from Google Play or send API requests

## how to build?

This module works with Windows, macOS or Linux. Download Go [1] and extract
archive. Then download Google Play Zip and extract archive. Then navigate to:

```
internal/play
```

and enter:

```
go build
```

1. https://go.dev/dl

## tool examples

[sign in](https://accounts.google.com/embedded/setup/v2/android) with your Google
Account. then get authorization code (`oauth_token`) cookie from
[browser&nbsp;storage][1]. should be valid for 10 minutes. then exchange
authorization code for refresh token (`aas_et`):

~~~
play -o oauth2_4/0Adeu5B...
~~~

[1]://firefox-source-docs.mozilla.org/devtools-user/storage_inspector

create a file containing `X-DFE-Device-ID` (GSF ID) for future requests:

~~~
play -d
~~~

get app details:

~~~
> play -i com.google.android.youtube
creator:  Google LLC
offer: 0 LBP
version:  19.16.38
downloads:  Apr 24, 2024
file: APK APK APK APK
android version: 8.0 and up
downloads: 16.22 billion
name: YouTube
size: 113.80 megabyte
version code: 1545729472
changelog:  For new features, look for in-product education &amp; notifications sharing the feature and how to use it!
~~~

acquire app. only needs to be done once per Google account:

~~~
play -i com.google.android.youtube -a
~~~

download APK. you need to specify any valid version code. the latest code is
provided by the previous details command. if APK is split, all pieces will be
downloaded:

~~~
play -i com.google.android.youtube -v 1540222400
~~~

## goals

1. [Pixel 6](//wikipedia.org/wiki/Pixel_6) (2021)
2. [Android 12](//wikipedia.org/wiki/Android_12) (2021)
3. [Google Play](//wikipedia.org/wiki/Google_Play) 29 (2022)

non goals:

email/password login. up to Android 4.4 (2013), the login is protected with TLS
fingerprinting, which is difficult but possible to bypass. since Android 5
(2014), Google uses bot-guard via JavaScript to protect the login. I do not
know how to reverse that, and I did not find any implementations online.

## thanks
- https://github.com/3052/google
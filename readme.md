# Google Play

Download APK from Google Play or send API requests

## How to install?

This module works with Windows, macOS or Linux. Download Go [1] and extract
archive. Then download Google Play Zip and extract archive. Then navigate to:

~~~
googleplay-main/cmd/play
~~~

and enter:

~~~
go build
~~~

1. https://go.dev/dl

## Tool examples

Before trying to Sign in, make sure your location is correct, to avoid
geo-blocking. You can test by logging into your Google account with a web
browser. Also, make sure the Google account you are using has logged into the
Play Store at least once before, using a physical or virtual Android device.
Create a file containing token (`aas_et`) for future requests:

~~~
play -email EMAIL -passwd PASSWORD
~~~

Create a file containing `X-DFE-Device-ID` (GSF ID) for future requests:

~~~
play -device
~~~

Get app details:

~~~
> play -d com.google.android.youtube
creator: Google LLC
file: APK APK APK APK
installation size: 49.53 megabyte
downloads: 14.05 billion
offer: 0 USD
title: YouTube
upload date: May 19, 2023
version: 18.20.34
version code: 1537856960
changelog: for new features, look in the education section
~~~

Purchase app. Only needs to be done once per Google account:

~~~
play -d com.google.android.youtube -purchase
~~~

Download APK. You need to specify any valid version code. The latest code is
provided by the previous details command. If APK is split, all pieces will be
downloaded:

~~~
play -d com.google.android.youtube -v 1537856960
~~~

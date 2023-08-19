# SDK version 22

1. Android Studio
2. Pixel 3a XL
3. API Level 22
4. Android 5.1 Google APIs image

then:

~~~
play -d com.android.vending -v 82011800
~~~

check version:

~~~
name='com.android.vending'
versionCode='82011800'
versionName='20.1.18-all [0] [PR] 311592326'
~~~

older versions are available, but they are buggy. You can also pull the APK from
the `google_apis_playstore` image, but it is buggy as well. start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_22 -writable-system
~~~

Install like this:

~~~
adb remount
adb push com.android.vending-82011800.apk /system/priv-app/Phonesky.apk
adb reboot
~~~

install system certificate. then:

~~~
mitmproxy
~~~

then set proxy. start Play Store. click Sign in. enter Email and click Next.
Enter password and click Next. if app closes for update, start again. click
Sign in. enter Email and click Next. Enter password and click Next. click I
agree. Under Use basic device backup, move slider to left. click ACCEPT. under
Apps, click an app. result:

~~~
POST /auth HTTP/1.1
Host: android.clients.google.com
Accept-Encoding: identity
Connection: Keep-Alive
User-Agent: GoogleAuth/1.4 (generic_x86 LMY48X); gzip
app: com.google.android.gms
content-type: application/x-www-form-urlencoded
device: 37ad95573...

ACCESS_TOKEN=1&
add_account=1&
device_country=us&
droidguard_results=CgZsBhLUyOzSEFsAAHdU9h5a1xWSAFpFJRNTS5M9AC4R4ZGh9VMLBADDJ6...&
google_play_services_version=11055270&
is_dev_key_gmscore=1&
lang=en_US&
sdk_version=22&
service=ac2dm&
Email=s...&
androidId=37ad95573...&
Token=oauth2_4%2F0Adeu5BWFf0FwBd5Ja8ZxOt3Ln7sjZnoaY2XMhYep6mcdzmuEB_puWAudaiPB...
~~~

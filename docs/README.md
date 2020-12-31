<p align="center">
  <img alt="kontrolio logo" src="../assets/logo.png" height="300" />
  <h3 align="center">kontrolio-cli</h3>
  <p align="center">Kontrolio's CLI time clock, clock card machine, punch clock, or time recorder.</p>
</p>

## ✈️ Offline mode

Kontrolio works offline by default. If you want to save your data on Kontrolio's platform, check the [online mode](#-online-mode)

## 🌐 Online mode

In order to register your data remotely on Kontrolio's database, you need to create an account on [kontrolio.com](https://kontrolio.com) and set the generated API Key on the [configuration file](#-configuration).

## 🧰 Configuration

Kontrolio has a [configuration](../.kontrolio.example.yaml) file that allows you to configure it. This is optional, you don't need to create this file unless you want to customize some default behavior.

The file must be named `.kontrolio.yaml`. Kontrolio looks for
this file in your home directory (`/home/marcelo` on Linux, `/Users/Marcelo` on macOs, `C:\Users\Marcelo` on Windows).

This is how it should look like:

```yaml
# Required if you want to save your data on Kontrolio's database.
api_key: "YOUR_API_KEY"
```

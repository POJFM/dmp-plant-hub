# 🌱 plant-hub

Ultra advanced smart zavlažovací systém.

## 🖥️ Server

- chad GoLang + GraphQL stack
- Postgres DB :5420
- communication with HW modules/sensors via GPIO

## 🖼️ Web app (client)

- React
- [Design on Figma](https://www.figma.com/file/7gMKRPDOrkKOT5GKmOmfsu/PlantHub?node-id=0%3A1)

## ✅ TODO

- [ ] podle public ip zjistit GPS souřadnice pro weather API

## ⚙️ Setup cross-compile on Arch

```
yay -S arm-linux-gnueabihf-glibc-headers
yay -S arm-linux-gnueabihf-gcc-stage2 arm-linux-gnueabihf-glibc
```

To anyone having issues building `arm-linux-gnueabihf-gcc-stage1` (`arm-linux-gnueabihf-glibc-headers` dependency), ensure that your makepkg.conf doesn't include "-Werror=format-security" in cflags. This might be causing the build to fail. <sup>[[1]](https://aur.archlinux.org/packages/arm-linux-gnueabihf-gcc-stage1/#pinned-806072)</sup>

## Links
[go requests](https://zetcode.com/golang/getpostrequest/)
[go method handlers](https://medium.com/geekculture/develop-rest-apis-in-go-using-gorilla-mux-5869b2179a18)
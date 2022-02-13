# 🌱 plant-hub

Ultra advanced smart irrigation system.

## 🖥️ Server and main program

- Chad GoLang
- Postgres DB :5420
- communication with HW modules/sensors via GPIO
- GraphQL API for DB data
- REST API for live data from sensors

## 🖼️ Web app (client)

- React
- Tailwind
- Material UI
- [Design on Figma](https://www.figma.com/file/7gMKRPDOrkKOT5GKmOmfsu/PlantHub?node-id=0%3A1)

## ⚙️ Setup cross-compile on Arch

```
yay -S arm-linux-gnueabihf-glibc-headers
yay -S arm-linux-gnueabihf-gcc-stage2 arm-linux-gnueabihf-glibc
```

To anyone having issues building `arm-linux-gnueabihf-gcc-stage1` (`arm-linux-gnueabihf-glibc-headers` dependency), ensure that your makepkg.conf doesn't include "-Werror=format-security" in cflags. This might be causing the build to fail. <sup>[[1]](https://aur.archlinux.org/packages/arm-linux-gnueabihf-gcc-stage1/#pinned-806072)</sup>
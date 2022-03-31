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

## 🏠 Setup local subnet

- Install dhcp package:
  `yay -Syu dhcpcd`
- Configure subnet in `/etc/dhcpd.conf`:

```
subnet 192.168.0.0 netmask 255.255.255.224 {
  range 192.168.0.10 192.168.0.20;
}
```

- Add your network card to subnet
  `sudo ip addr add 192.168.0.1/24 dev enp3s0`
- Restart dhcp daemon
  `systemctl restart dhcpd4`

## 📦 Build docker image for arm

```
# create bob the builder
docker buildx create --name bob
# switch to bob
docker buildx use bob
docker buildx inspect --bootstrap
docker login
docker buildx
# build and push image
docker buildx build --platform linux/arm64,linux/arm/v7 -t tassilobalbo/planthub-client --push .
```
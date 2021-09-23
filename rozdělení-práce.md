# Projekt PlantHub

## Popis: Chytrá péče o vaše rostliny a monitorování jejich životních funkcí.

## Doba trvání práce na projektu: 1 - 1,5 roku

## Náklady na projekt:

### Senzory a moduly:

- Raspberry Pi
- čerpadlo
- kapacitní senzor vlhkosti půdy
- senzor teploty a vlhkosti vzduchu
- senzor hladiny vody

### Rozdělené práce

#### Filip:

- Full Stack Web server
  - UI/Frontend - React.js, TypeScript (pro interaktivní ovládání a monitoring GardenBota)
  - Databáze - MongoDB
  - Backend - Golang (ještě není finální rozhodnutí)
  - Konfigurace RPI jakožto web server a IoT zařízení zároveň (znalosti z předmětu OSL)
  - Bezpečnost zařízení

#### Jakub:

- Sestavení a vytvoření modulu chytré správy zahrady (Dále jen “GardenBot”)
  - Návrh a sestavení obvodu plošných spojů a zbylých součástek hardwarové části projektu
  - Vytvoření programu pro ovládání GardenBota a nastavení RPI pro sběr dat ze senzorů

#### Společně:

- Nastavení dataflow a databáze
- interaktivní ovládání GardenBota z webového UI

# Maturitní práce  

## Školní rok:	2021/22 

## Autoři:		Filip Sikora, Jakub Vantuch 

## Název:		GardenBot 

## Zadání v bodech: 

### Vantuch
1. hardwarové řešení - stanovení cílů, volba hardware (řídící jednotka, senzory, akční členy) Vantuch 

2. návrh obvodu a plošného spoje Vantuch 

3. fyzická realizace Vantuch 

4. naprogramování řídící jednotky Vantuch 

### Sikora

5. softwarové řešení - stanovení cílů, volba sw platformy a konkrétního software SIkora 

6. konfigurace RPi jako webového serveru SIkora 

7. vytvoření databáze SIkora 

8. vytvoření front end SIkora 

9. vytvoření back end SIkora 

10. nastavení bezpečnosti SIkora 

11. ověření funkčnosti - společně 

## část maturitní práce Jakub Vantuch
### 1.Hardwarové řešení
#### Mé cíly
- Mým cílem je vytvořit funkční zavlažovací systém ovládaný mikropočítačem RaspberryPi, s automatickým spouštěním na základě vlhkosti půdy.
- Systém GardenBot dále získává informace o teplotě, vlhkosti a tlaku vzduchu a promítá je ve svém webovém rozhraní.
- Jelikož voda časem z nádrže dojde systém GardenBot snímá stav hladiny vody v nádrži a včas upozorní, že je třeba doplnit vodu.

#### Zvolený hardware
- mikropočítač RaspberryPi 3
- senzor teploty, vlhkosti a tlaku vzduchu DHT11
- kapacitní čidlo pro měření vlhkosti půdy
- senzor hladiny vody
- Ponorné mini čerpadlo eses
- tranzistor 2n2222
- LED diody
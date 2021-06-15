# Maturitní práce  

## Školní rok:	2021/22 

## Autoři:		Filip Sikora, Jakub Vantuch 

## Název:		GardenBot 

## Zadání v bodech: 

### Vantuch
1. hardwarové řešení - stanovení cílů, volba hardware (řídící jednotka, senzory, akční členy)

2. návrh obvodu a plošného spoje

3. fyzická realizace

4. naprogramování řídící jednotky

### Sikora

5. softwarové řešení - stanovení cílů, volba sw platformy a konkrétního software

6. konfigurace RPi jako webového serveru

7. vytvoření databáze

8. vytvoření front end

9. vytvoření back end

10. nastavení bezpečnosti

### Společně

11. interaktivní ovládání GardenBota z webového UI

12. ověření funkčnosti

## 1. část maturitní práce - Jakub Vantuch
#### Mé cíly
- Mým cílem je vytvořit funkční zavlažovací systém ovládaný mikropočítačem RaspberryPi, s automatickým spouštěním na základě vlhkosti půdy.
- Systém GardenBot dále získává informace o teplotě, vlhkosti a tlaku vzduchu a promítá je ve svém webovém rozhraní.
- Jelikož voda časem z nádrže dojde systém GardenBot snímá stav hladiny vody v nádrži a včas upozorní, že je třeba doplnit vodu.

### 1.Hardwarové řešení
#### Mikropočítač RaspberryPi 3
- RaspberryPi je jednodeskový mikropočítač s operačním systémem o velikosti platební karty. Neobsahuje displej ani úložiště pouze konektory USB, Ethernet, HDMI a konektor pro univerzální použití (GPIO). V projektu GardenBot představuje RaspberryPi hlavní řídící jednotku obvodu a zároveň webový server pro webové rozhraní. 

#### Senzory
  #### Senzor teploty a vlhkosti vzduchu DHT11 (Digital hum temp)
  - Senzor DHT11 se skládá z jednotky pro měření teploty, jednotky pro měření vlhkosti a převodníku.
  - Teplotu měří senzor thermistorem. Thermistor je keramický polovodič, který zmenšuje svou rezistivitu když se okolní teplota zvýší.
  - Vlhkost měří senzor na základě rezistivity substrátu umístěného mezi dvěma elektrodami. Tento substrát zachytává vlhkost a vytváří tak vodivé prostředí. 

  #### Kapacitní čidlo pro měření vlhkosti půdy
  - kapacitní čidlo se skládá ze dvou vodivých desek a převodníku. Čidlo funguje na způsob kapacitoru avšak jeho kapacita je ovlivněna vlhkostí, která ovlivňuje dielektrikum mezi dvouma deskama.

  #### Ultrasonický senzor
  - Ultrasonický senzor vydává zvukové vlny na vysoké frekvenci, neslyšitelné pro lidské ucho. Poté čeká, až se zvuk odrazí zpět, a vypočítá vzdálenost na základě času měřeného od vysílání zvukové vlny k zpětnému příjmutí.

  - Všechny naměřené údaje jsou v převodníku daného senzoru přepočítány na jednotky dané veličiny a odeslány analogovým signálem do řídící jednotky.

  #### Ponorné mini čerpadlo eses
  - Toto čerpadlo se skládá z DC motoru na němž je upevněna centrifuga pro čerpání vody a vlastního pouzdra, z kterého vede otvor pro napojení odtokové hadičky.

### 2. Návrh obvodu a plošného spoje
#### Zapojení
![alt text](./circuit.png)
  #### Senzor teploty a vlhkosti vzduchu DHT11
  - je zapojen do zdroje 5V a země a jeho signální pin je připojen k pinu GPIO23.

  #### Kapacitní čidlo pro měření vlhkosti půdy
  - je zapojeno do zdroje 3.3V a země a jeho signální pin je připojen k pinu GPIO22.

  #### Ultrasonický senzor
  - je zapojen do zdroje 5V a země a jeho piny jsou pomocí I2C sběrnice připojen k pinům GPIO2 a GPIO3.

  #### Ponorné mini čerpadlo eses
  - je zapojeno přes tranzistor do zdroje 5V a země. Jeho spuštění a vypnutí je ovládáno otevřením a zavřením tranzistoru, jehož báze je připojena k pinu GPIO18.

  #### LED dioda
  - je zapojena přes 1K ohmový rezistor do země a na pin GPIO27. LED dioda slouží jako přídávná signalizace nízké hladiny vody v nádrži.

#### Funkcionalita
  #### 1. Fáze inicializace
  - Půda musí být ze začátku suchá. Kapacitní čidlo pro měření vlhkosti půdy zastrčíme co nejhlouběji do půdy. RaspberryPi bude chvíli sbírat data a pak je zprůměruje do hodnoty, která bude sloužit jako limit pro spuštění čerpadla.
  - Ve webovém rozhraní jde navíc ještě manuálně nastavit hranice teploty a vlhkosti vzduchu pro spuštění čerpadla.
  - Nastavit se dá také množství vody, které bude přečerpáno při jednom spuštění a jaká je hranice pro přijatelnou výšku hladiny vody v nádrži. Pokud nejsou tyto hodnoty uvedeny čerpadlo bude vodu přečerpávat dokud se nezmění hodnota kapacitního čidla pro měření vlhkosti půdy a ultrasonický senzor použije výchozí nastavení.
  #### 2. Fáze měření
  - Kapacitní čidlo pro měření vlhkosti půdy a senzor teploty a vlhkosti vzduchu průběžně posílají naměřená data do RaspberryPi, kde se ukládají do databáze. Jestliže naměřené hodnoty překročí limitní hodnoty RaspberryPi pošle signál pro otevření tranzistoru což spustí čerpadlo, které začne čerpat vodu. Voda se čerpá tak dlouho dokud kapacitní čidlo pro měření vlhkosti půdy nezmění svou hodnotu nebo dokud není vyčerpán limit vody na jedno spuštění.
  #### 3. Fáze kontroly
  - Po ukončení přečerpávání se spustí ultrasonický senzor a změří výšku hladiny vody. Naměřená data poté odešlě do RaspberryPi kde se uloží do databáze. Pokud bude naměřená hodnota nižší než je limitní hodnota začne blikat LED dioda a RaspberryPi odešle upozornění o doplnění nádrže do webového rozhraní. Jakmile bude hladina doplněna signalizace se vypne.